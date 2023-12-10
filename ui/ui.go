package ui

import (
	"time"

	"github.com/kssilveira/idle-game-engine/data"
)

type Data struct {
	LastInput     data.ParsedInput
	Error         error
	Resources     []Resource
	Actions       []Action
	CustomActions []CustomAction
}

type Resource struct {
	Resource        data.Resource
	Rate            float64
	DurationToCap   time.Duration
	DurationToEmpty time.Duration
}

type Action struct {
	Name     string
	Type     string
	Quantity float64
	Locked   bool
	Costs    []Cost
	Adds     []Add
}

type Cost struct {
	Name     string
	Quantity float64
	Capacity float64
	Cost     float64
	Duration time.Duration
	Costs    []Cost
}

type Add struct {
	Name     string
	Quantity float64
	Capacity float64
}

type CustomAction struct {
	Name string
}
