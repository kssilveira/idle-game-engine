package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/kssilveira/idle-game-engine/kittens"
)

func main() {
	now := func() time.Time { return time.Now() }
	g := kittens.NewGame(now)
	logger := log.New(os.Stdout, "", 0 /* flags */)
	input := make(chan int)
	go func() {
		for {
			got := -1
			fmt.Scanf("%d", &got)
			input <- got
		}
	}()
	separator := "\033[H\033[2J"
	if err := g.Validate(); err != nil {
		log.Fatal(err)
	}
	g.Run(logger, separator, input, now)
}
