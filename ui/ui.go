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
	ShowAll       bool
	HideOverCap   bool
	HideCustom    bool
	HideAction    map[string]bool
	HideResource  map[string]bool
}

type Resource struct {
	Resource        data.Resource
	Rate            float64
	DurationToCap   time.Duration
	DurationToEmpty time.Duration
}

type Action struct {
	Name      string
	Type      string
	Count     float64
	IsLocked  bool
	IsHidden  bool
	IsOverCap bool
	Costs     []Cost
	Adds      []Add
}

type Cost struct {
	Name      string
	Count     float64
	Cap       float64
	Cost      float64
	IsOverCap bool
	Duration  time.Duration
	Costs     []Cost
}

type Add struct {
	Name  string
	Count float64
	Cap   float64
}

type CustomAction struct {
	Name string
}
