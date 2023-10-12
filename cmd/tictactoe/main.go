package main

import (
	"fmt"
	"os"

	"github.com/natanaelrusli/tic-tac-toe/internal/tictactoe"
)

func main() {
	game := tictactoe.NewTicTacToe()
	userInput := tictactoe.GetUserInput(os.Stdin)

	fmt.Println(userInput)
	if userInput != "" {
		game.StartGame(userInput)
	}
}
