package bships

import (
	"testing"
)

func TestVerticalShip(t *testing.T) {
	ship := Ship{Coord{'B', 5}, topToBottom, mid}
	game, _ := createTestGame(ship)

	checkHitOrMiss(game, Coord{'B', 5}, t, true, nil)
	checkHitOrMiss(game, Coord{'C', 5}, t, true, nil)
	checkHitOrMiss(game, Coord{'D', 5}, t, true, nil)
	checkHitOrMiss(game, Coord{'E', 5}, t, true, &ship)

	checkHitOrMiss(game, Coord{'B', 6}, t, false, nil)
	checkHitOrMiss(game, Coord{'B', 4}, t, false, nil)
	checkHitOrMiss(game, Coord{'A', 5}, t, false, nil)
	checkHitOrMiss(game, Coord{'F', 5}, t, false, nil)

	player, _ := game.players[playerName]

	if len(player.remaining) != 0 {
		t.Errorf("Expected remaining to be 0 for vertical ship, but was %v", len(player.remaining))
	}
}

func TestDuplicateHits(t *testing.T) {
	ship := Ship{Coord{'B', 5}, topToBottom, mid}
	game, _ := createTestGame(ship)

	checkHitOrMiss(game, Coord{'C', 5}, t, true, nil)
	checkHitOrMiss(game, Coord{'C', 5}, t, true, nil)

	player, _ := game.players[playerName]

	if len(player.remaining) != 3 {
		t.Errorf("Expected remaining to be 3 after dup hits, but was %v", len(player.remaining))
	}
}

func TestHorizontalShip(t *testing.T) {
	ship := Ship{Coord{'B', 2}, leftToRight, mid}
	game, _ := createTestGame(ship)

	checkHitOrMiss(game, Coord{'B', 2}, t, true, nil)
	checkHitOrMiss(game, Coord{'B', 3}, t, true, nil)
	checkHitOrMiss(game, Coord{'B', 4}, t, true, nil)
	checkHitOrMiss(game, Coord{'B', 5}, t, true, &ship)

	checkHitOrMiss(game, Coord{'B', 1}, t, false, nil)
	checkHitOrMiss(game, Coord{'B', 6}, t, false, nil)
	checkHitOrMiss(game, Coord{'A', 2}, t, false, nil)
	checkHitOrMiss(game, Coord{'C', 2}, t, false, nil)

	player, _ := game.players[playerName]

	if len(player.remaining) != 0 {
		t.Errorf("Expected remaining to be 0 for horizontal ship, but was %v", len(player.remaining))
	}
}

func TestSinkingShip(t *testing.T) {
	cfg := GameConfig{GridHeight: 10,
		GridWidth: 10,
		ShipTypes: map[ShipType]int{dinky: 1}}

	ship := Ship{dir: topToBottom, shipType: dinky, start: Coord{'A', 1}}
	game, _ := NewGame(cfg, playerName, ship)

	checkHitOrMiss(game, Coord{'A', 1}, t, true, nil)
	checkHitOrMiss(game, Coord{'B', 1}, t, true, nil)
	checkHitOrMiss(game, Coord{'C', 1}, t, true, &ship)

	// check still gives accurate information if repeated
	checkHitOrMiss(game, Coord{'C', 1}, t, true, &ship)
	checkHitOrMiss(game, Coord{'B', 1}, t, true, &ship)
}

// TODO General tidy-up of logic and use of pointers versus values.

// TODO Check things in the right place between files and structs.

// TODO Test multiple ships and winning.

func checkHitOrMiss(game *Game, move Coord, t *testing.T, expected bool, expSunk *Ship) {

	player, _ := game.players["player1"]
	hit, sunk, err := player.Attack(move)

	if err != nil {
		t.Errorf("Error returned when playing game")
	}

	if expected != hit {
		t.Errorf("Expected hit result to be %v for coord %v, but was %v", expected, move, hit)
	}

	if expSunk == nil && sunk == nil {
		// fine
	} else if expSunk == nil && sunk != nil {
		t.Errorf("Expected ship at coord %v not to be sunk", move)
	} else if expSunk != nil && sunk == nil {
		t.Errorf("Expected ship to be sunk at coord %v, but was nil", move)
	} else if *expSunk != *sunk {
		t.Errorf("Expected ship sunk result to be %v for coord %v, but was %v", expSunk, move, sunk)
	}
}
