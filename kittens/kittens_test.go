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
)

type iter struct {
	input   string
	elapsed int64
}

func TestRun(t *testing.T) {
	gather := "0"
	refine := "1"
	field := "2"
	sfield := "s2"
	hut := "3"
	shut := "s3"
	woodcutter := "4"
	swoodcutter := "s4"
	inputs := []struct {
		name      string
		iters     []iter
		resources map[string]float64
	}{{
		name: "gather catnip",
		iters: []iter{
			// gather 2 catnip
			{gather, 0}, {gather, 0},
			// end
			{"999", 0},
		},
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
			// end
			{"999", 0},
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
			// end
			{"999", 0},
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
			// end
			{"999", 0},
		},
	}, {
		name: "refine catnip 1",
		resources: map[string]float64{
			"catnip": 99,
		},
		iters: []iter{
			// refine catnip, not enough catnip
			{refine, 0},
			// gather 100th catnip
			{gather, 0},
			// refine catnip
			{refine, 0},
			// end
			{"999", 0},
		},
	}, {
		name: "refine catnip 2",
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
	}, {
		name: "hut 1",
		resources: map[string]float64{
			"catnip": 100,
			"wood":   4,
		},
		iters: []iter{
			// buy hut, not enough wood
			{hut, 0},
			// refine 5th wood
			{refine, 0},
			// buy hut
			{hut, 0},
			// wait 1 second and 10 seconds
			{"", 1}, {"", 10},
			// end
			{"999", 0},
		},
	}, {
		name: "hut 2",
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
			// end
			{"999", 0},
		},
	}, {
		name: "woodcutter 2",
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
			// end
			{"999", 0},
		},
	}, {
		name: "all",
		iters: append(
			repeat(all(), 100),
			// end
			iter{"999", 0}),
	}, {
		name: "order",
		iters: join(
			// gather 10 catnip
			repeat([]iter{{gather, 1}}, 10),
			// buy 55 catnip field
			repeat([]iter{{field, 1}, {sfield, 1}}, 55),
			// refine 5 wood
			repeat([]iter{{refine, 1}}, 5),
			// buy 1 hut, assign 2 woodcutter
			repeat([]iter{
				{hut, 1},
				{swoodcutter, 1}, {woodcutter, 1},
				{swoodcutter, 1}, {woodcutter, 1},
				{shut, 1}}, 5),
			// end
			[]iter{{"999", 0}}),
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
