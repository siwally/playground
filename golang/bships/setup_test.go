package bships

import (
	"testing"
)

func TestHorizontalShipSetup(t *testing.T) {

	// fine, up to right-hand edge
	ship := Ship{Coord{'D', 6}, leftToRight, dinky}

	if _, err := NewDefaultGame(ship); err != nil {
		t.Errorf("Should be able to create game with ship up to right-hand boundary; received err: %v", err)
	}

	// not fine, past right-hand edge
	ship = Ship{Coord{'D', 6}, leftToRight, mid}

	if _, err := NewDefaultGame(ship); err == nil {
		t.Errorf("Should not be able to create game with ship end past right-hand boundary")
	}

	// not fine, starting position too low
	ship = Ship{Coord{'D', 0}, leftToRight, mid}

	if _, err := NewDefaultGame(ship); err == nil {
		t.Errorf("Should not be able to create game with ship start point past left-hand boundary")
	}
}

func TestVerticalShipSetup(t *testing.T) {

	// fine, up to bottom edge
	ship := Ship{Coord{'D', 1}, topToBottom, grande}

	if _, err := NewDefaultGame(ship); err != nil {
		t.Errorf("Should be able to create game with ship up to bottom boundary; received err: %v", err)
	}

	// not fine, past right-hand edge
	ship = Ship{Coord{'E', 1}, topToBottom, grande}

	if _, err := NewDefaultGame(ship); err == nil {
		t.Errorf("Should not be able to create game with ship end past bottom boundary")
	}

	// not fine, starting position too low
	ship = Ship{Coord{'1', 1}, topToBottom, mid}

	if _, err := NewDefaultGame(ship); err == nil {
		t.Errorf("Should not be able to create game with ship start point past top boundary")
	}
}

func TestMultipleShips(t *testing.T) {

	// fine, not overlapping
	ship1 := Ship{Coord{'A', 2}, leftToRight, mid}
	ship2 := Ship{Coord{'C', 6}, topToBottom, dinky}

	if _, err := NewDefaultGame(ship1, ship2); err != nil {
		t.Errorf("Error returned when setting up game with multiple ships: %v", err)
	}

	// not fine, as ships overlap in their middles
	ship2 = Ship{Coord{'B', 3}, topToBottom, dinky}
	ship1 = Ship{Coord{'C', 2}, leftToRight, dinky}

	if _, err := NewDefaultGame(ship1, ship2); err == nil {
		t.Errorf("Error should have been returned when ship coords overlap")
	}
}

func TestArgsOutofBounds(t *testing.T) {
	ship := Ship{Coord{'B', 3}, -45, dinky}

	if _, err := NewDefaultGame(ship); err == nil {
		t.Errorf("Error should have been returned when ship Facing value out of bounds")
	}
}

func TestGameConfig(t *testing.T) {

	// fine, correct number of ships and ships the right lenght
	ship1 := Ship{Coord{'B', 1}, topToBottom, mid}
	ship2 := Ship{Coord{'A', 1}, leftToRight, grande}

	cfg := GameConfig{GridWidth: 8, GridHeight: 8}

	if _, err := NewGame(cfg, "player1", ship1, ship2); err != nil {
		t.Errorf("Should be able to create game with ships that conform to game config grid size; err=%v", err)
	}

	// not fine, ships don't fit on board of this size
	ship1 = Ship{Coord{'B', 1}, topToBottom, mid}
	ship2 = Ship{Coord{'A', 1}, leftToRight, grande}

	cfg = GameConfig{GridWidth: 4, GridHeight: 4}

	if _, err := NewGame(cfg, "player1", ship1, ship2); err == nil {
		t.Errorf("Error should have been returned to indicate ship won't fit on grid of this size")
	}

	// not fine, ships of wrong type even though enough ships
	ship1 = Ship{Coord{'B', 1}, topToBottom, mid}
	ship2 = Ship{Coord{'A', 1}, leftToRight, grande}

	shipTypes := map[ShipType]int{ShipType{name: "Destroyer", len: 5}: 2}
	cfg = GameConfig{GridWidth: 8, GridHeight: 8, ShipTypes: shipTypes}

	if _, err := NewGame(cfg, "player1", ship1, ship2); err == nil {
		t.Errorf("Error should have been returned to indicate not enough ships of specified type")
	}
}
