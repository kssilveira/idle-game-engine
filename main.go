package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/kssilveira/idle-game-engine/kittens"
)

var (
	auto        = flag.Bool("auto", false, "automatically trigger all actions")
	autoSleepMS = flag.Int("auto_sleep_ms", 1000, "sleep between auto actions")
)

func main() {
	flag.Parse()
	now := func() time.Time { return time.Now() }
	g := kittens.NewGame(now)
	g.GetResource("catnip").Quantity = 10
	logger := log.New(os.Stdout, "", 0 /* flags */)
	input := make(chan string)
	go func() {
		lastAuto := 0
		actions := []string{"1", "s1"}
		for {
			var got string
			if *auto {
				got = actions[lastAuto]
				lastAuto++
				lastAuto %= len(actions)
				time.Sleep(time.Second * time.Duration(*autoSleepMS) / 1000.)
			} else {
				fmt.Scanln(&got)
			}
			input <- got
		}
	}()
	separator := "\033[H\033[2J"
	if err := g.Validate(); err != nil {
		log.Fatal(err)
	}
	g.Run(logger, separator, input, now)
}
