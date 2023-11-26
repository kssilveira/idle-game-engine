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

func (r *Resource) AddQuantity(add float64) {
	r.Quantity += add
	if r.Quantity > r.Capacity && r.Capacity > 0 {
		r.Quantity = r.Capacity
	}
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

func (g *Game) Validate() error {
	for _, r := range g.Resources {
		if err := g.ValidateResource(r); err != nil {
			return err
		}
	}
	for _, a := range g.Actions {
		for _, r := range a.Add {
			if err := g.ValidateResource(&r); err != nil {
				return err
			}
		}
	}
	return nil
}

func (g *Game) ValidateResource(r *Resource) error {
	if _, ok := g.ResourceToIndex[r.Name]; !ok {
		return fmt.Errorf("invalid resource name %s", r.Name)
	}
	if _, ok := g.ResourceToIndex[r.ResourceFactor]; !ok && r.ResourceFactor != "" {
		return fmt.Errorf("invalid resource name %s", r.ResourceFactor)
	}
	for _, r := range r.Rate {
		if err := g.ValidateResource(&r); err != nil {
			return err
		}
	}
	return nil
}

func (g *Game) GetResource(name string) *Resource {
	return g.Resources[g.ResourceToIndex[name]]
}

func (g *Game) Update(now time.Time) {
	elapsed := now.Sub(g.Now)
	g.Now = now
	for _, resource := range g.Resources {
		factor := g.GetRate(resource)
		resource.AddQuantity(factor * elapsed.Seconds())
	}
}

func (g *Game) GetRate(resource *Resource) float64 {
	factor := 0.0
	for _, rate := range resource.Rate {
		one := g.GetResource(rate.Name).Quantity * rate.Factor
		if rate.ResourceFactor != "" {
			one *= g.GetResource(rate.ResourceFactor).Quantity
		}
		factor += one
	}
	return factor
}

func (g *Game) Act(index int) error {
	if index < 0 || index >= len(g.Actions) {
		return fmt.Errorf("invalid index %d", index)
	}
	for _, add := range g.Actions[index].Add {
		r := g.GetResource(add.Name)
		r.AddQuantity(add.Quantity)
	}
	return nil
}
