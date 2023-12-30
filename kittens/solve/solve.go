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
		"m Calendar", "m Agriculture", "s Hut", "m farmer",
		"m Barn", "m Catnip Field", "m Library",
		"m Archery", "s Hut", "s hunter", "s farmer",
		"m Animal Husbandry", "m Pasture",

		"40 s Send hunters", "Unic. Pasture", "10 s Unic. Pasture",

		"m Mining", "s Mine", "s Hut", "s miner", "s farmer", "m Mine",
		"m Workshop", "m Mineral Hoes", "m Mineral Axe", "m Bolas",
		"m Metal Working",

		"m Smelter", "s Hut", "s woodcutter", "s farmer", "Active Smelter",
		"m Iron Hoes", "m Iron Axe",

		"m Expanded Barns",
		"m Barn", "m Catnip Field", "m Library", "m Pasture", "m Mine", "m Workshop", "m Smelter",
		"m Hunting Armour",

		"m Civil Service", "m Mathematics",
		"m Academy",
		"m Celestial Mechanics",

		"m Construction", "m Catnip Enrichment", "m Composite Bow",
		"m Reinforced Barns",
		"11 s Warehouse",
		"m Barn", "m Catnip Field", "m Library", "m Pasture", "m Mine", "m Workshop", "m Smelter",
		"m Academy", "m Lumber Mill",
		"m Reinforced Saw",

		"m Engineering", "m Aqueduct",
		"m Currency", "m Gold Ore", "m Tradepost",

		"11 s Sharks", "m Ziggurat",
		"m Hunting Armour",

		"m Writing", "m Amphitheatre",
		"m Register",

		"m Philosophy", "100 s Sharks", "m Temple", "m Amphitheatre",

		"m Steel", "m Coal Furnace", "m Deep Mining", "m Steel Axe", "m Steel Armour", "m High Pressure Engine",

		"m Reinforced Warehouses",
		"m Barn", "m Catnip Field", "m Library", "m Pasture", "m Mine", "m Workshop", "m Smelter",
		"m Academy", "m Lumber Mill", "m Aqueduct", "m Tradepost",

		"m Machinery", "m Crossbow", "m Printing Press", "m Workshop Automation",
		"10 s Steamworks", "Active Steamworks", "m Temple",

		"m Theology", "s Hut", "s priest", "s farmer",
		"m Golden Spire", "m Solar Chant", "m Scholasticism",

		"m Astronomy", "m Observatory",

		"m Navigation",
		"3 s Hut", "3 s woodcutter", "3 s farmer",

		"11 s Harbour",
		"m Barn", "m Catnip Field", "m Library", "m Pasture", "m Mine", "m Workshop", "m Smelter", "m Academy", "m Lumber Mill", "m Aqueduct", "m Tradepost", "m Temple", "m Observatory", "m Sun Altar", "m Stained Glass", "m Golden Spire", "m Solar Chant", "m Scholasticism",
		"m Ironwood Huts", "m Solar Revolution",

		"m Architecture",
		"5 s Hut", "5 s woodcutter", "5 s farmer",
		"s Hut", "s miner", "s farmer",
		"13 s Hut", "13 s hunter", "13 s farmer",
		"5 Active Smelter",
		"m Mint", "Active Mint",
		"m Temple", "m Sun Altar", "m Stained Glass", "m Golden Spire", "m Solar Chant", "m Scholasticism", "m Basilica",

		"11 s Unicorn Tomb", "11 s Ivory Tower", "11 s Ivory Citadel", "11 s Sky Palace", "11 s Unicorn Utopia", "11 s Sunspire",
		"m Barn", "m Catnip Field", "m Library", "m Pasture", "m Mine", "m Workshop", "m Smelter", "m Academy", "m Lumber Mill", "m Aqueduct", "m Tradepost", "m Observatory", "m Sun Altar", "m Stained Glass", "m Golden Spire", "m Solar Chant", "m Scholasticism", "m Mint", "m Temple",
		"s Unicorn Utopia", "6 s Sunspire",
		"m Basilica",

		"2 s Physics", "m Steel Saw", "m Caravanserai", "m Pyrolysis", "m Pneumatic Press",

		"2 s Chemistry", "22 s Oil Well",
		"30 s Hut", "30 s miner", "30 s farmer",
		"100 s Zebras", "m Calciner", "Active Calciner", "m Calciner",
		"m Mine", "m Amphitheatre", "m Sunspire",
		"m Titanium Saw", "m Titanium Axe", "m Alloy Axe", "m Titanium Barns", "m Alloy Barns", "m Alloy Warehouses", "m Expanded Cargo", "m Silos", "m Alloy Armour", "m Astrolabe", "m Titanium Reflectors", "m Alloy Saw", "m Titanium Warehouses",

		"m Barn", "m Catnip Field", "m Library", "m Pasture", "m Mine", "m Workshop", "m Smelter", "m Academy", "m Lumber Mill", "m Aqueduct", "m Tradepost", "m Observatory", "m Sun Altar", "m Stained Glass", "m Golden Spire", "m Solar Chant", "m Scholasticism", "m Mint", "m Temple", "m Academy", "m Amphitheatre", "m Sunspire", "m Golden Spire", "m Basilica", "m Templars",
		"m Apocrypha",

		"m Acoustics", "m Chapel",

		"m Geology",
		"s Hut", "s geologist", "s farmer",
		"m Geodesy",
		"11 s Quarry",

		"m Electricity", "13 s Magneto", "Active Magneto",
		"m Industrialization",
		"m Barges", "m Advanced Automation", "m Logistics",
		"2 s Hut", "2 s woodcutter", "2 s farmer",
		"2 Active Smelter",

		"m Biology", "m Bio Lab", "Active Bio Lab",

		"m Barn", "m Catnip Field", "m Library", "m Pasture", "m Mine", "m Workshop", "m Smelter", "m Academy", "m Lumber Mill", "m Aqueduct", "m Tradepost", "m Observatory", "m Sun Altar", "m Stained Glass", "m Golden Spire", "m Solar Chant", "m Scholasticism", "m Mint", "m Temple", "m Academy", "m Amphitheatre", "m Sunspire", "m Golden Spire", "m Basilica", "m Templars",

		"m Drama and Poetry", "s Brewery", "s Festival",

		"m Mechanization", "m Concrete Pillars", "m Pumpjack", "m Concrete Warehouses", "m Concrete Barns",

		"m Barn", "m Catnip Field", "m Library", "m Pasture", "m Mine", "m Workshop", "m Smelter", "m Academy", "m Lumber Mill", "m Aqueduct", "m Tradepost", "m Observatory", "m Sun Altar", "m Stained Glass", "m Golden Spire", "m Solar Chant", "m Scholasticism", "m Mint", "m Temple", "m Academy", "m Amphitheatre", "m Sunspire", "m Golden Spire", "m Basilica", "m Templars", "m Library", "m Mine", "m Workshop", "m Pasture", "m Lumber Mill", "m Factory", "m Chapel", "m Bio Lab",

		"m Transcendence", "11 s Transcendence Level",

		"m Combustion", "m Tradepost", "m Offset Press", "m Fuel Injector", "m Oil Refinery",

		"m Metallurgy", "m Mining Drill", "m Electrolytic Smelting", "m Oxidation",
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
