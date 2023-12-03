package textui

import (
	"fmt"
	"log"
	"strings"

	"github.com/kssilveira/idle-game-engine/ui"
)

func Show(logger *log.Logger, separator string, data *ui.Data, isHTML bool) {
	if separator != "" {
		logger.Printf("%s", separator)
	}
	ShowResources(logger, data)
	ShowActions(logger, data, isHTML)
	logger.Printf("last input: %s\n", data.LastInput)
	if data.Error != nil {
		logger.Printf("error: %v\n", data.Error)
	}
}

func ShowResources(logger *log.Logger, data *ui.Data) {
	for _, d := range data.Resources {
		r := d.Resource
		if r.Quantity == 0 {
			continue
		}
		capacity := ""
		if r.Capacity > 0 {
			capacity = fmt.Sprintf("/%s", toString(r.Capacity))
		}
		extra := ""
		if d.Rate != 0 {
			capStr := ""
			if d.DurationToCap > 0 && r.Capacity > 0 {
				capStr = fmt.Sprintf(" %s to cap", d.DurationToCap)
			}
			if d.DurationToEmpty > 0 && r.StartQuantity == 0 {
				capStr = fmt.Sprintf(" %s to empty", d.DurationToEmpty)
			}
			rateStr := ""
			operator := "+"
			if r.StartQuantity > 0 {
				extra := ""
				if r.ProductionModulus > 0 {
					extra = fmt.Sprintf(" %% %d", r.ProductionModulus)
					if r.ProductionModulusEquals >= 0 {
						operator = "if"
						extra = fmt.Sprintf("%s == %d", extra, r.ProductionModulusEquals)
					}
				}
				rateStr = fmt.Sprintf("(%s %s %s%s)", toString(r.StartQuantity), operator, toString(d.Rate), extra)
			} else {
				rateStr = fmt.Sprintf("%s/s", toString(d.Rate))
			}
			extra = fmt.Sprintf(" %s%s", rateStr, capStr)
		}
		logger.Printf("[%s] %s %s%s%s\n", r.Type, r.Name, toString(r.Quantity), capacity, extra)
	}
}

func ShowActions(logger *log.Logger, data *ui.Data, isHTML bool) {
	for i, a := range data.Actions {
		if a.Locked {
			continue
		}
		name := fmt.Sprintf("%d: [%s] %s", i, a.Type, a.Name)
		if isHTML {
			name = fmt.Sprintf("%d: [%s] <a href='/%d'>%s</a> [<a href='/s%d'>skip</a>]", i, a.Type, i, a.Name, i)
		}
		parts := []string{name}
		costs := []string{}
		for _, c := range a.Costs {
			overCap := ""
			if c.Cost > c.Capacity && c.Capacity != -1 {
				overCap = "*"
			}
			out := fmt.Sprintf("%s/%s%s %s", toString(c.Quantity), toString(c.Cost), overCap, c.Duration)
			if c.Quantity >= c.Cost {
				out = fmt.Sprintf("%s", toString(c.Cost))
			}
			costs = append(costs, fmt.Sprintf("%s %s", c.Name, out))
		}
		if len(costs) > 0 {
			parts = append(parts, fmt.Sprintf(" -(%s)", strings.Join(costs, "")))
		}
		parts = append(parts, " +(")
		adds := []string{}
		for _, r := range a.Adds {
			one := fmt.Sprintf("%s %s", r.Name, toString(r.Quantity))
			if r.Quantity == 0 && r.Capacity > 0 {
				one = fmt.Sprintf("%s cap %s", r.Name, toString(r.Capacity))
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

func toString(n float64) string {
	res := fmt.Sprintf("%.2f", n)
	return strings.TrimSuffix(res, ".00")
}
