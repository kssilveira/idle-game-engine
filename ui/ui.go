package ui

import (
	"time"
)

type Data struct {
	LastInput     string
	Error         error
	Resources     []Resource
	Actions       []Action
	CustomActions []CustomAction
}

type Resource struct {
	Name            string
	Quantity        float64
	Capacity        float64
	Rate            float64
	DurationToCap   time.Duration
	DurationToEmpty time.Duration
	StartQuantity   float64
}

type Action struct {
	Name  string
	Costs []Cost
	Adds  []Add
}

type Cost struct {
	Name     string
	Quantity float64
	Capacity float64
	Cost     float64
	Duration time.Duration
}

type Add struct {
	Name     string
	Quantity float64
	Capacity float64
}

type CustomAction struct {
	Name string
}
