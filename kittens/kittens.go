package kittens

import (
	"fmt"
	"log"
	"time"

	"github.com/kssilveira/idle-game-engine/data"
	"github.com/kssilveira/idle-game-engine/game"
)

func NewGame(now game.Now) *game.Game {
	g := game.NewGame(now())
	g.AddResources([]data.Resource{{
		Name: "day", Type: "Calendar", IsHidden: true, Quantity: 0, Capacity: -1,
		Producers: []data.Resource{{
			Name: "", ProductionFactor: 0.5,
		}},
	}, {
		Name: "year", Type: "Calendar", Quantity: 1, StartQuantity: 1, Capacity: -1,
		Producers: []data.Resource{{
			Name: "day", ProductionFactor: 0.0025, ProductionFloor: true,
		}},
	}, {
		Name: "Spring", Type: "Calendar", Quantity: 1,
		StartQuantity: 1, ProductionModulus: 4, ProductionModulusEquals: 0, Capacity: -1,
		Producers: []data.Resource{{
			Name: "day", ProductionFactor: 0.01, ProductionFloor: true,
		}},
	}, {
		Name: "Summer", Type: "Calendar",
		StartQuantity: 1, ProductionModulus: 4, ProductionModulusEquals: 1, Capacity: -1,
		Producers: []data.Resource{{
			Name: "day", ProductionFactor: 0.01, ProductionFloor: true,
		}},
	}, {
		Name: "Autumn", Type: "Calendar",
		StartQuantity: 1, ProductionModulus: 4, ProductionModulusEquals: 2, Capacity: -1,
		Producers: []data.Resource{{
			Name: "day", ProductionFactor: 0.01, ProductionFloor: true,
		}},
	}, {
		Name: "Winter", Type: "Calendar",
		StartQuantity: 1, ProductionModulus: 4, ProductionModulusEquals: 3, Capacity: -1,
		Producers: []data.Resource{{
			Name: "day", ProductionFactor: 0.01, ProductionFloor: true,
		}},
	}, {
		Name: "day_of_year", Type: "Calendar", Quantity: 1,
		StartQuantity: 1, ProductionModulus: 400, ProductionModulusEquals: -1,
		Capacity: -1,
		Producers: []data.Resource{{
			Name: "day", ProductionFactor: 1, ProductionFloor: true,
		}},
	}, {
		Name: "catnip", Type: "Resource", Capacity: 5000,
		Producers: []data.Resource{{
			Name: "Catnip Field", ProductionFactor: 0.63 * (1 + 0.50), ProductionResourceFactor: "Spring",
		}, {
			Name: "Catnip Field", ProductionFactor: 0.63, ProductionResourceFactor: "Summer",
		}, {
			Name: "Catnip Field", ProductionFactor: 0.63, ProductionResourceFactor: "Autumn",
		}, {
			Name: "Catnip Field", ProductionFactor: 0.63 * (1 - 0.75), ProductionResourceFactor: "Winter",
		}, {
			Name: "kitten", ProductionFactor: -4.25, ProductionFloor: true, ProductionOnGone: true,
		}, {
			Name: "woodcutter", ProductionFactor: -4.25, ProductionOnGone: true,
		}, {
			Name: "scholar", ProductionFactor: -4.25, ProductionOnGone: true,
		}, {
			Name: "farmer", ProductionFactor: -4.25, ProductionOnGone: true,
		}, {
			Name: "hunter", ProductionFactor: -4.25, ProductionOnGone: true,
		}, {
			Name: "miner", ProductionFactor: -4.25, ProductionOnGone: true,
		}, {
			Name: "farmer", ProductionFactor: 5, ProductionResourceFactor: "happiness",
			ProductionBonus: []data.Resource{{
				Name: "Mineral Hoes", ProductionFactor: 0.5,
			}, {
				Name: "Iron Hoes", ProductionFactor: 0.3,
			}},
		}},
	}, {
		Name: "wood", Type: "Resource", Capacity: 200,
		Producers: []data.Resource{{
			Name: "woodcutter", ProductionFactor: 0.09, ProductionResourceFactor: "happiness",
		}},
	}, {
		Name: "science", Type: "Resource", Capacity: 250,
		Producers: []data.Resource{{
			Name: "scholar", ProductionFactor: 0.175, ProductionResourceFactor: "happiness",
		}},
		ProductionBonus: []data.Resource{{
			Name: "Library", ProductionFactor: 0.1,
		}},
	}, {
		Name: "catpower", Type: "Resource", Capacity: 250,
		Producers: []data.Resource{{
			Name: "hunter", ProductionFactor: 0.3, ProductionResourceFactor: "happiness",
		}},
	}, {
		Name: "minerals", Type: "Resource", Capacity: 250,
		Producers: []data.Resource{{
			Name: "miner", ProductionFactor: 0.25, ProductionResourceFactor: "happiness",
		}},
		ProductionBonus: []data.Resource{{
			Name: "Mine", ProductionFactor: 0.2,
		}},
	}, {
		Name: "iron", Type: "Resource", Capacity: 100,
	}, {
		Name: "gold", Type: "Resource", Capacity: 20,
	}, {
		Name: "kitten", Type: "Resource", Capacity: 0,
		Producers: []data.Resource{{
			Name: "", ProductionFactor: 0.05,
		}},
		OnGone: []data.Resource{{
			Name: "gone kitten", Quantity: 1,
		}},
	}, {
		Name: "all kittens", Type: "Resource", IsHidden: true, Capacity: -1, StartQuantity: 1,
		Producers: []data.Resource{{
			Name: "", ProductionFactor: -1,
		}, {
			Name: "kitten", ProductionFactor: 1, ProductionFloor: true,
		}, {
			Name: "woodcutter", ProductionFactor: 1,
		}, {
			Name: "scholar", ProductionFactor: 1,
		}, {
			Name: "farmer", ProductionFactor: 1,
		}, {
			Name: "hunter", ProductionFactor: 1,
		}, {
			Name: "miner", ProductionFactor: 1,
		}},
	}, {
		Name: "furs", Type: "Resource", Capacity: -1,
		Producers: []data.Resource{{
			Name: "all kittens", ProductionFactor: -0.05,
		}},
	}, {
		Name: "ivory", Type: "Resource", Capacity: -1,
		Producers: []data.Resource{{
			Name: "all kittens", ProductionFactor: -0.035,
		}},
	}, {
		Name: "spice", Type: "Resource", Capacity: -1,
		Producers: []data.Resource{{
			Name: "all kittens", ProductionFactor: -0.005,
		}},
	}, {
		Name: "unicorns", Type: "Resource", Capacity: -1,
	}, {
		Name: "blueprint", Type: "Resource", Capacity: -1,
	}, {
		Name: "gone kitten", Type: "Resource", Capacity: -1,
	}, {
		Name: "Catnip Field", Type: "Bonfire", IsHidden: true, Capacity: -1,
	}, {
		Name: "Hut", Type: "Bonfire", IsHidden: true, Capacity: -1,
	}, {
		Name: "Library", Type: "Bonfire", IsHidden: true, Capacity: -1,
	}, {
		Name: "Barn", Type: "Bonfire", IsHidden: true, Capacity: -1,
	}, {
		Name: "Mine", Type: "Bonfire", IsHidden: true, Capacity: -1,
	}, {
		Name: "Workshop", Type: "Bonfire", IsHidden: true, Capacity: -1,
	}, {
		Name: "woodcutter", Type: "Village", IsHidden: true, Capacity: -1,
		OnGone: []data.Resource{{
			Name: "gone kitten", Quantity: 1,
		}, {
			Name: "kitten", Capacity: 1,
		}},
	}, {
		Name: "scholar", Type: "Village", IsHidden: true, Capacity: -1,
		OnGone: []data.Resource{{
			Name: "gone kitten", Quantity: 1,
		}, {
			Name: "kitten", Capacity: 1,
		}},
	}, {
		Name: "farmer", Type: "Village", IsHidden: true, Capacity: -1,
		OnGone: []data.Resource{{
			Name: "gone kitten", Quantity: 1,
		}, {
			Name: "kitten", Capacity: 1,
		}},
	}, {
		Name: "hunter", Type: "Village", IsHidden: true, Capacity: -1,
		OnGone: []data.Resource{{
			Name: "gone kitten", Quantity: 1,
		}, {
			Name: "kitten", Capacity: 1,
		}},
	}, {
		Name: "miner", Type: "Village", IsHidden: true, Capacity: -1,
		OnGone: []data.Resource{{
			Name: "gone kitten", Quantity: 1,
		}, {
			Name: "kitten", Capacity: 1,
		}},
	}, {
		Name: "happiness", Type: "Village", StartQuantity: 1.1, Capacity: -1,
		Producers: []data.Resource{{
			Name: "all kittens", ProductionFactor: -0.02,
		}, {
			Name: "ivory", ProductionFactor: 0.1, ProductionBoolean: true,
		}, {
			Name: "furs", ProductionFactor: 0.1, ProductionBoolean: true,
		}, {
			Name: "spice", ProductionFactor: 0.1, ProductionBoolean: true,
		}, {
			Name: "unicorns", ProductionFactor: 0.1, ProductionBoolean: true,
		}},
	}, {
		Name: "Calendar", Type: "Science", IsHidden: true, Capacity: 1,
	}, {
		Name: "Agriculture", Type: "Science", IsHidden: true, Capacity: 1,
	}, {
		Name: "Archery", Type: "Science", IsHidden: true, Capacity: 1,
	}, {
		Name: "Mining", Type: "Science", IsHidden: true, Capacity: 1,
	}, {
		Name: "Animal Husbandry", Type: "Science", IsHidden: true, Capacity: 1,
	}, {
		Name: "Metal Working", Type: "Science", IsHidden: true, Capacity: 1,
	}, {
		Name: "Mineral Hoes", Type: "Workshop", IsHidden: true, Capacity: 1,
	}, {
		Name: "Iron Hoes", Type: "Workshop", IsHidden: true, Capacity: 1,
	}})
	g.Actions = []game.Action{{
		Name: "Gather catnip", Type: "Bonfire",
		LockedBy: data.Resource{Name: "Catnip Field"},
		Adds: []data.Resource{{
			Name: "catnip", Quantity: 1,
		}},
	}, {
		Name: "Refine catnip", Type: "Bonfire",
		UnlockedBy: data.Resource{Name: "catnip"},
		LockedBy:   data.Resource{Name: "woodcutter"},
		Costs: []data.Resource{{
			Name: "catnip", Quantity: 100,
		}},
		Adds: []data.Resource{{
			Name: "wood", Quantity: 1,
		}},
	}, {
		Name: "Catnip Field", Type: "Bonfire",
		UnlockedBy: data.Resource{Name: "catnip"},
		Costs: []data.Resource{{
			Name: "catnip", Quantity: 10, CostExponentBase: 1.12,
		}},
		Adds: []data.Resource{{
			Name: "Catnip Field", Quantity: 1,
		}},
	}, {
		Name: "Hut", Type: "Bonfire",
		UnlockedBy: data.Resource{Name: "wood"},
		Costs: []data.Resource{{
			Name: "wood", Quantity: 5, CostExponentBase: 2.5,
		}},
		Adds: []data.Resource{{
			Name: "Hut", Quantity: 1,
		}, {
			Name: "kitten", Capacity: 2,
		}},
	}, {
		Name: "Library", Type: "Bonfire",
		UnlockedBy: data.Resource{Name: "wood"},
		Costs: []data.Resource{{
			Name: "wood", Quantity: 25, CostExponentBase: 1.15,
		}},
		Adds: []data.Resource{{
			Name: "Library", Quantity: 1,
		}, {
			Name: "science", Capacity: 250,
		}},
	}, {
		Name: "Barn", Type: "Bonfire",
		UnlockedBy: data.Resource{Name: "Agriculture"},
		Costs: []data.Resource{{
			Name: "wood", Quantity: 50, CostExponentBase: 1.75,
		}},
		Adds: []data.Resource{{
			Name: "Barn", Quantity: 1,
		}, {
			Name: "catnip", Capacity: 5000,
		}, {
			Name: "wood", Capacity: 200,
		}, {
			Name: "minerals", Capacity: 250,
		}},
	}, {
		Name: "Mine", Type: "Bonfire",
		UnlockedBy: data.Resource{Name: "Mining"},
		Costs: []data.Resource{{
			Name: "wood", Quantity: 100, CostExponentBase: 1.15,
		}},
		Adds: []data.Resource{{
			Name: "Mine", Quantity: 1,
		}},
	}, {
		Name: "Workshop", Type: "Bonfire",
		UnlockedBy: data.Resource{Name: "Mining"},
		Costs: []data.Resource{{
			Name: "wood", Quantity: 100, CostExponentBase: 1.15,
		}, {
			Name: "minerals", Quantity: 400, CostExponentBase: 1.15,
		}},
		Adds: []data.Resource{{
			Name: "Workshop", Quantity: 1,
		}},
	}, {
		Name: "woodcutter", Type: "Village",
		UnlockedBy: data.Resource{Name: "Hut"},
		Costs: []data.Resource{{
			Name: "kitten", Quantity: 1, Capacity: 1,
		}},
		Adds: []data.Resource{{
			Name: "woodcutter", Quantity: 1,
		}},
	}, {
		Name: "scholar", Type: "Village",
		UnlockedBy: data.Resource{Name: "Library"},
		Costs: []data.Resource{{
			Name: "kitten", Quantity: 1, Capacity: 1,
		}},
		Adds: []data.Resource{{
			Name: "scholar", Quantity: 1,
		}},
	}, {
		Name: "farmer", Type: "Village",
		UnlockedBy: data.Resource{Name: "Agriculture"},
		Costs: []data.Resource{{
			Name: "kitten", Quantity: 1, Capacity: 1,
		}},
		Adds: []data.Resource{{
			Name: "farmer", Quantity: 1,
		}},
	}, {
		Name: "hunter", Type: "Village",
		UnlockedBy: data.Resource{Name: "Archery"},
		Costs: []data.Resource{{
			Name: "kitten", Quantity: 1, Capacity: 1,
		}},
		Adds: []data.Resource{{
			Name: "hunter", Quantity: 1,
		}},
	}, {
		Name: "miner", Type: "Village",
		UnlockedBy: data.Resource{Name: "Mine"},
		Costs: []data.Resource{{
			Name: "kitten", Quantity: 1, Capacity: 1,
		}},
		Adds: []data.Resource{{
			Name: "miner", Quantity: 1,
		}},
	}, {
		Name: "Send hunters", Type: "Village",
		UnlockedBy: data.Resource{Name: "Archery"},
		Costs: []data.Resource{{
			Name: "catpower", Quantity: 100,
		}},
		Adds: []data.Resource{{
			Name: "furs", Quantity: 39.5,
		}, {
			Name: "ivory", Quantity: 10.78,
		}, {
			Name: "unicorns", Quantity: 0.05,
		}},
	}, {
		Name: "Calendar", Type: "Science",
		UnlockedBy: data.Resource{Name: "Library"},
		LockedBy:   data.Resource{Name: "Calendar"},
		Costs: []data.Resource{{
			Name: "science", Quantity: 30,
		}},
		Adds: []data.Resource{{
			Name: "Calendar", Quantity: 1,
		}},
	}, {
		Name: "Agriculture", Type: "Science",
		UnlockedBy: data.Resource{Name: "Calendar"},
		LockedBy:   data.Resource{Name: "Agriculture"},
		Costs: []data.Resource{{
			Name: "science", Quantity: 100,
		}},
		Adds: []data.Resource{{
			Name: "Agriculture", Quantity: 1,
		}},
	}, {
		Name: "Archery", Type: "Science",
		UnlockedBy: data.Resource{Name: "Calendar"},
		LockedBy:   data.Resource{Name: "Archery"},
		Costs: []data.Resource{{
			Name: "science", Quantity: 300,
		}},
		Adds: []data.Resource{{
			Name: "Archery", Quantity: 1,
		}},
	}, {
		Name: "Mining", Type: "Science",
		UnlockedBy: data.Resource{Name: "Calendar"},
		LockedBy:   data.Resource{Name: "Mining"},
		Costs: []data.Resource{{
			Name: "science", Quantity: 500,
		}},
		Adds: []data.Resource{{
			Name: "Mining", Quantity: 1,
		}},
	}, {
		Name: "Animal Husbandry", Type: "Science",
		UnlockedBy: data.Resource{Name: "Archery"},
		LockedBy:   data.Resource{Name: "Animal Husbandry"},
		Costs: []data.Resource{{
			Name: "science", Quantity: 500,
		}},
		Adds: []data.Resource{{
			Name: "Animal Husbandry", Quantity: 1,
		}},
	}, {
		Name: "Metal Working", Type: "Science",
		UnlockedBy: data.Resource{Name: "Mining"},
		LockedBy:   data.Resource{Name: "Metal Working"},
		Costs: []data.Resource{{
			Name: "science", Quantity: 900,
		}},
		Adds: []data.Resource{{
			Name: "Metal Working", Quantity: 1,
		}},
	}, {
		Name: "Mineral Hoes", Type: "Workshop",
		UnlockedBy: data.Resource{Name: "Workshop"},
		LockedBy:   data.Resource{Name: "Mineral Hoes"},
		Costs: []data.Resource{{
			Name: "minerals", Quantity: 275,
		}, {
			Name: "science", Quantity: 100,
		}},
		Adds: []data.Resource{{
			Name: "Mineral Hoes", Quantity: 1,
		}},
	}, {
		Name: "Iron Hoes", Type: "Workshop",
		UnlockedBy: data.Resource{Name: "Workshop"},
		LockedBy:   data.Resource{Name: "Iron Hoes"},
		Costs: []data.Resource{{
			Name: "iron", Quantity: 25,
		}, {
			Name: "science", Quantity: 200,
		}},
		Adds: []data.Resource{{
			Name: "Iron Hoes", Quantity: 1,
		}},
	}, {
		Name: "Lizards", Type: "Trade",
		UnlockedBy: data.Resource{Name: "Archery"},
		Costs: []data.Resource{{
			Name: "catpower", Quantity: 50,
		}, {
			Name: "gold", Quantity: 15,
		}, {
			Name: "minerals", Quantity: 1000,
		}},
		Adds: []data.Resource{{
			Name: "wood", Quantity: 500,
		}, {
			Name: "blueprint", Quantity: 0.1,
		}, {
			Name: "spice", Quantity: 8.75,
		}},
	}}
	return g
}

const (
	gather = iota
	refine
	field
	hut
	library
	barn
	mine
	workshop
	woodcutter
	scholar
	farmer
	hunter
	miner
	sendhunters
	calendar
	agriculture
	archery
	mining
	animalhusbandry
	metalworking
	mineralhoes
	ironhoes
)

const (
	sdelta = 100
)

const (
	_ = iota + sdelta
	srefine
	sfield
	shut
	slibrary
	sbarn
	smine
	sworkshop
	swoodcutter
	sscholar
	sfarmer
	shunter
	sminer
	ssendhunters
	scalendar
	sagriculture
	sarchery
	smining
	sanimalhusbandry
	smetalworking
	smineralhoes
	sironhoes
)

func Solve(input chan string, sleepMS int) {
	for _, one := range []struct {
		cmds  []int
		count int
	}{
		{[]int{gather}, 10},
		{[]int{field}, 1},
		{[]int{sfield, field}, 58},
		{[]int{srefine, refine}, 5},
		{[]int{hut, swoodcutter, woodcutter}, 1},
		{[]int{slibrary, library, sscholar, scholar}, 1},
		{[]int{slibrary, library}, 14},
		{[]int{scalendar, calendar}, 1},
		{[]int{sagriculture, agriculture}, 1},

		{[]int{sbarn, barn}, 6},
		{[]int{sfield, field}, 25},
		{[]int{slibrary, library}, 15},
		{[]int{shut, hut, sfarmer, farmer}, 8},

		{[]int{sarchery, archery}, 1},
		{[]int{shut, hut, shunter, hunter}, 1},

		{[]int{smining, mining}, 1},
		{[]int{smine, mine}, 20},
		{[]int{shut, hut, sminer, miner}, 1},

		{[]int{sworkshop, workshop}, 20},
		{[]int{smineralhoes, mineralhoes}, 1},

		{[]int{sanimalhusbandry, animalhusbandry}, 1},
		{[]int{smetalworking, metalworking}, 1},

		{[]int{ssendhunters, sendhunters}, 10},
	} {
		for i := 0; i < one.count; i++ {
			for _, cmd := range one.cmds {
				input <- toInput(cmd)
				time.Sleep(time.Second * time.Duration(sleepMS) / 1000.)
			}
		}
	}
}

func toInput(cmd int) string {
	prefix := ""
	if cmd >= sdelta {
		prefix = "s"
		cmd -= sdelta
	}
	return fmt.Sprintf("%s%d", prefix, cmd)
}

func Graph(logger *log.Logger, g *game.Game) {
	logger.Printf("digraph {\n")
	typeToShape := map[string]string{
		"Resource": "cylinder",
		"Bonfire":  "box3d",
		"Village":  "house",
		"Science":  "diamond",
		"Workshop": "hexagon",
		"Trade":    "cds",
	}
	for _, r := range g.Resources {
		if r.Type == "Calendar" || r.Name == "gone kitten" || r.Name == "happiness" {
			continue
		}
		logger.Printf(`  "%s" [shape="%s"];`+"\n", r.Name, typeToShape[r.Type])
	}
	for _, a := range g.Actions {
		logger.Printf(`  "%s" [shape="%s"];`+"\n", a.Name, typeToShape[a.Type])
	}
	for _, r := range g.Resources {
		if r.Name == "happiness" {
			continue
		}
		last := ""
		for _, p := range r.Producers {
			if p.Name == "" || p.Name == "day" || p.Name == last {
				continue
			}
			last = p.Name
			if p.ProductionFactor < 0 {
				logger.Printf(`  "%s" -> "%s" [color="red"];`+"\n", r.Name, p.Name)
			} else {
				logger.Printf(`  "%s" -> "%s" [color="green"];`+"\n", p.Name, r.Name)
			}
			for _, b := range p.ProductionBonus {
				logger.Printf(`  "%s" -> "%s" [color="green"];`+"\n", b.Name, p.Name)
			}
		}
		for _, b := range r.ProductionBonus {
			logger.Printf(`  "%s" -> "%s" [color="green"];`+"\n", b.Name, r.Name)
		}
	}
	for _, a := range g.Actions {
		for _, c := range a.Costs {
			logger.Printf(`  "%s" -> "%s" [color="orange"];`+"\n", c.Name, a.Name)
		}
		for _, add := range a.Adds {
			if a.Name == add.Name {
				continue
			}
			logger.Printf(`  "%s" -> "%s" [color="limegreen"];`+"\n", a.Name, add.Name)
		}
		if a.UnlockedBy.Name != "" {
			logger.Printf(`  "%s" -> "%s" [color="blue"];`+"\n", a.UnlockedBy.Name, a.Name)
		}
	}
	logger.Printf("}\n")
}

func GraphEdges(logger *log.Logger, g *game.Game) {
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

func GraphNodes(logger *log.Logger, g *game.Game) {
	logger.Printf(`
digraph {
  "Resource" [shape="cylinder"];
  "Bonfire" [shape="box3d"];
  "Village" [shape="house"];
  "Science" [shape="diamond"];
  "Workshop" [shape="hexagon"];
  "Trade" [shape="cds"];
}
`)
}
