package bships

import (
	"testing"
)

func TestHorizontalShipSetup(t *testing.T) {

	// fine, up to right-hand edge
	ship := Ship{Coord{'D', 6}, leftToRight, 3}

	if _, err := NewDefaultGame(ship); err != nil {
		t.Errorf("Should be able to create game with ship up to right-hand boundary; received err: %v", err)
	}

	// not fine, past right-hand edge
	ship = Ship{Coord{'D', 6}, leftToRight, 4}

	if _, err := NewDefaultGame(ship); err == nil {
		t.Errorf("Should not be able to create game with ship end past right-hand boundary")
	}

	// not fine, starting position too low
	ship = Ship{Coord{'D', 0}, leftToRight, 4}

	if _, err := NewDefaultGame(ship); err == nil {
		t.Errorf("Should not be able to create game with ship start point past left-hand boundary")
	}
}

func TestVerticalShipSetup(t *testing.T) {

	// fine, up to bottom edge
	ship := Ship{Coord{'C', 1}, topToBottom, 6}

	if _, err := NewDefaultGame(ship); err != nil {
		t.Errorf("Should be able to create game with ship up to bottom boundary; received err: %v", err)
	}

	// not fine, past right-hand edge
	ship = Ship{Coord{'C', 1}, topToBottom, 7}

	if _, err := NewDefaultGame(ship); err == nil {
		t.Errorf("Should not be able to create game with ship end past bottom boundary")
	}

	// not fine, starting position too low
	ship = Ship{Coord{'1', 1}, topToBottom, 4}

	if _, err := NewDefaultGame(ship); err == nil {
		t.Errorf("Should not be able to create game with ship start point past top boundary")
	}
}

func TestMultipleShips(t *testing.T) {

	// fine, not overlapping
	ship1 := Ship{Coord{'A', 2}, leftToRight, 4}
	ship2 := Ship{Coord{'C', 6}, topToBottom, 3}

	if _, err := NewDefaultGame(ship1, ship2); err != nil {
		t.Errorf("Error returned when setting up game with multiple ships: %v", err)
	}

	// not fine, as ships overlap in their middles
	ship2 = Ship{Coord{'B', 3}, topToBottom, 3}
	ship1 = Ship{Coord{'C', 2}, leftToRight, 3}

	if _, err := NewDefaultGame(ship1, ship2); err == nil {
		t.Errorf("Error should have been returned when ship coords overlap")
	}
}

func TestArgsOutofBounds(t *testing.T) {
	ship := Ship{Coord{'B', 3}, -45, 3}

	if _, err := NewDefaultGame(ship); err == nil {
		t.Errorf("Error should have been returned when ship Facing value out of bounds")
	}
}

func TestGameConfig(t *testing.T) {

	// fine, correct number of ships and ships the right lenght
	ship1 := Ship{Coord{'B', 1}, topToBottom, 4}
	ship2 := Ship{Coord{'A', 1}, leftToRight, 5}

	cfg := GameConfig{gridStart: Coord{'A', 1}, gridWidth: 8, gridHeight: 8}

	if _, err := NewGame(cfg, ship1, ship2); err != nil {
		t.Errorf("Should be able to create game with ships that conform to game config; err=%v", err)
	}

	// TODO - not fine
}
