package solve

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/kssilveira/idle-game-engine/game"
)

func Solve(g *game.Game, input chan string, sleepMS int) error {
	precmds := []string{
		"h Charon", "h Umbra", "h Yarn", "h Helios", "h Cath", "h Redmoon", "h Dune", "h Piscine", "h Termogus",
		"h Spring", "h Summer", "h Autumn", "h Winter",
		"h day_of_year",
	}
	cmds := []string{
		// catnip
		"10 Gather catnip", "Catnip Field", "m Catnip Field",

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
		"s Sharks", "2 s Steamworks", "Active Steamworks",

		// priest
		"m Theology", "s Hut", "s priest", "s farmer",
		"m Amphitheatre", "m Temple",
		"m Golden Spire", "m Solar Chant", "m Scholasticism",
		"m Sun Altar", "m Stained Glass",
		"m Academy",
		//*/
	}
	for _, cmd := range precmds {
		if err := ToInput(g, cmd, input); err != nil {
			return err
		}
	}
	for _, cmd := range cmds {
		if err := ToInput(g, cmd, input); err != nil {
			return err
		}
		time.Sleep(time.Second * time.Duration(sleepMS) / 1000.)
	}
	return nil
}

func ToInput(g *game.Game, in string, input chan string) error {
	words := strings.Split(in, " ")
	prefix := ""
	count := 1
	if len(words) > 0 {
		cnt, err := strconv.Atoi(words[0])
		if err == nil {
			count = cnt
			words = words[1:]
		}
	}
	if len(words) > 0 {
		if len(words[0]) == 1 {
			prefix = words[0]
			words = words[1:]
		}
	}
	cmd := strings.Join(words, " ")
	if g.HasAction(cmd) {
		cmd = fmt.Sprintf("%d", g.GetActionIndex(cmd))
	} else if !g.HasResource(cmd) && in != "hc" && in != "hm" {
		return fmt.Errorf("invalid arg %s", cmd)
	}
	for i := 0; i < count; i++ {
		input <- fmt.Sprintf("%s%s", prefix, cmd)
	}
	return nil
}
