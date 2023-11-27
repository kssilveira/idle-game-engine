package kittens

import (
	"bytes"
	"log"
	"os"
	"testing"
	"time"
)

type iter struct {
	input   int
	elapsed int64
}

func TestRun(t *testing.T) {
	inputs := []struct {
		name  string
		iters []iter
	}{{
		name: "0",
		iters: []iter{
			// gather 9 catnip
			{0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0},
			// buy catnip field, catnip not enough
			{1, 0},
			// gather 10th catnip and buy catnip field
			{0, 0}, {1, 0},
			// wait 1 second and 10 seconds
			{-1, 1}, {-1, 10},
			// buy catnip field, catnip not enough
			{1, 0},
			// wait 1 second and buy catnip field
			{-1, 1}, {1, 0},
			// wait 1 second and 10 seconds
			{-1, 1}, {-1, 10},
			// end
			{999, 0},
		},
	}}
	for _, in := range inputs {
		var buf bytes.Buffer
		logger := log.New(&buf, "", 0 /* flags */)
		input := make(chan int)
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
		if err := g.Validate(); err != nil {
			t.Errorf("[%s] got err %v", in.name, err)
		}
		g.Run(logger, "###", input, nowfn)
		if err := os.WriteFile(in.name+".out", buf.Bytes(), 0644); err != nil {
			t.Errorf("[%s] got err %v", in.name, err)
		}
	}
}
