package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/kssilveira/idle-game-engine/game"
	"github.com/kssilveira/idle-game-engine/kittens"
	"github.com/kssilveira/idle-game-engine/textui"
	"github.com/kssilveira/idle-game-engine/ui"
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

	g, err := newGame(now)
	if err != nil {
		return err
	}

	input := make(chan string)
	output := make(chan *ui.Data)
	go g.Run(now, input, output)

	go func() {
		if err := handleInput(g, input); err != nil {
			log.Fatal(err)
		}
	}()
	var last string
	waiting := make(chan bool)
	refreshed := make(chan bool)
	go handleOutput(output, &last, waiting, refreshed)

	http.HandleFunc("/", handleHTTP(&last, input, *auto, *autoSleepMS, waiting, refreshed))
	return http.ListenAndServe(":8080", nil)
}

func newGame(now game.Now) (*game.Game, error) {
	g := kittens.NewGame(now)
	if err := updateResources(g, *resourceMap); err != nil {
		return nil, err
	}
	if err := g.Validate(); err != nil {
		return nil, err
	}
	return g, nil
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

func handleInput(g *game.Game, input game.Input) error {
	if *auto {
		if err := kittens.Solve(g, input, *autoSleepMS); err != nil {
			return err
		}
	}
	for {
		var got string
		fmt.Scanln(&got)
		input <- got
	}
	return nil
}

func handleOutput(output game.Output, last *string, waiting chan bool, refreshed chan bool) {
	logger := log.New(os.Stdout, "" /* prefix */, 0 /* flags */)
	separator := "\033[H\033[2J"
	for data := range output {
		textui.Show(logger, separator, data, false /* isHTML */, true /* showActionNumber */)
		var buf bytes.Buffer
		buflogger := log.New(&buf, "" /* prefix */, 0 /* flags */)
		textui.Show(buflogger, "" /* separator */, data, true /* isHTML */, true /* showActionNumber */)
		*last = buf.String()
		select {
		case <-waiting:
			refreshed <- true
		default:
		}
	}
}

func handleHTTP(last *string, input game.Input, auto bool, autoSleepMS int, waiting chan bool, refreshed chan bool) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if !auto {
			path := strings.TrimPrefix(r.URL.Path, "/")
			if path != "" && path != "favicon.ico" {
				input <- path
				waiting <- true
				<-refreshed
			}
		}
		fmt.Fprintf(w, `
<html>
<head>
  <meta http-equiv='refresh' content='%f; url=/'/>
  <style>
  body {
    font-family: monospace;
  }
  </style>
</head>
<body>
%s
</body>
</html>
`, float64(autoSleepMS)/1000, strings.Replace(*last, "\n", "<br>\n", -1))
	}
}
