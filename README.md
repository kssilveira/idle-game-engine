# idle-game-engine

Game engine for idle games like kittens game.

## Code

- [game.go](game/game.go) general game engine
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

- web server and web UI
- configurable game rules at runtime
- competition of best solutions
  - least actions 
  - least skipped time
- competition of bots

Features from kittens game:

- catnip
  - unlock Catnip Field
- wood
  - unlock Hut, Library
- Hut
  - unlock woodcutter
- Library
  - cost wood 25 28.75 33.06 rate 1.15
  - science production 10%
  - science cap 250
  - unlock scholar, Calendar
- scholar
  - science 0.175
- science cap 250
- Calendar
  - cost science 30
  - show year and day
    - day 2 seconds
    - season 100 days
    - year 400 days
  - unlock Agriculture
- Agriculture
  - cost science 100
  - unlock Archery, Mining, Barn, farmer
- Archery
  - cost science 300
  - TODO
- Mining
  - cost science 500
  - TODO
- Barn
  - cost wood 50 87.50 TODO
  - catnip cap 500
  - wood cap 200
- farmer
  - catnip 5
- TODO
- maybe
  - sell building for half price
