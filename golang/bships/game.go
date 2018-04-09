package bships

import (
	"fmt"
)

// Game represents a game of Batteships.
type Game struct {
	config  GameConfig
	Players map[string]*Player
}

// GameConfig holds details that don't change throughout the game, such as the size of the grid.
type GameConfig struct {
	GridWidth  int
	GridHeight int
	shipTypes  map[ShipType]int
}

// ShipType represents a type of ship and its properties, such as its length.
type ShipType struct {
	Name string
	Len  int
}

// Ship represents an individual ship and its properties.
type Ship struct {
	Start    Coord
	Dir      Facing
	ShipType ShipType
}

// Coord represents a single coordinate on a player's grid, e.g. {row: 'A', column: 3}.
type Coord struct {
	Row    rune
	Column int
}

// Facing indicates the direction of the ship on the grid, from its starting coordinate, e.g. from top to bottom.
type Facing int8

const (
	// TopToBottom represents a ship positioned vertically, from top to bottom
	TopToBottom Facing = iota

	// LeftToRight represents a ship positioned horizontally, from left to right
	LeftToRight Facing = iota
)

var gridStart = Coord{'A', 1}

// NewGame initialises and returns a new game.
func NewGame(cfg GameConfig) (game *Game) {

	return &Game{cfg, map[string]*Player{}}
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

	validateShipTypes(game.config.shipTypes, shipTypes)
	validateCoords(&game.config, coords)

	game.Players[playerName] = &player

	return
}

func (ship *Ship) getCoord(offset int) Coord {

	switch ship.Dir {
	case TopToBottom:
		return Coord{rune(int(ship.Start.Row) + offset), ship.Start.Column}
	case LeftToRight:
		return Coord{ship.Start.Row, ship.Start.Column + offset}
	default:
		panic(fmt.Sprintf("Unreachable condidtion in switch on ship.dir %v", ship.Dir))
	}
}

// +1 as ship placement is inclusive, i.e. a ship includes its starting position in its len
func validateCoords(cfg *GameConfig, coords map[Coord]*Ship) {

	for coord := range coords {

		if coord.Row < gridStart.Row {
			panic(fmt.Sprintf("Ship exceeds top boundary at coordinate %v", coord))
		}

		if int(coord.Row)-int(gridStart.Row)+1 > cfg.GridWidth {
			panic(fmt.Errorf("Ship exceeds lower boundary at coordinate %v", coord))
		}

		if coord.Column < 1 {
			panic(fmt.Errorf("Ship %v exceeds left-hand boundary", coord))
		}

		if coord.Column-gridStart.Column+1 > cfg.GridHeight {
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
