package main

import (
	"botmanager/botman"
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/navidys/tvxwidgets"
)

//NodeManager manages the nodes and the bots
type NodeManager struct {
	nodes        map[int]*Node
	botman       *botman.BotManager
	assgignments map[int]*botman.Bot
	in           io.WriteCloser
	out          io.Writer
	node_manager chan string
	bots_done    int8
	autoScale    bool
	interactive  bool
	quietmode    bool
	stepthrough  bool
	work         chan func(nm *NodeManager)

	capGraph *tvxwidgets.UtilModeGauge
	ui       *UIView
	data     map[string][]float64
	started  *time.Time
	sync.WaitGroup
	sync.RWMutex
}

//Busy checks whether the node is at capacity
func (node *Node) Busy() bool {
	for _, bot := range node.proc {
		if bot == nil || bot.Status == botman.READY {
			return false
		}
	}
	return true
}

//Idle returns true if the node is idle
func (node *Node) Idle() bool {
	if node.status != READY {
		return false
	}
	for _, bot := range node.proc {
		if bot == nil {
			continue
		}
		if bot.Working() || bot.WarmingUp() {
			return false
		}
	}
	return true
}

//Empty checks whether the node is not running any bots
func (node *Node) Empty() bool {
	for _, bot := range node.proc {
		if bot != nil {
			return false
		}
	}
	return true
}

//NewNodeMan generates a new NodeManager
func NewNodeMan(in io.WriteCloser) *NodeManager {

	bm := &NodeManager{
		nodes:        map[int]*Node{},
		node_manager: make(chan string),
		in:           in,
		work:         make(chan func(bm *NodeManager)),
		bots_done:    -1,
		data: map[string][]float64{
			"nodes":           []float64{0},
			"nodes.small":     []float64{0},
			"nodes.medium":    []float64{0},
			"nodes.large":     []float64{0},
			"emptyNodes":      []float64{0},
			"nodesWarmingUp":  []float64{0},
			"botRequests":     []float64{0},
			"idleBots":        []float64{0},
			"botsWarmingUp":   []float64{0},
			"busyBots":        []float64{0},
			"demand":          []float64{0},
			"slots.available": []float64{0},
			"slots.total":     []float64{0},
		},
	}

	bm.botman = botman.NewBotManager(bm)

	go bm.Run()

	return bm
}

//SetAutoScaling turns autoscaling off or on
func (nm *NodeManager) SetAutoScaling(on bool) {
	nm.autoScale = on
}

//SetQuietMode turns quietmode off or no
func (nm *NodeManager) SetQuietMode(on bool) {
	nm.quietmode = on
}

//SetInteractive turns interactive off or on
func (nm *NodeManager) SetUI(ui *UIView) {
	nm.ui = ui
}

//SetStout sets the out pipe
func (nm *NodeManager) SetStout(out io.Writer) {
	nm.out = out
}

//SetInteractive turns interactive off or on
func (nm *NodeManager) SetGraph(graph *tvxwidgets.UtilModeGauge) {
	nm.capGraph = graph
}

//SetInteractive turns interactive off or on
func (nm *NodeManager) SetInteractive(on bool) {
	nm.interactive = on
}

//SetInteractive turns interactive off or on
func (nm *NodeManager) SetStepThrough(on bool) {
	nm.stepthrough = on
}

//Node returns a node with the supplied ID
func (nm *NodeManager) Node(nodeID int) *Node {
	node, ok := nm.nodes[nodeID]
	if !ok {
		log.Panic("unable to retreive node")
	}
	return node
}

//NodeStats returns general stats for the Node
func (nm *NodeManager) NodeStats(nodeID int) map[string]string {
	node, ok := nm.nodes[nodeID]
	if !ok {
		log.Panic("attempting to retrieve stats for nonexistent node")
	}

	return map[string]string{
		"size": node.size,
		"cpu":  strconv.Itoa(node.cpu),
	}
}

//Allocate allocates a bot to a cpu core
func (nm *NodeManager) Allocate(nodeID int, proc int, bot *botman.Bot) bool {
	nm.RLock()
	defer nm.RUnlock()

	if _, ok := nm.nodes[nodeID]; !ok {
		log.Panic("attempting to allocate bots to nonexistent node")
	}
	if nm.nodes[nodeID].proc[proc] != nil && bot != nil {
		return false
	}
	nm.nodes[nodeID].proc[proc] = bot
	return true
}

//NodeProcs returns running processes
func (nm *NodeManager) NodeProcs(nodeID int) []*botman.Bot {
	return nm.nodes[nodeID].proc
}

//Run starts the simulation listener
func (nm *NodeManager) Run() {
	next_node_id := 1
	for {
		select {
		case fn := <-nm.work:
			nm.Lock()
			fn(nm)
			nm.Unlock()
		case size := <-nm.node_manager:
			nm.Lock()
			if DEBUG {
				nm.log(fmt.Sprintf("provisioning %s node\n", size))
			}
			node, ok := NodeTemplates[size]
			if !ok {
				panic("attempted to provision a node type that does not exist")
			}

			node.id = next_node_id
			next_node_id++

			node.proc = make([]*botman.Bot, node.cpu)
			node.status = PENDING

			nm.nodes[node.id] = &node
			nm.Unlock()
			nm.botman.MakeBots(node.id, node.proc).Wait()
		}
	}
}

//DoWork work scheduler
func (nm *NodeManager) DoWork(fn func(nm *NodeManager)) *sync.WaitGroup {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	nm.work <- func(nm *NodeManager) {
		fn(nm)
		wg.Done()
	}
	return wg
}

//Send sends a message to the simulation
func (nm *NodeManager) Send(msg interface{}) error {
	if nm.started == nil {
		now := time.Now()
		nm.started = &now
	}
	if msg == nil {
		_, err := io.WriteString(nm.in, "[]\n")
		return err
	}

	jsonData, err := json.Marshal(msg)
	if err != nil {
		nm.log(fmt.Sprintf("could not marshal json: %s\n", err))
		return err
	}
	nm.log(fmt.Sprintf("---> %s\n", string(jsonData)))

	if _, err := io.WriteString(nm.in, string(jsonData)+"\n"); err != nil {
		return err
	}

	return err
}

//RemoveBot removes a bot from the node and manager
func (nm *NodeManager) RemoveBot(bot *botman.Bot) {
	node, ok := nm.nodes[bot.NodeID]
	if !ok {
		panic("node not found")
	}
	for i, b := range node.proc {
		if b == bot {
			nm.Lock()
			node.proc[i] = nil
			nm.Unlock()
		}
	}
}

//NodeReady sets the node to ready
func (nm *NodeManager) NodeReady(nodeID int) {
	node, ok := nm.nodes[nodeID]
	if !ok {
		panic("node not found")
	}
	node.status = READY
}

var dots = 1

//log is a drop in replacement for std log
func (nm *NodeManager) log(text string) {
	if nm.quietmode {
		dotText := ""
		for i := 0; i < dots; i++ {
			dotText += "."
		}
		dots++
		if dots > 20 {
			dots = 0
		}
		fmt.Printf("\r%s", dotText)
		return
	}
	if nm.out != nil {
		if _, err := io.WriteString(nm.out, text); err != nil {
			panic(err)
		}
		return
	}
	log.Print(text)
}

//Receive receives messages from the simulation and processes them
func (nm *NodeManager) Receive(text string) {
	nm.log(fmt.Sprintf("<--- %s\n", text))
	var messages []ReceiveMsg
	err := json.Unmarshal([]byte(text), &messages)
	if err != nil {
		type Done struct {
			Kind string `json:"kind"`
			Data struct {
				AverageCost    float64 `json:"average_cost"`
				AverageLatency float64 `json:"average_latency"`
				Score          float64 `json:"score"`
			} `json:"data"`
		}
		var doneMsg Done
		err := json.Unmarshal([]byte(text), &doneMsg)
		if err != nil {
			nm.log(fmt.Sprintf("unable to parse json: %s\n", text))
			log.Panic(err)
		}
		now := time.Now()
		text := fmt.Sprintf(`
	Bot Wrangling Stats:
		Average Cost: %.2f
		Average Latency: %.2f
		Score: %.2f
		
		Real Runtime: %v
`, doneMsg.Data.AverageCost, doneMsg.Data.AverageLatency, doneMsg.Data.Score, now.Sub(*nm.started))
		nm.SetQuietMode(false)
		nm.log(text)
		if nm.ui == nil {
			nm.Done()
		} else {
			FinishModal(nm.ui, text)
		}
		return
	}
	for _, msg := range messages {
		switch msg.Kind {
		case Request.NODE_READY:
			if msg.Data.NodeID == nil {
				log.Panicf("node id missing in json: %s", text)
			}

			nodeID, err := strconv.Atoi(*msg.Data.NodeID)
			if err != nil {
				log.Panic(err, "not able to convert node id to int")
			}

			nm.NodeReady(nodeID)
		case Request.BOT_REQUEST:
			if msg.Data.BotRequestID == nil {
				log.Panicf("bot request id missing in json: %s", text)
			}
			nm.botman.AddRequest(*msg.Data.BotRequestID).Wait()
		case Request.BOT_READY:
			if msg.Data.BotID == nil {
				log.Panicf("bot id missing in json: %s", text)
			}
			botID, err := strconv.Atoi(*msg.Data.BotID)
			if err != nil {
				log.Panic(err, "not able to convert bot id to int")
			}
			nm.botman.Ready(botID).Wait()
		case Request.BOT_DONE:
			if msg.Data.BotID == nil {
				log.Panicf("bot id missing in json: %s", text)
			}
			botID, err := strconv.Atoi(*msg.Data.BotID)
			if err != nil {
				log.Panic(err, "not able to convert bot id to int")
			}

			bot := nm.botman.Bot(botID)
			nm.botman.Done(botID).Wait()
			nm.botman.MakeBots(bot.NodeID, nm.Node(bot.NodeID).proc).Wait()
		case Request.BOTS_DONE:
			nm.DoWork(func(nm *NodeManager) {
				nm.bots_done = 0
			}).Wait()
		}

	}
}

// ClusterStats all the stats you could ever need
type ClusterStats struct {
	slots     int
	available int
	taken     int
	pending   int
	requests  int
}

//Stats returns the full amount of cpu slots available and other important stats
func (nm *NodeManager) Stats() *ClusterStats {
	nm.RLock()
	defer nm.RUnlock()

	slots := 0
	available := 0
	taken := 0
	pending := 0
	for _, node := range nm.nodes {
		slots += node.cpu
		for _, b := range node.proc {
			if b == nil {
				continue
			}
			if b.Ready() {
				available++
			} else if b.Working() {
				taken++
			} else {
				pending++
			}
		}
	}
	return &ClusterStats{
		slots:     slots,
		available: available,
		taken:     taken,
		pending:   pending,
		requests:  nm.botman.Requests(),
	}
}

//NextTick send next batch of messages to the simulation
func (nm *NodeManager) NextTick() {
	send := []SendMsg{}

	waiting := false

	nm.botman.ForEach(func(bot *botman.Bot) bool {
		if bot.Done() {
			nm.botman.Delete(bot).Wait()
		} else if bot.Ready() && nm.botman.Requests() > 0 && bot.RequestID == nil {
			reqid, err := nm.botman.PopRequest()
			if err != nil {
				return true
			}
			bot.SetRequestID(&reqid)
			send = append(send, SendMsg{
				Kind: Command.BOT_ASSIGN,
				Data: &MessageData[int]{
					BotRequestID: bot.RequestID,
					BotID:        &bot.ID,
				},
			})
			nm.botman.Assign(bot.ID, bot.RequestID)

		} else if bot.Pending() {
			waiting = true
			if nm.Node(bot.NodeID).status != READY {
				return true
			}
			send = append(send, SendMsg{
				Kind: Command.BOT_CREATE,
				Data: &MessageData[int]{
					BotID:  &bot.ID,
					NodeID: &bot.NodeID,
				},
			})
			nm.botman.NotReady(bot.ID)
		} else if bot.WarmingUp() {
			waiting = true
		} else if bot.Working() {
			waiting = true
		}
		return true
	})

	stats := nm.Stats()
	frames_idle_for := 15
	//avg := getAvg(nm.data["demand"])
	rec_avg := 10
	if len(nm.data["demand"]) > 30 {
		rec_avg = getAvg(nm.data["demand"][len(nm.data["demand"])-30:])
	}

	powering_down := 0

	if nm.autoScale && nm.bots_done == -1 {
		if stats.available < 10 || stats.available-powering_down+stats.pending < rec_avg/10*2 {
			diff := rec_avg - (stats.available - powering_down + stats.pending)
			for i := 0; i < diff/2; i++ {
				nm.node_manager <- "medium"
			}
		}
	}

	nm.RLock()
	for _, node := range nm.nodes {
		if node.status == PENDING {
			send = append(send, SendMsg{
				Kind: Command.NODE_CREATE,
				Data: &MessageData[int]{
					NodeID:   &node.id,
					NodeSize: &node.size,
				},
			})
			node.status = NOT_READY
			waiting = true
		} else if node.status == READY {
			nm.RUnlock()
			if nm.autoScale && nm.bots_done == -1 {
				if node.Idle() {
					node.framesIdle++
				} else {
					node.framesIdle = 0
				}
				if node.Idle() && node.framesIdle > frames_idle_for {
					frames_idle_for++
					powering_down += node.cpu
					node.status = DESTROYING
					send = append(send, SendMsg{
						Kind: Command.NODE_DESTROY,
						Data: &MessageData[int]{
							NodeID: &node.id,
						},
					})
					for _, bot := range node.proc {
						if bot != nil {
							nm.botman.Delete(bot).Wait()
						}
					}
					nm.Lock()
					delete(nm.nodes, node.id)
					nm.Unlock()
				} else {
					nm.botman.MakeBots(node.id, node.proc).Wait()
				}
			} else {
				nm.botman.MakeBots(node.id, node.proc).Wait()
			}
			nm.RLock()
		} else {
			waiting = true
		}
	}
	nm.RUnlock()

	if nm.bots_done == 0 && nm.botman.Requests() == 0 && len(send) == 0 && !waiting {
		nm.bots_done = 1
		if err := nm.Send([]SendMsg{{
			Kind: Command.BOTS_ACK_DONE,
		}}); err != nil {
			panic(err)
		}
		return
	}

	if err := nm.Send(send); err != nil && nm.bots_done <= 0 {
		nm.log(err.Error())
	}
}

//Listen to stdout from the simulation
func (nm *NodeManager) Listen(stout io.ReadCloser) {
	scanner := bufio.NewScanner(stout)
	go func() {
		for scanner.Scan() {
			m := scanner.Text()
			nm.Receive(m)

			if !nm.stepthrough {
				nm.NextTick()
			}
			nm.TicKReport()
		}
	}()
}

//TicKReport reports on the current cluster's stats
func (nm *NodeManager) TicKReport() {
	stats := nm.Stats()
	busyNodes := map[string]int{
		"small":  0,
		"medium": 0,
		"large":  0,
	}
	nodes := map[string]int{
		"all":    len(nm.nodes),
		"small":  0,
		"medium": 0,
		"large":  0,
	}
	emptyNodes := 0
	nodesWarmingUp := 0
	botRequests := nm.botman.Requests()
	idleBots := 0
	botsWarmingUp := 0
	busyBots := 0

	nm.DoWork(func(nm *NodeManager) {
		for _, node := range nm.nodes {
			nodes[node.size]++
			if node.status == PENDING || node.status == NOT_READY {
				nodesWarmingUp++
			} else if node.Busy() {
				busyNodes[node.size]++
			} else if node.Empty() {
				emptyNodes++
			}
		}
	}).Wait()

	nm.botman.ForEach(func(bot *botman.Bot) bool {
		if bot.Ready() {
			idleBots++
		} else if bot.Pending() || bot.WarmingUp() {
			botsWarmingUp++
		} else if !bot.Done() && bot.RequestID != nil {
			busyBots++
		}
		return true
	})

	nm.data["nodes"] = append(nm.data["nodes"], float64(nodes["all"]))
	nm.data["nodes.small"] = append(nm.data["small"], float64(nodes["small"]))
	nm.data["nodes.medium"] = append(nm.data["medium"], float64(nodes["medium"]))
	nm.data["nodes.large"] = append(nm.data["large"], float64(nodes["large"]))
	nm.data["emptyNodes"] = append(nm.data["emptyNodes"], float64(emptyNodes))
	nm.data["nodesWarmingUp"] = append(nm.data["nodesWarmingUp"], float64(nodesWarmingUp))
	nm.data["botRequests"] = append(nm.data["botRequests"], float64(botRequests))
	nm.data["demand"] = append(nm.data["demand"], float64(botRequests+busyBots))
	nm.data["idleBots"] = append(nm.data["idleBots"], float64(idleBots))
	nm.data["botsWarmingUp"] = append(nm.data["botsWarmingUp"], float64(botsWarmingUp))
	nm.data["busyBots"] = append(nm.data["busyBots"], float64(busyBots))

	nm.data["slots.available"] = append(nm.data["slots.available"], float64(stats.available))
	nm.data["slots.total"] = append(nm.data["slots.total"], float64(stats.slots))

	if nm.ui != nil {
		perc := float64(stats.taken+nm.botman.Requests()) / float64(stats.slots) * float64(100)
		nm.capGraph.SetValue(perc)

		_, _, width, _ := nm.ui.botSparkline.GetInnerRect()

		if len(nm.data["slots.total"]) > width {
			nm.ui.botSparkline.SetData(nm.data["slots.total"][len(nm.data["slots.total"])-width:])
			nm.ui.nodeSparkline.SetData(nm.data["nodes"][len(nm.data["nodes"])-width:])
		} else {
			nm.ui.botSparkline.SetData(nm.data["slots.total"])
			nm.ui.nodeSparkline.SetData(nm.data["nodes"])
		}

		_, _, width, _ = nm.ui.lineChart.GetInnerRect()
		if len(nm.data["busyBots"]) > width {
			nm.ui.lineChart.SetData([][]float64{
				nm.data["slots.available"][len(nm.data["slots.available"])-width:],
				nm.data["botRequests"][len(nm.data["botRequests"])-width:],
			})
		} else {
			nm.ui.lineChart.SetData([][]float64{
				nm.data["slots.available"],
				nm.data["demand"],
			})
		}

		nm.ui.barChart.SetBarValue("sm", nodes["small"])
		nm.ui.barChart.SetBarValue("med", nodes["medium"])
		nm.ui.barChart.SetBarValue("lrg", nodes["large"])
		nm.ui.barChart.SetBarValue("busy", busyBots)
		nm.ui.barChart.SetBarValue("available", stats.available)
		nm.ui.barChart.SetBarValue("requests", botRequests)
		nm.ui.barChart.SetBarValue("total demand", botRequests+busyBots)

		rec_avg := 10
		if len(nm.data["demand"]) > 30 {
			rec_avg = getAvg(nm.data["demand"][len(nm.data["demand"])-30:])
		}

		nm.ui.barChart.SetBarValue("avg demand", rec_avg)

		nm.ui.app.Draw()
	}
}
