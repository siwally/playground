package bships

// Game represents a game of Batteships that is in progress.
type Game struct {
	ship            Ship
	remaining, hits map[coord]*Ship
}

type coord struct {
	row    rune
	column int
}

// Play executes a move against a player's grid
func (game *Game) Play(move coord) (bool, error) {

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
