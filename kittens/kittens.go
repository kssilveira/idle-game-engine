package kittens

import (
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
			Name: "woodcutter", ProductionFactor: -4.25,
		}},
	}, {
		Name: "wood", Capacity: 200,
		Producers: []game.Resource{{
			Name: "woodcutter", ProductionFactor: 0.09,
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
		Name: "Catnip Field", Capacity: -1,
	}, {
		Name: "Hut", Capacity: -1,
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
		Name: "Refine catnip",
		Costs: []game.Resource{{
			Name: "catnip", Quantity: 100, CostExponentBase: 1,
		}},
		Adds: []game.Resource{{
			Name: "wood", Quantity: 1,
		}},
	}, {
		Name: "Catnip Field",
		Costs: []game.Resource{{
			Name: "catnip", Quantity: 10, CostExponentBase: 1.12,
		}},
		Adds: []game.Resource{{
			Name: "Catnip Field", Quantity: 1,
		}},
	}, {
		Name: "Hut",
		Costs: []game.Resource{{
			Name: "wood", Quantity: 5, CostExponentBase: 2.5,
		}},
		Adds: []game.Resource{{
			Name: "Hut", Quantity: 1,
		}, {
			Name: "kitten", Capacity: 2,
		}},
	}, {
		Name: "woodcutter",
		Costs: []game.Resource{{
			Name: "kitten", Quantity: 1, Capacity: 1, CostExponentBase: 1,
		}},
		Adds: []game.Resource{{
			Name: "woodcutter", Quantity: 1,
		}},
	}}
	return g
}

/*

TODO

actions

buildings

resources

jobs

woodcutter reduced by happiness

happiness
- kittens 5 => 100% 98% 96% 94% 92% 90%

maybe

- sell building for half price

*/
