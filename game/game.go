package game

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/kssilveira/idle-game-engine/ui"
)

type Game struct {
	Resources []*Resource
	// maps resource name to index in Resources
	ResourceToIndex map[string]int
	Actions         []Action
	Now             time.Time
}

type Resource struct {
	Name     string
	Quantity float64
	Capacity float64

	Producers []Resource

	// quantity += producer.Quantity * ProductionFactor * elapsedTime
	ProductionFactor float64
	// quantity += producer.Quantity * ProductionFactor * elapsedTime * ProductionResourceFactor.Quantity
	ProductionResourceFactor string
	// quantity += floor(producer.Quantity) * ProductionFactor * elapsedTime
	ProductionFloor bool
	// quantity = StartQuantity + producer.Quantity * ProductionFactor
	StartQuantity float64

	// production *= 1 + bonus
	ProductionBonus []Resource

	OnGone []Resource

	// cost = Quantity * pow(CostExponentBase, add.Quantity)
	CostExponentBase float64
}

type Action struct {
	Name       string
	UnlockedBy Resource
	Costs      []Resource
	Adds       []Resource
}

type Input chan string
type Output chan *ui.Data
type Now func() time.Time

func NewGame(now time.Time) *Game {
	g := &Game{
		Now:             now,
		ResourceToIndex: map[string]int{},
	}
	g.AddResources([]Resource{{
		Name: "skip", Capacity: -1,
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
		if err := g.ValidateResource(&a.UnlockedBy); err != nil {
			return err
		}
		for _, r := range a.Costs {
			if err := g.ValidateResource(&r); err != nil {
				return err
			}
		}
		for _, r := range a.Adds {
			if err := g.ValidateResource(&r); err != nil {
				return err
			}
		}
	}
	return nil
}

func (g *Game) Run(now Now, input Input, output Output) {
	var in string
	var err error
	for {
		data := &ui.Data{
			LastInput: in,
			Error:     err,
		}
		g.PopulateUIResources(data)
		g.PopulateUIActions(data)
		output <- data
		select {
		case in = <-input:
			if in == "999" {
				close(output)
				return
			}
			g.Update(now())
			err = g.Act(in)
		case <-time.After(1 * time.Second):
			g.Update(now())
		}
	}
}

func (g *Game) PopulateUIResources(data *ui.Data) {
	for _, r := range g.Resources {
		data.Resources = append(data.Resources, ui.Resource{
			Name:            r.Name,
			Quantity:        r.Quantity,
			Capacity:        r.Capacity,
			Rate:            g.GetRate(r),
			DurationToCap:   g.GetDuration(r, r.Capacity),
			DurationToEmpty: g.GetDuration(r, 0),
			StartQuantity:   r.StartQuantity,
		})
	}
}

func (g *Game) PopulateUIActions(data *ui.Data) {
	for _, a := range g.Actions {
		action := ui.Action{
			Name:   a.Name,
			Locked: g.IsLocked(a),
		}
		for _, c := range a.Costs {
			cost := g.GetCost(a, c)
			r := g.GetResource(c.Name)
			action.Costs = append(action.Costs, ui.Cost{
				Name:     c.Name,
				Quantity: r.Quantity,
				Capacity: r.Capacity,
				Cost:     cost,
				Duration: g.GetDuration(r, cost),
			})
		}
		for _, r := range a.Adds {
			action.Adds = append(action.Adds, ui.Add{
				Name:     r.Name,
				Quantity: r.Quantity,
				Capacity: r.Capacity,
			})
		}
		data.Actions = append(data.Actions, action)
	}
	data.CustomActions = append(data.CustomActions, ui.CustomAction{
		Name: "sX: time skip until action X is available",
	})
}

func (g *Game) GetDuration(r *Resource, quantity float64) time.Duration {
	return time.Duration(((quantity - r.Quantity) / g.GetRate(r))) * time.Second
}

func (g *Game) Update(now time.Time) {
	elapsed := now.Sub(g.Now)
	g.Now = now
	for _, resource := range g.Resources {
		factor := g.GetRate(resource)
		if resource.StartQuantity != 0 {
			resource.Quantity = resource.StartQuantity + factor
		} else {
			resource.Add(Resource{Quantity: factor * elapsed.Seconds()})
		}
		if factor < 0 && resource.Quantity == 0 {
			g.UpdateRate(resource)
		}
	}
}

func (g *Game) Act(input string) error {
	skip, a, err := g.ParseInput(input)
	if err != nil {
		return err
	}
	if g.IsLocked(a) {
		return fmt.Errorf("action %s is locked", a.Name)
	}
	if err := g.CheckMax(a); err != nil {
		return err
	}
	skipTime, err := g.GetSkipTime(a, skip)
	if err != nil {
		return err
	}
	if skip && skipTime > 0 {
		g.TimeSkip(skipTime)
		return nil
	}
	for _, c := range a.Costs {
		r := g.GetResource(c.Name)
		r.Quantity -= g.GetCost(a, c)
		r.Capacity -= c.Capacity
	}
	for _, add := range a.Adds {
		r := g.GetResource(add.Name)
		r.Add(add)
	}
	return nil
}

func (g *Game) ParseInput(input string) (bool, Action, error) {
	skip := false
	if strings.HasPrefix(input, "s") {
		skip = true
		input = input[1:]
	}
	index, err := strconv.Atoi(input)
	if err != nil {
		return false, Action{}, err
	}
	if index < 0 || index >= len(g.Actions) {
		return false, Action{}, fmt.Errorf("invalid index %d", index)
	}
	return skip, g.Actions[index], nil
}

func (g *Game) IsLocked(a Action) bool {
	name := a.UnlockedBy.Name
	return name != "" && g.GetResource(a.UnlockedBy.Name).Quantity <= 0
}

func (g *Game) CheckMax(a Action) error {
	found := false
	for _, add := range a.Adds {
		r := g.GetResource(add.Name)
		if r.Quantity < r.Capacity || r.Capacity == -1 || add.Capacity > 0 {
			found = true
			break
		}
	}
	if !found {
		return fmt.Errorf("added resources already at max")
	}
	return nil
}

func (g *Game) GetSkipTime(a Action, skip bool) (time.Duration, error) {
	var skipTime time.Duration
	for _, c := range a.Costs {
		r := g.GetResource(c.Name)
		cost := g.GetCost(a, c)
		if r.Quantity < cost {
			if skip && g.GetRate(r) > 0 && (r.Capacity == -1 || r.Quantity < r.Capacity) {
				duration := g.GetDuration(r, cost) + time.Second
				if duration > skipTime {
					skipTime = duration
				}
			} else {
				return 0, fmt.Errorf("not enough %s", c.Name)
			}
		}
	}
	return skipTime, nil
}

func (g *Game) TimeSkip(skip time.Duration) {
	g.GetResource("skip").Quantity += float64(skip / time.Second)
	now := g.Now
	g.Now = time.Time(now.Add(-skip))
	g.Update(now)
}

func (g *Game) ValidateResource(r *Resource) error {
	if r.Name != "" && !g.HasResource(r.Name) {
		return fmt.Errorf("invalid resource name %s", r.Name)
	}
	if r.ProductionResourceFactor != "" && !g.HasResource(r.ProductionResourceFactor) {
		return fmt.Errorf("invalid resource name %s", r.ProductionResourceFactor)
	}
	for _, r := range r.Producers {
		if err := g.ValidateResource(&r); err != nil {
			return err
		}
	}
	for _, r := range r.OnGone {
		if err := g.ValidateResource(&r); err != nil {
			return err
		}
	}
	return nil
}

func (r *Resource) Add(add Resource) {
	r.Capacity += add.Capacity
	r.Quantity += add.Quantity
	if r.Quantity > r.Capacity && r.Capacity >= 0 {
		r.Quantity = r.Capacity
	}
	if r.Quantity < 0 {
		r.Quantity = 0
	}
}

func (g *Game) GetRate(resource *Resource) float64 {
	factor := 0.0
	for _, p := range resource.Producers {
		factor += g.GetOneRate(p)
	}
	bonus := 1.0
	for _, p := range resource.ProductionBonus {
		bonus += g.GetOneRate(p)
	}
	return factor * bonus
}

func (g *Game) GetOneRate(p Resource) float64 {
	one := g.GetQuantityForRate(p) * p.ProductionFactor
	if p.ProductionResourceFactor != "" {
		one *= g.GetQuantityForRate(*g.GetResource(p.ProductionResourceFactor))
	}
	return one
}

func (g *Game) GetQuantityForRate(p Resource) float64 {
	quantity := 1.0
	if p.Name != "" {
		quantity = g.GetResource(p.Name).Quantity
	}
	if p.ProductionFloor {
		quantity = math.Floor(quantity)
	}
	return quantity
}

func (g *Game) UpdateRate(resource *Resource) {
	for _, p := range resource.Producers {
		one := g.GetOneRate(p)
		if one < 0 {
			r := g.GetResource(p.Name)
			r.Quantity--
			for _, onGone := range r.OnGone {
				gone := g.GetResource(onGone.Name)
				gone.Quantity += onGone.Quantity
				gone.Capacity += onGone.Capacity
			}
			return
		}
	}
}

func (g *Game) GetResource(name string) *Resource {
	return g.Resources[g.ResourceToIndex[name]]
}

func (g *Game) GetCost(a Action, c Resource) float64 {
	return c.Quantity * math.Pow(c.CostExponentBase, g.GetResource(a.Adds[0].Name).Quantity)
}

func (g *Game) HasResource(name string) bool {
	_, ok := g.ResourceToIndex[name]
	return ok
}
