package bships

import (
	"fmt"
)

// Facing indicates the direction of the ship on the grid, from its starting coordinate, e.g. from top to bottom.
type Facing int8

const (
	topToBottom Facing = iota
	leftToRight Facing = iota
	gridHeight  int    = 8
	gridWidth   int    = 8
)

var gridStart = coord{'A', 1}

// Ship represents the location of a ship.
type Ship struct {
	start coord
	dir   Facing
	len   int
}

// NewGame initialises and returns a new game, set up with the ships provided at their specified locations.
func NewGame(ship Ship) (*Game, error) {

	shipCoords := map[coord]*Ship{}

	for i := 0; i < ship.len; i++ {

		pos := getCoord(&ship, i)

		if err := validateCoord(pos); err != nil {
			return nil, err
		}

		shipCoords[pos] = &ship
	}

	return &Game{ship, shipCoords, map[coord]*Ship{}}, nil
}

// +1 as ship placement is inclusive, i.e. a ship includes its starting position in its len
func validateCoord(coord coord) error {

	if coord.row < gridStart.row {
		return fmt.Errorf("Ship exceeds top boundary at coordinate %v", coord)
	}

	if int(coord.row)-int(gridStart.row)+1 > gridWidth {
		return fmt.Errorf("Ship exceeds lower boundary at coordinate %v", coord)
	}

	if coord.column < 1 {
		return fmt.Errorf("Ship %v exceeds left-hand boundary", coord)
	}

	if coord.column-gridStart.column+1 > gridHeight {
		return fmt.Errorf("Ship %v exceeds right-hand boundary", coord)
	}

	return nil
}

func getCoord(ship *Ship, offset int) coord {

	switch ship.dir {
	case topToBottom:
		return coord{rune(int(ship.start.row) + offset), ship.start.column}
	case leftToRight:
		return coord{ship.start.row, ship.start.column + offset}
	default:
		panic("Unreachable condidtion in switch on ship.dir")
	}
}
