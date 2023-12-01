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

	go handleInput(input)
	var last string
	go handleOutput(output, &last)

	http.HandleFunc("/", handleHTTP(&last, *autoSleepMS))
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

func handleInput(input game.Input) {
	if *auto {
		kittens.Solve(input, *autoSleepMS)
		return
	}
	for {
		var got string
		fmt.Scanln(&got)
		input <- got
	}
}

func handleOutput(output game.Output, last *string) {
	logger := log.New(os.Stdout, "" /* prefix */, 0 /* flags */)
	separator := "\033[H\033[2J"
	for data := range output {
		textui.Show(logger, separator, data)
		var buf bytes.Buffer
		buflogger := log.New(&buf, "" /* prefix */, 0 /* flags */)
		textui.Show(buflogger, "" /* separator */, data)
		*last = buf.String()
	}
}

func handleHTTP(last *string, autoSleepMS int) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `
<html>
<head>
  <meta http-equiv='refresh' content='%f'/>
</head>
<body>
  <pre>
%s
  </pre>
</body>
</html>
`, float64(autoSleepMS)/1000, *last)
	}
}
