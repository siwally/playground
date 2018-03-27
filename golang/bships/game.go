package bships

// Game represents a game of Batteships that is in progress.
type Game struct {
	config  GameConfig
	players map[string]Player
}

// Player represents a player in a Game, with a record of the ships they are defending and hits against them.
type Player struct {
	remaining, hits map[Coord]*Ship
	hitsByShip      map[*Ship]int
}

// Attack executes a move against a player's grid.
func (player *Player) Attack(move Coord) (bool, *Ship, error) {

	if ship, alreadyHit := player.hits[move]; alreadyHit {
		return true, player.sunk(ship), nil
	}

	ship, hit := player.remaining[move]

	if !hit {
		return false, nil, nil
	}

	// o.k., it's a new hit!
	delete(player.remaining, move)
	player.hits[move] = ship
	player.hitsByShip[ship]++

	return true, player.sunk(ship), nil
}

func (player *Player) sunk(ship *Ship) (sunk *Ship) {

	if player.hitsByShip[ship] == ship.shipType.len {
		sunk = ship
	}

	return
}
