package bships

import (
	"testing"
)

func TestVerticalShip(t *testing.T) {
	ship := Ship{coord{'B', 5}, topToBottom, 4}
	game, _ := NewGame(ship)

	checkHitOrMiss(game, coord{'B', 5}, t, true)
	checkHitOrMiss(game, coord{'C', 5}, t, true)
	checkHitOrMiss(game, coord{'D', 5}, t, true)
	checkHitOrMiss(game, coord{'E', 5}, t, true)

	checkHitOrMiss(game, coord{'B', 6}, t, false)
	checkHitOrMiss(game, coord{'B', 4}, t, false)
	checkHitOrMiss(game, coord{'A', 5}, t, false)
	checkHitOrMiss(game, coord{'F', 5}, t, false)

	if len(game.remaining) != 0 {
		t.Errorf("Expected remaining to be 0 for vertical ship, but was %v", len(game.remaining))
	}
}

func TestDuplicateHits(t *testing.T) {
	ship := Ship{coord{'B', 5}, topToBottom, 4}
	game, _ := NewGame(ship)

	checkHitOrMiss(game, coord{'C', 5}, t, true)
	checkHitOrMiss(game, coord{'C', 5}, t, true)

	if len(game.remaining) != 3 {
		t.Errorf("Expected remaining to be 3 after dup hits, but was %v", len(game.remaining))
	}
}

func TestHorizontalShip(t *testing.T) {
	ship := Ship{coord{'B', 2}, leftToRight, 4}
	game, _ := NewGame(ship)

	checkHitOrMiss(game, coord{'B', 2}, t, true)
	checkHitOrMiss(game, coord{'B', 3}, t, true)
	checkHitOrMiss(game, coord{'B', 4}, t, true)
	checkHitOrMiss(game, coord{'B', 5}, t, true)

	checkHitOrMiss(game, coord{'B', 1}, t, false)
	checkHitOrMiss(game, coord{'B', 6}, t, false)
	checkHitOrMiss(game, coord{'A', 2}, t, false)
	checkHitOrMiss(game, coord{'C', 2}, t, false)

	if len(game.remaining) != 0 {
		t.Errorf("Expected remaining to be 0 for horizontal ship, but was %v", len(game.remaining))
	}
}

func TestShipBoundaries(t *testing.T) {

	// fine, up to right-hand edge
	ship := Ship{coord{'D', 6}, leftToRight, 3}

	if _, err := NewGame(ship); err != nil {
		t.Errorf("Should be able to create game with ship up righ-hand boundary; received err: %v", err)
	}

	// not fine, past right-hand edge
	ship = Ship{coord{'D', 6}, leftToRight, 4}

	if _, err := NewGame(ship); err == nil {
		t.Errorf("Should not be able to create game with ship end past right-hand boundary")
	}

	// not fine, starting position too low
	ship = Ship{coord{'D', 0}, leftToRight, 4}

	if _, err := NewGame(ship); err == nil {
		t.Errorf("Should not be able to create game with ship start point past left-hand boundary")
	}

}

// TODO Test multiple ships

// TODO Test multiple ships overlapping

// TODO Test ship lengths, types and number of ships - game definition or game config

// Gameplay features
// TODO Test sinking a ship
// TODO Test ending a game when the ship is sunk
// TODO Test game stats - ships hit, ships remaining

func checkHitOrMiss(game *Game, move coord, t *testing.T, expected bool) {

	res, err := game.Play(move)

	if err != nil {
		t.Errorf("Error returned when playing game")
	}

	if expected != res {
		t.Errorf("Expected hit result to be %v for coord %v, but was %v", expected, move, res)
	}
}
