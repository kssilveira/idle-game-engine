package game

import (
	"fmt"
	"time"
)

type Resource struct {
	Name           string
	Quantity       float64
	Capacity       float64
	Rate           []Resource
	Factor         float64
	ResourceFactor string
}

type Action struct {
	Name string
	Add  []Resource
}

type Game struct {
	Resources       []*Resource
	ResourceToIndex map[string]int
	Actions         []Action
	Now             time.Time
}

func NewGame(now time.Time) *Game {
	return &Game{
		Now:             now,
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

func (g *Game) Update(now time.Time) error {
	elapsed := now.Sub(g.Now)
	g.Now = now
	for _, resource := range g.Resources {
		factor, err := g.GetRate(resource)
		if err != nil {
			return err
		}
		resource.Quantity += factor * elapsed.Seconds()
	}
	return nil
}

func (g *Game) GetRate(resource *Resource) (float64, error) {
	factor := 0.0
	for _, rate := range resource.Rate {
		rateResource, err := g.GetResource(rate.Name)
		if err != nil {
			return 0, err
		}
		one := rateResource.Quantity * rate.Factor
		if rate.ResourceFactor != "" {
			resourceFactor, err := g.GetResource(rate.ResourceFactor)
			if err != nil {
				return 0, err
			}
			one *= resourceFactor.Quantity
		}
		factor += one
	}
	return factor, nil
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
