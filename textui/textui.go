package textui

import (
	"fmt"
	"log"
	"strings"

	"github.com/kssilveira/idle-game-engine/ui"
)

type Config struct {
	Logger            *log.Logger
	Separator         string
	IsHTML            bool
	HideActionNumbers bool
	HideCustomActions bool
}

func Show(cfg Config, data *ui.Data) {
	if cfg.Separator != "" {
		cfg.Logger.Printf("%s", cfg.Separator)
	}
	showResources(cfg, data)
	showActions(cfg, data)
	prefix := ""
	if data.LastInput.IsSkip {
		prefix = "skip "
	}
	if data.LastInput.IsCreate {
		prefix = "create "
	}
	if data.LastInput.IsMax {
		prefix = "max "
	}
	cfg.Logger.Printf("last action: %s%s\n", prefix, data.LastInput.Action.Name)
	if data.Error != nil {
		cfg.Logger.Printf("error: %v\n", data.Error)
	}
}

func showResources(cfg Config, data *ui.Data) {
	for _, d := range data.Resources {
		r := d.Resource
		if r.IsHidden || d.Rate == 0 || r.Count == 0 {
			continue
		}
		status := ""
		capacity := ""
		if r.Cap > 0 {
			capacity = fmt.Sprintf("/%s", toString(r.Cap))
		}
		extra := ""
		if d.Rate != 0 {
			capStr := ""
			if d.DurationToCap > 0 && r.Cap > 0 {
				capStr = fmt.Sprintf(" %s to cap", d.DurationToCap)
			}
			if d.DurationToEmpty > 0 && r.StartCount == 0 {
				capStr = fmt.Sprintf(" %s to empty", d.DurationToEmpty)
				status = "[-] "
			}
			rateStr := ""
			operator := "+"
			if r.StartCount > 0 {
				extra := ""
				if r.ProductionModulus > 0 {
					extra = fmt.Sprintf(" %% %d", r.ProductionModulus)
					if r.ProductionModulusEquals >= 0 {
						operator = "if"
						extra = fmt.Sprintf("%s == %d", extra, r.ProductionModulusEquals)
					}
				}
				rateStr = fmt.Sprintf("(%s %s %s%s)", toString(r.StartCount), operator, toString(d.Rate), extra)
			} else {
				rateStr = fmt.Sprintf("%s/s", toString(d.Rate))
			}
			extra = fmt.Sprintf(" %s%s", rateStr, capStr)
		}
		cfg.Logger.Printf("%s[%s] %s %s%s%s\n", status, r.Type, r.Name, toString(r.Count), capacity, extra)
	}
}

func showActions(cfg Config, data *ui.Data) {
	for i, a := range data.Actions {
		if a.IsLocked || a.IsHidden {
			continue
		}
		status := ""
		quantity := ""
		if a.Count > 0 {
			quantity = fmt.Sprintf(" (%s)", toString(a.Count))
		}
		name := fmt.Sprintf("%s%s", a.Name, quantity)
		if cfg.IsHTML {
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
			one := fmt.Sprintf("%s %s", r.Name, toString(r.Count))
			if r.Count == 0 && r.Cap > 0 {
				one = fmt.Sprintf("%s cap %s", r.Name, toString(r.Cap))
			}
			adds = append(adds, one)
		}
		parts = append(parts, strings.Join(adds, ", "))
		number := fmt.Sprintf("%3d", i)
		if cfg.HideActionNumbers {
			number = "X"
		}
		cfg.Logger.Printf("%s: %s[%s] %s)\n", number, status, a.Type, strings.Join(parts, ""))
	}
	if !cfg.HideCustomActions {
		for _, a := range data.CustomActions {
			cfg.Logger.Printf("%s\n", a.Name)
		}
	}
}

func getCosts(costs []ui.Cost, status *string) string {
	res := []string{}
	for _, c := range costs {
		if c.Cost == 0 {
			continue
		}
		overCap := ""
		if c.Cost > c.Cap && c.Cap != -1 {
			overCap = "*"
			*status = "[*] "
		}
		duration := ""
		if c.Duration != 0 {
			duration = fmt.Sprintf(" %s", c.Duration)
		}
		out := fmt.Sprintf("%s/%s%s%s", toString(c.Count), toString(c.Cost), overCap, duration)
		if c.Count >= c.Cost {
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
	if n == 0 {
		return "0"
	}
	for precision := 2; precision < 10; precision++ {
		res := fmt.Sprintf("%.*f", precision, n)
		res = strings.TrimRight(res, "0")
		res = strings.TrimRight(res, ".")
		if res != "0" {
			return res
		}
	}
	res := fmt.Sprintf("%f", n)
	res = strings.TrimRight(res, "0")
	res = strings.TrimRight(res, ".")
	return res
}
