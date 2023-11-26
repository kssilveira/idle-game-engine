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

func (g *Game) AddResources(resources []Resource) {
	for _, resource := range resources {
		g.ResourceToIndex[resource.Name] = len(g.Resources)
		cp := resource
		g.Resources = append(g.Resources, &cp)
	}
}

func (g *Game) GetResource(name string) (*Resource, error) {
	index, ok := g.ResourceToIndex[name]
	if !ok {
		return nil, fmt.Errorf("invalid resource name %s", name)
	}
	return g.Resources[index], nil
}

func (g *Game) Act(index int) error {
	if index < 0 || index >= len(g.Actions) {
		return fmt.Errorf("invalid index %d", index)
	}
	for _, add := range g.Actions[index].Add {
		r, err := g.GetResource(add.Name)
		if err != nil {
			return err
		}
		r.Quantity += add.Quantity
		if r.Capacity > 0 && r.Quantity > r.Capacity {
			r.Quantity = r.Capacity
		}
	}
	return nil
}
