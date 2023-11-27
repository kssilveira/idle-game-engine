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
		}},
	}, {
		Name: "Catnip Field",
	}, {
		Name: "Spring", Quantity: 1,
	}, {
		Name: "Summer",
	}, {
		Name: "Autumn",
	}, {
		Name: "Winter",
	}})
	g.Actions = []game.Action{{
		Name: "Gather catnip",
		Adds: []game.Resource{{
			Name: "catnip", Quantity: 1,
		}},
	}, {
		Name: "Catnip Field",
		Costs: []game.Resource{{
			Name: "catnip", Quantity: 10, CostExponentBase: 1.12,
		}},
		Adds: []game.Resource{{
			Name: "Catnip Field", Quantity: 1,
		}},
	}}
	return g
}

/*

TODO

- move error to bottom
- show last action
- time skip action
- auto act with time skip until the end

actions

- "Refine catnip" cost 100 catnip add 1 wood

buildings

- "Hut" cost 5 wood add 2 kitten
  - cost 10 12.5 31.25 78.13 195.31 488.28

resources

- wood cap 200
- kitten
  - eat 4.25 catnip/s

jobs

woodcuttter 0.09 wood/s
  - reduced by happiness

happiness
- kittens 5 => 100% 98% 96% 94% 92% 90%

maybe

- sell building for half price

*/
