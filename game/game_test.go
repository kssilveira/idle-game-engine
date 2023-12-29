package game

import (
	"testing"
	"time"

	"github.com/kssilveira/idle-game-engine/data"
)

func TestAct(t *testing.T) {
	inputs := []struct {
		name          string
		resources     []data.Resource
		actions       []data.Action
		inputs        []string
		want          []int
		wantCap       int
		wantResources map[string]int
	}{{
		name: "add 1",
		resources: []data.Resource{{
			Name: "resource", Cap: 2,
		}},
		actions: []data.Action{{
			Name: "add 1",
			Adds: []data.Resource{{
				Name: "resource", Count: 1,
			}},
		}},
		inputs: []string{"0", "0"},
		want:   []int{1, 2},
	}, {
		name: "cost",
		resources: []data.Resource{{
			Name: "resource", Count: 100, Cap: -1,
		}, {
			Name: "producer", Cap: -1,
		}},
		actions: []data.Action{{
			Name:             "producer",
			CostExponentBase: 2,
			Costs: []data.Resource{{
				Name: "resource", Count: 1,
			}},
			Adds: []data.Resource{{
				Name: "producer", Count: 1,
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
			Name: "resource", Count: 1, Cap: -1,
			Producers: []data.Resource{{
				Name: "producer", Factor: 1,
			}},
		}, {
			Name: "producer", Cap: -1,
		}},
		actions: []data.Action{{
			Name:             "producer",
			CostExponentBase: 2,
			Costs: []data.Resource{{
				Name: "resource", Count: 1,
			}},
			Adds: []data.Resource{{
				Name: "producer", Count: 1,
			}},
		}},
		inputs: []string{"0", "s0", "s0"},
		want:   []int{0, 1, 1},
		wantResources: map[string]int{
			"producer": 3,
		},
	}, {
		name: "skip producer action",
		resources: []data.Resource{{
			Name: "resource", Count: 1, Cap: -1,
			ProducerAction: "make resource",
		}, {
			Name: "nested", Cap: -1,
			Producers: []data.Resource{{
				Name: "producer", Factor: 1,
			}},
		}, {
			Name: "producer", Cap: -1,
		}},
		actions: []data.Action{{
			Name: "producer",
			Costs: []data.Resource{{
				Name: "resource", Count: 1,
			}},
			Adds: []data.Resource{{
				Name: "producer", Count: 1,
			}},
		}, {
			Name: "make resource",
			Costs: []data.Resource{{
				Name: "nested", Count: 1,
			}},
			Adds: []data.Resource{{
				Name: "resource", Count: 1,
			}},
		}},
		inputs: []string{"0", "s0", "1", "0", "s0", "1", "0"},
		want:   []int{0, 0, 1, 0, 0, 1, 0},
	}, {
		name: "make producer action",
		resources: []data.Resource{{
			Name: "resource", Count: 1, Cap: -1,
			ProducerAction: "make resource",
		}, {
			Name: "nested", Cap: -1,
			Producers: []data.Resource{{
				Name: "producer", Factor: 1,
			}},
		}, {
			Name: "producer", Cap: -1,
		}},
		actions: []data.Action{{
			Name:             "producer",
			CostExponentBase: 2,
			Costs: []data.Resource{{
				Name: "resource", Count: 1,
			}},
			Adds: []data.Resource{{
				Name: "producer", Count: 1,
			}},
		}, {
			Name: "make resource",
			Costs: []data.Resource{{
				Name: "nested", Count: 1,
			}},
			Adds: []data.Resource{{
				Name: "resource", Count: 1,
			}},
		}},
		inputs: []string{"0", "s0", "s0"},
		want:   []int{0, 0, 0},
		wantResources: map[string]int{
			"producer": 3,
		},
	}, {
		name: "make partial producer action",
		resources: []data.Resource{{
			Name: "resource", Count: 1, Cap: -1,
			ProducerAction: "make resource",
		}, {
			Name: "nested", Cap: 2,
			Producers: []data.Resource{{
				Name: "producer", Factor: 1,
			}},
		}, {
			Name: "producer", Cap: -1,
		}},
		actions: []data.Action{{
			Name:             "producer",
			CostExponentBase: 2,
			Costs: []data.Resource{{
				Name: "resource", Count: 1,
			}},
			Adds: []data.Resource{{
				Name: "producer", Count: 1,
			}},
		}, {
			Name: "make resource",
			Costs: []data.Resource{{
				Name: "nested", Count: 1,
			}},
			Adds: []data.Resource{{
				Name: "resource", Count: 1,
			}},
		}},
		inputs: []string{"0", "s0", "s0", "s0"},
		want:   []int{0, 0, 2, 0},
		wantResources: map[string]int{
			"producer": 3,
		},
	}, {
		name: "add 1 capacity",
		resources: []data.Resource{{
			Name: "resource",
		}},
		actions: []data.Action{{
			Name: "add 1",
			Adds: []data.Resource{{
				Name: "resource", Cap: 1,
			}},
		}},
		inputs:  []string{"0", "0"},
		want:    []int{0, 0},
		wantCap: 2,
	}, {
		name: "cost 1 capacity",
		resources: []data.Resource{{
			Name: "resource", Cap: 4,
		}, {
			Name: "other", Cap: -1,
		}},
		actions: []data.Action{{
			Name: "cost 1",
			Costs: []data.Resource{{
				Name: "resource", Cap: 1,
			}},
			Adds: []data.Resource{{
				Name: "other", Count: 1,
			}},
		}},
		inputs:  []string{"0", "0"},
		want:    []int{0, 0},
		wantCap: 2,
	}, {
		name: "add production bonus",
		resources: []data.Resource{{
			Name: "resource", Count: 2, Cap: -1,
		}},
		actions: []data.Action{{
			Name: "add",
			Adds: []data.Resource{{
				Name: "resource", Count: 3,
				Bonus: []data.Resource{{
					Name: "resource", Factor: 4,
				}},
			}},
		}},
		inputs: []string{"0", "0"},
		want: []int{
			2 + 3*(1+2*4),
			2 + 3*(1+2*4) + 3*(1+(2+3*(1+2*4))*4),
		},
	}, {
		name: "max",
		resources: []data.Resource{{
			Name: "resource", Count: 1, Cap: 10,
			Producers: []data.Resource{{
				Name: "producer", Factor: 1,
			}},
		}, {
			Name: "producer", Cap: -1,
		}},
		actions: []data.Action{{
			Name:             "producer",
			CostExponentBase: 2,
			Costs: []data.Resource{{
				Name: "resource", Count: 1,
			}},
			Adds: []data.Resource{{
				Name: "producer", Count: 1,
			}},
		}},
		inputs: []string{"0", "m0"},
		want:   []int{0, 10},
		wantResources: map[string]int{
			"producer": 4, // cost 1, 2, 4, 8, 16
		},
	}}
	for _, in := range inputs {
		g := NewGame(time.Unix(0, 0))
		g.AddResources(in.resources)
		g.AddActions(in.actions)
		if err := g.Validate(); err != nil {
			t.Errorf("[%s] Validate got err %v", in.name, err)
		}
		for index, input := range in.inputs {
			if _, err := g.act(input); err != nil {
				t.Errorf("[%s] index %d got err %v", in.name, index, err)
			}
			want := in.want[index]
			got := int(g.GetResource("resource").Count)
			if got != want {
				t.Errorf("[%s] index %d want %d got %d", in.name, index, want, got)
			}
		}
		got := int(g.GetResource("resource").Cap)
		if got != in.wantCap && in.wantCap != 0 {
			t.Errorf("[%s] capacity want %d got %d", in.name, in.wantCap, got)
		}
		for name, want := range in.wantResources {
			got := int(g.GetResource(name).Count)
			if got != want {
				t.Errorf("[%s] resource %s want %d got %d", in.name, name, want, got)
			}
		}
	}
}

func TestUpdate(t *testing.T) {
	inputs := []struct {
		name          string
		resources     []data.Resource
		times         []int
		want          []int
		wantResources map[string]int
		wantCap       []int
	}{{
		name: "one input",
		resources: []data.Resource{{
			Name: "resource", Cap: -1,
			Producers: []data.Resource{{
				Name: "input", Factor: 2,
			}},
		}, {
			Name: "input", Count: 3, Cap: -1,
		}},
		times: []int{4, 5, 6},
		want: []int{
			2 * 3 * 4,
			2 * 3 * 5,
			2 * 3 * 6,
		},
	}, {
		name: "over cap",
		resources: []data.Resource{{
			Name: "resource", Cap: 28,
			Producers: []data.Resource{{
				Name: "input", Factor: 2,
			}},
		}, {
			Name: "input", Count: 3, Cap: -1,
		}},
		times: []int{4, 5, 6},
		want: []int{
			2 * 3 * 4,
			28,
			28,
		},
	}, {
		name: "two producers",
		resources: []data.Resource{{
			Name: "resource", Cap: -1,
			Producers: []data.Resource{{
				Name: "input 1", Factor: 2,
			}, {
				Name: "input 2", Factor: 3,
			}},
		}, {
			Name: "input 1", Count: 4, Cap: -1,
		}, {
			Name: "input 2", Count: 5, Cap: -1,
		}},
		times: []int{6, 7, 8},
		want: []int{
			(2*4 + 3*5) * 6,
			(2*4 + 3*5) * 7,
			(2*4 + 3*5) * 8,
		},
	}, {
		name: "production floor",
		resources: []data.Resource{{
			Name: "resource", Cap: -1,
			Producers: []data.Resource{{
				Name: "input", Factor: 2, ProductionFloor: true,
			}},
		}, {
			Name: "input", Count: 0, Cap: -1,
			Producers: []data.Resource{{
				Name: "", Factor: 0.5,
			}},
		}},
		times: []int{1, 2, 3, 4, 5},
		want: []int{
			0,
			0,
			2,
			2 + 2,
			2 + 2 + 4,
		},
	}, {
		name: "production boolean",
		resources: []data.Resource{{
			Name: "resource", Cap: -1,
			Producers: []data.Resource{{
				Name: "input", Factor: 1, ProductionBoolean: true,
			}},
		}, {
			Name: "input", Count: 0, Cap: -1,
			Producers: []data.Resource{{
				Name: "", Factor: 0.25,
			}},
		}},
		times: []int{1, 3, 4, 5, 7, 8, 9},
		want:  []int{0, 2, 3, 4, 6, 7, 8},
	}, {
		name: "negative production",
		resources: []data.Resource{{
			Name: "resource", Count: 2, Cap: -1,
			Producers: []data.Resource{{
				Name: "input", Factor: -0.2, ProductionOnGone: true,
			}},
		}, {
			Name: "input", Count: 5, Cap: -1,
			OnGone: []data.Resource{{
				Name: "gone input", Count: 1,
			}},
		}, {
			Name: "gone input", Cap: -1,
		}},
		times: []int{0, 1, 2, 3},
		want: []int{
			2, 1, 0, 0,
		},
		wantResources: map[string]int{
			"input":      3,
			"gone input": 2,
		},
	}, {
		name: "negative production not ongone",
		resources: []data.Resource{{
			Name: "resource", Count: 2, Cap: -1,
			Producers: []data.Resource{{
				Name: "input", Factor: -0.2,
			}},
		}, {
			Name: "input", Count: 5, Cap: -1,
			OnGone: []data.Resource{{
				Name: "gone input", Count: 1,
			}},
		}, {
			Name: "gone input", Cap: -1,
		}},
		times: []int{0, 1, 2, 3},
		want: []int{
			2, 1, 0, 0,
		},
		wantResources: map[string]int{
			"input":      5,
			"gone input": 0,
		},
	}, {
		name: "start quantity",
		resources: []data.Resource{{
			Name: "resource", StartCount: 10, Cap: -1,
			Producers: []data.Resource{{
				Name: "input", Factor: 2,
			}},
		}, {
			Name: "input", Count: 3, Cap: -1,
		}},
		times: []int{4, 5, 6},
		want: []int{
			10 + 2*3,
			10 + 2*3,
			10 + 2*3,
		},
	}, {
		name: "start quantity modulus",
		resources: []data.Resource{{
			Name: "resource", StartCount: 10, ProductionModulus: 4, ProductionModulusEquals: -1, Cap: -1,
			Producers: []data.Resource{{
				Name: "input", Factor: 2,
			}},
		}, {
			Name: "input", Count: 3, Cap: -1,
		}},
		times: []int{4, 5, 6},
		want: []int{
			10 + 2*3%4,
			10 + 2*3%4,
			10 + 2*3%4,
		},
	}, {
		name: "start quantity modulus equals",
		resources: []data.Resource{{
			Name: "input", Count: 0, Cap: -1,
			Producers: []data.Resource{{
				Name: "", Factor: 1,
			}},
		}, {
			Name: "resource", StartCount: 1, ProductionModulus: 2, ProductionModulusEquals: 1,
			Cap: -1,
			Producers: []data.Resource{{
				Name: "input", Factor: 1,
			}},
		}},
		times: []int{0, 1, 2, 3},
		want:  []int{0, 1, 0, 1},
	}, {
		name: "time",
		resources: []data.Resource{{
			Name: "resource", StartCount: 1, Cap: -1,
			Producers: []data.Resource{{
				Name: "time", Factor: 1,
			}},
		}},
		times: []int{4, 5, 6},
		want: []int{
			1 + 4,
			1 + 5,
			1 + 6,
		},
	}, {
		name: "production bonus",
		resources: []data.Resource{{
			Name: "resource", Cap: -1,
			Producers: []data.Resource{{
				Name: "input", Factor: 2,
			}},
			Bonus: []data.Resource{{
				Name: "input", Factor: 7,
			}},
		}, {
			Name: "input", Count: 3, Cap: -1,
		}},
		times: []int{4, 5, 6},
		want: []int{
			2 * 3 * 4 * (1 + 3*7),
			2 * 3 * 5 * (1 + 3*7),
			2 * 3 * 6 * (1 + 3*7),
		},
	}, {
		name: "production bonus inside producer",
		resources: []data.Resource{{
			Name: "resource", Cap: -1,
			Producers: []data.Resource{{
				Name: "input", Factor: 2,
				Bonus: []data.Resource{{
					Name: "input", Factor: 7,
				}},
			}},
		}, {
			Name: "input", Count: 3, Cap: -1,
		}},
		times: []int{4, 5, 6},
		want: []int{
			2 * 3 * 4 * (1 + 3*7),
			2 * 3 * 5 * (1 + 3*7),
			2 * 3 * 6 * (1 + 3*7),
		},
	}, {
		name: "two producers production bonus",
		resources: []data.Resource{{
			Name: "resource", Cap: -1,
			Producers: []data.Resource{{
				Name: "input 1", Factor: 2,
			}, {
				Name: "input 2", Factor: 3,
			}},
			Bonus: []data.Resource{{
				Name: "input 1", Factor: 9,
			}, {
				Name: "input 2", Factor: 10,
			}},
		}, {
			Name: "input 1", Count: 4, Cap: -1,
		}, {
			Name: "input 2", Count: 5, Cap: -1,
		}},
		times: []int{6, 7, 8},
		want: []int{
			(2*4 + 3*5) * (1 + 4*9 + 5*10) * 6,
			(2*4 + 3*5) * (1 + 4*9 + 5*10) * 7,
			(2*4 + 3*5) * (1 + 4*9 + 5*10) * 8,
		},
	}, {
		name: "two producers production bonus inside producers",
		resources: []data.Resource{{
			Name: "resource", Cap: -1,
			Producers: []data.Resource{{
				Name: "input 1", Factor: 2,
				Bonus: []data.Resource{{
					Name: "input 2", Factor: 10,
				}},
			}, {
				Name: "input 2", Factor: 3,
				Bonus: []data.Resource{{
					Name: "input 1", Factor: 9,
				}},
			}},
		}, {
			Name: "input 1", Count: 4, Cap: -1,
		}, {
			Name: "input 2", Count: 5, Cap: -1,
		}},
		times: []int{6, 7, 8},
		want: []int{
			(2*4*(1+5*10) + 3*5*(1+4*9)) * 6,
			(2*4*(1+5*10) + 3*5*(1+4*9)) * 7,
			(2*4*(1+5*10) + 3*5*(1+4*9)) * 8,
		},
	}}
	for _, in := range inputs {
		g := NewGame(time.Unix(0, 0))
		g.AddResources(in.resources)
		if err := g.Validate(); err != nil {
			t.Errorf("[%s] Validate got err %v", in.name, err)
		}
		for index, one := range in.times {
			g.update(time.Unix(int64(one), 0))
			if len(in.want) > 0 {
				want := in.want[index]
				got := int(g.GetResource("resource").Count)
				if got != want {
					t.Errorf("[%s] index %d want %d got %d", in.name, index, want, got)
				}
			}
			if len(in.wantCap) > 0 {
				want := in.wantCap[index]
				got := int(g.GetResource("resource").Cap)
				if got != want {
					t.Errorf("[%s] index %d want capacity %d got %d", in.name, index, want, got)
				}
			}
		}
		for name, want := range in.wantResources {
			got := int(g.GetResource(name).Count)
			if got != want {
				t.Errorf("[%s] resource %s want %d got %d", in.name, name, want, got)
			}
		}
	}
}
