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

## Ideas

General features:

- shared web server
- configurable game rules at runtime
- competition of best solutions
  - least actions 
  - least skipped time
- competition of bots

Features from kittens game:

- Animal Husbandry
  - unlocked by Archery
  - science 500
  - TODO
- hunter
  - unlocked by Archery
  - catpower 0.3
- catpower
  - cap 250
- Metal Working
  - unlocked by Mining
  - science 900
  - TODO
- Mine
  - unlocked by Mining
  - wood 100 115 132.25 rate 1.15
  - minerals 20%
- miner
  - unlocked by Mine
  - minerals 0.25
- minerals
  - cap 750
- Workshop
  - unlocked by miner
  - wood 100
  - minerals 400
  - craft 6%
  - TODO
- Send hunters
  - unlocked by hunter
  - catpower 100
  - random furs ivory TODO
- furs
  - demand 0.1 TODO
  - happiness TODO
- ivory
  - demand 0.07 TODO
  - happiness TODO
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
- TODO
- maybe
  - rX: revert action X
    - sell building
    - unassign worker
