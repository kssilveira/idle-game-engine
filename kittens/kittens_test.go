package kittens

import (
	"bytes"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

type iter struct {
	input   string
	elapsed int64
}

func TestRun(t *testing.T) {
	gather := "0"
	refine := "1"
	srefine := "s1"
	field := "2"
	sfield := "s2"
	inputs := []struct {
		name      string
		iters     []iter
		resources map[string]float64
	}{{
		name: "gather",
		iters: []iter{
			// gather 2 catnip
			{gather, 0}, {gather, 0},
			// end
			{"999", 0},
		},
	}, {
		name: "buy one",
		resources: map[string]float64{
			"catnip": 9,
		},
		iters: []iter{
			// buy catnip field, catnip not enough
			{field, 0},
			// gather 10th catnip
			{gather, 0},
			// buy catnip field
			{field, 0},
			// wait 1 second and 10 seconds
			{"", 1}, {"", 10},
			// end
			{"999", 0},
		},
	}, {
		name: "buy two",
		resources: map[string]float64{
			"catnip": 100,
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
			// end
			{"999", 0},
		},
	}, {
		name: "skip",
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
			// end
			{"999", 0},
		},
	}, {
		name: "skip until max",
		resources: map[string]float64{
			"catnip": 10,
		},
		iters: append(
			append(
				// field
				repeat([]iter{{field, 1}, {sfield, 1}}, 56),
				// wood
				repeat([]iter{{srefine, 1}, {refine, 1}}, 142)...),
			// end
			[]iter{{sfield, 1}, {"999", 0}}...),
	}, {
		name: "refine one",
		resources: map[string]float64{
			"catnip": 99,
		},
		iters: []iter{
			// refine catnip, catnip not enough
			{refine, 0},
			// gather 100th catnip
			{gather, 0},
			// refine catnip
			{refine, 0},
			// end
			{"999", 0},
		},
	}, {
		name: "refine two",
		resources: map[string]float64{
			"catnip": 200,
		},
		iters: []iter{
			// refine catnip
			{refine, 0},
			// refine catnip,
			{refine, 0},
			// end
			{"999", 0},
		},
	}}
	for _, in := range inputs {
		var buf bytes.Buffer
		logger := log.New(&buf, "", 0 /* flags */)
		input := make(chan string)
		go func() {
			for _, one := range in.iters {
				input <- one.input
			}
		}()
		timeIndex := 0
		now := time.Unix(0, 0)
		nowfn := func() time.Time {
			res := now
			now = now.Add(time.Duration(in.iters[timeIndex].elapsed) * time.Second)
			timeIndex++
			return res
		}
		g := NewGame(nowfn)
		for name, quantity := range in.resources {
			g.GetResource(name).Quantity = quantity
		}
		if err := g.Validate(); err != nil {
			t.Errorf("[%s] got err %v", in.name, err)
		}
		g.Run(logger, "###", input, nowfn)
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
