package solve

import (
	"fmt"
	"strings"
	"time"

	"github.com/kssilveira/idle-game-engine/game"
	"github.com/kssilveira/idle-game-engine/ui"
)

type Config struct {
	Game      *game.Game
	Input     chan string
	LastData  *ui.Data
	Waiting   chan bool
	Refreshed chan bool
	Type      string
	SleepMS   int
	PermFn    func(int) []int
}

func Solve(cfg Config) error {
	switch cfg.Type {
	case "smart":
		return solveSmart(cfg)
	case "random":
		return solveRandom(cfg)
	case "fixed":
		fallthrough
	case "":
		return solveFixed(cfg)
	default:
		return fmt.Errorf("invalid type %s", cfg.Type)
	}
	return nil
}

func solveSmart(cfg Config) error {
	cfg.Input <- "hc"
	cfg.Input <- "hm"
	reset := 0
	firstResource := map[string]bool{}
	start := time.Now()
	for iter := 0; ; iter++ {
		if iter == 120 {
			iter = 0
			cfg.Input <- "r"
			cfg.Input <- "S"
			reset += 1

			fmt.Printf("\n==========\n%s: reset %d\n==========\n\n", time.Since(start), reset)

			for _, resource := range cfg.LastData.Resources {
				if !firstResource[resource.Resource.Name] {
					fmt.Printf("missing resource %s\n", resource.Resource.Name)
				}
				if resource.Resource.Name == "paragon" {
					fmt.Printf("\nresource paragon %f\n\n", resource.Resource.Count)
				}
			}

			fmt.Printf("\n==========\n%s: reset %d\n==========\n\n", time.Since(start), reset)
		}
		cfg.Waiting <- true
		<-cfg.Refreshed

		for _, resource := range cfg.LastData.Resources {
			if resource.Resource.Count > 0 && !firstResource[resource.Resource.Name] {
				firstResource[resource.Resource.Name] = true
				fmt.Printf("%s: resource %s\n", time.Since(start), resource.Resource.Name)
			}
		}

		kittenActions := []ui.Action{}
		kittenCounts := map[int]float64{}
		kittenIndexes := map[string]int{}
		isKitten := map[int]bool{}
		minJobCount := 999.0
		for _, action := range cfg.LastData.Actions {
			if action.IsLocked || action.IsHidden {
				continue
			}
			if action.Type != "Job" {
				continue
			}
			if action.Count > 0 && action.Count < minJobCount {
				minJobCount = action.Count
			}
		}
		for index, action := range cfg.LastData.Actions {
			if action.IsLocked || action.IsHidden || action.Name == "Burn Paragon" {
				continue
			}
			if action.IsOverCap && action.Type != "Job" {
				continue
			}
			if strings.HasPrefix(action.Name, "Active ") && action.Count > 0 {
				continue
			}
			for _, add := range action.Adds {
				if add.Name == "kitten" {
					kittenActions = append(kittenActions, action)
					isKitten[index] = true
					kittenCounts[index] = add.Cap
					kittenIndexes[action.Name] = index
					break
				}
			}
			if action.Type == "Job" && action.Count == 0 && len(kittenActions) > 0 {
				kittenAction := kittenActions[0]
				kittenIndex := kittenIndexes[kittenAction.Name]
				cfg.Input <- fmt.Sprintf("s %d", kittenIndex)
				if kittenCounts[kittenIndex] == 1 {
					cfg.Input <- fmt.Sprintf("s %d", kittenIndex)
				}
				cfg.Input <- fmt.Sprintf("s %d", index)
				cfg.Input <- fmt.Sprintf("s farmer")
			}
			if isKitten[index] && iter < 100 {
				if action.Count == 0 {
					cfg.Input <- fmt.Sprintf("s %d", index)
				}
			} else if action.Type == "Job" {
				if action.Count == minJobCount {
					cfg.Input <- fmt.Sprintf("s %d", index)
				}
			} else {
				cfg.Input <- fmt.Sprintf("111 s %d", index)
			}
		}
		time.Sleep(time.Second * time.Duration(cfg.SleepMS) / 1000.)
	}
	return nil
}

func solveRandom(cfg Config) error {
	cfg.Input <- "hc"
	cfg.Input <- "hm"
	for {
		cfg.Waiting <- true
		<-cfg.Refreshed
		options := []int{}
		for i, a := range cfg.LastData.Actions {
			if a.IsLocked || a.IsHidden || a.IsOverCap {
				continue
			}
			options = append(options, i)
		}
		for _, index := range cfg.PermFn(len(options)) {
			cfg.Input <- fmt.Sprintf("s %d", options[index])
		}
		time.Sleep(time.Second * time.Duration(cfg.SleepMS) / 1000.)
	}
	return nil
}

func solveFixed(cfg Config) error {
	precmds := []string{
		"h Charon", "h Umbra", "h Yarn", "h Helios", "h Cath", "h Redmoon", "h Dune", "h Piscine", "h Termogus",
		"h Spring", "h Summer", "h Autumn", "h Winter",
		"h day_of_year",
	}
	cmds := []string{
		// catnip
		"10 s Gather catnip", "Catnip Field", "m Catnip Field",

		// woodcutter
		"5 s Refine catnip", "Hut", "s woodcutter",

		// scholar
		"s Library", "s scholar", "m Library",

		// farmer
		"m Calendar", "m Agriculture", "s Hut", "2 m farmer",
		"m Barn", "m Catnip Field", "m Library",

		// hunter
		"m Archery", "s Hut", "s hunter", "s farmer",

		// unicorn
		"m Animal Husbandry", "m Pasture",
		"40 s Send hunters", "Unic. Pasture", "10 s Unic. Pasture",

		"hc",

		// miner
		"m Mining", "s Mine", "s Hut", "s miner", "s farmer", "m Mine",
		"m Workshop", "m Mineral Hoes", "m Mineral Axe", "m Bolas",

		// iron
		"m Metal Working", "m Smelter",
		"s Hut", "s woodcutter", "s farmer",
		"Active Smelter",
		"m Iron Hoes", "m Iron Axe",
		"m Expanded Barns",
		"m Barn", "m Catnip Field", "m Library", "m Mine", "m Workshop", "m Smelter", "m Pasture",
		"m Hunting Armour",

		"m Civil Service", "m Mathematics", "m Academy", "m Celestial Mechanics",

		// craft
		"m Construction", "m Catnip Enrichment", "m Composite Bow",
		"m Reinforced Barns", "33 s Warehouse",
		"m Barn", "m Catnip Field", "m Library", "m Mine", "m Workshop", "m Smelter", "m Pasture", "m Academy", "m Lumber Mill",
		"m Reinforced Saw",

		"m Engineering", "m Aqueduct",

		// gold
		"m Currency", "m Gold Ore", "m Tradepost",

		"h Send hunters", "h Lizards", "h Griffins",

		// culture
		"m Writing", "m Register", "10 s Sharks", "m Amphitheatre",

		"hm", "h Nagas",

		"m Philosophy", "20 s Sharks", "m Temple",

		"m Steel", "m Coal Furnace", "m Deep Mining", "m Steel Axe", "m Steel Armour", "m High Pressure Engine",

		"m Reinforced Warehouses", "m Ironwood Huts", "30 s Sharks", "m High Pressure Engine",
		"m Mine", "m Workshop", "m Aqueduct", "m Lumber Mill", "m Tradepost",

		// manuscript
		"m Machinery", "m Crossbow", "m Printing Press", "m Workshop Automation",
		"2 s Hut", "2 s woodcutter", "2 s farmer",
		"30 s Sharks", "2 s Steamworks", "Active Steamworks",

		// priest
		"m Theology",
		"s Hut", "s priest", "s farmer",
		"m Amphitheatre", "m Temple",
		"m Golden Spire", "m Solar Chant", "m Scholasticism",
		"m Sun Altar", "m Stained Glass",
		"m Academy",
		//*/
	}
	for _, cmd := range precmds {
		cfg.Input <- cmd
	}
	for _, cmd := range cmds {
		cfg.Input <- cmd
		time.Sleep(time.Second * time.Duration(cfg.SleepMS) / 1000.)
	}
	return nil
}
