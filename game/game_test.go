package game

import (
	"testing"
	"time"

	"github.com/kssilveira/idle-game-engine/data"
)

func TestAct(t *testing.T) {
	inputs := []struct {
		name         string
		resources    []data.Resource
		actions      []Action
		inputs       []string
		want         []int
		wantCapacity int
	}{{
		name: "add 1",
		resources: []data.Resource{{
			Name: "resource", Capacity: 2,
		}},
		actions: []Action{{
			Name: "add 1",
			Adds: []data.Resource{{
				Name: "resource", Quantity: 1,
			}},
		}},
		inputs: []string{"0", "0"},
		want:   []int{1, 2},
	}, {
		name: "cost",
		resources: []data.Resource{{
			Name: "resource", Quantity: 100, Capacity: -1,
		}, {
			Name: "producer", Capacity: -1,
		}},
		actions: []Action{{
			Name: "producer",
			Costs: []data.Resource{{
				Name: "resource", Quantity: 1, CostExponentBase: 2,
			}},
			Adds: []data.Resource{{
				Name: "producer", Quantity: 1,
			}},
		}},
		inputs: []string{"0", "0", "0"},
		want: []int{
			100 - (1),
			100 - (1 + 2),
			100 - (1 + 2 + 4),
		},
	}, {
		name: "skip",
		resources: []data.Resource{{
			Name: "resource", Quantity: 1, Capacity: -1,
			Producers: []data.Resource{{
				Name: "producer", ProductionFactor: 1,
			}},
		}, {
			Name: "producer", Capacity: -1,
		}, {
			Name: "skip", Capacity: -1,
		}},
		actions: []Action{{
			Name: "producer",
			Costs: []data.Resource{{
				Name: "resource", Quantity: 1, CostExponentBase: 2,
			}},
			Adds: []data.Resource{{
				Name: "producer", Quantity: 1,
			}},
		}},
		inputs: []string{"0", "s0", "0", "s0", "0"},
		want:   []int{0, 3, 1, 5, 1},
	}, {
		name: "add 1 capacity",
		resources: []data.Resource{{
			Name: "resource",
		}},
		actions: []Action{{
			Name: "add 1",
			Adds: []data.Resource{{
				Name: "resource", Capacity: 1,
			}},
		}},
		inputs:       []string{"0", "0"},
		want:         []int{0, 0},
		wantCapacity: 2,
	}, {
		name: "cost 1 capacity",
		resources: []data.Resource{{
			Name: "resource", Capacity: 4,
		}, {
			Name: "other", Capacity: -1,
		}},
		actions: []Action{{
			Name: "cost 1",
			Costs: []data.Resource{{
				Name: "resource", Capacity: 1,
			}},
			Adds: []data.Resource{{
				Name: "other", Quantity: 1,
			}},
		}},
		inputs:       []string{"0", "0"},
		want:         []int{0, 0},
		wantCapacity: 2,
	}}
	for _, in := range inputs {
		g := NewGame(time.Unix(0, 0))
		g.AddResources(in.resources)
		g.Actions = in.actions
		if err := g.Validate(); err != nil {
			t.Errorf("[%s] Validate got err %v", in.name, err)
		}
		for index, input := range in.inputs {
			if err := g.Act(input); err != nil {
				t.Errorf("[%s] index %d got err %v", in.name, index, err)
			}
			want := in.want[index]
			got := int(g.GetResource("resource").Quantity)
			if got != want {
				t.Errorf("[%s] index %d want %d got %d", in.name, index, want, got)
			}
		}
		got := int(g.GetResource("resource").Capacity)
		if got != in.wantCapacity && in.wantCapacity != 0 {
			t.Errorf("[%s] capacity want %d got %d", in.name, in.wantCapacity, got)
		}
	}
}

func TestUpdate(t *testing.T) {
	inputs := []struct {
		name          string
		resources     []data.Resource
		times         []int64
		want          []int
		wantResources map[string]int
	}{{
		name: "one input",
		resources: []data.Resource{{
			Name: "resource", Capacity: -1,
			Producers: []data.Resource{{
				Name: "input", ProductionFactor: 2,
			}},
		}, {
			Name: "input", Quantity: 3, Capacity: -1,
		}},
		times: []int64{4, 5, 6},
		want: []int{
			2 * 3 * 4,
			2 * 3 * 5,
			2 * 3 * 6,
		},
	}, {
		name: "over cap",
		resources: []data.Resource{{
			Name: "resource", Capacity: 28,
			Producers: []data.Resource{{
				Name: "input", ProductionFactor: 2,
			}},
		}, {
			Name: "input", Quantity: 3, Capacity: -1,
		}},
		times: []int64{4, 5, 6},
		want: []int{
			2 * 3 * 4,
			28,
			28,
		},
	}, {
		name: "two producers",
		resources: []data.Resource{{
			Name: "resource", Capacity: -1,
			Producers: []data.Resource{{
				Name: "input 1", ProductionFactor: 2,
			}, {
				Name: "input 2", ProductionFactor: 3,
			}},
		}, {
			Name: "input 1", Quantity: 4, Capacity: -1,
		}, {
			Name: "input 2", Quantity: 5, Capacity: -1,
		}},
		times: []int64{6, 7, 8},
		want: []int{
			(2*4 + 3*5) * 6,
			(2*4 + 3*5) * 7,
			(2*4 + 3*5) * 8,
		},
	}, {
		name: "one resource factor",
		resources: []data.Resource{{
			Name: "resource", Capacity: -1,
			Producers: []data.Resource{{
				Name: "input", ProductionFactor: 2, ProductionResourceFactor: "resource factor",
			}},
		}, {
			Name: "input", Quantity: 3, Capacity: -1,
		}, {
			Name: "resource factor", Quantity: 4, Capacity: -1,
		}},
		times: []int64{5, 6, 7},
		want: []int{
			2 * 3 * 4 * 5,
			2 * 3 * 4 * 6,
			2 * 3 * 4 * 7,
		},
	}, {
		name: "production floor",
		resources: []data.Resource{{
			Name: "resource", Capacity: -1,
			Producers: []data.Resource{{
				Name: "input", ProductionFactor: 2, ProductionFloor: true,
			}},
		}, {
			Name: "input", Quantity: 0, Capacity: -1,
			Producers: []data.Resource{{
				Name: "", ProductionFactor: 0.5,
			}},
		}},
		times: []int64{1, 2, 3, 4, 5},
		want: []int{
			0,
			0,
			2,
			2 + 2,
			2 + 2 + 4,
		},
	}, {
		name: "negative production",
		resources: []data.Resource{{
			Name: "resource", Quantity: 2, Capacity: -1,
			Producers: []data.Resource{{
				Name: "input", ProductionFactor: -0.2,
			}},
		}, {
			Name: "input", Quantity: 5, Capacity: -1,
			OnGone: []data.Resource{{
				Name: "gone input", Quantity: 1,
			}},
		}, {
			Name: "gone input", Capacity: -1,
		}},
		times: []int64{0, 1, 2, 3},
		want: []int{
			2, 1, 0, 0,
		},
		wantResources: map[string]int{
			"input":      3,
			"gone input": 2,
		},
	}, {
		name: "start quantity",
		resources: []data.Resource{{
			Name: "resource", StartQuantity: 10, Capacity: -1,
			Producers: []data.Resource{{
				Name: "input", ProductionFactor: 2,
			}},
		}, {
			Name: "input", Quantity: 3, Capacity: -1,
		}},
		times: []int64{4, 5, 6},
		want: []int{
			10 + 2*3,
			10 + 2*3,
			10 + 2*3,
		},
	}, {
		name: "start quantity modulus",
		resources: []data.Resource{{
			Name: "resource", StartQuantity: 10, ProductionModulus: 4, Capacity: -1,
			Producers: []data.Resource{{
				Name: "input", ProductionFactor: 2,
			}},
		}, {
			Name: "input", Quantity: 3, Capacity: -1,
		}},
		times: []int64{4, 5, 6},
		want: []int{
			10 + 2*3%4,
			10 + 2*3%4,
			10 + 2*3%4,
		},
	}}
	for _, in := range inputs {
		g := NewGame(time.Unix(0, 0))
		g.AddResources(in.resources)
		if err := g.Validate(); err != nil {
			t.Errorf("[%s] Validate got err %v", in.name, err)
		}
		for index, one := range in.times {
			g.Update(time.Unix(one, 0))
			want := in.want[index]
			got := int(g.GetResource("resource").Quantity)
			if got != want {
				t.Errorf("[%s] index %d want %d got %d", in.name, index, want, got)
			}
		}
		for name, want := range in.wantResources {
			got := int(g.GetResource(name).Quantity)
			if got != want {
				t.Errorf("[%s] resource %s want %d got %d", in.name, name, want, got)
			}
		}
	}
}
