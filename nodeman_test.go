package main

import (
	"botmanager/botman"
	"bytes"
	"io"
	"os/exec"
	"reflect"
	"sync"
	"testing"
	"time"

	"github.com/navidys/tvxwidgets"
)

func TestFinishModal(t *testing.T) {
	type args struct {
		ui    *UIView
		score string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			FinishModal(tt.args.ui, tt.args.score)
		})
	}
}

func TestGraph(t *testing.T) {
	tests := []struct {
		name string
		want *tvxwidgets.UtilModeGauge
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Graph(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Graph() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewLineChart(t *testing.T) {
	tests := []struct {
		name string
		want *tvxwidgets.Plot
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLineChart(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLineChart() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewNodeMan(t *testing.T) {
	type args struct {
		in io.WriteCloser
	}
	tests := []struct {
		name string
		args args
		want *NodeManager
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNodeMan(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNodeMan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNodeManager_Allocate(t *testing.T) {
	type fields struct {
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
		capGraph     *tvxwidgets.UtilModeGauge
		ui           *UIView
		data         map[string][]float64
		started      *time.Time
		Cmd          *exec.Cmd
		WaitGroup    sync.WaitGroup
		RWMutex      sync.RWMutex
	}
	type args struct {
		nodeID int
		proc   int
		bot    *botman.Bot
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nm := &NodeManager{
				nodes:        tt.fields.nodes,
				botman:       tt.fields.botman,
				assgignments: tt.fields.assgignments,
				in:           tt.fields.in,
				out:          tt.fields.out,
				node_manager: tt.fields.node_manager,
				bots_done:    tt.fields.bots_done,
				autoScale:    tt.fields.autoScale,
				interactive:  tt.fields.interactive,
				quietmode:    tt.fields.quietmode,
				stepthrough:  tt.fields.stepthrough,
				work:         tt.fields.work,
				capGraph:     tt.fields.capGraph,
				ui:           tt.fields.ui,
				data:         tt.fields.data,
				started:      tt.fields.started,
				Cmd:          tt.fields.Cmd,
				WaitGroup:    tt.fields.WaitGroup,
				RWMutex:      tt.fields.RWMutex,
			}
			if got := nm.Allocate(tt.args.nodeID, tt.args.proc, tt.args.bot); got != tt.want {
				t.Errorf("Allocate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNodeManager_DoWork(t *testing.T) {
	type fields struct {
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
		capGraph     *tvxwidgets.UtilModeGauge
		ui           *UIView
		data         map[string][]float64
		started      *time.Time
		Cmd          *exec.Cmd
		WaitGroup    sync.WaitGroup
		RWMutex      sync.RWMutex
	}
	type args struct {
		fn func(nm *NodeManager)
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *sync.WaitGroup
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nm := &NodeManager{
				nodes:        tt.fields.nodes,
				botman:       tt.fields.botman,
				assgignments: tt.fields.assgignments,
				in:           tt.fields.in,
				out:          tt.fields.out,
				node_manager: tt.fields.node_manager,
				bots_done:    tt.fields.bots_done,
				autoScale:    tt.fields.autoScale,
				interactive:  tt.fields.interactive,
				quietmode:    tt.fields.quietmode,
				stepthrough:  tt.fields.stepthrough,
				work:         tt.fields.work,
				capGraph:     tt.fields.capGraph,
				ui:           tt.fields.ui,
				data:         tt.fields.data,
				started:      tt.fields.started,
				Cmd:          tt.fields.Cmd,
				WaitGroup:    tt.fields.WaitGroup,
				RWMutex:      tt.fields.RWMutex,
			}
			if got := nm.DoWork(tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DoWork() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNodeManager_Listen(t *testing.T) {
	type fields struct {
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
		capGraph     *tvxwidgets.UtilModeGauge
		ui           *UIView
		data         map[string][]float64
		started      *time.Time
		Cmd          *exec.Cmd
		WaitGroup    sync.WaitGroup
		RWMutex      sync.RWMutex
	}
	type args struct {
		stout io.ReadCloser
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nm := &NodeManager{
				nodes:        tt.fields.nodes,
				botman:       tt.fields.botman,
				assgignments: tt.fields.assgignments,
				in:           tt.fields.in,
				out:          tt.fields.out,
				node_manager: tt.fields.node_manager,
				bots_done:    tt.fields.bots_done,
				autoScale:    tt.fields.autoScale,
				interactive:  tt.fields.interactive,
				quietmode:    tt.fields.quietmode,
				stepthrough:  tt.fields.stepthrough,
				work:         tt.fields.work,
				capGraph:     tt.fields.capGraph,
				ui:           tt.fields.ui,
				data:         tt.fields.data,
				started:      tt.fields.started,
				Cmd:          tt.fields.Cmd,
				WaitGroup:    tt.fields.WaitGroup,
				RWMutex:      tt.fields.RWMutex,
			}
			nm.Listen(tt.args.stout)
		})
	}
}

func TestNodeManager_NextTick(t *testing.T) {
	type fields struct {
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
		capGraph     *tvxwidgets.UtilModeGauge
		ui           *UIView
		data         map[string][]float64
		started      *time.Time
		Cmd          *exec.Cmd
		WaitGroup    sync.WaitGroup
		RWMutex      sync.RWMutex
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nm := &NodeManager{
				nodes:        tt.fields.nodes,
				botman:       tt.fields.botman,
				assgignments: tt.fields.assgignments,
				in:           tt.fields.in,
				out:          tt.fields.out,
				node_manager: tt.fields.node_manager,
				bots_done:    tt.fields.bots_done,
				autoScale:    tt.fields.autoScale,
				interactive:  tt.fields.interactive,
				quietmode:    tt.fields.quietmode,
				stepthrough:  tt.fields.stepthrough,
				work:         tt.fields.work,
				capGraph:     tt.fields.capGraph,
				ui:           tt.fields.ui,
				data:         tt.fields.data,
				started:      tt.fields.started,
				Cmd:          tt.fields.Cmd,
				WaitGroup:    tt.fields.WaitGroup,
				RWMutex:      tt.fields.RWMutex,
			}
			nm.NextTick()
		})
	}
}

func TestNodeManager_Node(t *testing.T) {
	type fields struct {
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
		capGraph     *tvxwidgets.UtilModeGauge
		ui           *UIView
		data         map[string][]float64
		started      *time.Time
		Cmd          *exec.Cmd
		WaitGroup    sync.WaitGroup
		RWMutex      sync.RWMutex
	}
	type args struct {
		nodeID int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Node
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nm := &NodeManager{
				nodes:        tt.fields.nodes,
				botman:       tt.fields.botman,
				assgignments: tt.fields.assgignments,
				in:           tt.fields.in,
				out:          tt.fields.out,
				node_manager: tt.fields.node_manager,
				bots_done:    tt.fields.bots_done,
				autoScale:    tt.fields.autoScale,
				interactive:  tt.fields.interactive,
				quietmode:    tt.fields.quietmode,
				stepthrough:  tt.fields.stepthrough,
				work:         tt.fields.work,
				capGraph:     tt.fields.capGraph,
				ui:           tt.fields.ui,
				data:         tt.fields.data,
				started:      tt.fields.started,
				Cmd:          tt.fields.Cmd,
				WaitGroup:    tt.fields.WaitGroup,
				RWMutex:      tt.fields.RWMutex,
			}
			if got := nm.Node(tt.args.nodeID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Node() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNodeManager_NodeProcs(t *testing.T) {
	type fields struct {
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
		capGraph     *tvxwidgets.UtilModeGauge
		ui           *UIView
		data         map[string][]float64
		started      *time.Time
		Cmd          *exec.Cmd
		WaitGroup    sync.WaitGroup
		RWMutex      sync.RWMutex
	}
	type args struct {
		nodeID int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []*botman.Bot
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nm := &NodeManager{
				nodes:        tt.fields.nodes,
				botman:       tt.fields.botman,
				assgignments: tt.fields.assgignments,
				in:           tt.fields.in,
				out:          tt.fields.out,
				node_manager: tt.fields.node_manager,
				bots_done:    tt.fields.bots_done,
				autoScale:    tt.fields.autoScale,
				interactive:  tt.fields.interactive,
				quietmode:    tt.fields.quietmode,
				stepthrough:  tt.fields.stepthrough,
				work:         tt.fields.work,
				capGraph:     tt.fields.capGraph,
				ui:           tt.fields.ui,
				data:         tt.fields.data,
				started:      tt.fields.started,
				Cmd:          tt.fields.Cmd,
				WaitGroup:    tt.fields.WaitGroup,
				RWMutex:      tt.fields.RWMutex,
			}
			if got := nm.NodeProcs(tt.args.nodeID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NodeProcs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNodeManager_NodeReady(t *testing.T) {
	type fields struct {
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
		capGraph     *tvxwidgets.UtilModeGauge
		ui           *UIView
		data         map[string][]float64
		started      *time.Time
		Cmd          *exec.Cmd
		WaitGroup    sync.WaitGroup
		RWMutex      sync.RWMutex
	}
	type args struct {
		nodeID int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nm := &NodeManager{
				nodes:        tt.fields.nodes,
				botman:       tt.fields.botman,
				assgignments: tt.fields.assgignments,
				in:           tt.fields.in,
				out:          tt.fields.out,
				node_manager: tt.fields.node_manager,
				bots_done:    tt.fields.bots_done,
				autoScale:    tt.fields.autoScale,
				interactive:  tt.fields.interactive,
				quietmode:    tt.fields.quietmode,
				stepthrough:  tt.fields.stepthrough,
				work:         tt.fields.work,
				capGraph:     tt.fields.capGraph,
				ui:           tt.fields.ui,
				data:         tt.fields.data,
				started:      tt.fields.started,
				Cmd:          tt.fields.Cmd,
				WaitGroup:    tt.fields.WaitGroup,
				RWMutex:      tt.fields.RWMutex,
			}
			nm.NodeReady(tt.args.nodeID)
		})
	}
}

func TestNodeManager_NodeStats(t *testing.T) {
	type fields struct {
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
		capGraph     *tvxwidgets.UtilModeGauge
		ui           *UIView
		data         map[string][]float64
		started      *time.Time
		Cmd          *exec.Cmd
		WaitGroup    sync.WaitGroup
		RWMutex      sync.RWMutex
	}
	type args struct {
		nodeID int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   map[string]string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nm := &NodeManager{
				nodes:        tt.fields.nodes,
				botman:       tt.fields.botman,
				assgignments: tt.fields.assgignments,
				in:           tt.fields.in,
				out:          tt.fields.out,
				node_manager: tt.fields.node_manager,
				bots_done:    tt.fields.bots_done,
				autoScale:    tt.fields.autoScale,
				interactive:  tt.fields.interactive,
				quietmode:    tt.fields.quietmode,
				stepthrough:  tt.fields.stepthrough,
				work:         tt.fields.work,
				capGraph:     tt.fields.capGraph,
				ui:           tt.fields.ui,
				data:         tt.fields.data,
				started:      tt.fields.started,
				Cmd:          tt.fields.Cmd,
				WaitGroup:    tt.fields.WaitGroup,
				RWMutex:      tt.fields.RWMutex,
			}
			if got := nm.NodeStats(tt.args.nodeID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NodeStats() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNodeManager_Receive(t *testing.T) {
	type fields struct {
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
		capGraph     *tvxwidgets.UtilModeGauge
		ui           *UIView
		data         map[string][]float64
		started      *time.Time
		Cmd          *exec.Cmd
		WaitGroup    sync.WaitGroup
		RWMutex      sync.RWMutex
	}
	type args struct {
		text string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nm := &NodeManager{
				nodes:        tt.fields.nodes,
				botman:       tt.fields.botman,
				assgignments: tt.fields.assgignments,
				in:           tt.fields.in,
				out:          tt.fields.out,
				node_manager: tt.fields.node_manager,
				bots_done:    tt.fields.bots_done,
				autoScale:    tt.fields.autoScale,
				interactive:  tt.fields.interactive,
				quietmode:    tt.fields.quietmode,
				stepthrough:  tt.fields.stepthrough,
				work:         tt.fields.work,
				capGraph:     tt.fields.capGraph,
				ui:           tt.fields.ui,
				data:         tt.fields.data,
				started:      tt.fields.started,
				Cmd:          tt.fields.Cmd,
				WaitGroup:    tt.fields.WaitGroup,
				RWMutex:      tt.fields.RWMutex,
			}
			nm.Receive(tt.args.text)
		})
	}
}

func TestNodeManager_RemoveBot(t *testing.T) {
	type fields struct {
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
		capGraph     *tvxwidgets.UtilModeGauge
		ui           *UIView
		data         map[string][]float64
		started      *time.Time
		Cmd          *exec.Cmd
		WaitGroup    sync.WaitGroup
		RWMutex      sync.RWMutex
	}
	type args struct {
		bot *botman.Bot
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nm := &NodeManager{
				nodes:        tt.fields.nodes,
				botman:       tt.fields.botman,
				assgignments: tt.fields.assgignments,
				in:           tt.fields.in,
				out:          tt.fields.out,
				node_manager: tt.fields.node_manager,
				bots_done:    tt.fields.bots_done,
				autoScale:    tt.fields.autoScale,
				interactive:  tt.fields.interactive,
				quietmode:    tt.fields.quietmode,
				stepthrough:  tt.fields.stepthrough,
				work:         tt.fields.work,
				capGraph:     tt.fields.capGraph,
				ui:           tt.fields.ui,
				data:         tt.fields.data,
				started:      tt.fields.started,
				Cmd:          tt.fields.Cmd,
				WaitGroup:    tt.fields.WaitGroup,
				RWMutex:      tt.fields.RWMutex,
			}
			nm.RemoveBot(tt.args.bot)
		})
	}
}

func TestNodeManager_Run(t *testing.T) {
	type fields struct {
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
		capGraph     *tvxwidgets.UtilModeGauge
		ui           *UIView
		data         map[string][]float64
		started      *time.Time
		Cmd          *exec.Cmd
		WaitGroup    sync.WaitGroup
		RWMutex      sync.RWMutex
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nm := &NodeManager{
				nodes:        tt.fields.nodes,
				botman:       tt.fields.botman,
				assgignments: tt.fields.assgignments,
				in:           tt.fields.in,
				out:          tt.fields.out,
				node_manager: tt.fields.node_manager,
				bots_done:    tt.fields.bots_done,
				autoScale:    tt.fields.autoScale,
				interactive:  tt.fields.interactive,
				quietmode:    tt.fields.quietmode,
				stepthrough:  tt.fields.stepthrough,
				work:         tt.fields.work,
				capGraph:     tt.fields.capGraph,
				ui:           tt.fields.ui,
				data:         tt.fields.data,
				started:      tt.fields.started,
				Cmd:          tt.fields.Cmd,
				WaitGroup:    tt.fields.WaitGroup,
				RWMutex:      tt.fields.RWMutex,
			}
			nm.Run()
		})
	}
}

func TestNodeManager_Send(t *testing.T) {
	type fields struct {
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
		capGraph     *tvxwidgets.UtilModeGauge
		ui           *UIView
		data         map[string][]float64
		started      *time.Time
		Cmd          *exec.Cmd
		WaitGroup    sync.WaitGroup
		RWMutex      sync.RWMutex
	}
	type args struct {
		msg interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nm := &NodeManager{
				nodes:        tt.fields.nodes,
				botman:       tt.fields.botman,
				assgignments: tt.fields.assgignments,
				in:           tt.fields.in,
				out:          tt.fields.out,
				node_manager: tt.fields.node_manager,
				bots_done:    tt.fields.bots_done,
				autoScale:    tt.fields.autoScale,
				interactive:  tt.fields.interactive,
				quietmode:    tt.fields.quietmode,
				stepthrough:  tt.fields.stepthrough,
				work:         tt.fields.work,
				capGraph:     tt.fields.capGraph,
				ui:           tt.fields.ui,
				data:         tt.fields.data,
				started:      tt.fields.started,
				Cmd:          tt.fields.Cmd,
				WaitGroup:    tt.fields.WaitGroup,
				RWMutex:      tt.fields.RWMutex,
			}
			if err := nm.Send(tt.args.msg); (err != nil) != tt.wantErr {
				t.Errorf("Send() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNodeManager_SetAutoScaling(t *testing.T) {
	type fields struct {
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
		capGraph     *tvxwidgets.UtilModeGauge
		ui           *UIView
		data         map[string][]float64
		started      *time.Time
		Cmd          *exec.Cmd
		WaitGroup    sync.WaitGroup
		RWMutex      sync.RWMutex
	}
	type args struct {
		on bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nm := &NodeManager{
				nodes:        tt.fields.nodes,
				botman:       tt.fields.botman,
				assgignments: tt.fields.assgignments,
				in:           tt.fields.in,
				out:          tt.fields.out,
				node_manager: tt.fields.node_manager,
				bots_done:    tt.fields.bots_done,
				autoScale:    tt.fields.autoScale,
				interactive:  tt.fields.interactive,
				quietmode:    tt.fields.quietmode,
				stepthrough:  tt.fields.stepthrough,
				work:         tt.fields.work,
				capGraph:     tt.fields.capGraph,
				ui:           tt.fields.ui,
				data:         tt.fields.data,
				started:      tt.fields.started,
				Cmd:          tt.fields.Cmd,
				WaitGroup:    tt.fields.WaitGroup,
				RWMutex:      tt.fields.RWMutex,
			}
			nm.SetAutoScaling(tt.args.on)
		})
	}
}

func TestNodeManager_SetGraph(t *testing.T) {
	type fields struct {
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
		capGraph     *tvxwidgets.UtilModeGauge
		ui           *UIView
		data         map[string][]float64
		started      *time.Time
		Cmd          *exec.Cmd
		WaitGroup    sync.WaitGroup
		RWMutex      sync.RWMutex
	}
	type args struct {
		graph *tvxwidgets.UtilModeGauge
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nm := &NodeManager{
				nodes:        tt.fields.nodes,
				botman:       tt.fields.botman,
				assgignments: tt.fields.assgignments,
				in:           tt.fields.in,
				out:          tt.fields.out,
				node_manager: tt.fields.node_manager,
				bots_done:    tt.fields.bots_done,
				autoScale:    tt.fields.autoScale,
				interactive:  tt.fields.interactive,
				quietmode:    tt.fields.quietmode,
				stepthrough:  tt.fields.stepthrough,
				work:         tt.fields.work,
				capGraph:     tt.fields.capGraph,
				ui:           tt.fields.ui,
				data:         tt.fields.data,
				started:      tt.fields.started,
				Cmd:          tt.fields.Cmd,
				WaitGroup:    tt.fields.WaitGroup,
				RWMutex:      tt.fields.RWMutex,
			}
			nm.SetGraph(tt.args.graph)
		})
	}
}

func TestNodeManager_SetInteractive(t *testing.T) {
	type fields struct {
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
		capGraph     *tvxwidgets.UtilModeGauge
		ui           *UIView
		data         map[string][]float64
		started      *time.Time
		Cmd          *exec.Cmd
		WaitGroup    sync.WaitGroup
		RWMutex      sync.RWMutex
	}
	type args struct {
		on bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nm := &NodeManager{
				nodes:        tt.fields.nodes,
				botman:       tt.fields.botman,
				assgignments: tt.fields.assgignments,
				in:           tt.fields.in,
				out:          tt.fields.out,
				node_manager: tt.fields.node_manager,
				bots_done:    tt.fields.bots_done,
				autoScale:    tt.fields.autoScale,
				interactive:  tt.fields.interactive,
				quietmode:    tt.fields.quietmode,
				stepthrough:  tt.fields.stepthrough,
				work:         tt.fields.work,
				capGraph:     tt.fields.capGraph,
				ui:           tt.fields.ui,
				data:         tt.fields.data,
				started:      tt.fields.started,
				Cmd:          tt.fields.Cmd,
				WaitGroup:    tt.fields.WaitGroup,
				RWMutex:      tt.fields.RWMutex,
			}
			nm.SetInteractive(tt.args.on)
		})
	}
}

func TestNodeManager_SetQuietMode(t *testing.T) {
	type fields struct {
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
		capGraph     *tvxwidgets.UtilModeGauge
		ui           *UIView
		data         map[string][]float64
		started      *time.Time
		Cmd          *exec.Cmd
		WaitGroup    sync.WaitGroup
		RWMutex      sync.RWMutex
	}
	type args struct {
		on bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nm := &NodeManager{
				nodes:        tt.fields.nodes,
				botman:       tt.fields.botman,
				assgignments: tt.fields.assgignments,
				in:           tt.fields.in,
				out:          tt.fields.out,
				node_manager: tt.fields.node_manager,
				bots_done:    tt.fields.bots_done,
				autoScale:    tt.fields.autoScale,
				interactive:  tt.fields.interactive,
				quietmode:    tt.fields.quietmode,
				stepthrough:  tt.fields.stepthrough,
				work:         tt.fields.work,
				capGraph:     tt.fields.capGraph,
				ui:           tt.fields.ui,
				data:         tt.fields.data,
				started:      tt.fields.started,
				Cmd:          tt.fields.Cmd,
				WaitGroup:    tt.fields.WaitGroup,
				RWMutex:      tt.fields.RWMutex,
			}
			nm.SetQuietMode(tt.args.on)
		})
	}
}

func TestNodeManager_SetStepThrough(t *testing.T) {
	type fields struct {
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
		capGraph     *tvxwidgets.UtilModeGauge
		ui           *UIView
		data         map[string][]float64
		started      *time.Time
		Cmd          *exec.Cmd
		WaitGroup    sync.WaitGroup
		RWMutex      sync.RWMutex
	}
	type args struct {
		on bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nm := &NodeManager{
				nodes:        tt.fields.nodes,
				botman:       tt.fields.botman,
				assgignments: tt.fields.assgignments,
				in:           tt.fields.in,
				out:          tt.fields.out,
				node_manager: tt.fields.node_manager,
				bots_done:    tt.fields.bots_done,
				autoScale:    tt.fields.autoScale,
				interactive:  tt.fields.interactive,
				quietmode:    tt.fields.quietmode,
				stepthrough:  tt.fields.stepthrough,
				work:         tt.fields.work,
				capGraph:     tt.fields.capGraph,
				ui:           tt.fields.ui,
				data:         tt.fields.data,
				started:      tt.fields.started,
				Cmd:          tt.fields.Cmd,
				WaitGroup:    tt.fields.WaitGroup,
				RWMutex:      tt.fields.RWMutex,
			}
			nm.SetStepThrough(tt.args.on)
		})
	}
}

func TestNodeManager_SetStout(t *testing.T) {
	type fields struct {
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
		capGraph     *tvxwidgets.UtilModeGauge
		ui           *UIView
		data         map[string][]float64
		started      *time.Time
		Cmd          *exec.Cmd
		WaitGroup    sync.WaitGroup
		RWMutex      sync.RWMutex
	}
	tests := []struct {
		name    string
		fields  fields
		wantOut string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nm := &NodeManager{
				nodes:        tt.fields.nodes,
				botman:       tt.fields.botman,
				assgignments: tt.fields.assgignments,
				in:           tt.fields.in,
				out:          tt.fields.out,
				node_manager: tt.fields.node_manager,
				bots_done:    tt.fields.bots_done,
				autoScale:    tt.fields.autoScale,
				interactive:  tt.fields.interactive,
				quietmode:    tt.fields.quietmode,
				stepthrough:  tt.fields.stepthrough,
				work:         tt.fields.work,
				capGraph:     tt.fields.capGraph,
				ui:           tt.fields.ui,
				data:         tt.fields.data,
				started:      tt.fields.started,
				Cmd:          tt.fields.Cmd,
				WaitGroup:    tt.fields.WaitGroup,
				RWMutex:      tt.fields.RWMutex,
			}
			out := &bytes.Buffer{}
			nm.SetStout(out)
			if gotOut := out.String(); gotOut != tt.wantOut {
				t.Errorf("SetStout() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}

func TestNodeManager_SetUI(t *testing.T) {
	type fields struct {
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
		capGraph     *tvxwidgets.UtilModeGauge
		ui           *UIView
		data         map[string][]float64
		started      *time.Time
		Cmd          *exec.Cmd
		WaitGroup    sync.WaitGroup
		RWMutex      sync.RWMutex
	}
	type args struct {
		ui *UIView
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nm := &NodeManager{
				nodes:        tt.fields.nodes,
				botman:       tt.fields.botman,
				assgignments: tt.fields.assgignments,
				in:           tt.fields.in,
				out:          tt.fields.out,
				node_manager: tt.fields.node_manager,
				bots_done:    tt.fields.bots_done,
				autoScale:    tt.fields.autoScale,
				interactive:  tt.fields.interactive,
				quietmode:    tt.fields.quietmode,
				stepthrough:  tt.fields.stepthrough,
				work:         tt.fields.work,
				capGraph:     tt.fields.capGraph,
				ui:           tt.fields.ui,
				data:         tt.fields.data,
				started:      tt.fields.started,
				Cmd:          tt.fields.Cmd,
				WaitGroup:    tt.fields.WaitGroup,
				RWMutex:      tt.fields.RWMutex,
			}
			nm.SetUI(tt.args.ui)
		})
	}
}

func TestNodeManager_Stats(t *testing.T) {
	type fields struct {
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
		capGraph     *tvxwidgets.UtilModeGauge
		ui           *UIView
		data         map[string][]float64
		started      *time.Time
		Cmd          *exec.Cmd
		WaitGroup    sync.WaitGroup
		RWMutex      sync.RWMutex
	}
	tests := []struct {
		name   string
		fields fields
		want   *ClusterStats
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nm := &NodeManager{
				nodes:        tt.fields.nodes,
				botman:       tt.fields.botman,
				assgignments: tt.fields.assgignments,
				in:           tt.fields.in,
				out:          tt.fields.out,
				node_manager: tt.fields.node_manager,
				bots_done:    tt.fields.bots_done,
				autoScale:    tt.fields.autoScale,
				interactive:  tt.fields.interactive,
				quietmode:    tt.fields.quietmode,
				stepthrough:  tt.fields.stepthrough,
				work:         tt.fields.work,
				capGraph:     tt.fields.capGraph,
				ui:           tt.fields.ui,
				data:         tt.fields.data,
				started:      tt.fields.started,
				Cmd:          tt.fields.Cmd,
				WaitGroup:    tt.fields.WaitGroup,
				RWMutex:      tt.fields.RWMutex,
			}
			if got := nm.Stats(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Stats() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNodeManager_TicKReport(t *testing.T) {
	type fields struct {
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
		capGraph     *tvxwidgets.UtilModeGauge
		ui           *UIView
		data         map[string][]float64
		started      *time.Time
		Cmd          *exec.Cmd
		WaitGroup    sync.WaitGroup
		RWMutex      sync.RWMutex
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nm := &NodeManager{
				nodes:        tt.fields.nodes,
				botman:       tt.fields.botman,
				assgignments: tt.fields.assgignments,
				in:           tt.fields.in,
				out:          tt.fields.out,
				node_manager: tt.fields.node_manager,
				bots_done:    tt.fields.bots_done,
				autoScale:    tt.fields.autoScale,
				interactive:  tt.fields.interactive,
				quietmode:    tt.fields.quietmode,
				stepthrough:  tt.fields.stepthrough,
				work:         tt.fields.work,
				capGraph:     tt.fields.capGraph,
				ui:           tt.fields.ui,
				data:         tt.fields.data,
				started:      tt.fields.started,
				Cmd:          tt.fields.Cmd,
				WaitGroup:    tt.fields.WaitGroup,
				RWMutex:      tt.fields.RWMutex,
			}
			nm.TicKReport()
		})
	}
}

func TestNodeManager_log(t *testing.T) {
	type fields struct {
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
		capGraph     *tvxwidgets.UtilModeGauge
		ui           *UIView
		data         map[string][]float64
		started      *time.Time
		Cmd          *exec.Cmd
		WaitGroup    sync.WaitGroup
		RWMutex      sync.RWMutex
	}
	type args struct {
		text string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nm := &NodeManager{
				nodes:        tt.fields.nodes,
				botman:       tt.fields.botman,
				assgignments: tt.fields.assgignments,
				in:           tt.fields.in,
				out:          tt.fields.out,
				node_manager: tt.fields.node_manager,
				bots_done:    tt.fields.bots_done,
				autoScale:    tt.fields.autoScale,
				interactive:  tt.fields.interactive,
				quietmode:    tt.fields.quietmode,
				stepthrough:  tt.fields.stepthrough,
				work:         tt.fields.work,
				capGraph:     tt.fields.capGraph,
				ui:           tt.fields.ui,
				data:         tt.fields.data,
				started:      tt.fields.started,
				Cmd:          tt.fields.Cmd,
				WaitGroup:    tt.fields.WaitGroup,
				RWMutex:      tt.fields.RWMutex,
			}
			nm.log(tt.args.text)
		})
	}
}

func TestNode_Busy(t *testing.T) {
	type fields struct {
		id         int
		size       string
		status     NodeStatus
		framesIdle int
		cpu        int
		cost       int
		proc       []*botman.Bot
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := &Node{
				id:         tt.fields.id,
				size:       tt.fields.size,
				status:     tt.fields.status,
				framesIdle: tt.fields.framesIdle,
				cpu:        tt.fields.cpu,
				cost:       tt.fields.cost,
				proc:       tt.fields.proc,
			}
			if got := node.Busy(); got != tt.want {
				t.Errorf("Busy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_Empty(t *testing.T) {
	type fields struct {
		id         int
		size       string
		status     NodeStatus
		framesIdle int
		cpu        int
		cost       int
		proc       []*botman.Bot
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := &Node{
				id:         tt.fields.id,
				size:       tt.fields.size,
				status:     tt.fields.status,
				framesIdle: tt.fields.framesIdle,
				cpu:        tt.fields.cpu,
				cost:       tt.fields.cost,
				proc:       tt.fields.proc,
			}
			if got := node.Empty(); got != tt.want {
				t.Errorf("Empty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_Idle(t *testing.T) {
	type fields struct {
		id         int
		size       string
		status     NodeStatus
		framesIdle int
		cpu        int
		cost       int
		proc       []*botman.Bot
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := &Node{
				id:         tt.fields.id,
				size:       tt.fields.size,
				status:     tt.fields.status,
				framesIdle: tt.fields.framesIdle,
				cpu:        tt.fields.cpu,
				cost:       tt.fields.cost,
				proc:       tt.fields.proc,
			}
			if got := node.Idle(); got != tt.want {
				t.Errorf("Idle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUI(t *testing.T) {
	type args struct {
		nm *NodeManager
	}
	tests := []struct {
		name string
		args args
		want *UIView
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UI(tt.args.nm); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UI() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getAvg(t *testing.T) {
	type args struct {
		arr []float64
	}
	tests := []struct {
		name    string
		args    args
		wantAvg int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAvg := getAvg(tt.args.arr); gotAvg != tt.wantAvg {
				t.Errorf("getAvg() = %v, want %v", gotAvg, tt.wantAvg)
			}
		})
	}
}

func Test_getMode(t *testing.T) {
	type args struct {
		arr []float64
	}
	tests := []struct {
		name     string
		args     args
		wantMode int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMode := getMode(tt.args.arr); gotMode != tt.wantMode {
				t.Errorf("getMode() = %v, want %v", gotMode, tt.wantMode)
			}
		})
	}
}

func Test_isFlagPassed(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isFlagPassed(tt.args.name); got != tt.want {
				t.Errorf("isFlagPassed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newBarChart(t *testing.T) {
	tests := []struct {
		name string
		want *tvxwidgets.BarChart
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newBarChart(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newBarChart() = %v, want %v", got, tt.want)
			}
		})
	}
}
