package kittens

import "bytes"
import "log"
import "testing"
import "strings"

func TestRun(t *testing.T) {
	inputs := []struct {
		name   string
		inputs []int
		want   string
	}{{
		name:   "0 0 0",
		inputs: []int{0, 1},
		want: `
catnip 0/5000
Catnip Field 0
0: 'Gather catnip' (catnip + 1)
1: 'Catnip Field' (Catnip Field + 1)
catnip 1/5000
Catnip Field 0
0: 'Gather catnip' (catnip + 1)
1: 'Catnip Field' (Catnip Field + 1)
catnip 1/5000
Catnip Field 1
0: 'Gather catnip' (catnip + 1)
1: 'Catnip Field' (Catnip Field + 1)
`,
	}}
	for _, in := range inputs {
		var buf bytes.Buffer
		logger := log.New(&buf, "", 0 /* flags */)
		index := 0
		Run(logger, func() int {
			if index >= len(in.inputs) {
				return 999
			}
			res := in.inputs[index]
			index++
			return res
		})
		got := buf.String()
		if strings.TrimSpace(got) != strings.TrimSpace(in.want) {
			t.Errorf("[%s] want:\n%s\ngot:\n%s\n", in.name, in.want, got)
		}
	}
}
