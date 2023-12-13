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
- Make all the inputs for an action

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

- remove duplication from kittens.NewGame
- change 'make' command to also make other resources (remove premature return)
- change production resource factor to production bonus
- create m command that skips and buys the max
- validate capacityproducers with startcapacity and without capacity
- change production factor to just factor
- change production bonus to bonus
- change solve commands to string
- fix test performance
- dont set quantity and startquantity
- split kittens file
- make functions and attributes private

- Bonfire building upgrades
- Bonfire effects for
  - Steamworks, Magneto, Reactor, Tradepost, Mint, Brewery, Chronosphere, Broadcast Tower
- Bonfire active buildings
- Job engineer
- Metaphysics

- https://wiki.kittensgame.com/en/home
  - https://wiki.kittensgame.com/en/game-tabs/workshop
    - TODO

- maybe
  - rX: revert action X
    - sell building
    - unassign worker
    - deactivate building
  - energy
