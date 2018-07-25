package main

import "fmt"

var board = [10][10]bool{}

func main() {

	// Start by playing against computer, with ships hard-coded onto grid
	// Ask user for moves, then send to server and show if they've hit or not

	// TODO Create basic game with one player, to try to hit.

	fmt.Println("Their grid...")
	printGrid(board)
	move := askForMove()

	fmt.Printf("Move is %v\n", move)
}

func printGrid(board [10][10]bool) {

	fmt.Println()

	for _, row := range board {

		for _, b := range row {
			if b {
				fmt.Printf("%s", " X ")
			} else {
				fmt.Printf("%s", " - ")
			}
		}

		fmt.Println()
	}

	fmt.Println()
}

func askForMove() (move string) {
	fmt.Print("Enter move: ")

	fmt.Scanln(&move)

	return
}
