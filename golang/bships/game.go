package bships

import (
	"errors"
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
func (game *Game) AddPlayer(playerName string, ships ...Ship) error {

	player := Player{}

	shipTypes, coords, err := player.PlotShips(&game.config, ships...)

	if err = validatePlottedShips(game, shipTypes, coords, err); err != nil {
		return fmt.Errorf("Error adding player %v when validating plotted ships: %v", playerName, err)
	}

	game.Players[playerName] = &player

	return nil
}

func validatePlottedShips(game *Game, shipTypes map[ShipType]int, coords map[Coord]*Ship, prevErr error) error {

	if prevErr != nil {
		return prevErr
	}

	if err := validateShipTypes(game.config.shipTypes, shipTypes); err != nil {
		return err
	}

	return validateCoords(&game.config, coords)
}

// +1 as ship placement is inclusive, i.e. a ship includes its starting position in its len
func validateCoords(cfg *GameConfig, coords map[Coord]*Ship) error {

	for coord := range coords {

		if coord.Row < gridStart.Row {
			return fmt.Errorf("Ship exceeds top boundary at coordinate %v", coord)
		}

		if int(coord.Row)-int(gridStart.Row)+1 > cfg.GridWidth {
			return fmt.Errorf("Ship exceeds lower boundary at coordinate %v", coord)
		}

		if coord.Column < 1 {
			return fmt.Errorf("Ship %v exceeds left-hand boundary", coord)
		}

		if coord.Column-gridStart.Column+1 > cfg.GridHeight {
			return fmt.Errorf("Ship %v exceeds right-hand boundary", coord)
		}
	}

	return nil
}

func validateShipTypes(required, actual map[ShipType]int) error {

	if required == nil {
		return nil
	}

	errMsg := "Ship types and numbers in GameConfig do not match ships supplied for Player"

	if len(required) != len(actual) {
		return errors.New(errMsg)
	}

	for k, v := range required {

		v2, found := actual[k]

		if !found || v != v2 {
			return errors.New(errMsg)
		}
	}

	return nil
}

func (ship *Ship) getCoord(offset int) Coord {

	if ship.Dir == TopToBottom {
		return Coord{rune(int(ship.Start.Row) + offset), ship.Start.Column}
	}

	// Assume LeftToRight otherwise
	return Coord{ship.Start.Row, ship.Start.Column + offset}
}
