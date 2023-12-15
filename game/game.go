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
	Resources []*data.Resource `json:",omitempty"`
	Actions   []data.Action    `json:",omitempty"`

	// maps resource name to index in Resources
	resourceToIndex map[string]int
	// maps action name to index in Actions
	actionToIndex map[string]int
	now           time.Time
	errors        []error
}

type Input chan string
type Output chan *ui.Data
type Now func() time.Time

func NewGame(now time.Time) *Game {
	g := &Game{
		now:             now,
		resourceToIndex: map[string]int{},
		actionToIndex:   map[string]int{},
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
		g.AddResource(resource)
	}
}

func (g *Game) AddResource(resource data.Resource) {
	if _, ok := g.resourceToIndex[resource.Name]; ok {
		g.errors = append(g.errors, fmt.Errorf("duplicate resource %s", resource.Name))
	}
	g.resourceToIndex[resource.Name] = len(g.Resources)
	cp := resource
	g.Resources = append(g.Resources, &cp)
}

func (g *Game) AddActions(actions []data.Action) {
	for _, action := range actions {
		g.AddAction(action)
	}
}

func (g *Game) AddAction(action data.Action) {
	if _, ok := g.actionToIndex[action.Name]; ok {
		g.errors = append(g.errors, fmt.Errorf("duplicate action %s", action.Name))
	}
	g.actionToIndex[action.Name] = len(g.Actions)
	cp := action
	g.Actions = append(g.Actions, cp)
}

func (g *Game) Validate() error {
	if len(g.errors) > 0 {
		return fmt.Errorf("%v", g.errors)
	}
	for _, r := range g.Resources {
		if err := g.validateResource(r); err != nil {
			return err
		}
	}
	for _, a := range g.Actions {
		for _, list := range append([][]data.Resource{}, a.Costs, a.Adds) {
			for _, r := range list {
				if err := g.validateResource(&r); err != nil {
					return err
				}
			}
			for _, name := range []string{a.UnlockedBy, a.LockedBy} {
				if err := g.validateResourceName(name); err != nil {
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
		g.update(now())
		data := &ui.Data{
			LastInput: parsedInput,
			Error:     err,
		}
		g.populateUIResources(data)
		g.populateUIActions(data)
		output <- data
		select {
		case in = <-input:
			if in == "999" {
				close(output)
				return
			}
			parsedInput, err = g.act(in)
		case <-time.After(1 * time.Second):
		}
	}
}

func (g *Game) populateUIResources(data *ui.Data) {
	for _, r := range g.Resources {
		data.Resources = append(data.Resources, ui.Resource{
			Resource:        *r,
			Rate:            g.getRate(r),
			DurationToCap:   g.getDuration(r, r.Capacity),
			DurationToEmpty: g.getDuration(r, 0),
		})
	}
}

func (g *Game) populateUIActions(data *ui.Data) {
	for _, a := range g.Actions {
		action := ui.Action{
			Name:   a.Name,
			Type:   a.Type,
			Locked: g.isLocked(a),
		}
		if g.HasResource(a.Name) {
			action.Quantity = g.GetResource(a.Name).Quantity
		}
		action.Costs = g.populateUICosts(a, false /* isNested */)
		for _, r := range a.Adds {
			action.Adds = append(action.Adds, ui.Add{
				Name:     r.Name,
				Quantity: g.getActionAdd(r).Quantity,
				Capacity: r.Capacity,
			})
		}
		data.Actions = append(data.Actions, action)
	}
	data.CustomActions = []ui.CustomAction{{
		Name: "sX: time skip, create inputs and buy action X",
	}, {
		Name: "cX: create inputs and buy action X",
	}, {
		Name: "mX: max action X (skip, create, buy)",
	}}
}

func (g *Game) populateUICosts(a data.Action, isNested bool) []ui.Cost {
	res := []ui.Cost{}
	for _, c := range a.Costs {
		cost := g.getCost(a, c)
		r := g.GetResource(c.Name)
		one := ui.Cost{
			Name:     c.Name,
			Quantity: r.Quantity,
			Capacity: r.Capacity,
			Cost:     cost,
			Duration: g.getDuration(r, cost),
		}
		if isNested {
			one.Capacity = -1
		}
		if r.ProducerAction != "" {
			one.Costs = g.populateUICosts(g.getNestedAction(a, c), true /* isNested */)
		}
		res = append(res, one)
	}
	return res
}

func (g *Game) getDuration(r *data.Resource, quantity float64) time.Duration {
	return time.Duration(((quantity - r.Quantity) / g.getRate(r))) * time.Second
}

func (g *Game) update(now time.Time) {
	elapsed := now.Sub(g.now)
	g.now = now
	g.GetResource("time").Quantity += float64(elapsed / time.Second)
	for _, resource := range g.Resources {
		if resource.StartCapacity > 0 {
			resource.Capacity = resource.StartCapacity + g.getCapacityRate(resource)
		}
		factor := g.getRate(resource)
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
			g.updateRate(resource)
		}
	}
}

func (g *Game) act(in string) (data.ParsedInput, error) {
	input, err := g.parseInput(in)
	if err != nil {
		return input, err
	}
	if g.isLocked(input.Action) {
		return input, fmt.Errorf("action %s is locked", input.Action.Name)
	}
	if err := g.checkMax(input.Action); err != nil {
		return input, err
	}
	if input.IsSkip {
		skipTime, err := g.getSkipTime(input.Action)
		if err != nil {
			return input, err
		}
		if skipTime > 0 {
			g.timeSkip(skipTime)
		}
	}
	if input.IsSkip || input.IsCreate {
		for _, c := range input.Action.Costs {
			r := g.GetResource(c.Name)
			if r.ProducerAction == "" {
				continue
			}
			nested := fmt.Sprintf("%d", g.actionToIndex[r.ProducerAction])
			need := int(g.getNeededNestedAction(input.Action, c))
			for i := 0; i < need; i++ {
				if _, err := g.act(nested); err != nil {
					break
				}
			}
		}
	}
	if input.IsMax {
		prefixes := []string{"s", "s", "c", ""}
		for {
			errors := 0
			for _, prefix := range prefixes {
				if _, err := g.act(fmt.Sprintf("%s%d", prefix, input.Index)); err != nil {
					errors++
				}
			}
			if errors > 3 || g.isLocked(input.Action) {
				break
			}
		}
		return input, nil
	}
	for _, c := range input.Action.Costs {
		if g.GetResource(c.Name).Quantity < g.getCost(input.Action, c) {
			if input.IsSkip {
				return input, nil
			}
			return input, fmt.Errorf("not enough %s", c.Name)
		}
	}
	for _, c := range input.Action.Costs {
		r := g.GetResource(c.Name)
		r.Quantity -= g.getCost(input.Action, c)
		r.Capacity -= c.Capacity
	}
	for _, add := range input.Action.Adds {
		r := g.GetResource(add.Name)
		r.Add(g.getActionAdd(add))
	}
	return input, nil
}

func (g *Game) getActionAdd(add data.Resource) data.Resource {
	bonus := 1.0
	for _, p := range add.Bonus {
		bonus += g.getOneRate(p)
	}
	add.Quantity *= bonus
	return add
}

func (g *Game) parseInput(in string) (data.ParsedInput, error) {
	res := data.ParsedInput{}
	if strings.HasPrefix(in, "s") {
		res.IsSkip = true
		in = in[1:]
	}
	if strings.HasPrefix(in, "c") {
		res.IsCreate = true
		in = in[1:]
	}
	if strings.HasPrefix(in, "m") {
		res.IsMax = true
		in = in[1:]
	}
	index, err := strconv.Atoi(in)
	if err != nil {
		return res, err
	}
	if index < 0 || index >= len(g.Actions) {
		return res, fmt.Errorf("invalid index %d", index)
	}
	res.Index = index
	res.Action = g.Actions[index]
	return res, nil
}

func (g *Game) isLocked(a data.Action) bool {
	return (a.UnlockedBy != "" && g.GetResource(a.UnlockedBy).Quantity <= 0) ||
		(a.LockedBy != "" && g.GetResource(a.LockedBy).Quantity > 0)
}

func (g *Game) checkMax(a data.Action) error {
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

func (g *Game) getSkipTime(a data.Action) (time.Duration, error) {
	var skipTime time.Duration
	for _, c := range a.Costs {
		r := g.GetResource(c.Name)
		cost := g.getCost(a, c)
		if r.Quantity >= cost {
			continue
		}
		if g.getRate(r) > 0 && (r.Capacity == -1 || r.Quantity < r.Capacity) {
			duration := g.getDuration(r, cost) + time.Second
			if duration > skipTime {
				skipTime = duration
			}
			continue
		}
		if r.ProducerAction != "" {
			duration, err := g.getSkipTime(g.getNestedAction(a, c))
			if err == nil {
				if duration > skipTime {
					skipTime = duration
				}
				continue
			}
		}
		return 0, fmt.Errorf("not enough %s", c.Name)
	}
	return skipTime, nil
}

func (g *Game) getNestedAction(a data.Action, c data.Resource) data.Action {
	r := g.GetResource(c.Name)
	if r.ProducerAction == "" {
		return data.Action{}
	}
	need := g.getNeededNestedAction(a, c)
	res := data.Action{
		Adds: []data.Resource{{}},
	}
	action := g.GetAction(r.ProducerAction)
	if g.isLocked(action) {
		return data.Action{}
	}
	for _, c := range action.Costs {
		cost := g.getCost(action, c) * need
		res.Costs = append(res.Costs, data.Resource{
			Name:     c.Name,
			Quantity: cost,
		})
	}
	return res
}

func (g *Game) getNeededNestedAction(a data.Action, c data.Resource) float64 {
	r := g.GetResource(c.Name)
	if r.ProducerAction == "" {
		return 0
	}
	cost := g.getCost(a, c) - r.Quantity
	if cost < 0 {
		return 0
	}
	action := g.GetAction(r.ProducerAction)
	res := math.Ceil(cost / g.getActionAdd(action.Adds[0]).Quantity)
	return res
}

func (g *Game) timeSkip(skip time.Duration) {
	g.GetResource("skip").Quantity += float64(skip / time.Second)
	now := g.now
	g.now = time.Time(now.Add(-skip))
	g.update(now)
}

func (g *Game) validateResource(r *data.Resource) error {
	for _, name := range []string{r.Name} {
		if err := g.validateResourceName(name); err != nil {
			return err
		}
	}
	for _, name := range []string{r.ProducerAction} {
		if err := g.validateActionName(name); err != nil {
			return err
		}
	}
	for _, list := range append(
		[][]data.Resource{}, r.Producers, r.CapacityProducers, r.Bonus, r.OnGone) {
		for _, r := range list {
			if err := g.validateResource(&r); err != nil {
				return err
			}
		}
	}
	if r.StartQuantity != 0 {
		if r.Quantity != 0 {
			return fmt.Errorf("resource %s has StartQuantity and Quantity", r.Name)
		}
		if len(r.Producers) == 0 {
			return fmt.Errorf("resource %s has StartQuantity and no Producers", r.Name)
		}
	}
	if r.StartCapacity != 0 || len(r.CapacityProducers) > 0 {
		if r.Capacity != 0 {
			return fmt.Errorf("resource %s should not set Capacity", r.Name)
		}
		if r.StartCapacity == 0 || len(r.CapacityProducers) == 0 {
			return fmt.Errorf("resource %s should set StartCapacity and CapacityProducers", r.Name)
		}
	}
	return nil
}

func (g *Game) validateResourceName(name string) error {
	if name != "" && !g.HasResource(name) {
		return fmt.Errorf("invalid resource name %s", name)
	}
	return nil
}

func (g *Game) validateActionName(name string) error {
	if name != "" && !g.HasAction(name) {
		return fmt.Errorf("invalid action name %s", name)
	}
	return nil
}

func (g *Game) getRate(resource *data.Resource) float64 {
	factor := 0.0
	for _, p := range resource.Producers {
		factor += g.getOneRate(p)
	}
	bonus := 1.0
	for _, p := range resource.Bonus {
		bonus += g.getOneRate(p)
	}
	return factor * bonus
}

func (g *Game) getCapacityRate(resource *data.Resource) float64 {
	factor := 0.0
	for _, p := range resource.CapacityProducers {
		factor += g.getOneRate(p)
	}
	return factor
}

func (g *Game) getOneRate(resource data.Resource) float64 {
	one := g.getQuantityForRate(resource) * resource.Factor
	bonus := 1.0
	for _, p := range resource.Bonus {
		bonus += g.getOneRate(p)
	}
	return one * bonus
}

func (g *Game) getQuantityForRate(p data.Resource) float64 {
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

func (g *Game) updateRate(resource *data.Resource) {
	for _, p := range resource.Producers {
		if !p.ProductionOnGone {
			continue
		}
		one := g.getOneRate(p)
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
	return g.Resources[g.resourceToIndex[name]]
}

func (g *Game) GetAction(name string) data.Action {
	return g.Actions[g.actionToIndex[name]]
}

func (g *Game) GetActionIndex(name string) int {
	return g.actionToIndex[name]
}

func (g *Game) getCost(a data.Action, c data.Resource) float64 {
	base := c.CostExponentBase
	if base == 0 {
		base = 1
	}
	return c.Quantity * math.Pow(base, g.GetResource(a.Adds[0].Name).Quantity)
}

func (g *Game) HasResource(name string) bool {
	_, ok := g.resourceToIndex[name]
	return ok
}

func (g *Game) HasAction(name string) bool {
	_, ok := g.actionToIndex[name]
	return ok
}
