package textui

import (
	"fmt"
	"log"
	"strings"

	"github.com/kssilveira/idle-game-engine/ui"
)

func Show(logger *log.Logger, separator string, data *ui.Data, isHTML, showActionNumber bool) {
	if separator != "" {
		logger.Printf("%s", separator)
	}
	ShowResources(logger, data)
	ShowActions(logger, data, isHTML, showActionNumber)
	skip := ""
	if data.LastInput.IsSkip {
		skip = "skip "
	}
	isMake := ""
	if data.LastInput.IsMake {
		isMake = "make "
	}
	logger.Printf("last action: %s%s%s\n", skip, isMake, data.LastInput.Action.Name)
	if data.Error != nil {
		logger.Printf("error: %v\n", data.Error)
	}
}

func ShowResources(logger *log.Logger, data *ui.Data) {
	for _, d := range data.Resources {
		r := d.Resource
		if r.IsHidden || r.Quantity == 0 {
			continue
		}
		status := ""
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
				status = "[-] "
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
		logger.Printf("%s[%s] %s %s%s%s\n", status, r.Type, r.Name, toString(r.Quantity), capacity, extra)
	}
}

func ShowActions(logger *log.Logger, data *ui.Data, isHTML, showActionNumber bool) {
	for i, a := range data.Actions {
		if a.Locked {
			continue
		}
		status := ""
		quantity := ""
		if a.Quantity > 0 {
			quantity = fmt.Sprintf(" (%s)", toString(a.Quantity))
		}
		name := fmt.Sprintf("%s%s", a.Name, quantity)
		if isHTML {
			name = fmt.Sprintf("<a href='/%d'>%s%s</a> [<a href='/s%d'>skip</a>]", i, a.Name, quantity, i)
		}
		parts := []string{name}
		costs := getCosts(a.Costs, &status)
		if costs != "" {
			parts = append(parts, fmt.Sprintf(" -(%s)", costs))
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
		number := "XX"
		if showActionNumber {
			number = fmt.Sprintf("%2d", i)
		}
		logger.Printf("%s: %s[%s] %s)\n", number, status, a.Type, strings.Join(parts, ""))
	}
	for _, a := range data.CustomActions {
		logger.Printf("%s\n", a.Name)
	}
}

func getCosts(costs []ui.Cost, status *string) string {
	res := []string{}
	for _, c := range costs {
		overCap := ""
		if c.Cost > c.Capacity && c.Capacity != -1 {
			overCap = "*"
			*status = "[*] "
		}
		duration := ""
		if c.Duration != 0 {
			duration = fmt.Sprintf(" %s", c.Duration)
		}
		out := fmt.Sprintf("%s/%s%s%s", toString(c.Quantity), toString(c.Cost), overCap, duration)
		if c.Quantity >= c.Cost {
			out = fmt.Sprintf("%s", toString(c.Cost))
		}
		var status string
		nested := getCosts(c.Costs, &status)
		extra := ""
		if nested != "" {
			extra = fmt.Sprintf(" (%s)", nested)
		}
		res = append(res, fmt.Sprintf("%s %s%s", c.Name, out, extra))
	}
	return strings.Join(res, ", ")
}

func toString(n float64) string {
	res := fmt.Sprintf("%.2f", n)
	return strings.TrimSuffix(res, ".00")
}
