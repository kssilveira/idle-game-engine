package main

import "fmt"

type Resource struct {
	Name string
	Quantity int
	Capacity int
}

type Action struct {
	Name string
	Add []Resource
}

type Game struct {
	Resources []*Resource
	ResourceToIndex map[string]int
	Actions []Action
}

func NewGame() *Game {
	return &Game{
		ResourceToIndex: map[string]int{},
	}
}

func (g *Game) AddResource(r *Resource) {
	g.ResourceToIndex[r.Name] = len(g.Resources)
	g.Resources = append(g.Resources, r)
}

func main() {
	game := NewGame()
	game.AddResource(&Resource{
			Name: "catnip",
			Capacity: 5,
		})
	game.Actions = []Action{{
		Name: "Gather catnip",
		Add: []Resource{{
			Name: "catnip",
			Quantity: 1,
		}},
	}}
	for ;; {
		for _, r := range game.Resources {
			fmt.Printf("%s %d/%d\n", r.Name, r.Quantity, r.Capacity)
		}
		for i, a := range game.Actions {
			fmt.Printf("%d: '%s' (", i, a.Name)
			for _, r := range a.Add  {
				fmt.Printf("%s + %d", r.Name, r.Quantity)
			}
			fmt.Printf(")\n")
		}
		input := -1
		fmt.Printf("> ")
		fmt.Scanf("%d", &input)
		if input < 0 || input >= len(game.Actions) {
			fmt.Printf("invalid input\n")
			continue
		}
		for _, add := range game.Actions[input].Add {
			r := game.Resources[game.ResourceToIndex[add.Name]]
			r.Quantity += add.Quantity
			if r.Quantity > r.Capacity {
				r.Quantity = r.Capacity
			}
		}
	}
}

/*

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
