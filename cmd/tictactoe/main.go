package main

import (
	"github.com/natanaelrusli/tic-tac-toe/internal/tictactoe"
)

func main() {
	game := tictactoe.NewTicTacToe()
	game.StartGame()
}
