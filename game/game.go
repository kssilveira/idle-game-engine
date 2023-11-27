package game

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
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

type Input chan string
type Now func() time.Time

func NewGame(now time.Time) *Game {
	g := &Game{
		Now:             now,
		ResourceToIndex: map[string]int{},
	}
	g.AddResources([]Resource{{
		Name: "skip",
	}})
	return g
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

func (g *Game) Run(logger *log.Logger, separator string, input Input, now Now) {
	var in string
	var err error
	for {
		logger.Printf("%s", separator)
		g.ShowResources(logger)
		g.ShowActions(logger)
		logger.Printf("last input: %s\n", in)
		if err != nil {
			logger.Printf("error: %v\n", err)
		}
		select {
		case in = <-input:
			if in == "999" {
				return
			}
			g.Update(now())
			err = g.Act(in)
		case <-time.After(1 * time.Second):
			g.Update(now())
		}
	}
}

func (g *Game) ShowResources(logger *log.Logger) {
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
				capStr = fmt.Sprintf(", %s to cap", g.GetDuration(r, r.Capacity))
			}
			rateStr = fmt.Sprintf(" (%.2f/s%s)", rate, capStr)
		}
		logger.Printf("%s %.2f%s%s\n", r.Name, r.Quantity, capacity, rateStr)
	}
}

func (g *Game) ShowActions(logger *log.Logger) {
	for i, a := range g.Actions {
		parts := []string{
			fmt.Sprintf("%d: '%s' (", i, a.Name),
		}
		for _, c := range a.Costs {
			cost := g.GetCost(a, c)
			r := g.GetResource(c.Name)
			out := fmt.Sprintf("%.2f/%.2f %s", r.Quantity, cost, g.GetDuration(r, cost))
			if r.Quantity >= cost {
				out = fmt.Sprintf("%.2f", cost)
			}
			parts = append(parts, fmt.Sprintf("%s %s", c.Name, out))
		}
		parts = append(parts, ") (")
		for _, r := range a.Adds {
			parts = append(parts, fmt.Sprintf("%s + %.0f", r.Name, r.Quantity))
		}
		logger.Printf("%s)\n", strings.Join(parts, ""))
	}
	logger.Printf("sX: time skip until action X is available\n")
}

func (g *Game) GetDuration(r *Resource, quantity float64) time.Duration {
	return time.Duration(((quantity - r.Quantity) / g.GetRate(r))) * time.Second
}

func (g *Game) Update(now time.Time) {
	elapsed := now.Sub(g.Now)
	g.Now = now
	for _, resource := range g.Resources {
		factor := g.GetRate(resource)
		resource.AddQuantity(factor * elapsed.Seconds())
	}
}

func (g *Game) Act(input string) error {
	skip := false
	var skipTime time.Duration
	if strings.HasPrefix(input, "s") {
		skip = true
		input = input[1:]
	}
	index, err := strconv.Atoi(input)
	if err != nil {
		return err
	}
	if index < 0 || index >= len(g.Actions) {
		return fmt.Errorf("invalid index %d", index)
	}
	a := g.Actions[index]
	for _, c := range a.Costs {
		r := g.GetResource(c.Name)
		cost := g.GetCost(a, c)
		if r.Quantity < cost {
			if skip && g.GetRate(r) > 0 && r.Quantity < r.Capacity {
				duration := g.GetDuration(r, cost) + time.Second
				if duration > skipTime {
					skipTime = duration
				}
			} else {
				return fmt.Errorf("resource %s not enough", c.Name)
			}
		}
	}
	if skip && skipTime > 0 {
		g.TimeSkip(skipTime)
		return nil
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

func (g *Game) TimeSkip(skip time.Duration) {
	g.GetResource("skip").Quantity += float64(skip / time.Second)
	now := g.Now
	g.Now = time.Time(now.Add(-skip))
	g.Update(now)
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
