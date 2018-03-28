package bships

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

	if _, hit := player.remaining[move]; !hit {
		return false, nil, nil
	}

	ship, _ := player.remaining[move]
	player.recordHit(move, ship)

	return true, player.sunk(ship), nil
}

func (player *Player) sunk(ship *Ship) (sunk *Ship) {

	if player.hitsByShip[ship] == ship.shipType.len {
		sunk = ship
	}

	return
}

func (player *Player) recordHit(move Coord, ship *Ship) {
	delete(player.remaining, move)
	player.hits[move] = ship
	player.hitsByShip[ship]++
}
