package bships

import (
	"fmt"
)

// Game represents a game of Batteships.
type Game struct {
	config  GameConfig
	players map[string]Player
}

// GameConfig holds details that don't change throughout the game, such as the size of the grid.
type GameConfig struct {
	GridWidth  int
	GridHeight int
	ShipTypes  map[ShipType]int
}

// ShipType represents a type of ship and its properties, such as its length.
type ShipType struct {
	name string
	len  int
}

// Ship represents an individual ship and its properties.
type Ship struct {
	start    Coord
	dir      Facing
	shipType ShipType
}

// Coord represents a single coordinate on a player's grid, e.g. {row: 'A', column: 3}.
type Coord struct {
	row    rune
	column int
}

// Facing indicates the direction of the ship on the grid, from its starting coordinate, e.g. from top to bottom.
type Facing int8

const (
	topToBottom Facing = iota
	leftToRight Facing = iota
)

var gridStart = Coord{'A', 1}

// NewGame initialises and returns a new game.
func NewGame(cfg GameConfig) (game *Game) {

	return &Game{cfg, map[string]Player{}}
}

// AddPlayer adds a player and their grid of ships into the game.
func (game *Game) AddPlayer(playerName string, ships ...Ship) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Error adding player %v: %v", playerName, r)
		}
	}()

	player := Player{}
	shipTypes, coords := player.PlotShips(&game.config, ships...)

	validateShipTypes(game.config.ShipTypes, shipTypes)
	validateCoords(&game.config, coords)

	game.players[playerName] = player

	return
}

func (ship *Ship) getCoord(offset int) Coord {

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
func validateCoords(cfg *GameConfig, coords map[Coord]*Ship) {

	for coord := range coords {

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
}

func validateShipTypes(required, actual map[ShipType]int) {

	if required == nil {
		return
	}

	errMsg := "Ship types and numbers in GameConfig do not match ships supplied for Player"

	if len(required) != len(actual) {
		panic(errMsg)
	}

	for k, v := range required {

		v2, found := actual[k]

		if !found || v != v2 {
			panic(errMsg)
		}
	}
}
