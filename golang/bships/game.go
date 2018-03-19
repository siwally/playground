package bships

// Game represents a game of Batteships that is in progress.
type Game struct {
	config          GameConfig
	ships           []Ship
	remaining, hits map[Coord]*Ship
	hitsByShip      map[*Ship]int
}

// Result is the result of a player's move.
type Result struct {
	hit      bool
	sunkShip *Ship
}

// Play executes a move against a player's grid.
func (game *Game) Play(move Coord) (Result, error) {

	if _, alreadyHit := game.hits[move]; alreadyHit {
		return Result{true, game.getShipIfSunk(game.hits[move])}, nil
	}

	ship, hit := game.remaining[move]

	if !hit {
		return Result{false, nil}, nil
	}

	delete(game.remaining, move)
	game.hits[move] = ship
	game.hitsByShip[ship]++

	sunk := game.getShipIfSunk(ship)

	return Result{true, sunk}, nil
}

func (game *Game) getShipIfSunk(ship *Ship) (sunk *Ship) {

	if game.hitsByShip[ship] == ship.shipType.len {
		sunk = ship
	}

	return
}
