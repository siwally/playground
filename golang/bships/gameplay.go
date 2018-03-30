package bships

import "fmt"

// Player represents a player in a Game, with a record of the ships they are defending and hits against them.
type Player struct {
	remaining, hits map[Coord]*Ship
	hitsByShip      map[*Ship]int
}

// PlotShips plots the ships onto the player's grid and returns the ship types found and the full ship coordinates.
func (player *Player) PlotShips(cfg *GameConfig, ships ...Ship) (types map[ShipType]int, coords map[Coord]*Ship) {
	types = map[ShipType]int{}
	coords = map[Coord]*Ship{}

	hitsByShip := map[*Ship]int{}

	for _, ship := range ships {
		types[ship.shipType]++

		hitsByShip[&ship] = 0
		plotCoords(cfg, &ship, coords)
	}

	player.remaining = coords
	player.hits = map[Coord]*Ship{}
	player.hitsByShip = hitsByShip

	return
}

func plotCoords(cfg *GameConfig, ship *Ship, shipCoords map[Coord]*Ship) (coords []Coord) {

	coords = make([]Coord, ship.shipType.len)

	for i := 0; i < ship.shipType.len; i++ {

		pos := ship.getCoord(i)
		coords[i] = pos

		if _, dup := shipCoords[pos]; dup {
			panic(fmt.Sprintf("Unable to place ship at %v, as ship already at this coordinate", pos))
		}

		shipCoords[pos] = ship
	}

	return
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
