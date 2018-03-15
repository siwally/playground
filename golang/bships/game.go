package bships

import (
	"fmt"
)

// Facing indicates the direction of the ship on the grid, from its starting coordinate, e.g. from top to bottom.
type Facing int8

const (
	topToBottom Facing = iota
	leftToRight Facing = iota
	gridLen     int    = 8
	gridWidth   int    = 8
)

var gridTopLeft = coord{'A', 1}

// Ship represents the location of a ship.
type Ship struct {
	start coord
	dir   Facing
	len   int
}

// Game represents a game of Batteships that is in progress.
type Game struct {
	ship            Ship // TODO separate game set-up from gameplay
	remaining, hits map[coord]*Ship
}

// NewGame initialises and returns a new game, set up with the ships provided at their specified locations.
func NewGame(ship Ship) (*Game, error) {

	shipCoords := map[coord]*Ship{}

	switch ship.dir {
	case topToBottom:

		for i := 0; i < ship.len; i++ {
			coord := coord{rune(int(ship.start.row) + i), ship.start.column}
			shipCoords[coord] = &ship
		}

	case leftToRight:

		if ship.start.column+(ship.len-1) > gridLen {
			return nil, fmt.Errorf("Ship %v exceeds right-hand boundary", ship)
		}

		if ship.start.column+(ship.len-1) > gridLen {
			return nil, fmt.Errorf("Ship %v exceeds right-hand boundary", ship)
		}

		for i := 0; i < ship.len; i++ {
			coord := coord{ship.start.row, ship.start.column + i}
			shipCoords[coord] = &ship
		}
	}

	return &Game{ship, shipCoords, map[coord]*Ship{}}, nil
}

type coord struct {
	row    rune
	column int
}

// Play executes a move against a player's battleships grid
func (game *Game) Play(move coord) (bool, error) {

	if _, alreadyHit := game.hits[move]; alreadyHit {
		return true, nil
	}

	v, hit := game.remaining[move]

	if hit {
		delete(game.remaining, move)
		game.hits[move] = v
	}

	return hit, nil
}
