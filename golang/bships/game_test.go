package bships

import (
	"testing"
)

func TestVerticalShip(t *testing.T) {
	ship := Ship{Coord{'B', 5}, topToBottom, mid}
	game, _ := NewDefaultGame(ship)

	checkHitOrMiss(game, Coord{'B', 5}, t, true, nil)
	checkHitOrMiss(game, Coord{'C', 5}, t, true, nil)
	checkHitOrMiss(game, Coord{'D', 5}, t, true, nil)
	checkHitOrMiss(game, Coord{'E', 5}, t, true, &ship)

	checkHitOrMiss(game, Coord{'B', 6}, t, false, nil)
	checkHitOrMiss(game, Coord{'B', 4}, t, false, nil)
	checkHitOrMiss(game, Coord{'A', 5}, t, false, nil)
	checkHitOrMiss(game, Coord{'F', 5}, t, false, nil)

	if len(game.remaining) != 0 {
		t.Errorf("Expected remaining to be 0 for vertical ship, but was %v", len(game.remaining))
	}
}

func TestDuplicateHits(t *testing.T) {
	ship := Ship{Coord{'B', 5}, topToBottom, mid}
	game, _ := NewDefaultGame(ship)

	checkHitOrMiss(game, Coord{'C', 5}, t, true, nil)
	checkHitOrMiss(game, Coord{'C', 5}, t, true, nil)

	if len(game.remaining) != 3 {
		t.Errorf("Expected remaining to be 3 after dup hits, but was %v", len(game.remaining))
	}
}

func TestHorizontalShip(t *testing.T) {
	ship := Ship{Coord{'B', 2}, leftToRight, mid}
	game, _ := NewDefaultGame(ship)

	checkHitOrMiss(game, Coord{'B', 2}, t, true, nil)
	checkHitOrMiss(game, Coord{'B', 3}, t, true, nil)
	checkHitOrMiss(game, Coord{'B', 4}, t, true, nil)
	checkHitOrMiss(game, Coord{'B', 5}, t, true, &ship)

	checkHitOrMiss(game, Coord{'B', 1}, t, false, nil)
	checkHitOrMiss(game, Coord{'B', 6}, t, false, nil)
	checkHitOrMiss(game, Coord{'A', 2}, t, false, nil)
	checkHitOrMiss(game, Coord{'C', 2}, t, false, nil)

	if len(game.remaining) != 0 {
		t.Errorf("Expected remaining to be 0 for horizontal ship, but was %v", len(game.remaining))
	}
}

func TestSinkingShip(t *testing.T) {
	cfg := GameConfig{gridHeight: 10,
		gridWidth: 10,
		shipTypes: map[ShipType]int{dinky: 1}}

	ship := Ship{dir: topToBottom, shipType: dinky, start: Coord{'A', 1}}
	game, _ := NewGame(cfg, ship)

	checkHitOrMiss(game, Coord{'A', 1}, t, true, nil)
	checkHitOrMiss(game, Coord{'B', 1}, t, true, nil)
	checkHitOrMiss(game, Coord{'C', 1}, t, true, &ship)

	// check still gives accurate information if repeated
	checkHitOrMiss(game, Coord{'C', 1}, t, true, &ship)
	checkHitOrMiss(game, Coord{'B', 1}, t, true, &ship)
}

// TODO Test multiple ships and winning

// TODO General tidy-up of logic and use of pointers versus values

// TOOD Introduce players and introduce stats - who owns what?

func checkHitOrMiss(game *Game, move Coord, t *testing.T, expected bool, expSunk *Ship) {

	res, err := game.Play(move)

	if err != nil {
		t.Errorf("Error returned when playing game")
	}

	if expected != res.hit {
		t.Errorf("Expected hit result to be %v for coord %v, but was %v", expected, move, res.hit)
	}

	if expSunk == nil && res.sunkShip == nil {
		// fine
	} else if expSunk == nil && res.sunkShip != nil {
		t.Errorf("Expected ship at coord %v not to be sunk", move)
	} else if expSunk != nil && res.sunkShip == nil {
		t.Errorf("Expected ship to be sunk at coord %v, but was nil", move)
	} else if *expSunk != *res.sunkShip {
		t.Errorf("Expected ship sunk result to be %v for coord %v, but was %v", expSunk, move, res.sunkShip)
	}
}
