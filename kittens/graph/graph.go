package graph

import (
	"fmt"
	"log"
	"sort"

	"github.com/kssilveira/idle-game-engine/data"
	"github.com/kssilveira/idle-game-engine/game"
)

var (
	typeToShape = map[string]string{
		"Resource": "cylinder",
		"Calendar": "cylinder",
		"Building": "box3d",
		"Job":      "house",
		"Tech":     "diamond",
		"Craft":    "cds",
		"Trade":    "cds",
	}
)

func Graph(logger *log.Logger, g *game.Game, colors map[string]bool) {
	logger.Printf("digraph {\n")
	nodes := map[string]bool{}
	edges := map[string]bool{}
	excluded := map[string]bool{
		"happiness":   true,
		"":            true,
		"day":         true,
		"Spring":      true,
		"Winter":      true,
		"gone kitten": true,
		"science":     true,
	}
	counts := map[string]int{}
	edgefn := func(from, to, color string) {
		edge(logger, nodes, edges, colors, excluded, counts, from, to, color)
	}
	for _, r := range g.Resources {
		for _, p := range r.Producers {
			if p.Factor < 0 {
				edgefn(r.Name, p.Name, "red")
			} else {
				edgefn(p.Name, r.Name, "green")
			}
			graphBonus(edgefn, p)
		}
		edgefn(r.CapResource, r.Name, "limegreen")
		edgefn(r.ResetResource, r.Name, "limegreen")
		graphBonus(edgefn, *r)
	}
	for _, a := range g.Actions {
		for _, c := range a.Costs {
			edgefn(c.Name, a.Name, "orange")
			c.Factor = game.GetFactor(c.Factor) * -1
			graphBonus(edgefn, c)
		}
		for _, add := range a.Adds {
			edgefn(a.Name, add.Name, "limegreen")
			graphBonus(edgefn, add)
		}
		edgefn(a.UnlockedBy, a.Name, "blue")
		edgefn(a.CostExponentBaseResource.Name, a.Name, "orange")
		if a.CostExponentBaseResource.Name == "" {
			a.CostExponentBaseResource.Name = a.Name
		}
		a.CostExponentBaseResource.Factor = game.GetFactor(a.CostExponentBaseResource.Factor) * -1
		graphBonus(edgefn, a.CostExponentBaseResource)
	}
	for _, r := range g.Resources {
		if !nodes[r.Name] {
			continue
		}
		nodes[r.Name] = false
		logger.Printf(`  "%s" [shape="%s"];`+"\n", r.Name, typeToShape[r.Type])
	}
	for _, a := range g.Actions {
		if !nodes[a.Name] {
			continue
		}
		nodes[a.Name] = false
		logger.Printf(`  "%s" [shape="%s"];`+"\n", a.Name, typeToShape[a.Type])
	}
	logger.Printf("}\n")

	keys := []string{}
	for key := range counts {
		keys = append(keys, key)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		ki := keys[i]
		kj := keys[j]
		if counts[ki] > counts[kj] {
			return true
		}
		if counts[ki] < counts[kj] {
			return false
		}
		return ki < kj
	})
	for i, k := range keys {
		if i > 5 {
			break
		}
		logger.Printf("# %s: %d\n", k, counts[k])
	}
}

func GraphEdges(logger *log.Logger, g *game.Game, colors map[string]bool) {
	logger.Printf(`
digraph {
  node [label="" width=0 style=invis];
  { rank="same"; n0; n1; n2; n3; n4; n5; n6; n7; n8; n9; }
  n0 -> n1 [color="red" label="feeds"];
  n2 -> n3 [color="green" label="produces"];
  n4 -> n5 [color="orange" label="buys"];
  n6 -> n7 [color="limegreen" label="adds"];
  n8 -> n9 [color="blue" label="unlocks"];
}
`)
}

func GraphNodes(logger *log.Logger, g *game.Game, colors map[string]bool) {
	logger.Printf(`
digraph {
`)
	keys := []string{}
	for key := range typeToShape {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		logger.Printf(`  "%s" [shape="%s"];`+"\n", key, typeToShape[key])
	}
	logger.Printf(`
}
`)
}

func edge(logger *log.Logger, nodes map[string]bool, edges map[string]bool, colors map[string]bool, excluded map[string]bool, counts map[string]int, from, to, color string) {
	if from == to {
		return
	}
	if len(colors) > 0 && !colors[color] {
		return
	}
	if excluded[from] || excluded[to] {
		return
	}
	key := fmt.Sprintf("%s+%s+%s", from, to, color)
	if edges[key] {
		return
	}
	edges[key] = true
	nodes[from] = true
	nodes[to] = true
	counts[from]++
	counts[to]++
	logger.Printf(`  "%s" -> "%s" [color="%s"];`+"\n", from, to, color)
}

func graphBonus(edgefn func(from, to, color string), r data.Resource) {
	for _, b := range r.Bonus {
		color := "green"
		if game.GetFactor(r.Factor)*game.GetFactor(b.Factor) < 0 {
			color = "red"
		}
		if b.Name == "" {
			b.Name = r.Name
		}
		edgefn(b.Name, r.Name, color)
		graphBonus(edgefn, b)
	}
}
