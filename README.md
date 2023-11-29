# idle-game-engine

Game engine for idle games like kittens game.

## Code

- [game/game.go] general game engine
- [ui/ui.go] representation of the UI
- [textui/textui.go] text UI
- [kittens/kittens.go] kittens game implemented using the game engine
- [kittens/testdata/solve.out] text UI output for the kittens game solution
- [main.go] run kittens game

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

- woodcutter reduced by happiness
  - kittens 5 => 100% 98% 96% 94% 92% 90%
- maybe sell building for half price
- etc.
