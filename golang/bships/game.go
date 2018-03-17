package bships

// Game represents a game of Batteships that is in progress.
type Game struct {
	config          GameConfig
	ships           []Ship
	remaining, hits map[Coord]*Ship
}

// Play executes a move against a player's grid
func (game *Game) Play(move Coord) (bool, error) {

	if _, alreadyHit := game.hits[move]; alreadyHit {
		return true, nil
	}

	v, hit := game.remaining[move]

	if hit {
		delete(game.remaining, move)
		game.hits[move] = v
	}

	return hit, nil
}
