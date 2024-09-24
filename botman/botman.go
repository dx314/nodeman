package botman

import (
	"errors"
	"log"
	"sync"
)

var DEBUG bool = false

type NodeManager interface {
	Allocate(nodeID int, proc int, bot *Bot) bool
	NodeProcs(nodeID int) []*Bot
	NodeStats(nodeID int) map[string]string
	RemoveBot(bot *Bot)
}

type Node interface {
}

type BotStatus string

const (
	PENDING   BotStatus = "PENDING"   // not yet sent to simulation
	NOT_READY BotStatus = "NOT_READY" // sent but not ready
	READY     BotStatus = "READY"     // ready status received fromm simulation
	ASSIGNED  BotStatus = "ASSIGNED"  // assignment has been confirmed by simulation
	DONE      BotStatus = "DONE"      // assignment complete

)

type Bot struct {
	ID        int
	Status    BotStatus
	RequestID *int
	NodeID    int
	sync.RWMutex
}

func (bot *Bot) Ready() bool {
	bot.RLock()
	defer bot.RUnlock()
	return bot.Status == READY
}

func (bot *Bot) Pending() bool {
	bot.RLock()
	defer bot.RUnlock()
	return bot.Status == PENDING
}

func (bot *Bot) Working() bool {
	bot.RLock()
	defer bot.RUnlock()
	return bot.Status == ASSIGNED || bot.Status == READY && bot.RequestID != nil
}

func (bot *Bot) Exists() bool {
	bot.RLock()
	defer bot.RUnlock()
	return bot.Status != PENDING && bot.Status != DONE
}

func (bot *Bot) Done() bool {
	bot.RLock()
	defer bot.RUnlock()
	return bot.Status == DONE
}

func (bot *Bot) WarmingUp() bool {
	bot.RLock()
	defer bot.RUnlock()
	return bot.Status == NOT_READY || bot.Status == PENDING
}

func (bot *Bot) SetRequestID(reqID *int) {
	bot.Lock()
	defer bot.Unlock()
	bot.RequestID = reqID
}

type BotManager struct {
	bots      map[int]*Bot
	requests  []int
	nodeman   NodeManager
	work      chan func()
	paused    bool
	nextBotId int
}

func NewBotManager(nodeman NodeManager) *BotManager {
	bm := &BotManager{
		bots:      map[int]*Bot{},
		requests:  []int{},
		nodeman:   nodeman,
		work:      make(chan func()),
		paused:    false,
		nextBotId: 1,
	}
	go bm.Run()
	return bm
}

func (bm *BotManager) doWork(cmd string, fn func()) *sync.WaitGroup {
	if DEBUG {
		log.Println("bot cmd", cmd)
	}
	wg := &sync.WaitGroup{}
	wg.Add(1)

	bm.work <- func() {
		fn()
		wg.Done()
		if DEBUG {
			log.Println("bot cmd finished", cmd)
		}
	}

	return wg
}

func (bm *BotManager) Requests() int {
	return len(bm.requests)
}

func (bm *BotManager) ForEach(fn func(*Bot) bool) {
	bots := map[int]*Bot{}
	bm.doWork("forEach", func() {
		for key, value := range bm.bots {
			bots[key] = value
		}
	}).Wait()
	for _, bot := range bots {
		if !fn(bot) {
			return
		}
	}
}

//makeBot fills nodes empty procs with bots, returns created count
func (bm *BotManager) makeBot(nodeID int, procs []*Bot) int {
	count := 0
	skipped := 0
	for i, proc := range procs {
		if proc != nil {
			skipped++
			continue
		}
		bot := &Bot{
			Status: PENDING,
			ID:     bm.nextBotId,
			NodeID: nodeID,
		}
		bm.nextBotId++
		bm.bots[bot.ID] = bot
		bm.nodeman.Allocate(nodeID, i, bot)
		count++
	}
	if skipped == len(procs) {
		return 0
	}
	stats := bm.nodeman.NodeStats(nodeID)
	if DEBUG {
		log.Printf("%d bots on node %d (%s, cpu: %s) pending creation\n", count, nodeID, stats["size"], stats["cpu"])
	}
	return count
}

func (bm *BotManager) FindIdle() *Bot {
	var bot *Bot
	bm.ForEach(func(b *Bot) bool {
		if b.Status == READY {
			bot = b
			return false
		}
		return true
	})
	return bot
}

func (bm *BotManager) Ready(botID int) *sync.WaitGroup {
	return bm.doWork("Ready", func() {
		bm.bots[botID].Status = READY
	})
}

func (bm *BotManager) Delete(bot *Bot) *sync.WaitGroup {
	if DEBUG {
		log.Printf("deleting bot %d on node %d\n", bot.ID, bot.NodeID)
	}
	return bm.doWork("Delete", func() {
		delete(bm.bots, bot.ID)
		bm.nodeman.RemoveBot(bot)
	})
}

func (bm *BotManager) NotReady(botID int) *sync.WaitGroup {
	return bm.doWork("NotReady", func() {
		bot := bm.bots[botID]
		bot.Lock()
		defer bot.Unlock()
		bot.Status = NOT_READY
	})
}

func (bm *BotManager) Done(botID int) *sync.WaitGroup {
	return bm.doWork("Done", func() {
		bot := bm.bots[botID]
		bot.Lock()
		defer bot.Unlock()
		bot.Status = DONE
	})
}

func (bm *BotManager) Bot(botID int) *Bot {
	return bm.bots[botID]
}
func (bm *BotManager) Assign(botID int, reqID *int) *sync.WaitGroup {
	return bm.doWork("Assign", func() {
		bot := bm.bots[botID]
		bot.Lock()
		defer bot.Unlock()
		bot.Status = ASSIGNED
		bot.RequestID = reqID
	})
}

//MakeBots fills empty processor slots with bots
func (bm *BotManager) MakeBots(nodeID int, procs []*Bot) *sync.WaitGroup {
	if DEBUG {
		log.Printf("making bots on node %d\n", nodeID)
	}

	return bm.doWork("MakeBots", func() {
		bm.makeBot(nodeID, procs)
	})
}

var ErrNoRequests = errors.New("no requests available")

func (bm *BotManager) PopRequest() (int, error) {
	var reqid int
	wg := bm.doWork("PopRequest", func() {
		if len(bm.requests) == 0 {
			reqid = -1
			return
		}
		reqid, bm.requests = bm.requests[0], bm.requests[1:]
	})
	wg.Wait()
	if reqid == -1 {
		return -1, ErrNoRequests
	}
	return reqid, nil
}
func (bm *BotManager) AddRequest(reqID int) *sync.WaitGroup {
	return bm.doWork("AddRequest", func() {
		bm.requests = append(bm.requests, reqID)
	})
}

func (bm *BotManager) Run() {
	for {
		select {
		case fn := <-bm.work:
			fn()
		}
	}
}
