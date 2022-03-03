package main

import "fmt"

// Make a 3x3 board with each cell being a string of length 1
type Board [9]string

func (print *Board) print() {
	for i := 0; i < 9; i += 3 {
		println(print[i] + "|" + print[i+1] + "|" + print[i+2])
	}
}

func main() {
	// Initialize the board
	var board Board = [9]string{"-", "-", "-", "-", "-", "-", "-", "-", "-"}
	// Print the board
	for {
		// Get the input
		getInput(&board)
		// Print the board
		board.print()
		// Run the ai
		ai(&board)
		// Print the board
		board.print()
		// Check if the game is over
		if checkWin(&board) {
			break
		}
	}
}

func checkWin(board *Board) bool {

	// Check if there is a win
	// Check rows
	for i := 0; i < 3; i++ {
		if board[i*3] == board[i*3+1] && board[i*3+1] == board[i*3+2] {
			if board[i*3] == "X" {
				fmt.Println("Player wins!")
				return true
			} else if board[i*3] == "O" {
				fmt.Println("AI wins!")
				return true
			}
		}
	}
	// Check columns
	for i := 0; i < 3; i++ {
		if board[i] == board[i+3] && board[i+3] == board[i+6] {
			if board[i] == "X" {
				fmt.Println("Player wins!")
				return true
			} else if board[i] == "O" {
				fmt.Println("AI wins!")
				return true
			}
		}
	}
	// Check diagonals
	if board[0] == board[4] && board[4] == board[8] {
		if board[0] == "X" {
			fmt.Println("Player wins!")
			return true
		} else if board[0] == "O" {
			fmt.Println("AI wins!")
			return true
		}
	}
	if board[2] == board[4] && board[4] == board[6] {
		if board[2] == "X" {
			fmt.Println("Player wins!")
			return true
		} else if board[2] == "O" {
			fmt.Println("AI wins!")
			return true
		}
	}
	for i := 0; i < 9; i++ {
		if board[i] == "-" {
			return false
		}
	}
	// If the game is over and the AI nor the player won, it is a draw
	fmt.Println("Draw!")
	return true

}

// Our AI simply loops over the board and fills the first empty cell
func ai(board *Board) {
	println("AI's turn!")
	for i := 0; i < 9; i++ {
		if board[i] == "-" {
			board[i] = "O"
			break
		}
	}
}

func getInput(board *Board) {
	// Gets the input from the user
	// input MUST be an integer between 1 and 9
	// The cell must be empty
	var input int = 0
	fmt.Print("Enter a number between 1 and 9: ")
	fmt.Scan(&input)
	// Check if the input is valid
	if input < 1 || input > 9 {
		fmt.Println("Invalid input")
		getInput(board)
	}
	// Check if the cell is empty
	if board[input-1] != "-" {
		fmt.Println("Cell is not empty")
		getInput(board)
	} else { // NOTE: This else statement is needed otherwise the line bellow is still reached before starting the recusrion
		// Set the cell to X
		board[input-1] = "X"
	}
}
