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

// Gameplay features
// TODO Test sinking a ship, multiple ships
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
