package game

import (
	"fmt"
	"math"
	"time"
)

type Game struct {
	Resources       []*Resource
	ResourceToIndex map[string]int
	Actions         []Action
	Now             time.Time
}

type Resource struct {
	Name                     string
	Quantity                 float64
	Capacity                 float64
	Producers                []Resource
	ProductionFactor         float64
	ProductionResourceFactor string
	CostExponentBase         float64
}

type Action struct {
	Name  string
	Costs []Resource
	Adds  []Resource
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
		for _, r := range a.Adds {
			if err := g.ValidateResource(&r); err != nil {
				return err
			}
		}
	}
	return nil
}

func (g *Game) Update(now time.Time) {
	elapsed := now.Sub(g.Now)
	g.Now = now
	for _, resource := range g.Resources {
		factor := g.GetRate(resource)
		resource.AddQuantity(factor * elapsed.Seconds())
	}
}

func (g *Game) Act(index int) error {
	if index < 0 || index >= len(g.Actions) {
		return fmt.Errorf("invalid index %d", index)
	}
	a := g.Actions[index]
	for _, c := range a.Costs {
		r := g.GetResource(c.Name)
		if r.Quantity < g.GetCost(a, c) {
			return fmt.Errorf("resource %s not enough", c.Name)
		}
	}
	for _, c := range a.Costs {
		r := g.GetResource(c.Name)
		r.Quantity -= g.GetCost(a, c)
	}
	for _, add := range a.Adds {
		r := g.GetResource(add.Name)
		r.AddQuantity(add.Quantity)
	}
	return nil
}

func (g *Game) ValidateResource(r *Resource) error {
	if _, ok := g.ResourceToIndex[r.Name]; !ok {
		return fmt.Errorf("invalid resource name %s", r.Name)
	}
	if _, ok := g.ResourceToIndex[r.ProductionResourceFactor]; !ok && r.ProductionResourceFactor != "" {
		return fmt.Errorf("invalid resource name %s", r.ProductionResourceFactor)
	}
	for _, r := range r.Producers {
		if err := g.ValidateResource(&r); err != nil {
			return err
		}
	}
	return nil
}

func (r *Resource) AddQuantity(add float64) {
	r.Quantity += add
	if r.Quantity > r.Capacity && r.Capacity > 0 {
		r.Quantity = r.Capacity
	}
}

func (g *Game) GetRate(resource *Resource) float64 {
	factor := 0.0
	for _, p := range resource.Producers {
		one := g.GetResource(p.Name).Quantity * p.ProductionFactor
		if p.ProductionResourceFactor != "" {
			one *= g.GetResource(p.ProductionResourceFactor).Quantity
		}
		factor += one
	}
	return factor
}

func (g *Game) GetResource(name string) *Resource {
	return g.Resources[g.ResourceToIndex[name]]
}

func (g *Game) GetCost(a Action, c Resource) float64 {
	return c.Quantity * math.Pow(c.CostExponentBase, g.GetResource(a.Adds[0].Name).Quantity)
}
