package main

import (
	"botmanager/botman"
	"flag"
	"fmt"
	"github.com/rivo/tview"
	"log"
	"os"
	"os/exec"
	"time"
)

var DEBUG bool = false

var NodeTemplates = map[string]Node{"small": {
	size: "small",
	cpu:  1,
	cost: 4,
}, "medium": {
	size: "medium",
	cpu:  2,
	cost: 6,
}, "large": {
	size: "large",
	cpu:  4,
	cost: 12,
}}

var Command = struct {
	NODE_CREATE   string
	NODE_DESTROY  string
	BOT_CREATE    string
	BOT_ASSIGN    string
	BOTS_ACK_DONE string
}{
	NODE_CREATE:   "cloud_provider.create_node",
	NODE_DESTROY:  "cloud_provider.destroy_node",
	BOT_CREATE:    "bot_provider.create_bot",
	BOT_ASSIGN:    "bot_requester.fulfill_bot_request",
	BOTS_ACK_DONE: "bot_requester.ack_done",
}

var Request = struct {
	NODE_READY  string
	BOT_READY   string
	BOT_DONE    string
	BOT_REQUEST string
	BOTS_DONE   string
}{
	NODE_READY:  "cloud_provider.node_ready",
	BOT_READY:   "bot_provider.bot_ready",
	BOT_DONE:    "bot_provider.bot_completed",
	BOT_REQUEST: "bot_requester.bot_request",
	BOTS_DONE:   "bot_requester.done",
}

type NodeStatus string

const (
	PENDING    NodeStatus = "PENDING"    // not yet sent to simulation
	NOT_READY  NodeStatus = "NOT_READY"  // sent but not ready
	READY      NodeStatus = "READY"      // ready status received fromm simulation
	DESTROYING NodeStatus = "DESTROYING" // destroying bot
	DESTROYED  NodeStatus = "DESTROYED"  // bot destroyed
)

type Node struct {
	id         int
	size       string
	status     NodeStatus
	framesIdle int
	cpu        int
	cost       int
	proc       []*botman.Bot
}

type MessageData[T any] struct {
	BotRequestID *int    `json:"bot_request_id,omitempty"`
	BotID        *T      `json:"bot_id,omitempty"`
	NodeID       *T      `json:"node_id,omitempty"`
	NodeSize     *string `json:"node_size,omitempty"`
}

//Message covers all messages sent from the simulation
type ReceiveMsg struct {
	Kind string               `json:"kind"`
	Data *MessageData[string] `json:"data,omitempty"`
}

type SendMsg struct {
	Kind string            `json:"kind"`
	Data *MessageData[int] `json:"data,omitempty"`
}

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

func main() {
	flag.Bool("no-autoscale", false, "disable autoscaling")
	flag.Bool("quietmode", false, "don't log io")
	flag.Bool("debug", false, "enable debug logging")
	flag.Bool("interactive", false, "start with the terminal UI")
	flag.Bool("stepthrough", false, "enable enter to step through io")
	flag.Parse()

	filePath := os.Args[len(os.Args)-1]

	var suff string = "autoscaling"

	autoScale := true
	if isFlagPassed("no-autoscale") {
		suff = "autoscale disabled"
		autoScale = false
	}

	stepthrough := false
	if isFlagPassed("stepthrough") {
		stepthrough = true
	}

	if isFlagPassed("debug") {
		DEBUG = true
		botman.DEBUG = true
	}

	quietmode := false
	if isFlagPassed("quietmode") {
		quietmode = true
	}

	interactive := false
	if isFlagPassed("interactive") {
		interactive = true
	}

	cmd := exec.Command("python", "-u", "../simulation.py", filePath)
	in, _ := cmd.StdinPipe()
	cmd.Stderr = os.Stderr

	stout, _ := cmd.StdoutPipe()

	nm := NewNodeMan(in)
	nm.Listen(stout)

	nm.SetStepThrough(stepthrough)
	nm.SetAutoScaling(autoScale)
	nm.SetQuietMode(quietmode)

	if interactive {
		nm.SetInteractive(true)

		ui := UI(nm)
		nm.SetUI(ui)

		w := tview.ANSIWriter(ui.textView)
		log.SetOutput(w)

		nm.SetStout(ui.textView)
		nm.SetGraph(ui.capGraph)
	}
	nm.log(fmt.Sprintf("Spinning up node manager with %s\n", suff))

	if err := cmd.Start(); err != nil {
		log.Panic(err)
	}

	nm.node_manager <- "large"
	//nm.node_manager <- "large"

	if nm.stepthrough {
		go func() {
			for {
				fmt.Scanln()
				nm.NextTick()
			}
		}()
	}

	if !nm.interactive && !nm.stepthrough {
		time.Sleep(time.Second * 1)
		nm.NextTick()
	}

	if nm.interactive {
		err := nm.ui.app.Run()
		if err != nil {
			panic(err)
		}
		if !nm.interactive && !nm.stepthrough {
			log.SetOutput(os.Stdout)
			nm.ui = nil
			nm.Add(1)
			nm.NextTick()
			_ = cmd.Wait()
			nm.Wait()
		}
	} else if !nm.interactive {
		nm.Add(1)
		_ = cmd.Wait()
		nm.Wait()
	}
}
