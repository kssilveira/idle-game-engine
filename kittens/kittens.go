package kittens

import (
	"time"

	"github.com/kssilveira/idle-game-engine/game"
)

func NewGame(now game.Now) *game.Game {
	g := game.NewGame(now())
	g.AddResources([]game.Resource{{
		Name: "catnip", Capacity: 5000,
		Producers: []game.Resource{{
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
			Name: "woodcutter", ProductionFactor: -4.25, ProductionFloor: true,
		}, {
			Name: "scholar", ProductionFactor: -4.25, ProductionFloor: true,
		}},
	}, {
		Name: "wood", Capacity: 200,
		Producers: []game.Resource{{
			Name: "woodcutter", ProductionFactor: 0.09, ProductionResourceFactor: "happiness",
		}},
	}, {
		Name: "science", Capacity: 250,
		Producers: []game.Resource{{
			Name: "scholar", ProductionFactor: 0.175, ProductionResourceFactor: "happiness",
		}},
		ProductionBonus: []game.Resource{{
			Name: "Library", ProductionFactor: 0.1,
		}},
	}, {
		Name: "kitten", Capacity: 0,
		Producers: []game.Resource{{
			Name: "", ProductionFactor: 0.05,
		}},
		OnGone: []game.Resource{{
			Name: "gone kitten", Quantity: 1,
		}},
	}, {
		Name: "gone kitten", Capacity: -1,
	}, {
		Name: "woodcutter", Capacity: -1,
		OnGone: []game.Resource{{
			Name: "gone kitten", Quantity: 1,
		}, {
			Name: "kitten", Capacity: 1,
		}},
	}, {
		Name: "scholar", Capacity: -1,
		OnGone: []game.Resource{{
			Name: "gone kitten", Quantity: 1,
		}, {
			Name: "kitten", Capacity: 1,
		}},
	}, {
		Name: "happiness", StartQuantity: 1.1, Capacity: -1,
		Producers: []game.Resource{{
			Name: "kitten", ProductionFactor: -0.02, ProductionFloor: true,
		}, {
			Name: "woodcutter", ProductionFactor: -0.02, ProductionFloor: true,
		}, {
			Name: "scholar", ProductionFactor: -0.02, ProductionFloor: true,
		}},
	}, {
		Name: "Catnip Field", Capacity: -1,
	}, {
		Name: "Hut", Capacity: -1,
	}, {
		Name: "Library", Capacity: -1,
	}, {
		Name: "Spring", Quantity: 1, Capacity: -1,
	}, {
		Name: "Summer", Capacity: -1,
	}, {
		Name: "Autumn", Capacity: -1,
	}, {
		Name: "Winter", Capacity: -1,
	}})
	g.Actions = []game.Action{{
		Name: "Gather catnip",
		Adds: []game.Resource{{
			Name: "catnip", Quantity: 1,
		}},
	}, {
		Name:       "Refine catnip",
		UnlockedBy: game.Resource{Name: "catnip"},
		Costs: []game.Resource{{
			Name: "catnip", Quantity: 100, CostExponentBase: 1,
		}},
		Adds: []game.Resource{{
			Name: "wood", Quantity: 1,
		}},
	}, {
		Name:       "Catnip Field",
		UnlockedBy: game.Resource{Name: "catnip"},
		Costs: []game.Resource{{
			Name: "catnip", Quantity: 10, CostExponentBase: 1.12,
		}},
		Adds: []game.Resource{{
			Name: "Catnip Field", Quantity: 1,
		}},
	}, {
		Name:       "Hut",
		UnlockedBy: game.Resource{Name: "wood"},
		Costs: []game.Resource{{
			Name: "wood", Quantity: 5, CostExponentBase: 2.5,
		}},
		Adds: []game.Resource{{
			Name: "Hut", Quantity: 1,
		}, {
			Name: "kitten", Capacity: 2,
		}},
	}, {
		Name:       "Library",
		UnlockedBy: game.Resource{Name: "wood"},
		Costs: []game.Resource{{
			Name: "wood", Quantity: 25, CostExponentBase: 1.15,
		}},
		Adds: []game.Resource{{
			Name: "Library", Quantity: 1,
		}, {
			Name: "science", Capacity: 250,
		}},
	}, {
		Name:       "woodcutter",
		UnlockedBy: game.Resource{Name: "Hut"},
		Costs: []game.Resource{{
			Name: "kitten", Quantity: 1, Capacity: 1, CostExponentBase: 1,
		}},
		Adds: []game.Resource{{
			Name: "woodcutter", Quantity: 1,
		}},
	}, {
		Name:       "scholar",
		UnlockedBy: game.Resource{Name: "Library"},
		Costs: []game.Resource{{
			Name: "kitten", Quantity: 1, Capacity: 1, CostExponentBase: 1,
		}},
		Adds: []game.Resource{{
			Name: "scholar", Quantity: 1,
		}},
	}}
	return g
}

const (
	gather      = "0"
	refine      = "1"
	field       = "2"
	sfield      = "s2"
	hut         = "3"
	shut        = "s3"
	library     = "4"
	slibrary    = "s4"
	woodcutter  = "5"
	swoodcutter = "s5"
	scholar     = "6"
	sscholar    = "s6"
)

func Solve(input chan string, sleepMS int) {
	for _, one := range []struct {
		cmds  []string
		count int
	}{
		// gather 10 catnip
		{[]string{gather}, 10},
		// buy 55 catnip field
		{[]string{field, sfield}, 55},
		// refine 5 wood
		{[]string{refine}, 5},
		// buy 4 hut, assign 8 woodcutter
		{[]string{
			hut,
			swoodcutter, woodcutter,
			swoodcutter, woodcutter,
			shut}, 4},
		// buy 1 hut
		{[]string{hut, shut}, 1},
		// buy 1 library
		{[]string{library, slibrary}, 1},
		// assign 2 scholar
		{[]string{scholar}, 2},
		// buy 14 library
		{[]string{library, slibrary}, 13},
		// done
		{[]string{"done"}, 1},
	} {
		for i := 0; i < one.count; i++ {
			for _, cmd := range one.cmds {
				input <- cmd
				time.Sleep(time.Second * time.Duration(sleepMS) / 1000.)
			}
		}
	}
}
