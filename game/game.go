package game

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/kssilveira/idle-game-engine/data"
	"github.com/kssilveira/idle-game-engine/ui"
)

type Game struct {
	Resources []*data.Resource
	// maps resource name to index in Resources
	ResourceToIndex map[string]int
	Actions         []data.Action
	// maps action name to index in Actions
	ActionToIndex map[string]int
	Now           time.Time
}

type Input chan string
type Output chan *ui.Data
type Now func() time.Time

func NewGame(now time.Time) *Game {
	g := &Game{
		Now:             now,
		ResourceToIndex: map[string]int{},
		ActionToIndex:   map[string]int{},
	}
	g.AddResources([]data.Resource{{
		Name: "time", Type: "Calendar", Capacity: -1,
	}, {
		Name: "skip", Type: "Calendar", Capacity: -1,
	}})
	return g
}

func (g *Game) AddResources(resources []data.Resource) {
	for _, resource := range resources {
		g.ResourceToIndex[resource.Name] = len(g.Resources)
		cp := resource
		g.Resources = append(g.Resources, &cp)
	}
}

func (g *Game) AddActions(actions []data.Action) {
	for _, action := range actions {
		g.ActionToIndex[action.Name] = len(g.Actions)
		cp := action
		g.Actions = append(g.Actions, cp)
	}
}

func (g *Game) Validate() error {
	for _, r := range g.Resources {
		if err := g.ValidateResource(r); err != nil {
			return err
		}
	}
	for _, a := range g.Actions {
		for _, list := range append([][]data.Resource{}, a.Costs, a.Adds) {
			for _, r := range list {
				if err := g.ValidateResource(&r); err != nil {
					return err
				}
			}
			for _, name := range []string{a.UnlockedBy, a.LockedBy} {
				if err := g.ValidateResourceName(name); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (g *Game) Run(now Now, input Input, output Output) {
	var in string
	var parsedInput data.ParsedInput
	var err error
	for {
		g.Update(now())
		data := &ui.Data{
			LastInput: parsedInput,
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
			parsedInput, err = g.Act(in)
		case <-time.After(1 * time.Second):
		}
	}
}

func (g *Game) PopulateUIResources(data *ui.Data) {
	for _, r := range g.Resources {
		data.Resources = append(data.Resources, ui.Resource{
			Resource:        *r,
			Rate:            g.GetRate(r),
			DurationToCap:   g.GetDuration(r, r.Capacity),
			DurationToEmpty: g.GetDuration(r, 0),
		})
	}
}

func (g *Game) PopulateUIActions(data *ui.Data) {
	for _, a := range g.Actions {
		action := ui.Action{
			Name:   a.Name,
			Type:   a.Type,
			Locked: g.IsLocked(a),
		}
		if g.HasResource(a.Name) {
			action.Quantity = g.GetResource(a.Name).Quantity
		}
		action.Costs = g.PopulateUICosts(a, false /* isNested */)
		for _, r := range a.Adds {
			action.Adds = append(action.Adds, ui.Add{
				Name:     r.Name,
				Quantity: g.GetActionAdd(r).Quantity,
				Capacity: r.Capacity,
			})
		}
		data.Actions = append(data.Actions, action)
	}
	data.CustomActions = []ui.CustomAction{{
		Name: "sX: time skip until action X is available",
	}, {
		Name: "mX: make inputs for action X",
	}}
}

func (g *Game) PopulateUICosts(a data.Action, isNested bool) []ui.Cost {
	res := []ui.Cost{}
	for _, c := range a.Costs {
		cost := g.GetCost(a, c)
		r := g.GetResource(c.Name)
		one := ui.Cost{
			Name:     c.Name,
			Quantity: r.Quantity,
			Capacity: r.Capacity,
			Cost:     cost,
			Duration: g.GetDuration(r, cost),
		}
		if isNested {
			one.Capacity = -1
		}
		if r.ProducerAction != "" {
			one.Costs = g.PopulateUICosts(g.GetNestedAction(a, c), true /* isNested */)
		}
		res = append(res, one)
	}
	return res
}

func (g *Game) GetDuration(r *data.Resource, quantity float64) time.Duration {
	return time.Duration(((quantity - r.Quantity) / g.GetRate(r))) * time.Second
}

func (g *Game) Update(now time.Time) {
	elapsed := now.Sub(g.Now)
	g.Now = now
	g.GetResource("time").Quantity += float64(elapsed / time.Second)
	for _, resource := range g.Resources {
		if resource.StartCapacity > 0 {
			resource.Capacity = resource.StartCapacity + g.GetCapacityRate(resource)
		}
		factor := g.GetRate(resource)
		if resource.ProductionModulus != 0 {
			factor = float64(int(factor) % resource.ProductionModulus)
		}
		if resource.StartQuantity != 0 {
			if resource.ProductionModulus != 0 && resource.ProductionModulusEquals >= 0 {
				if int(factor) == resource.ProductionModulusEquals {
					resource.Quantity = resource.StartQuantity
				} else {
					resource.Quantity = 0
				}
			} else {
				resource.Quantity = resource.StartQuantity + factor
			}
		} else {
			resource.Add(data.Resource{Quantity: factor * elapsed.Seconds()})
		}
		if factor < 0 && resource.Quantity == 0 {
			g.UpdateRate(resource)
		}
	}
}

func (g *Game) Act(in string) (data.ParsedInput, error) {
	input, err := g.ParseInput(in)
	if err != nil {
		return input, err
	}
	if g.IsLocked(input.Action) {
		return input, fmt.Errorf("action %s is locked", input.Action.Name)
	}
	if err := g.CheckMax(input.Action); err != nil {
		return input, err
	}
	skipTime, err := g.GetSkipTime(input.Action, input)
	if err != nil {
		return input, err
	}
	if input.IsSkip && skipTime > 0 {
		g.TimeSkip(skipTime)
		return input, nil
	}
	if input.IsMake {
		for _, c := range input.Action.Costs {
			r := g.GetResource(c.Name)
			if r.ProducerAction == "" {
				continue
			}
			nested := fmt.Sprintf("%d", g.ActionToIndex[r.ProducerAction])
			for i := 0; i < int(g.GetNeededNestedAction(input.Action, c)); i++ {
				if _, err := g.Act(nested); err != nil {
					return input, err
				}
			}
		}
		return input, nil
	}
	for _, c := range input.Action.Costs {
		r := g.GetResource(c.Name)
		r.Quantity -= g.GetCost(input.Action, c)
		r.Capacity -= c.Capacity
	}
	for _, add := range input.Action.Adds {
		r := g.GetResource(add.Name)
		r.Add(g.GetActionAdd(add))
	}
	return input, nil
}

func (g *Game) GetActionAdd(add data.Resource) data.Resource {
	bonus := 1.0
	for _, p := range add.ProductionBonus {
		bonus += g.GetOneRate(p)
	}
	add.Quantity *= bonus
	return add
}

func (g *Game) ParseInput(in string) (data.ParsedInput, error) {
	res := data.ParsedInput{}
	if strings.HasPrefix(in, "s") {
		res.IsSkip = true
		in = in[1:]
	}
	if strings.HasPrefix(in, "m") {
		res.IsMake = true
		in = in[1:]
	}
	index, err := strconv.Atoi(in)
	if err != nil {
		return res, err
	}
	if index < 0 || index >= len(g.Actions) {
		return res, fmt.Errorf("invalid index %d", index)
	}
	res.Action = g.Actions[index]
	return res, nil
}

func (g *Game) IsLocked(a data.Action) bool {
	return (a.UnlockedBy != "" && g.GetResource(a.UnlockedBy).Quantity <= 0) ||
		(a.LockedBy != "" && g.GetResource(a.LockedBy).Quantity > 0)
}

func (g *Game) CheckMax(a data.Action) error {
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

func (g *Game) GetSkipTime(a data.Action, input data.ParsedInput) (time.Duration, error) {
	var skipTime time.Duration
	for _, c := range a.Costs {
		r := g.GetResource(c.Name)
		cost := g.GetCost(a, c)
		if r.Quantity >= cost {
			continue
		}
		if input.IsSkip && g.GetRate(r) > 0 && (r.Capacity == -1 || r.Quantity < r.Capacity) {
			duration := g.GetDuration(r, cost) + time.Second
			if duration > skipTime {
				skipTime = duration
			}
			continue
		}
		if (input.IsSkip || input.IsMake) && r.ProducerAction != "" {
			nested, err := g.GetSkipTime(g.GetNestedAction(a, c), input)
			if err == nil {
				if input.IsSkip {
					duration := nested + time.Second
					if duration > skipTime {
						skipTime = duration
					}
				}
				continue
			}
		}
		return 0, fmt.Errorf("not enough %s", c.Name)
	}
	return skipTime, nil
}

func (g *Game) GetNestedAction(a data.Action, c data.Resource) data.Action {
	r := g.GetResource(c.Name)
	if r.ProducerAction == "" {
		return data.Action{}
	}
	need := g.GetNeededNestedAction(a, c)
	res := data.Action{
		Adds: []data.Resource{{}},
	}
	action := g.GetAction(r.ProducerAction)
	for _, c := range action.Costs {
		cost := g.GetCost(action, c) * need
		res.Costs = append(res.Costs, data.Resource{
			Name:     c.Name,
			Quantity: cost,
		})
	}
	return res
}

func (g *Game) GetNeededNestedAction(a data.Action, c data.Resource) float64 {
	r := g.GetResource(c.Name)
	if r.ProducerAction == "" {
		return 0
	}
	cost := g.GetCost(a, c)
	action := g.GetAction(r.ProducerAction)
	return math.Ceil(cost / g.GetActionAdd(action.Adds[0]).Quantity)
}

func (g *Game) TimeSkip(skip time.Duration) {
	g.GetResource("skip").Quantity += float64(skip / time.Second)
	now := g.Now
	g.Now = time.Time(now.Add(-skip))
	g.Update(now)
}

func (g *Game) ValidateResource(r *data.Resource) error {
	for _, name := range []string{r.Name, r.ProductionResourceFactor} {
		if err := g.ValidateResourceName(name); err != nil {
			return err
		}
	}
	for _, name := range []string{r.ProducerAction} {
		if err := g.ValidateActionName(name); err != nil {
			return err
		}
	}
	for _, list := range append(
		[][]data.Resource{}, r.Producers, r.CapacityProducers, r.ProductionBonus, r.OnGone) {
		for _, r := range list {
			if err := g.ValidateResource(&r); err != nil {
				return err
			}
		}
	}
	return nil
}

func (g *Game) ValidateResourceName(name string) error {
	if name != "" && !g.HasResource(name) {
		return fmt.Errorf("invalid resource name %s", name)
	}
	return nil
}

func (g *Game) ValidateActionName(name string) error {
	if name != "" && !g.HasAction(name) {
		return fmt.Errorf("invalid action name %s", name)
	}
	return nil
}

func (g *Game) GetRate(resource *data.Resource) float64 {
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

func (g *Game) GetCapacityRate(resource *data.Resource) float64 {
	factor := 0.0
	for _, p := range resource.CapacityProducers {
		factor += g.GetOneRate(p)
	}
	return factor
}

func (g *Game) GetOneRate(resource data.Resource) float64 {
	one := g.GetQuantityForRate(resource) * resource.ProductionFactor
	if resource.ProductionResourceFactor != "" {
		one *= g.GetQuantityForRate(*g.GetResource(resource.ProductionResourceFactor))
	}
	bonus := 1.0
	for _, p := range resource.ProductionBonus {
		bonus += g.GetOneRate(p)
	}
	return one * bonus
}

func (g *Game) GetQuantityForRate(p data.Resource) float64 {
	quantity := 1.0
	if p.Name != "" {
		quantity = g.GetResource(p.Name).Quantity
	}
	if p.ProductionFloor {
		quantity = math.Floor(quantity)
	}
	if p.ProductionBoolean {
		if quantity > 0 {
			quantity = 1
		}
	}
	return quantity
}

func (g *Game) UpdateRate(resource *data.Resource) {
	for _, p := range resource.Producers {
		if !p.ProductionOnGone {
			continue
		}
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

func (g *Game) GetResource(name string) *data.Resource {
	return g.Resources[g.ResourceToIndex[name]]
}

func (g *Game) GetAction(name string) data.Action {
	return g.Actions[g.ActionToIndex[name]]
}

func (g *Game) GetCost(a data.Action, c data.Resource) float64 {
	base := c.CostExponentBase
	if base == 0 {
		base = 1
	}
	return c.Quantity * math.Pow(base, g.GetResource(a.Adds[0].Name).Quantity)
}

func (g *Game) HasResource(name string) bool {
	_, ok := g.ResourceToIndex[name]
	return ok
}

func (g *Game) HasAction(name string) bool {
	_, ok := g.ActionToIndex[name]
	return ok
}
