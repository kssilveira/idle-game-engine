package textui

import (
	"fmt"
	"log"
	"strings"

	"github.com/kssilveira/idle-game-engine/ui"
)

func Show(logger *log.Logger, separator string, data *ui.Data) {
	logger.Printf("%s", separator)
	ShowResources(logger, data)
	ShowActions(logger, data)
	logger.Printf("last input: %s\n", data.LastInput)
	if data.Error != nil {
		logger.Printf("error: %v\n", data.Error)
	}
}

func ShowResources(logger *log.Logger, data *ui.Data) {
	for _, r := range data.Resources {
		if r.Quantity == 0 {
			continue
		}
		capacity := ""
		if r.Capacity > 0 {
			capacity = fmt.Sprintf("/%.0f", r.Capacity)
		}
		extra := ""
		if r.Rate != 0 {
			capStr := ""
			if r.DurationToCap > 0 && r.Capacity > 0 {
				capStr = fmt.Sprintf(", %s to cap", r.DurationToCap)
			}
			if r.DurationToEmpty > 0 && r.StartQuantity == 0 {
				capStr = fmt.Sprintf(", %s to empty", r.DurationToEmpty)
			}
			rateStr := ""
			if r.StartQuantity > 0 {
				rateStr = fmt.Sprintf("%.2f + %.2f", r.StartQuantity, r.Rate)
			} else {
				rateStr = fmt.Sprintf("%.2f/s", r.Rate)
			}
			extra = fmt.Sprintf(" (%s%s)", rateStr, capStr)
		}
		logger.Printf("%s %.2f%s%s\n", r.Name, r.Quantity, capacity, extra)
	}
}

func ShowActions(logger *log.Logger, data *ui.Data) {
	for i, a := range data.Actions {
		parts := []string{
			fmt.Sprintf("%d: '%s' (", i, a.Name),
		}
		for _, c := range a.Costs {
			out := fmt.Sprintf("%.2f/%.2f %s", c.Quantity, c.Cost, c.Duration)
			if c.Quantity >= c.Cost {
				out = fmt.Sprintf("%.2f", c.Cost)
			}
			parts = append(parts, fmt.Sprintf("%s %s", c.Name, out))
		}
		parts = append(parts, ") (")
		adds := []string{}
		for _, r := range a.Adds {
			one := fmt.Sprintf("%s + %.0f", r.Name, r.Quantity)
			if r.Quantity == 0 && r.Capacity > 0 {
				one = fmt.Sprintf("%s cap + %.0f", r.Name, r.Capacity)
			}
			adds = append(adds, one)
		}
		parts = append(parts, strings.Join(adds, ", "))
		logger.Printf("%s)\n", strings.Join(parts, ""))
	}
	for _, a := range data.CustomActions {
		logger.Printf("%s\n", a.Name)
	}
}
