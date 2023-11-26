package game

import "testing"

func TestGame(t *testing.T) {
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
		g := NewGame()
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
			got := r.Quantity
			if got != want {
				t.Errorf("[%s] index %d want %d got %d", in.name, index, want, got)
			}
		}
	}
}
