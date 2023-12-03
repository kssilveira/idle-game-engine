package kittens

import (
	"fmt"
	"time"

	"github.com/kssilveira/idle-game-engine/data"
	"github.com/kssilveira/idle-game-engine/game"
)

func NewGame(now game.Now) *game.Game {
	g := game.NewGame(now())
	g.AddResources([]data.Resource{{
		Name: "day", Type: "Calendar", Quantity: 0, Capacity: -1,
		Producers: []data.Resource{{
			Name: "", ProductionFactor: 0.5,
		}},
	}, {
		Name: "year", Type: "Calendar", Quantity: 1, StartQuantity: 1, Capacity: -1,
		Producers: []data.Resource{{
			Name: "day", ProductionFactor: 0.0025, ProductionFloor: true,
		}},
	}, {
		Name: "Spring", Type: "Calendar", Quantity: 1,
		StartQuantity: 1, ProductionModulus: 4, ProductionModulusEquals: 0, Capacity: -1,
		Producers: []data.Resource{{
			Name: "day", ProductionFactor: 0.01, ProductionFloor: true,
		}},
	}, {
		Name: "Summer", Type: "Calendar",
		StartQuantity: 1, ProductionModulus: 4, ProductionModulusEquals: 1, Capacity: -1,
		Producers: []data.Resource{{
			Name: "day", ProductionFactor: 0.01, ProductionFloor: true,
		}},
	}, {
		Name: "Autumn", Type: "Calendar",
		StartQuantity: 1, ProductionModulus: 4, ProductionModulusEquals: 2, Capacity: -1,
		Producers: []data.Resource{{
			Name: "day", ProductionFactor: 0.01, ProductionFloor: true,
		}},
	}, {
		Name: "Winter", Type: "Calendar",
		StartQuantity: 1, ProductionModulus: 4, ProductionModulusEquals: 3, Capacity: -1,
		Producers: []data.Resource{{
			Name: "day", ProductionFactor: 0.01, ProductionFloor: true,
		}},
	}, {
		Name: "day_of_year", Type: "Calendar", Quantity: 1,
		StartQuantity: 1, ProductionModulus: 400, ProductionModulusEquals: -1,
		Capacity: -1,
		Producers: []data.Resource{{
			Name: "day", ProductionFactor: 1, ProductionFloor: true,
		}},
	}, {
		Name: "catnip", Type: "Resource", Capacity: 5000,
		Producers: []data.Resource{{
			Name: "Catnip Field", ProductionFactor: 0.63 * (1 + 0.50), ProductionResourceFactor: "Spring",
		}, {
			Name: "Catnip Field", ProductionFactor: 0.63, ProductionResourceFactor: "Summer",
		}, {
			Name: "Catnip Field", ProductionFactor: 0.63, ProductionResourceFactor: "Autumn",
		}, {
			Name: "Catnip Field", ProductionFactor: 0.63 * (1 - 0.75), ProductionResourceFactor: "Winter",
		}, {
			Name: "kitten", ProductionFactor: -4.25, ProductionFloor: true,
		}, {
			Name: "woodcutter", ProductionFactor: -4.25,
		}, {
			Name: "scholar", ProductionFactor: -4.25,
		}, {
			Name: "farmer", ProductionFactor: -4.25,
		}, {
			Name: "farmer", ProductionFactor: 5,
		}},
	}, {
		Name: "wood", Type: "Resource", Capacity: 200,
		Producers: []data.Resource{{
			Name: "woodcutter", ProductionFactor: 0.09, ProductionResourceFactor: "happiness",
		}},
	}, {
		Name: "science", Type: "Resource", Capacity: 250,
		Producers: []data.Resource{{
			Name: "scholar", ProductionFactor: 0.175, ProductionResourceFactor: "happiness",
		}},
		ProductionBonus: []data.Resource{{
			Name: "Library", ProductionFactor: 0.1,
		}},
	}, {
		Name: "kitten", Type: "Resource", Capacity: 0,
		Producers: []data.Resource{{
			Name: "", ProductionFactor: 0.05,
		}},
		OnGone: []data.Resource{{
			Name: "gone kitten", Quantity: 1,
		}},
	}, {
		Name: "gone kitten", Type: "Resource", Capacity: -1,
	}, {
		Name: "Catnip Field", Type: "Bonfire", Capacity: -1,
	}, {
		Name: "Hut", Type: "Bonfire", Capacity: -1,
	}, {
		Name: "Library", Type: "Bonfire", Capacity: -1,
	}, {
		Name: "Barn", Type: "Bonfire", Capacity: -1,
	}, {
		Name: "woodcutter", Type: "Village", Capacity: -1,
		OnGone: []data.Resource{{
			Name: "gone kitten", Quantity: 1,
		}, {
			Name: "kitten", Capacity: 1,
		}},
	}, {
		Name: "scholar", Type: "Village", Capacity: -1,
		OnGone: []data.Resource{{
			Name: "gone kitten", Quantity: 1,
		}, {
			Name: "kitten", Capacity: 1,
		}},
	}, {
		Name: "farmer", Type: "Village", Capacity: -1,
		OnGone: []data.Resource{{
			Name: "gone kitten", Quantity: 1,
		}, {
			Name: "kitten", Capacity: 1,
		}},
	}, {
		Name: "happiness", Type: "Village", StartQuantity: 1.1, Capacity: -1,
		Producers: []data.Resource{{
			Name: "kitten", ProductionFactor: -0.02, ProductionFloor: true,
		}, {
			Name: "woodcutter", ProductionFactor: -0.02,
		}, {
			Name: "scholar", ProductionFactor: -0.02,
		}, {
			Name: "farmer", ProductionFactor: -0.02,
		}},
	}, {
		Name: "Calendar", Type: "Science", Capacity: 1,
	}, {
		Name: "Agriculture", Type: "Science", Capacity: 1,
	}, {
		Name: "Archery", Type: "Science", Capacity: 1,
	}, {
		Name: "Mining", Type: "Science", Capacity: 1,
	}})
	g.Actions = []game.Action{{
		Name: "Gather catnip", Type: "Bonfire",
		LockedBy: data.Resource{Name: "Catnip Field"},
		Adds: []data.Resource{{
			Name: "catnip", Quantity: 1,
		}},
	}, {
		Name: "Refine catnip", Type: "Bonfire",
		UnlockedBy: data.Resource{Name: "catnip"},
		LockedBy:   data.Resource{Name: "woodcutter"},
		Costs: []data.Resource{{
			Name: "catnip", Quantity: 100, CostExponentBase: 1,
		}},
		Adds: []data.Resource{{
			Name: "wood", Quantity: 1,
		}},
	}, {
		Name: "Catnip Field", Type: "Bonfire",
		UnlockedBy: data.Resource{Name: "catnip"},
		Costs: []data.Resource{{
			Name: "catnip", Quantity: 10, CostExponentBase: 1.12,
		}},
		Adds: []data.Resource{{
			Name: "Catnip Field", Quantity: 1,
		}},
	}, {
		Name: "Hut", Type: "Bonfire",
		UnlockedBy: data.Resource{Name: "wood"},
		Costs: []data.Resource{{
			Name: "wood", Quantity: 5, CostExponentBase: 2.5,
		}},
		Adds: []data.Resource{{
			Name: "Hut", Quantity: 1,
		}, {
			Name: "kitten", Capacity: 2,
		}},
	}, {
		Name: "Library", Type: "Bonfire",
		UnlockedBy: data.Resource{Name: "wood"},
		Costs: []data.Resource{{
			Name: "wood", Quantity: 25, CostExponentBase: 1.15,
		}},
		Adds: []data.Resource{{
			Name: "Library", Quantity: 1,
		}, {
			Name: "science", Capacity: 250,
		}},
	}, {
		Name: "Barn", Type: "Bonfire",
		UnlockedBy: data.Resource{Name: "Agriculture"},
		Costs: []data.Resource{{
			Name: "wood", Quantity: 50, CostExponentBase: 1.75,
		}},
		Adds: []data.Resource{{
			Name: "Barn", Quantity: 1,
		}, {
			Name: "catnip", Capacity: 5000,
		}, {
			Name: "wood", Capacity: 200,
		}},
	}, {
		Name: "woodcutter", Type: "Village",
		UnlockedBy: data.Resource{Name: "Hut"},
		Costs: []data.Resource{{
			Name: "kitten", Quantity: 1, Capacity: 1, CostExponentBase: 1,
		}},
		Adds: []data.Resource{{
			Name: "woodcutter", Quantity: 1,
		}},
	}, {
		Name: "scholar", Type: "Village",
		UnlockedBy: data.Resource{Name: "Library"},
		Costs: []data.Resource{{
			Name: "kitten", Quantity: 1, Capacity: 1, CostExponentBase: 1,
		}},
		Adds: []data.Resource{{
			Name: "scholar", Quantity: 1,
		}},
	}, {
		Name: "farmer", Type: "Village",
		UnlockedBy: data.Resource{Name: "Agriculture"},
		Costs: []data.Resource{{
			Name: "kitten", Quantity: 1, Capacity: 1, CostExponentBase: 1,
		}},
		Adds: []data.Resource{{
			Name: "farmer", Quantity: 1,
		}},
	}, {
		Name: "Calendar", Type: "Science",
		UnlockedBy: data.Resource{Name: "Library"},
		LockedBy:   data.Resource{Name: "Calendar"},
		Costs: []data.Resource{{
			Name: "science", Quantity: 30, CostExponentBase: 1,
		}},
		Adds: []data.Resource{{
			Name: "Calendar", Quantity: 1,
		}},
	}, {
		Name: "Agriculture", Type: "Science",
		UnlockedBy: data.Resource{Name: "Calendar"},
		LockedBy:   data.Resource{Name: "Agriculture"},
		Costs: []data.Resource{{
			Name: "science", Quantity: 100, CostExponentBase: 1,
		}},
		Adds: []data.Resource{{
			Name: "Agriculture", Quantity: 1,
		}},
	}, {
		Name: "Archery", Type: "Science",
		UnlockedBy: data.Resource{Name: "Calendar"},
		LockedBy:   data.Resource{Name: "Archery"},
		Costs: []data.Resource{{
			Name: "science", Quantity: 300, CostExponentBase: 1,
		}},
		Adds: []data.Resource{{
			Name: "Archery", Quantity: 1,
		}},
	}, {
		Name: "Mining", Type: "Science",
		UnlockedBy: data.Resource{Name: "Calendar"},
		LockedBy:   data.Resource{Name: "Mining"},
		Costs: []data.Resource{{
			Name: "science", Quantity: 500, CostExponentBase: 1,
		}},
		Adds: []data.Resource{{
			Name: "Mining", Quantity: 1,
		}},
	}}
	return g
}

const (
	gather = iota
	refine
	field
	hut
	library
	barn
	woodcutter
	scholar
	farmer
	calendar
	agriculture
	archery
	mining
)

const (
	sdelta = 100
)

const (
	_ = iota + sdelta
	srefine
	sfield
	shut
	slibrary
	sbarn
	swoodcutter
	sscholar
	sfarmer
	scalendar
	sagriculture
	sarchery
	smining
)

func Solve(input chan string, sleepMS int) {
	for _, one := range []struct {
		cmds  []int
		count int
	}{
		{[]int{gather}, 10},
		{[]int{field}, 1},
		{[]int{sfield, field}, 58},
		{[]int{srefine, refine}, 5},
		{[]int{hut, swoodcutter, woodcutter}, 1},
		{[]int{slibrary, library, sscholar, scholar}, 1},
		{[]int{slibrary, library}, 14},
		{[]int{scalendar, calendar}, 1},
		{[]int{sagriculture, agriculture}, 1},
		{[]int{sarchery, archery}, 1},
		{[]int{smining, mining}, 1},
		{[]int{sbarn, barn}, 6},

		{[]int{sfield, field}, 25},
		{[]int{slibrary, library}, 15},
		{[]int{shut, hut, sfarmer, farmer}, 10},
	} {
		for i := 0; i < one.count; i++ {
			for _, cmd := range one.cmds {
				input <- toInput(cmd)
				time.Sleep(time.Second * time.Duration(sleepMS) / 1000.)
			}
		}
	}
}

func toInput(cmd int) string {
	prefix := ""
	if cmd >= sdelta {
		prefix = "s"
		cmd -= sdelta
	}
	return fmt.Sprintf("%s%d", prefix, cmd)
}
