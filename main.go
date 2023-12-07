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
	graph        = flag.Bool("graph", false, "show graph")
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

	if *graph {
		return showGraph(g)
	}

	input := make(chan string)
	output := make(chan *ui.Data)
	go g.Run(now, input, output)

	go handleInput(input)
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

func handleInput(input game.Input) {
	if *auto {
		kittens.Solve(input, *autoSleepMS)
	}
	for {
		var got string
		fmt.Scanln(&got)
		input <- got
	}
}

func handleOutput(output game.Output, last *string, waiting chan bool, refreshed chan bool) {
	logger := log.New(os.Stdout, "" /* prefix */, 0 /* flags */)
	separator := "\033[H\033[2J"
	for data := range output {
		textui.Show(logger, separator, data, false /* isHTML */)
		var buf bytes.Buffer
		buflogger := log.New(&buf, "" /* prefix */, 0 /* flags */)
		textui.Show(buflogger, "" /* separator */, data, true /* isHTML */)
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

func showGraph(g *game.Game) error {
	fmt.Printf("digraph {\n")
	typeToShape := map[string]string{
		"Resource": "cylinder",
		"Bonfire": "box3d",
		"Village": "house",
		"Science": "diamond",
	}
	for _, r := range g.Resources {
		if r.Type == "Calendar" || r.Name == "gone kitten" {
			continue
		}
		fmt.Printf(`  "%s" [shape="%s"];` + "\n", r.Name, typeToShape[r.Type])
	}
	for _, a := range g.Actions {
		fmt.Printf(`  "%s" [shape="%s"];` + "\n", a.Name, typeToShape[a.Type])
	}
	for _, r := range g.Resources {
		last := ""
		for _, p := range r.Producers {
			if p.Name == "" || p.Name == "day" || p.Name == last {
				continue
			}
			last = p.Name
			if p.ProductionFactor < 0 {
				fmt.Printf(`  "%s" -> "%s" [color="red"];` + "\n", r.Name, p.Name)
			} else {
				fmt.Printf(`  "%s" -> "%s" [color="green"];` + "\n", p.Name, r.Name)
			}
		}
	}
	for _, a := range g.Actions {
		for _, c := range a.Costs {
			fmt.Printf(`  "%s" -> "%s" [color="orange"];` + "\n", c.Name, a.Name)
		}
		for _, add := range a.Adds {
			if a.Name == add.Name {
				continue
			}
			fmt.Printf(`  "%s" -> "%s" [color="limegreen"];` + "\n", a.Name, add.Name)
		}
		if a.UnlockedBy.Name != "" {
				fmt.Printf(`  "%s" -> "%s" [color="blue"];` + "\n", a.UnlockedBy.Name, a.Name)
		}
	}
	fmt.Printf(`
subgraph cluster_01 {
  label = "Arrows";
  node [shape=point, style=invis]
  n0 -> n1 [color="red" label="consumes"]
  n2 -> n3 [color="green" label="produces"]
  n4 -> n5 [color="orange" label="costs"]
  n6 -> n7 [color="limegreen" label="adds"]
  n8 -> n9 [color="blue" label="unlocks"]
}

subgraph cluster_02 {
  label = "Shapes";
  "Resource" [shape="cylinder"];
  "Bonfire" [shape="box3d"];
  "Village" [shape="house"];
  "Science" [shape="diamond"];
}
`)
	fmt.Printf("}\n")
	return nil
}
