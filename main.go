package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/kssilveira/idle-game-engine/kittens"
)

func main() {
	logger := log.New(os.Stdout, "", 0 /* flags */)
	kittens.Run(logger, func() int {
		input := -1
		fmt.Printf("> ")
		fmt.Scanf("%d", &input)
		return input
	}, func() time.Time {
		return time.Now()
	})
}
