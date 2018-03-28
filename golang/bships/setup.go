package bships

import (
	"fmt"
)

// Facing indicates the direction of the ship on the grid, from its starting coordinate, e.g. from top to bottom.
type Facing int8

const (
	topToBottom Facing = iota
	leftToRight Facing = iota
)

// GameConfig holds details that don't change throughout the game, such as the size of the grid.
type GameConfig struct {
	GridWidth  int
	GridHeight int
	ShipTypes  map[ShipType]int
}

// Coord represents a single coordinate on the grid, e.g. {row: 'A', column: 3.
type Coord struct {
	row    rune
	column int
}

// ShipType represents a type of ship and its properties, such as its size.
type ShipType struct {
	name string
	len  int
}

// Ship represents an individual ship and its location.
type Ship struct {
	start    Coord
	dir      Facing
	shipType ShipType
}

var gridStart = Coord{'A', 1}

// NewGame initialises and returns a new game, set up with the ships provided at their specified locations.
func NewGame(cfg GameConfig, playerName string, ships ...Ship) (game *Game, err error) {

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Error setting up new game: %v", r)
		}
	}()

	player := createPlayer(&cfg, playerName, ships...)

	return &Game{cfg, map[string]Player{playerName: player}}, nil
}

func createPlayer(cfg *GameConfig, playerName string, ships ...Ship) Player {
	shipsPlotted := map[ShipType]int{}

	shipCoords := map[Coord]*Ship{}
	hitsByShip := map[*Ship]int{}

	for _, ship := range ships {
		shipsPlotted[ship.shipType]++

		hitsByShip[&ship] = 0
		plotCoords(cfg, &ship, shipCoords)
	}

	if cfg.ShipTypes != nil {
		validateShipTypes(cfg.ShipTypes, shipsPlotted)
	}

	return Player{remaining: shipCoords, hits: map[Coord]*Ship{}, hitsByShip: hitsByShip}
}

// NewDefaultGame sets up a game using a standard configuration for the grid and no restriction on ship types.
func NewDefaultGame(ships ...Ship) (game *Game, err error) {

	cfg := GameConfig{GridWidth: 8, GridHeight: 8}

	return NewGame(cfg, "player1", ships...)
}

func plotCoords(cfg *GameConfig, ship *Ship, shipCoords map[Coord]*Ship) {

	for i := 0; i < ship.shipType.len; i++ {

		pos := getCoord(ship, i)
		validateCoord(cfg, pos)

		if _, dup := shipCoords[pos]; dup {
			panic(fmt.Sprintf("Unable to place ship at %v, as ship already at this coordinate", pos))
		}

		shipCoords[pos] = ship
	}
}

func getCoord(ship *Ship, offset int) Coord {

	switch ship.dir {
	case topToBottom:
		return Coord{rune(int(ship.start.row) + offset), ship.start.column}
	case leftToRight:
		return Coord{ship.start.row, ship.start.column + offset}
	default:
		panic(fmt.Sprintf("Unreachable condidtion in switch on ship.dir %v", ship.dir))
	}
}

// +1 as ship placement is inclusive, i.e. a ship includes its starting position in its len
func validateCoord(cfg *GameConfig, coord Coord) {

	if coord.row < gridStart.row {
		panic(fmt.Sprintf("Ship exceeds top boundary at coordinate %v", coord))
	}

	if int(coord.row)-int(gridStart.row)+1 > cfg.GridWidth {
		panic(fmt.Errorf("Ship exceeds lower boundary at coordinate %v", coord))
	}

	if coord.column < 1 {
		panic(fmt.Errorf("Ship %v exceeds left-hand boundary", coord))
	}

	if coord.column-gridStart.column+1 > cfg.GridHeight {
		panic(fmt.Errorf("Ship %v exceeds right-hand boundary", coord))
	}
}

func validateShipTypes(types1, types2 map[ShipType]int) {

	errMsg := "Ship types and numbers in game config do not match ships supplied for game"

	if len(types1) != len(types2) {
		panic(errMsg)
	}

	for k, v := range types1 {

		v2, found := types2[k]

		if !found || v != v2 {
			panic(errMsg)
		}
	}
}
