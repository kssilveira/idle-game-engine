package kittens

import (
	"bytes"
	"encoding/json"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/kssilveira/idle-game-engine/game"
	"github.com/kssilveira/idle-game-engine/kittens/graph"
	"github.com/kssilveira/idle-game-engine/kittens/solve"
	"github.com/kssilveira/idle-game-engine/textui"
	"github.com/kssilveira/idle-game-engine/ui"
)

type iter struct {
	input   string
	elapsed int64
}

func TestRun(t *testing.T) {
	inputs := []struct {
		name      string
		iters     []iter
		isHTML    bool
		resources map[string]int
	}{{
		name: "gather catnip",
		iters: []iter{
			{"Gather catnip", 0}, {"Gather catnip", 0},
		},
	}, {
		name: "gather catnip html",
		iters: []iter{
			{"Gather catnip", 0}, {"Gather catnip", 0},
		},
		isHTML: true,
	}, {
		name: "catnip field 1",
		resources: map[string]int{
			"catnip": 9,
		},
		iters: []iter{
			{"Catnip Field", 0}, {"Gather catnip", 0}, {"Catnip Field", 0},
			{"Gather catnip", 1}, {"Gather catnip", 10},
			{"Gather catnip", 200}, {"Gather catnip", 200}, {"Gather catnip", 200}, {"Gather catnip", 200},
		},
	}, {
		name: "catnip field 2",
		resources: map[string]int{
			"catnip": 200,
		},
		iters: []iter{
			{"Catnip Field", 0}, {"Gather catnip", 1},
			{"Catnip Field", 0}, {"Gather catnip", 1},
		},
	}, {
		name: "catnip field skip",
		resources: map[string]int{
			"catnip": 10,
		},
		iters: []iter{
			{"Catnip Field", 0}, {"Gather catnip", 1},
			{"s Catnip Field", 0},
			{"Gather catnip", 1},
		},
	}, {
		name: "refine catnip",
		resources: map[string]int{
			"catnip": 200,
		},
		iters: []iter{
			{"Refine catnip", 0}, {"Refine catnip", 0},
		},
	}, {
		name: "hut",
		resources: map[string]int{
			"catnip": 1000,
			"wood":   100,
		},
		iters: []iter{
			{"Hut", 0}, {"Gather catnip", 1},
			{"Hut", 0}, {"Gather catnip", 1},
			{"Gather catnip", 100},
		},
	}, {
		name: "library",
		resources: map[string]int{
			"catnip": 1000,
			"wood":   100,
			"kitten": 2,
		},
		iters: []iter{
			{"Library", 0}, {"scholar", 0}, {"Gather catnip", 1},
			{"Library", 0}, {"Gather catnip", 1},
			{"scholar", 0}, {"Gather catnip", 1},
		},
	}, {
		name: "woodcutter",
		resources: map[string]int{
			"catnip": 1000,
			"kitten": 2,
			"Hut":    1,
		},
		iters: []iter{
			{"woodcutter", 0}, {"Gather catnip", 1},
			{"woodcutter", 0}, {"Gather catnip", 1},
		},
	}, {
		name: "farmer",
		resources: map[string]int{
			"catnip":      1000,
			"kitten":      2,
			"Agriculture": 1,
		},
		iters: []iter{
			{"farmer", 0}, {"Gather catnip", 1},
			{"farmer", 0}, {"Gather catnip", 1},
		},
	}, {
		name: "gone",
		resources: map[string]int{
			"catnip":      1000,
			"kitten":      4,
			"Hut":         1,
			"Library":     1,
			"Agriculture": 1,
		},
		iters: []iter{
			{"woodcutter", 0}, {"Refine catnip", 1},
			{"scholar", 0}, {"Refine catnip", 1},
			{"farmer", 0}, {"Refine catnip", 1},
			{"Refine catnip", 79},
			{"Refine catnip", 1}, {"Refine catnip", 1}, {"Refine catnip", 1}, {"Refine catnip", 1},
		},
	}, {
		name: "barn",
		resources: map[string]int{
			"catnip":      1,
			"wood":        200,
			"Agriculture": 1,
		},
		iters: []iter{
			{"Barn", 0}, {"Gather catnip", 1},
			{"Barn", 0}, {"Gather catnip", 1},
		},
	}, {
		name: "solve",
	}}
	{
		g := NewGame(func() time.Time { return time.Unix(0, 0) })
		if err := g.Validate(); err != nil {
			t.Errorf("Validate got err %v", err)
		}
	}
	for _, in := range inputs {
		var buf bytes.Buffer
		input := make(chan string)
		timeIndex := 0
		now := time.Unix(0, 0)
		nowfn := func() time.Time {
			res := now
			if len(in.iters) == 0 {
				now = now.Add(time.Second)
			} else {
				now = now.Add(time.Duration(append(in.iters, iter{"Gather catnip", 0}, iter{"Gather catnip", 0})[timeIndex].elapsed) * time.Second)
				timeIndex++
			}
			return res
		}
		g := NewGame(nowfn)
		go func() {
			if len(in.iters) == 0 {
				if err := solve.Solve(g, input, 0 /* sleepMS */); err != nil {
					t.Errorf("[%s] Solve err %v", in.name, err)
				}
			}
			for _, one := range in.iters {
				if err := solve.ToInput(g, one.input, input); err != nil {
					t.Errorf("[%s] got input err %v", in.name, err)
				}
			}
			input <- "999"
		}()
		for name, quantity := range in.resources {
			if !g.HasResource(name) {
				t.Errorf("[%s] missing resource %s", in.name, name)
			}
			r := g.GetResource(name)
			r.Quantity = float64(quantity)
			if r.Capacity != -1 && r.Capacity < r.Quantity {
				r.Capacity = r.Quantity
			}
		}
		output := make(chan *ui.Data)
		go g.Run(nowfn, input, output)
		cfg := textui.Config{
			Logger:            log.New(&buf, "", 0 /* flags */),
			Separator:         "###",
			IsHTML:            in.isHTML,
			HideActionNumbers: true,
			HideCustomActions: true,
		}
		for data := range output {
			textui.Show(cfg, data)
		}
		name := filepath.Join("testdata", strings.Replace(in.name, " ", "_", -1)+".out")
		if err := os.WriteFile(name, buf.Bytes(), 0644); err != nil {
			t.Errorf("[%s] got err %v", in.name, err)
		}
	}
}

func repeat(iters []iter, count int) []iter {
	res := []iter{}
	for i := 0; i < count; i++ {
		res = append(res, iters...)
	}
	return res
}

func TestGraph(t *testing.T) {
	inputs := []struct {
		name   string
		fn     func(*log.Logger, *game.Game, map[string]bool)
		colors map[string]bool
	}{{
		name: "graph",
		fn:   graph.Graph,
	}, {
		name: "graph edges",
		fn:   graph.GraphEdges,
	}, {
		name: "graph nodes",
		fn:   graph.GraphNodes,
	}, {
		name: "graph_blue",
		fn:   graph.Graph,
		colors: map[string]bool{
			"blue": true,
		},
	}}
	for _, in := range inputs {
		var buf bytes.Buffer
		logger := log.New(&buf, "", 0 /* flags */)
		g := NewGame(func() time.Time { return time.Unix(0, 0) })
		in.fn(logger, g, in.colors)
		name := strings.Replace(in.name, " ", "_", -1)
		dot := filepath.Join("testdata", name+".dot")
		if err := os.WriteFile(dot, buf.Bytes(), 0644); err != nil {
			t.Errorf("TestGraph.Graph got err %v", err)
		}
		if false {
			svg := filepath.Join("testdata", name+".svg")
			cmd := exec.Command("dot", "-Tsvg", "-o", svg, dot)
			if err := cmd.Run(); err != nil {
				t.Errorf("[%s] got err %v", in.name, err)
			}
		}
	}
}

func TestNew(t *testing.T) {
	g := NewGame(func() time.Time { return time.Unix(0, 0) })
	content, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		t.Errorf("TestNew got err %v", err)
	}
	if err := os.WriteFile(filepath.Join("testdata", "game.out"), content, 0644); err != nil {
		t.Errorf("TestNew got err %v", err)
	}
}
