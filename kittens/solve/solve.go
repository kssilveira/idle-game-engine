package solve

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/kssilveira/idle-game-engine/game"
)

func Solve(g *game.Game, input chan string, sleepMS int) error {
	cmds := []string{
		"Gather catnip", "Catnip Field", "m Catnip Field", "h Catnip Field", "h catnip",
		"h Charon", "h Umbra", "h Yarn", "h Helios", "h Cath", "h Redmoon", "h Dune", "h Piscine", "h Termogus",
		"h Autumn", "h Spring", "h Summer", "h Winter",
		"h day_of_year", "hc",

		"s Refine catnip", "Hut", "s woodcutter", "h wood",

		"s Library", "s scholar", "m Library", "h Library", "h science", "h woodcutter", "h scholar",

		"m Calendar", "m Agriculture", "s Hut", "2 m farmer", "h farmer",
		"m Barn",

		"m Archery", "s Hut", "s hunter", "s farmer", "h hunter", "h catpower", "h Lizards", "h Griffins",

		"m Animal Husbandry", "m Pasture", "h Pasture",

		"40 s Send hunters", "Unic. Pasture", "10 s Unic. Pasture", "h Unic. Pasture", "h Send hunters", "h unicorn",

		"m Mining", "s Mine", "s Hut", "s miner", "s farmer", "m Mine", "h Mine", "h mineral", "h miner",

		"m Workshop", "m Mineral Hoes", "m Mineral Axe", "m Bolas", "h Workshop",

		"m Metal Working", "m Smelter",
		"s Hut", "s woodcutter", "s farmer",
		"Active Smelter", "h Smelter", "h Active Smelter", "h iron",
		"m Iron Hoes", "m Iron Axe",
		"m Expanded Barns", "m Barn", "m Hunting Armour",

		"m Civil Service", "m Mathematics", "m Academy", "m Celestial Mechanics", "h Academy",

		"m Construction", "m Catnip Enrichment", "m Composite Bow",
		"m Reinforced Barns", "11 s Warehouse", "m Barn", "m Lumber Mill", "m Reinforced Saw", "h Lumber Mill", "h Warehouse",

		"m Engineering", "m Aqueduct", "h Aqueduct",

		"m Currency", "m Gold Ore", "m Tradepost", "h Tradepost",

		"m Writing", "m Amphitheatre", "m Register", "h Nagas",

		"m Philosophy",

		"m Steel", "m Coal Furnace", "m Deep Mining", "m Steel Axe", "m Steel Armour", "m High Pressure Engine",

		"m Reinforced Warehouses",

		"m Machinery", "m Crossbow", "m Printing Press", "m Workshop Automation",
		"s Hut", "s woodcutter", "s farmer",
		"11 s Sharks", "2 s Steamworks", "Active Steamworks",
		"m Library", "m Academy",

		"m Theology", "s Hut", "s priest", "s farmer",
		"m Amphitheatre", "m Temple",
		"m Golden Spire", "m Solar Chant", "m Scholasticism",
		"h gold", "h culture", "h faith", "h beam", "h slab", "h manuscript", "h Steamworks", "h Active Steamworks", "h Temple", "h priest", "h Sharks", "h coal",
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
	} else if !g.HasResource(cmd) && in != "hc" {
		return fmt.Errorf("invalid arg %s", cmd)
	}
	for i := 0; i < count; i++ {
		input <- fmt.Sprintf("%s%s", prefix, cmd)
	}
	return nil
}
