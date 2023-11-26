package kittens

import "fmt"
import "log"
import "strings"
import "time"
import "github.com/kssilveira/idle-game-engine/game"

type Input func() int
type Now func() time.Time

func Run(logger *log.Logger, input Input, now Now) {
	g := game.NewGame(now())
	g.AddResources([]game.Resource{{
		Name:     "catnip",
		Capacity: 5000,
		Rate: []game.Resource{{
			Name:           "Catnip Field",
			Factor:         0.63 * (1 + 0.50),
			ResourceFactor: "Spring",
		}, {
			Name:           "Catnip Field",
			Factor:         0.63,
			ResourceFactor: "Summer",
		}, {
			Name:           "Catnip Field",
			Factor:         0.63,
			ResourceFactor: "Autumn",
		}, {
			Name:           "Catnip Field",
			Factor:         0.63 * (1 - 0.75),
			ResourceFactor: "Winter",
		}},
	}, {
		Name: "Catnip Field",
	}, {
		Name:     "Spring",
		Quantity: 1,
	}, {
		Name: "Summer",
	}, {
		Name: "Autumn",
	}, {
		Name: "Winter",
	}})
	g.Actions = []game.Action{{
		Name: "Gather catnip",
		Add: []game.Resource{{
			Name:     "catnip",
			Quantity: 1,
		}},
	}, {
		Name: "Catnip Field",
		Add: []game.Resource{{
			Name:     "Catnip Field",
			Quantity: 1,
		}},
	}}
	for {
		for _, r := range g.Resources {
			if r.Quantity == 0 {
				continue
			}
			capacity := ""
			if r.Capacity > 0 {
				capacity = fmt.Sprintf("/%.0f", r.Capacity)
			}
			rateStr := ""
			rate, err := g.GetRate(r)
			if err != nil {
				logger.Printf("%v\n", err)
			}
			if rate != 0 {
				rateStr = fmt.Sprintf(" (%.2f/s)", rate)
			}
			logger.Printf("%s %.2f%s%s\n", r.Name, r.Quantity, capacity, rateStr)
		}
		for i, a := range g.Actions {
			parts := []string{
				fmt.Sprintf("%d: '%s' (", i, a.Name),
			}
			for _, r := range a.Add {
				parts = append(parts, fmt.Sprintf("%s + %.0f", r.Name, r.Quantity))
			}
			logger.Printf("%s)\n", strings.Join(parts, ""))
		}
		in := input()
		if in == 999 {
			break
		}
		if err := g.Update(now()); err != nil {
			logger.Printf("%v\n", err)
		}
		if err := g.Act(in); err != nil {
			logger.Printf("%v\n", err)
		}
	}
}

/*

TODO

- fix import syntax
- validate Game to simplify error handling

actions

- "Gather catnip" add 1 catnip
- "Refine catnip" cost 100 catnip add 1 wood

buildings

- count X resource X cost X time Xh (over cap)
  - sell for half price
- "Catnip Field" cost 10 catnip add 0.63 catnip/s
	- cost 10 11.20 12.54 14.05 15.74 17.62
	- add
		- winter -75% 0.16/s
		- spring +50% 0.94/s

- "Hut" cost 5 wood add 2 kitten
  - cost 10 12.5 31.25 78.13 195.31 488.28

resources

- count X cap X rate X/s to cap Xh
- catnip cap 5000
- wood cap 200
- kitten cap 0
  - eat 4.25 catnip/s

jobs

woodcuttter 0.09 wood/s
  - reduced by happiness

happiness
- kittens 5 => 100% 98% 96% 94% 92% 90%

*/
