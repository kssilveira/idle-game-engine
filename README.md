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

- move trade before single use actions
- rename types
  - Village to Job and Management
  - Workshop to Upgrade

- Bonfire Hut
  - catpower cap 75
- Bonfire Workshop
  - craft 6%
- Science Civil Service
  - unlocked by Animal Husbandry
  - science 1500
- Science Mathematics
  - unlocked by Animal Husbandry
  - science 1000
- Science Construction
  - unlocked by Animal Husbandry
  - science 1300
- Science Currency
  - unlocked by Civil Service
  - science 2200
  - TODO
- Bonfire Academy
  - unlocked by Mathematics
  - wood 50 57.5 rate 1.15
  - minerals 70 80.5
  - science 100 115
  - science bonus 20%
  - science cap 500
- Upgrade Celestial Mechanics
  - unlocked by Mathematics
  - science 250
- Science Engineering
  - unlocked by Construction
  - science 1500
  - TODO
- Bonfire Warehouse
  - unlocked by Construction
  - beam 1.5 TODO
  - slab 2
  - wood cap 262.5
  - minerals cap 350
  - iron cap 43.75
- Bonfire Log House
  - wood 200 TODO
  - minerals 250
  - catpower cap 50
  - kittens cap 1
- Upgrade Reinforced Saw
  - unlocked by Construction
  - iron 1000
  - science 2500
  - lumber mill 20%
- Upgrade Composite Bow
  - unlocked by Construction
  - wood 200
  - iron 100
  - science 500
  - hunter 50%
- Upgrade Catnip Enrichment
  - unlocked by Construction
  - catnip 5000
  - science 500
  - catnip refine 100%
- Craft beam
  - unlocked by Construction
  - wood 175
- Craft slab
  - unlocked by Construction
  - minerals 250
- Craft plate
  - unlocked by Construction
  - iron 125
- Craft gear
  - unlocked by Construction
  - steel 15
- Craft scaffold
  - unlocked by Construction
  - beam 50
- Craft manuscript
  - unlocked by Construction
  - culture 400
  - parchment 25
- Craft megalith
  - unlocked by Construction
  - beam 25
  - slab 50
  - plate 5
- TODO
- maybe
  - rX: revert action X
    - sell building
    - unassign worker
    - deactivate building
