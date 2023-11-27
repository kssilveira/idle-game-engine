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
	auto = flag.Bool("auto", false, "automatically trigger all actions")
)

func main() {
	flag.Parse()
	now := func() time.Time { return time.Now() }
	g := kittens.NewGame(now)
	logger := log.New(os.Stdout, "", 0 /* flags */)
	input := make(chan int)
	go func() {
		lastAuto := 0
		for {
			got := -1

			if *auto {
				got = lastAuto
				lastAuto++
				lastAuto %= len(g.Actions)
				time.Sleep(time.Second)
			} else {
				fmt.Scanf("%d", &got)
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
