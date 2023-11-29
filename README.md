# idle-game-engine

Game engine for idle games like kittens game.

## Code

- [game.go](game/game.go) general game engine
- [ui.go](ui/ui.go) representation of the UI
- [textui.go](textui/textui.go) text UI
- [kittens.go](kittens/kittens.go) kittens game implemented using the game engine
- [solve.out](kittens/testdata/solve.out) text UI output for the kittens game solution
- [main.go](main.go) run kittens game

## Development

Run tests and build:

```
$ go test ./... && go build
```

Run interactive kittens game:

```
$ ./idle-game-engine
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
- Mining
  - cost science 500
- Barn
  - cost wood 50 87.50
  - catnip cap 500
  - wood cap 200
- farmer
  - catnip 5
- etc.
- maybe
  - sell building for half price
