package solve

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/kssilveira/idle-game-engine/game"
)

func Solve(g *game.Game, input chan string, sleepMS int) error {
	input <- "h"
	cmds := []string{
		"Gather catnip", "Catnip Field", "m Catnip Field",
		"s Refine catnip", "Hut", "s woodcutter",
		"s Library", "s scholar", "m Library",
		"s Calendar", "s Agriculture", "s Hut", "s farmer", "s farmer",
		"m Barn", "m Catnip Field", "m Library",
		"s Archery", "s Hut", "s hunter", "s farmer",
		"s Animal Husbandry", "m Pasture",

		"40 s Send hunters", "Unic. Pasture", "10 s Unic. Pasture",

		"s Mining", "s Mine", "s Hut", "s miner", "s farmer", "m Mine",
		"m Workshop", "s Mineral Hoes", "s Mineral Axe", "s Bolas",
		"s Metal Working",

		"s Smelter", "s Hut", "s woodcutter", "s farmer", "Active Smelter",
		"s Iron Hoes", "s Iron Axe",

		"s Expanded Barns",
		"m Barn", "m Catnip Field", "m Library", "m Pasture", "m Mine", "m Workshop", "m Smelter",
		"s Hunting Armour",

		"s Civil Service", "s Mathematics",
		"m Academy",
		"s Celestial Mechanics",

		"s Construction", "s Catnip Enrichment", "s Composite Bow",
		"m Reinforced Barns",
		"11 s Warehouse",
		"m Barn", "m Catnip Field", "m Library", "m Pasture", "m Mine", "m Workshop", "m Smelter",
		"m Academy", "m Lumber Mill",
		"s Reinforced Saw",

		"s Engineering", "m Aqueduct",
		"s Currency", "s Gold Ore", "m Tradepost",

		"11 s Sharks", "m Ziggurat",
		"s Hunting Armour",

		"s Writing", "m Amphitheatre",
		"s Register",

		"s Philosophy", "100 s Sharks", "m Temple", "m Amphitheatre",

		"s Steel", "s Coal Furnace", "s Deep Mining", "s Steel Axe", "s Steel Armour", "s High Pressure Engine",

		"s Reinforced Warehouses",
		"m Barn", "m Catnip Field", "m Library", "m Pasture", "m Mine", "m Workshop", "m Smelter",
		"m Academy", "m Lumber Mill", "m Aqueduct", "m Tradepost",

		"s Machinery", "s Crossbow", "s Printing Press", "s Workshop Automation",
		"10 s Steamworks", "Active Steamworks", "m Temple",

		"s Theology", "s Hut", "s priest", "s farmer",
		"m Golden Spire", "m Solar Chant", "m Scholasticism",
		"10 s Praise the sun!",

		"s Astronomy", "m Observatory",

		"s Navigation",
		"3 s Hut", "3 s woodcutter", "3 s farmer",

		"11 s Harbour",
		"m Barn", "m Catnip Field", "m Library", "m Pasture", "m Mine", "m Workshop", "m Smelter", "m Academy", "m Lumber Mill", "m Aqueduct", "m Tradepost", "m Temple", "m Observatory", "m Sun Altar", "m Stained Glass", "m Golden Spire", "m Solar Chant", "m Scholasticism",
		"s Ironwood Huts", "s Solar Revolution",

		"s Architecture",
		"5 s Hut", "5 s woodcutter", "5 s farmer",
		"s Hut", "s miner", "s farmer",
		"13 s Hut", "13 s hunter", "13 s farmer",
		"5 Active Smelter",
		"m Mint", "Active Mint",
		"m Temple", "m Sun Altar", "m Stained Glass", "m Golden Spire", "m Solar Chant", "m Scholasticism", "m Basilica",

		"11 s Unicorn Tomb", "11 s Ivory Tower", "11 s Ivory Citadel", "11 s Sky Palace", "11 s Unicorn Utopia", "11 s Sunspire",
		"m Barn", "m Catnip Field", "m Library", "m Pasture", "m Mine",
		"m Workshop", "m Smelter", "m Academy", "m Lumber Mill", "m Aqueduct", "m Tradepost", "m Observatory", "m Sun Altar", "m Stained Glass", "m Golden Spire", "m Solar Chant", "m Scholasticism", "m Mint", "m Temple",
		"s Unicorn Utopia", "6 s Sunspire",
		"m Basilica",

		"2 s Physics", "s Steel Saw", "s Caravanserai", "s Pyrolysis", "3 s Pneumatic Press",

		"2 s Chemistry", "22 s Oil Well",
		"30 s Hut", "30 s miner", "30 s farmer",
		"100 s Zebras", "s Calciner", "Active Calciner", "m Calciner",
	}
	for _, cmd := range cmds {
		if err := ToInput(g, cmd, input); err != nil {
			return err
		}
		time.Sleep(time.Second * time.Duration(sleepMS) / 1000.)
	}
	return nil
}

func ToInput(g *game.Game, cmd string, input chan string) error {
	words := strings.Split(cmd, " ")
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
	cmd = strings.Join(words, " ")
	if !g.HasAction(cmd) {
		return fmt.Errorf("invalid action %s", cmd)
	}
	for i := 0; i < count; i++ {
		input <- fmt.Sprintf("%s%d", prefix, g.GetActionIndex(cmd))
	}
	return nil
}
