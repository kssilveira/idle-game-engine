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

- create resource 'free kittens'
- maybe sell building for half price
- etc.
