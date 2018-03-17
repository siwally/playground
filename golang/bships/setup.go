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
	gridWidth  int
	gridHeight int
	gridStart  Coord
}

// Coord represents a single coordinate on the grid, e.g. {row: 'A', column: 3{}
type Coord struct {
	row    rune
	column int
}

// Ship represents an individual ship and its location.
type Ship struct {
	start Coord
	dir   Facing
	len   int
}

// NewGame initialises and returns a new game, set up with the ships provided at their specified locations.
func NewGame(cfg GameConfig, ships ...Ship) (game *Game, err error) {

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Unexpected error setting up new game: %v", r)
		}
	}()

	shipCoords := map[Coord]*Ship{}

	for _, ship := range ships {
		if err := plotCoords(&cfg, &ship, shipCoords); err != nil {
			return nil, err
		}
	}

	return &Game{cfg, ships, shipCoords, map[Coord]*Ship{}}, nil
}

// NewDefaultGame sets up a game using a standard configuration for the grid and ships
func NewDefaultGame(ships ...Ship) (game *Game, err error) {

	cfg := GameConfig{gridStart: Coord{'A', 1}, gridWidth: 8, gridHeight: 8}

	return NewGame(cfg, ships...)
}

func plotCoords(cfg *GameConfig, ship *Ship, shipCoords map[Coord]*Ship) error {

	for i := 0; i < ship.len; i++ {

		pos := getCoord(ship, i)

		if err := validateCoord(cfg, pos); err != nil {
			return err
		}

		if _, dup := shipCoords[pos]; dup {
			return fmt.Errorf("Unable to place ship at %v, as ship already at this coordinate", pos)
		}

		shipCoords[pos] = ship
	}

	return nil
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
func validateCoord(cfg *GameConfig, coord Coord) error {

	if coord.row < cfg.gridStart.row {
		return fmt.Errorf("Ship exceeds top boundary at coordinate %v", coord)
	}

	if int(coord.row)-int(cfg.gridStart.row)+1 > cfg.gridWidth {
		return fmt.Errorf("Ship exceeds lower boundary at coordinate %v", coord)
	}

	if coord.column < 1 {
		return fmt.Errorf("Ship %v exceeds left-hand boundary", coord)
	}

	if coord.column-cfg.gridStart.column+1 > cfg.gridHeight {
		return fmt.Errorf("Ship %v exceeds right-hand boundary", coord)
	}

	return nil
}
