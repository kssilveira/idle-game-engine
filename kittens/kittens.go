package kittens

import "fmt"
import "github.com/kssilveira/idle-game-engine/game"

type Input func()int

func Run(input Input) {
	g := game.NewGame()
	g.AddResource(&game.Resource{
			Name: "catnip",
			Capacity: 5000,
		})
	g.Actions = []game.Action{{
		Name: "Gather catnip",
		Add: []game.Resource{{
			Name: "catnip",
			Quantity: 1,
		}},
	}}
	for ;; {
		for _, r := range g.Resources {
			fmt.Printf("%s %d/%d\n", r.Name, r.Quantity, r.Capacity)
		}
		for i, a := range g.Actions {
			fmt.Printf("%d: '%s' (", i, a.Name)
			for _, r := range a.Add  {
				fmt.Printf("%s + %d", r.Name, r.Quantity)
			}
			fmt.Printf(")\n")
		}
		if err := g.Act(input()); err != nil {
			fmt.Printf("%v\n", err)
		}
	}
}
