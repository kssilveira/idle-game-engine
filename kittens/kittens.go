package kittens

import "fmt"
import "log"
import "strings"
import "github.com/kssilveira/idle-game-engine/game"

type Input func() int

func Run(logger *log.Logger, input Input) {
	g := game.NewGame()
	g.AddResource(&game.Resource{
		Name:     "catnip",
		Capacity: 5000,
	})
	g.Actions = []game.Action{{
		Name: "Gather catnip",
		Add: []game.Resource{{
			Name:     "catnip",
			Quantity: 1,
		}},
	}}
	for {
		for _, r := range g.Resources {
			logger.Printf("%s %d/%d\n", r.Name, r.Quantity, r.Capacity)
		}
		for i, a := range g.Actions {
			parts := []string{
				fmt.Sprintf("%d: '%s' (", i, a.Name),
			}
			for _, r := range a.Add {
				parts = append(parts, fmt.Sprintf("%s + %d", r.Name, r.Quantity))
			}
			logger.Printf("%s)\n", strings.Join(parts, ""))
		}
		in := input()
		if in == 999 {
			break
		}
		if err := g.Act(in); err != nil {
			logger.Printf("%v\n", err)
		}
	}
}
