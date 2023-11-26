package main

import "fmt"
import "log"
import "os"
import "github.com/kssilveira/idle-game-engine/kittens"

func main() {
	logger := log.New(os.Stdout, "", 0 /* flags */)
	kittens.Run(logger, func() int {
		input := -1
		fmt.Printf("> ")
		fmt.Scanf("%d", &input)
		return input
	})
}

/*

actions

- "Gather catnip" add 1 catnip
- "Refine catnip" cost 100 catnip add 1 wood

buildings

- count X resource X cost X time Xh (over cap)
  - sell for half price
- "Catnip Field" cost 10 catnip add 0.63 catnip/s
	- cost 10 11.20 12.54 14.05 15.74 17.62
	- add
		- winter -75% 0.16/s
		- spring +50% 0.94/s

- "Hut" cost 5 wood add 2 kitten
  - cost 10 12.5 31.25 78.13 195.31 488.28

resources

- count X cap X rate X/s to cap Xh
- catnip cap 5000
- wood cap 200
- kitten cap 0
  - eat 4.25 catnip/s

jobs

woodcuttter 0.09 wood/s
  - reduced by happiness

happiness
- kittens 5 => 100% 98% 96% 94% 92% 90%

*/
