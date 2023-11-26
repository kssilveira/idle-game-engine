package game

import "fmt"

type Resource struct {
	Name     string
	Quantity int
	Capacity int
}

type Action struct {
	Name string
	Add  []Resource
}

type Game struct {
	Resources       []*Resource
	ResourceToIndex map[string]int
	Actions         []Action
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

func (g *Game) Act(index int) error {
	if index < 0 || index >= len(g.Actions) {
		return fmt.Errorf("invalid index %d", index)
	}
	for _, add := range g.Actions[index].Add {
		r := g.Resources[g.ResourceToIndex[add.Name]]
		r.Quantity += add.Quantity
		if r.Quantity > r.Capacity {
			r.Quantity = r.Capacity
		}
	}
	return nil
}
