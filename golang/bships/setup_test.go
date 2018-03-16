package bships

import (
	"testing"
)

func TestHorizontalShipSetup(t *testing.T) {

	// fine, up to right-hand edge
	ship := Ship{coord{'D', 6}, leftToRight, 3}

	if _, err := NewGame(ship); err != nil {
		t.Errorf("Should be able to create game with ship up to right-hand boundary; received err: %v", err)
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

func TestVerticalShipSetup(t *testing.T) {

	// fine, up to bottom edge
	ship := Ship{coord{'C', 1}, topToBottom, 6}

	if _, err := NewGame(ship); err != nil {
		t.Errorf("Should be able to create game with ship up to bottom boundary; received err: %v", err)
	}

	// not fine, past right-hand edge
	ship = Ship{coord{'C', 1}, topToBottom, 7}

	if _, err := NewGame(ship); err == nil {
		t.Errorf("Should not be able to create game with ship end past bottom boundary")
	}

	// not fine, starting position too low
	ship = Ship{coord{'1', 1}, topToBottom, 4}

	if _, err := NewGame(ship); err == nil {
		t.Errorf("Should not be able to create game with ship start point past top boundary")
	}
}

// TODO Test error rather than panic with invalid Facing value for ship

// TODO Test with multiple ships, including overlapping

// TODO Test ship lengths, types and number of ships - game definition or game metadta
