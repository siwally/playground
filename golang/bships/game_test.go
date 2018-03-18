package bships

import (
	"testing"
)

func TestVerticalShip(t *testing.T) {
	ship := Ship{Coord{'B', 5}, topToBottom, mid}
	game, _ := NewDefaultGame(ship)

	checkHitOrMiss(game, Coord{'B', 5}, t, true)
	checkHitOrMiss(game, Coord{'C', 5}, t, true)
	checkHitOrMiss(game, Coord{'D', 5}, t, true)
	checkHitOrMiss(game, Coord{'E', 5}, t, true)

	checkHitOrMiss(game, Coord{'B', 6}, t, false)
	checkHitOrMiss(game, Coord{'B', 4}, t, false)
	checkHitOrMiss(game, Coord{'A', 5}, t, false)
	checkHitOrMiss(game, Coord{'F', 5}, t, false)

	if len(game.remaining) != 0 {
		t.Errorf("Expected remaining to be 0 for vertical ship, but was %v", len(game.remaining))
	}
}

func TestDuplicateHits(t *testing.T) {
	ship := Ship{Coord{'B', 5}, topToBottom, mid}
	game, _ := NewDefaultGame(ship)

	checkHitOrMiss(game, Coord{'C', 5}, t, true)
	checkHitOrMiss(game, Coord{'C', 5}, t, true)

	if len(game.remaining) != 3 {
		t.Errorf("Expected remaining to be 3 after dup hits, but was %v", len(game.remaining))
	}
}

func TestHorizontalShip(t *testing.T) {
	ship := Ship{Coord{'B', 2}, leftToRight, mid}
	game, _ := NewDefaultGame(ship)

	checkHitOrMiss(game, Coord{'B', 2}, t, true)
	checkHitOrMiss(game, Coord{'B', 3}, t, true)
	checkHitOrMiss(game, Coord{'B', 4}, t, true)
	checkHitOrMiss(game, Coord{'B', 5}, t, true)

	checkHitOrMiss(game, Coord{'B', 1}, t, false)
	checkHitOrMiss(game, Coord{'B', 6}, t, false)
	checkHitOrMiss(game, Coord{'A', 2}, t, false)
	checkHitOrMiss(game, Coord{'C', 2}, t, false)

	if len(game.remaining) != 0 {
		t.Errorf("Expected remaining to be 0 for horizontal ship, but was %v", len(game.remaining))
	}
}

// Gameplay features
// TODO Test ending a game when a ship is sunk
// TODO Test sinking multiple ships and ending game
// TODO Test game stats - ships hit, ships remaining
// TOOD Think about multiple players and how this would work - check the rules :)

func checkHitOrMiss(game *Game, move Coord, t *testing.T, expected bool) {

	res, err := game.Play(move)

	if err != nil {
		t.Errorf("Error returned when playing game")
	}

	if expected != res {
		t.Errorf("Expected hit result to be %v for coord %v, but was %v", expected, move, res)
	}
}
