package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/kssilveira/idle-game-engine/data"
	"github.com/kssilveira/idle-game-engine/game"
	"github.com/kssilveira/idle-game-engine/textui"
	"github.com/kssilveira/idle-game-engine/ui"
)

func main() {
	g := game.NewGame(time.Now())
	g.AddResources([]data.Resource{{
		Name: "catnip", Count: 10, Cap: 100,
		Producers: []data.Resource{{
			Name: "Catnip Field", Factor: 0.63,
		}},
	}, {
		Name: "Catnip Field", Cap: -1,
	}})
	g.AddActions([]data.Action{{
		Name: "Catnip Field",
		Costs: []data.Resource{{
			Name: "catnip", Count: 10, CostExponentBase: 1.12,
		}},
		Adds: []data.Resource{{
			Name: "Catnip Field", Count: 1,
		}},
	}})
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

	cfg := textui.Config{
		Logger:    log.New(os.Stdout, "", 0 /* flags */),
		Separator: "\033[H\033[2J",
	}
	for data := range output {
		textui.Show(cfg, data)
	}
}
