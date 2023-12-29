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

		"11 s Sharks", "m Ziggurat", "Sacrifice Unicorns",
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
