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
		Name: "day", Quantity: 0, Capacity: -1,
		Producers: []data.Resource{{
			Name: "", ProductionFactor: 0.5,
		}},
	}, {
		Name: "year", Quantity: 1, StartQuantity: 1, Capacity: -1,
		Producers: []data.Resource{{
			Name: "day", ProductionFactor: 0.0025, ProductionFloor: true,
		}},
	}, {
		Name: "Spring", Quantity: 1,
		StartQuantity: 1, ProductionModulus: 4, ProductionModulusEquals: 0, Capacity: -1,
		Producers: []data.Resource{{
			Name: "day", ProductionFactor: 0.01, ProductionFloor: true,
		}},
	}, {
		Name: "Summer", StartQuantity: 1, ProductionModulus: 4, ProductionModulusEquals: 1, Capacity: -1,
		Producers: []data.Resource{{
			Name: "day", ProductionFactor: 0.01, ProductionFloor: true,
		}},
	}, {
		Name: "Autumn", StartQuantity: 1, ProductionModulus: 4, ProductionModulusEquals: 2, Capacity: -1,
		Producers: []data.Resource{{
			Name: "day", ProductionFactor: 0.01, ProductionFloor: true,
		}},
	}, {
		Name: "Winter", StartQuantity: 1, ProductionModulus: 4, ProductionModulusEquals: 3, Capacity: -1,
		Producers: []data.Resource{{
			Name: "day", ProductionFactor: 0.01, ProductionFloor: true,
		}},
	}, {
		Name: "day_of_year", Quantity: 1,
		StartQuantity: 1, ProductionModulus: 400, ProductionModulusEquals: -1,
		Capacity: -1,
		Producers: []data.Resource{{
			Name: "day", ProductionFactor: 1, ProductionFloor: true,
		}},
	}, {
		Name: "catnip", Capacity: 5000,
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
		}},
	}, {
		Name: "wood", Capacity: 200,
		Producers: []data.Resource{{
			Name: "woodcutter", ProductionFactor: 0.09, ProductionResourceFactor: "happiness",
		}},
	}, {
		Name: "science", Capacity: 250,
		Producers: []data.Resource{{
			Name: "scholar", ProductionFactor: 0.175, ProductionResourceFactor: "happiness",
		}},
		ProductionBonus: []data.Resource{{
			Name: "Library", ProductionFactor: 0.1,
		}},
	}, {
		Name: "kitten", Capacity: 0,
		Producers: []data.Resource{{
			Name: "", ProductionFactor: 0.05,
		}},
		OnGone: []data.Resource{{
			Name: "gone kitten", Quantity: 1,
		}},
	}, {
		Name: "gone kitten", Capacity: -1,
	}, {
		Name: "woodcutter", Capacity: -1,
		OnGone: []data.Resource{{
			Name: "gone kitten", Quantity: 1,
		}, {
			Name: "kitten", Capacity: 1,
		}},
	}, {
		Name: "scholar", Capacity: -1,
		OnGone: []data.Resource{{
			Name: "gone kitten", Quantity: 1,
		}, {
			Name: "kitten", Capacity: 1,
		}},
	}, {
		Name: "happiness", StartQuantity: 1.1, Capacity: -1,
		Producers: []data.Resource{{
			Name: "kitten", ProductionFactor: -0.02, ProductionFloor: true,
		}, {
			Name: "woodcutter", ProductionFactor: -0.02,
		}, {
			Name: "scholar", ProductionFactor: -0.02,
		}},
	}, {
		Name: "Catnip Field", Capacity: -1,
	}, {
		Name: "Hut", Capacity: -1,
	}, {
		Name: "Library", Capacity: -1,
	}, {
		Name: "Barn", Capacity: -1,
	}, {
		Name: "Calendar", Capacity: 1,
	}, {
		Name: "Agriculture", Capacity: 1,
	}, {
		Name: "Archery", Capacity: 1,
	}, {
		Name: "Mining", Capacity: 1,
	}})
	g.Actions = []game.Action{{
		Name:     "Gather catnip",
		LockedBy: data.Resource{Name: "Catnip Field"},
		Adds: []data.Resource{{
			Name: "catnip", Quantity: 1,
		}},
	}, {
		Name:       "Refine catnip",
		UnlockedBy: data.Resource{Name: "catnip"},
		LockedBy:   data.Resource{Name: "woodcutter"},
		Costs: []data.Resource{{
			Name: "catnip", Quantity: 100, CostExponentBase: 1,
		}},
		Adds: []data.Resource{{
			Name: "wood", Quantity: 1,
		}},
	}, {
		Name:       "Catnip Field",
		UnlockedBy: data.Resource{Name: "catnip"},
		Costs: []data.Resource{{
			Name: "catnip", Quantity: 10, CostExponentBase: 1.12,
		}},
		Adds: []data.Resource{{
			Name: "Catnip Field", Quantity: 1,
		}},
	}, {
		Name:       "Hut",
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
		Name:       "Library",
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
		Name:       "Barn",
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
		Name:       "woodcutter",
		UnlockedBy: data.Resource{Name: "Hut"},
		Costs: []data.Resource{{
			Name: "kitten", Quantity: 1, Capacity: 1, CostExponentBase: 1,
		}},
		Adds: []data.Resource{{
			Name: "woodcutter", Quantity: 1,
		}},
	}, {
		Name:       "scholar",
		UnlockedBy: data.Resource{Name: "Library"},
		Costs: []data.Resource{{
			Name: "kitten", Quantity: 1, Capacity: 1, CostExponentBase: 1,
		}},
		Adds: []data.Resource{{
			Name: "scholar", Quantity: 1,
		}},
	}, {
		Name:       "Calendar",
		UnlockedBy: data.Resource{Name: "Library"},
		LockedBy:   data.Resource{Name: "Calendar"},
		Costs: []data.Resource{{
			Name: "science", Quantity: 30, CostExponentBase: 1,
		}},
		Adds: []data.Resource{{
			Name: "Calendar", Quantity: 1,
		}},
	}, {
		Name:       "Agriculture",
		UnlockedBy: data.Resource{Name: "Calendar"},
		LockedBy:   data.Resource{Name: "Agriculture"},
		Costs: []data.Resource{{
			Name: "science", Quantity: 100, CostExponentBase: 1,
		}},
		Adds: []data.Resource{{
			Name: "Agriculture", Quantity: 1,
		}},
	}, {
		Name:       "Archery",
		UnlockedBy: data.Resource{Name: "Calendar"},
		LockedBy:   data.Resource{Name: "Archery"},
		Costs: []data.Resource{{
			Name: "science", Quantity: 300, CostExponentBase: 1,
		}},
		Adds: []data.Resource{{
			Name: "Archery", Quantity: 1,
		}},
	}, {
		Name:       "Mining",
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
		{[]int{field, sfield}, 58},
		{[]int{refine}, 5},
		{[]int{hut, swoodcutter, woodcutter}, 1},
		{[]int{slibrary, library, sscholar, scholar}, 1},
		{[]int{slibrary, library}, 14},
		{[]int{scalendar, calendar}, 1},
		{[]int{sagriculture, agriculture}, 1},
		{[]int{sarchery, archery}, 1},
		{[]int{smining, mining}, 1},
		{[]int{sbarn, barn}, 6},

		{[]int{field, sfield}, 25},
		{[]int{slibrary, library}, 15},
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
