package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/kssilveira/idle-game-engine/game"
	"github.com/kssilveira/idle-game-engine/kittens"
)

var (
	auto        = flag.Bool("auto", false, "automatically trigger all actions")
	autoSleepMS = flag.Int("auto_sleep_ms", 1000, "sleep between auto actions")
	resourceMap = flag.String("resource_map", "", "map of resource quantities, e.g. 'catnip:1,Catnip Field:2,wood:3")
)

func main() {
	flag.Parse()
	if err := all(); err != nil {
		log.Fatal(err)
	}
}

func all() error {
	now := func() time.Time { return time.Now() }
	g := kittens.NewGame(now)
	if err := updateResources(g, *resourceMap); err != nil {
		return err
	}
	logger := log.New(os.Stdout, "", 0 /* flags */)
	input := make(chan string)
	go func() {
		if *auto {
			kittens.Solve(input, *autoSleepMS)
			return
		}
		for {
			var got string
			fmt.Scanln(&got)
			input <- got
		}
	}()
	separator := "\033[H\033[2J"
	if err := g.Validate(); err != nil {
		return err
	}
	g.Run(logger, separator, input, now)
	return nil
}

func updateResources(g *game.Game, resourceMap string) error {
	if resourceMap == "" {
		return nil
	}
	for _, one := range strings.Split(strings.TrimSpace(resourceMap), ",") {
		words := strings.Split(strings.TrimSpace(one), ":")
		if len(words) != 2 {
			return fmt.Errorf("--resource_map has invalid value '%s'", one)
		}
		if !g.HasResource(words[0]) {
			return fmt.Errorf("--resource_map has invalid resource '%s'", words[0])
		}
		value, err := strconv.ParseFloat(words[1], 64)
		if err != nil {
			return fmt.Errorf("--resource_map has invalid value '%s' err %v", words[0], err)
		}
		r := g.GetResource(words[0])
		r.Quantity = value
		if r.Capacity != -1 && r.Capacity < value {
			r.Capacity = value
		}
	}
	return nil
}
