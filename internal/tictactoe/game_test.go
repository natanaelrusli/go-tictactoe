package tictactoe

import (
	"testing"
)

func TestSplitIntoBoard(t *testing.T) {
	given := "XOXOOOXOO"
	want := [3][3]string{{"X", "O", "X"}, {"O", "O", "O"}, {"X", "O", "O"}}

	game := NewTicTacToe()
	res := game.SplitIntoBoard(given)

	if res != want {
		t.Fatalf(`SplitIntoBoard("XOXOOOXOX") = %q want: %#q`, res, want)
	}
}
