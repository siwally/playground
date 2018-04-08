package bships

import (
	"testing"
)

const playerName = "player1"

var dinky = ShipType{Name: "dinky", Len: 3}
var mid = ShipType{Name: "mid", Len: 4}
var grande = ShipType{Name: "grande", Len: 5}

func TestHorizontalShipSetup(t *testing.T) {

	// fine, up to right-hand edge
	ship := Ship{Coord{'D', 6}, LeftToRight, dinky}

	if _, err := createTestGame(ship); err != nil {
		t.Errorf("Should be able to create game with ship up to right-hand boundary; received err: %v", err)
	}

	// not fine, past right-hand edge
	ship = Ship{Coord{'D', 6}, LeftToRight, mid}

	if _, err := createTestGame(ship); err == nil {
		t.Errorf("Should not be able to create game with ship end past right-hand boundary")
	}

	// not fine, starting position too low
	ship = Ship{Coord{'D', 0}, LeftToRight, mid}

	if _, err := createTestGame(ship); err == nil {
		t.Errorf("Should not be able to create game with ship start point past left-hand boundary")
	}
}

func TestVerticalShipSetup(t *testing.T) {

	// fine, up to bottom edge
	ship := Ship{Coord{'D', 1}, TopToBottom, grande}

	if _, err := createTestGame(ship); err != nil {
		t.Errorf("Should be able to create game with ship up to bottom boundary; received err: %v", err)
	}

	// not fine, past right-hand edge
	ship = Ship{Coord{'E', 1}, TopToBottom, grande}

	if _, err := createTestGame(ship); err == nil {
		t.Errorf("Should not be able to create game with ship end past bottom boundary")
	}

	// not fine, starting position too low
	ship = Ship{Coord{'1', 1}, TopToBottom, mid}

	if _, err := createTestGame(ship); err == nil {
		t.Errorf("Should not be able to create game with ship start point past top boundary")
	}
}

func TestMultipleShips(t *testing.T) {

	// fine, not overlapping
	ship1 := Ship{Coord{'A', 2}, LeftToRight, mid}
	ship2 := Ship{Coord{'C', 6}, TopToBottom, dinky}

	if _, err := createTestGame(ship1, ship2); err != nil {
		t.Errorf("Error returned when setting up game with multiple ships: %v", err)
	}

	// not fine, as ships overlap in their middles
	ship2 = Ship{Coord{'B', 3}, TopToBottom, dinky}
	ship1 = Ship{Coord{'C', 2}, LeftToRight, dinky}

	if _, err := createTestGame(ship1, ship2); err == nil {
		t.Errorf("Error should have been returned when ship coords overlap")
	}
}

func TestArgsOutofBounds(t *testing.T) {
	ship := Ship{Coord{'B', 3}, -45, dinky}

	if _, err := createTestGame(ship); err == nil {
		t.Errorf("Error should have been returned when ship Facing value out of bounds")
	}
}

func TestGameConfig(t *testing.T) {

	// fine, correct number of ships and ships the right lenght
	ship1 := Ship{Coord{'B', 1}, TopToBottom, mid}
	ship2 := Ship{Coord{'A', 1}, LeftToRight, grande}

	cfg := GameConfig{GridWidth: 8, GridHeight: 8}

	game := NewGame(cfg)
	if err := game.AddPlayer(playerName, ship1, ship2); err != nil {
		t.Errorf("Should be able to create game with ships that conform to game config grid size; err=%v", err)
	}

	// not fine, ships don't fit on board of this size
	ship1 = Ship{Coord{'B', 1}, TopToBottom, mid}
	ship2 = Ship{Coord{'A', 1}, LeftToRight, grande}

	cfg = GameConfig{GridWidth: 4, GridHeight: 4}
	game = NewGame(cfg)
	if err := game.AddPlayer(playerName, ship1, ship2); err == nil {
		t.Errorf("Error should have been returned to indicate ship won't fit on grid of this size")
	}

	// not fine, ships of wrong type even though enough ships
	ship1 = Ship{Coord{'B', 1}, TopToBottom, mid}
	ship2 = Ship{Coord{'A', 1}, LeftToRight, grande}

	shipTypes := map[ShipType]int{ShipType{Name: "Destroyer", Len: 5}: 2}
	cfg = GameConfig{GridWidth: 8, GridHeight: 8, shipTypes: shipTypes}

	game = NewGame(cfg)
	if err := game.AddPlayer(playerName, ship1, ship2); err == nil {
		t.Errorf("Error should have been returned to indicate not enough ships of specified type")
	}
}

// Creates a game using a standard configuration for the grid and no restriction on ship types.
func createTestGame(ships ...Ship) (game *Game, err error) {

	cfg := GameConfig{GridWidth: 8, GridHeight: 8}

	game = NewGame(cfg)
	err = game.AddPlayer(playerName, ships...)

	return
}
