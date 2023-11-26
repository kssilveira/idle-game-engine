package kittens

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/kssilveira/idle-game-engine/game"
)

type Input chan int
type Now func() time.Time

func Run(logger *log.Logger, separator string, input Input, now Now) {
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
	var err error
	if err = g.Validate(); err != nil {
		logger.Fatalf("%v\n", err)
	}
	for {
		logger.Printf("%s", separator)
		if err != nil {
			logger.Printf("%v\n", err)
		}
		for _, r := range g.Resources {
			if r.Quantity == 0 {
				continue
			}
			capacity := ""
			if r.Capacity > 0 {
				capacity = fmt.Sprintf("/%.0f", r.Capacity)
			}
			rateStr := ""
			rate := g.GetRate(r)
			if rate != 0 {
				capStr := ""
				if r.Capacity > 0 {
					duration := time.Duration(((r.Capacity - r.Quantity) / rate)) * time.Second
					capStr = fmt.Sprintf(", %s to cap", duration)
				}
				rateStr = fmt.Sprintf(" (%.2f/s%s)", rate, capStr)
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
		select {
		case in := <-input:
			if in == 999 {
				return
			}
			g.Update(now())
			if err = g.Act(in); err != nil {
				logger.Printf("%v\n", err)
			}
		case <-time.After(1 * time.Second):
			g.Update(now())
		}
	}
}

/*

TODO

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
