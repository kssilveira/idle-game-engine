package kittens

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

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
		resources map[string]float64
	}{{
		name: "gather catnip",
		iters: []iter{
			// gather 2 catnip
			{gather, 0}, {gather, 0},
		},
	}, {
		name: "gather catnip html",
		iters: []iter{
			// gather 2 catnip
			{gather, 0}, {gather, 0},
		},
		isHTML: true,
	}, {
		name: "catnip field 1",
		resources: map[string]float64{
			"catnip": 9,
		},
		iters: []iter{
			// buy catnip field, not enough catnip
			{field, 0},
			// gather 10th catnip
			{gather, 0},
			// buy catnip field
			{field, 0},
			// wait 1 second and 10 seconds
			{"", 1}, {"", 10},
		},
	}, {
		name: "catnip field 2",
		resources: map[string]float64{
			"catnip": 200,
		},
		iters: []iter{
			// buy 1st catnip field
			{field, 0},
			// wait 1 second and 10 seconds
			{"", 1}, {"", 10},
			// buy 2nd catnip field
			{field, 0},
			// wait 1 second and 10 seconds
			{"", 1}, {"", 10},
		},
	}, {
		name: "catnip field skip",
		resources: map[string]float64{
			"catnip": 10,
		},
		iters: []iter{
			// buy catnip field
			{field, 0},
			// wait 1 second
			{"", 1},
			// skip to buy catnip field and buy it
			{sfield, 0}, {field, 0},
			// wait 1 second and 10 seconds
			{"", 1}, {"", 10},
		},
	}, {
		name: "refine catnip",
		resources: map[string]float64{
			"catnip": 200,
		},
		iters: []iter{
			// refine catnip
			{refine, 0},
			// refine catnip,
			{refine, 0},
		},
	}, {
		name: "hut",
		resources: map[string]float64{
			"wood": 100,
		},
		iters: []iter{
			// buy 1st hut
			{hut, 0},
			// wait 1 second and 10 seconds
			{"", 1}, {"", 10},
			// buy 2nd hut
			{hut, 0},
			// wait 1 second and 10 seconds
			{"", 1}, {"", 10},
		},
	}, {
		name: "library",
		resources: map[string]float64{
			"catnip": 1000,
			"wood":   100,
			"kitten": 2,
		},
		iters: []iter{
			// buy 1st library, assign 1st scholar
			{library, 0}, {scholar, 0},
			// wait 1 second and 10 seconds
			{"", 1}, {"", 10},
			// buy 2nd library
			{library, 0},
			// wait 1 second and 10 seconds
			{"", 1}, {"", 10},
			// assign 2nd scholar
			{scholar, 0},
			// wait 1 second and 10 seconds
			{"", 1}, {"", 10},
		},
	}, {
		name: "woodcutter",
		resources: map[string]float64{
			"catnip": 1000,
			"kitten": 2,
		},
		iters: []iter{
			// 1st woodcutter
			{woodcutter, 0},
			// wait 1 second and 10 seconds
			{"", 1}, {"", 10},
			// 2nd woodcutter
			{woodcutter, 0},
			// wait 1 second and 10 seconds
			{"", 1}, {"", 10},
		},
	}, {
		name: "gone",
		resources: map[string]float64{
			"catnip": 1000,
			"kitten": 3,
		},
		iters: []iter{
			// woodcutter
			{woodcutter, 0},
			// wait 1 second and 10 seconds
			{"", 1}, {"", 10},
			// scholar
			{scholar, 0},
			// wait 1 second and 10 seconds
			{"", 1}, {"", 10},
			// wait 1000 seconds
			{"", 1000},
		},
	}, {
		name: "solve",
	}}
	for _, in := range inputs {
		var buf bytes.Buffer
		logger := log.New(&buf, "", 0 /* flags */)
		input := make(chan string)
		go func() {
			if len(in.iters) == 0 {
				Solve(input, 0 /* sleepMS */)
			}
			for _, one := range in.iters {
				input <- one.input
			}
			input <- "999"
		}()
		timeIndex := 0
		now := time.Unix(0, 0)
		nowfn := func() time.Time {
			res := now
			if len(in.iters) == 0 {
				now = now.Add(time.Second)
			} else {
				now = now.Add(time.Duration(append(in.iters, iter{"", 0})[timeIndex].elapsed) * time.Second)
				timeIndex++
			}
			return res
		}
		g := NewGame(nowfn)
		for name, quantity := range in.resources {
			if !g.HasResource(name) {
				t.Errorf("[%s] missing resource %s", in.name, name)
			}
			r := g.GetResource(name)
			r.Quantity = quantity
			if r.Capacity != -1 && r.Capacity < quantity {
				r.Capacity = quantity
			}
		}
		if err := g.Validate(); err != nil {
			t.Errorf("[%s] got err %v", in.name, err)
		}
		output := make(chan *ui.Data)
		go g.Run(nowfn, input, output)
		for data := range output {
			textui.Show(logger, "###", data, in.isHTML)
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

func join(iters ...[]iter) []iter {
	res := []iter{}
	for _, iter := range iters {
		res = append(res, iter...)
	}
	return res
}

func all() []iter {
	res := []iter{}
	n := len(NewGame(func() time.Time { return time.Unix(0, 0) }).Actions)
	for i := 0; i < n; i++ {
		res = append(res, []iter{
			{fmt.Sprintf("s%d", i), 1},
			{fmt.Sprintf("%d", i), 1},
		}...)
	}
	return res
}
