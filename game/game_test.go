package game

import "testing"
import "time"

func TestAct(t *testing.T) {
	inputs := []struct {
		name      string
		resources []Resource
		actions   []Action
		inputs    []int
		want      []int
	}{{
		name: "add 1",
		resources: []Resource{{
			Name:     "resource",
			Capacity: 2,
		}},
		actions: []Action{{
			Name: "add 1",
			Add: []Resource{{
				Name:     "resource",
				Quantity: 1,
			}},
		}},
		inputs: []int{0, 0, 0},
		want:   []int{1, 2, 2},
	}, {
		name: "no cap",
		resources: []Resource{{
			Name: "resource",
		}},
		actions: []Action{{
			Name: "add 1",
			Add: []Resource{{
				Name:     "resource",
				Quantity: 1,
			}},
		}},
		inputs: []int{0, 0, 0},
		want:   []int{1, 2, 3},
	}}
	for _, in := range inputs {
		g := NewGame(time.Unix(0, 0))
		g.AddResources(in.resources)
		g.Actions = in.actions
		for index, input := range in.inputs {
			if err := g.Act(input); err != nil {
				t.Errorf("[%s] index %d got err %v", in.name, index, err)
			}
			want := in.want[index]
			r, err := g.GetResource("resource")
			if err != nil {
				t.Errorf("[%s] index %d got err %v", in.name, index, err)
			}
			got := int(r.Quantity)
			if got != want {
				t.Errorf("[%s] index %d want %d got %d", in.name, index, want, got)
			}
		}
	}
}

func TestUpdate(t *testing.T) {
	inputs := []struct {
		name      string
		resources []Resource
		times     []int64
		want      []int
	}{{
		name: "one input",
		resources: []Resource{{
			Name: "resource",
			Rate: []Resource{{
				Name:   "input",
				Factor: 2,
			}},
		}, {
			Name:     "input",
			Quantity: 3,
		}},
		times: []int64{4, 5, 6},
		want:  []int{2 * 3 * 4, 2 * 3 * 5, 2 * 3 * 6},
	}, {
		name: "two inputs",
		resources: []Resource{{
			Name: "resource",
			Rate: []Resource{{
				Name:   "input 1",
				Factor: 2,
			}, {
				Name:   "input 2",
				Factor: 3,
			}},
		}, {
			Name:     "input 1",
			Quantity: 4,
		}, {
			Name:     "input 2",
			Quantity: 5,
		}},
		times: []int64{6, 7, 8},
		want:  []int{(2*4 + 3*5) * 6, (2*4 + 3*5) * 7, (2*4 + 3*5) * 8},
	}, {
		name: "one resource factor",
		resources: []Resource{{
			Name: "resource",
			Rate: []Resource{{
				Name:           "input",
				Factor:         2,
				ResourceFactor: "resource factor",
			}},
		}, {
			Name:     "input",
			Quantity: 3,
		}, {
			Name:     "resource factor",
			Quantity: 4,
		}},
		times: []int64{5, 6, 7},
		want:  []int{2 * 3 * 4 * 5, 2 * 3 * 4 * 6, 2 * 3 * 4 * 7},
	}}
	for _, in := range inputs {
		g := NewGame(time.Unix(0, 0))
		g.AddResources(in.resources)
		for index, one := range in.times {
			if err := g.Update(time.Unix(one, 0)); err != nil {
				t.Errorf("[%s] index %d got err %v", in.name, index, err)
			}
			want := in.want[index]
			r, err := g.GetResource("resource")
			if err != nil {
				t.Errorf("[%s] index %d got err %v", in.name, index, err)
			}
			got := int(r.Quantity)
			if got != want {
				t.Errorf("[%s] index %d want %d got %d", in.name, index, want, got)
			}
		}
	}
}
