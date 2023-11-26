package kittens

import (
	"bytes"
	"log"
	"strings"
	"testing"
	"time"
)

func TestRun(t *testing.T) {
	inputs := []struct {
		name   string
		inputs []int
		times  []int64
		want   string
	}{{
		name:   "0 0 0",
		inputs: []int{1, 0},
		times:  []int64{1, 2, 3},
		want: `
###
Spring 1.00
0: 'Gather catnip' (catnip + 1)
1: 'Catnip Field' (Catnip Field + 1)
###
Catnip Field 1.00
Spring 1.00
0: 'Gather catnip' (catnip + 1)
1: 'Catnip Field' (Catnip Field + 1)
###
catnip 1.94/5000 (0.94/s, 1h28m8s to cap)
Catnip Field 1.00
Spring 1.00
0: 'Gather catnip' (catnip + 1)
1: 'Catnip Field' (Catnip Field + 1)
`,
	}}
	for _, in := range inputs {
		var buf bytes.Buffer
		logger := log.New(&buf, "", 0 /* flags */)
		input := make(chan int)
		go func() {
			for _, one := range in.inputs {
				input <- one
			}
			input <- 999
		}()
		timeIndex := 0
		Run(logger, "###", input, func() time.Time {
			res := time.Unix(in.times[timeIndex], 0)
			timeIndex++
			return res
		})
		got := buf.String()
		if strings.TrimSpace(got) != strings.TrimSpace(in.want) {
			t.Errorf("[%s] want:\n%s\ngot:\n%s\n", in.name, in.want, got)
		}
	}
}
