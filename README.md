# idle-game-engine

Game engine for idle games like kittens game.

## Code

- [game.go](game/game.go) general game engine
- [data.go](data/data.go) representation of the data
- [ui.go](ui/ui.go) representation of the UI
- [textui.go](textui/textui.go) text UI
- [simple.go](examples/simple/simple.go) run simple game
- [kittens.go](kittens/kittens.go) kittens game
- [main.go](main.go) run kittens game
- [solve.out](kittens/testdata/solve.out) text UI output for kittens game solution
- [graph.svg](kittens/testdata/graph.svg) graph of kittens game dependencies

## Special Actions

- Time skip until a game action is available

## Dev

Run tests and build:

```
$ go test ./... && go build
```

## Run

Run interactive kittens game:

```
$ ./idle-game-engine
```

Use text UI or http://localhost:8080/.

Set starting resources:

```
$ ./idle-game-engine --resource_map='catnip:100,Catnip Field:1'
```

Watch kittens game solution:

```
$ ./idle-game-engine --auto
```

Faster solution:

```
./idle-game-engine --auto --auto_sleep_ms=100
```

## General Ideas

- shared web server
- configurable game rules at runtime
- competition of best solutions
  - least actions 
  - least skipped time
- competition of bots

## Kittens

Graph of current features:

![graph](kittens/testdata/graph.svg)

Nodes:

![graph nodes](kittens/testdata/graph_nodes.svg)

Edges:

![graph edges](kittens/testdata/graph_edges.svg)

More features from kittens game:

- Workshop
  - craft 6%
- Lizards
  - unlocked by hunter
  - catpower 50
  - gold 15
  - minerals 1000
  - wood 648 - 702
- Send explorers
  - unlocked by hunter
  - catpower 1000
  - TODO
- Mineral Hoes
  - unlocked by Workshop
  - minerals 275
  - science 100
  - farmer 50%
    - before 2 * 5 * 0.98 = 9.8
    - after 2 * 5 * 0.98 * (1 + 0.5) = 14.7
- Iron Hoes
  - unlocked by Workshop
  - iron 25
  - science 200
  - farmer 30%
- Mineral Axe
  - unlocked by Workshop
  - minerals 500
  - science 100
  - woodcutter 70%
- Iron Axe
  - unlocked by Workshop
  - iron 50
  - science 200
  - woodcutter 50%
- Expanded Barns
  - unlocked by Workshop
  - wood 1000
  - minerals 750
  - iron 50
  - science 500
  - barn 75%
- Reinforced Barns
  - unlocked by Workshop
  - iron 100
  - science 800
  - beam 25
  - slab 10
  - barn 80%
- Bolas
  - unlocked by Workshop
  - wood 50
  - minerals 250
  - science 1000
  - hunter 100%
- Smelter
  - unlocked by Metal Working
  - minerals 200 230 rate 1.15
  - wood -0.25/s
  - mineral -0.5/s
  - iron 0.1/s
- iron
  - cap 50
  - barn cap 50
- Hunting Armor
  - unlocked by Metal Working
  - iron 750
  - science 2000
  - hunter 200%
- Civil Service
  - unlocked by Animal Husbandry
  - science 1500
  - TODO
- Mathematics
  - unlocked by Animal Husbandry
  - science 1000
  - TODO
- Construction
  - unlocked by Animal Husbandry
  - science 1300
  - TODO
- Pasture
  - unlocked by Animal Husbandry
  - catnip 100 115 rate 1.15
  - wood 10 11.5
  - catnip demand -0.5%
- Unic. Pasture
  - unlocked by Animal Husbandry
  - unicorns 2 TODO
  - unicorns 0.005/s
  - capnip consumption -0.15%
- TODO
- maybe
  - rX: revert action X
    - sell building
    - unassign worker
