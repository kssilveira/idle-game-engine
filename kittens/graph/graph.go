package graph

import (
	"fmt"
	"log"

	"github.com/kssilveira/idle-game-engine/data"
	"github.com/kssilveira/idle-game-engine/game"
)

func Graph(logger *log.Logger, g *game.Game, colors map[string]bool) {
	logger.Printf("digraph {\n")
	typeToShape := map[string]string{
		"Resource": "cylinder",
		"Building": "box3d",
		"Village":  "house",
		"Tech":     "diamond",
		"Upgrade":  "hexagon",
		"Craft":    "cds",
		"Trade":    "cds",
	}
	nodes := map[string]bool{}
	edges := map[string]bool{}
	excluded := map[string]bool{
		"happiness":   true,
		"":            true,
		"day":         true,
		"Spring":      true,
		"Winter":      true,
		"gone kitten": true,
	}
	edgefn := func(from, to, color string) {
		edge(logger, nodes, edges, colors, excluded, from, to, color)
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
		for _, p := range r.CapacityProducers {
			edgefn(p.Name, r.Name, "limegreen")
			graphBonus(edgefn, p)
		}
		graphBonus(edgefn, *r)
	}
	for _, a := range g.Actions {
		for _, c := range a.Costs {
			edgefn(c.Name, a.Name, "orange")
			graphBonus(edgefn, c)
		}
		for _, add := range a.Adds {
			edgefn(a.Name, add.Name, "limegreen")
			graphBonus(edgefn, add)
		}
		edgefn(a.UnlockedBy, a.Name, "blue")
	}
	for _, r := range g.Resources {
		if !nodes[r.Name] {
			continue
		}
		logger.Printf(`  "%s" [shape="%s"];`+"\n", r.Name, typeToShape[r.Type])
	}
	for _, a := range g.Actions {
		if !nodes[a.Name] {
			continue
		}
		logger.Printf(`  "%s" [shape="%s"];`+"\n", a.Name, typeToShape[a.Type])
	}
	logger.Printf("}\n")
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
  "Resource" [shape="cylinder"];
  "Building" [shape="box3d"];
  "Village" [shape="house"];
  "Tech" [shape="diamond"];
  "Upgrade" [shape="hexagon"];
  "Craft" [shape="cds"];
  "Trade" [shape="cds"];
}
`)
}

func edge(logger *log.Logger, nodes map[string]bool, edges map[string]bool, colors map[string]bool, excluded map[string]bool, from, to, color string) {
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
	logger.Printf(`  "%s" -> "%s" [color="%s"];`+"\n", from, to, color)
}

func graphBonus(edgefn func(from, to, color string), r data.Resource) {
	for _, b := range r.Bonus {
		color := "green"
		if r.Factor*b.Factor < 0 {
			color = "red"
		}
		edgefn(b.Name, r.Name, color)
		graphBonus(edgefn, b)
	}
}
