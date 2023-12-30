package solve

import (
	"fmt"
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
	IsSmart   bool
	SleepMS   int
}

func Solve(cfg Config) error {
	if cfg.IsSmart {
		return solveSmart(cfg)
	}
	return solve(cfg)
}

func solveSmart(cfg Config) error {
	cfg.Input <- "hc"
	cfg.Input <- "hm"

	for {
		cfg.Waiting <- true
		<-cfg.Refreshed
		for i, a := range cfg.LastData.Actions {
			if a.IsLocked || a.IsHidden || a.IsOverCap {
				continue
			}
			cfg.Input <- fmt.Sprintf("s %d", i)
		}
		time.Sleep(time.Second * time.Duration(cfg.SleepMS) / 1000.)
	}

	cfg.Input <- "999"
	return nil
}

func solve(cfg Config) error {
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
