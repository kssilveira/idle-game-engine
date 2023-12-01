package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/kssilveira/idle-game-engine/game"
	"github.com/kssilveira/idle-game-engine/textui"
	"github.com/kssilveira/idle-game-engine/ui"
)

func main() {
	g := game.NewGame(time.Now())
	g.AddResources([]game.Resource{{
		Name: "catnip", Quantity: 10, Capacity: 100,
		Producers: []game.Resource{{
			Name: "Catnip Field", ProductionFactor: 0.63,
		}},
	}, {
		Name: "Catnip Field", Capacity: -1,
	}})
	g.Actions = []game.Action{{
		Name: "Catnip Field",
		Costs: []game.Resource{{
			Name: "catnip", Quantity: 10, CostExponentBase: 1.12,
		}},
		Adds: []game.Resource{{
			Name: "Catnip Field", Quantity: 1,
		}},
	}}
	if err := g.Validate(); err != nil {
		log.Fatal(err)
	}

	now := func() time.Time { return time.Now() }
	input := make(chan string)
	output := make(chan *ui.Data)
	go g.Run(now, input, output)

	go func() {
		for {
			var got string
			fmt.Scanln(&got)
			input <- got
		}
	}()

	logger := log.New(os.Stdout, "", 0 /* flags */)
	separator := "\033[H\033[2J"
	for data := range output {
		textui.Show(logger, separator, data)
	}
}
