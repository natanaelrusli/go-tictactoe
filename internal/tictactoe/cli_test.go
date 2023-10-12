package tictactoe

import (
	"strings"
	"testing"
)

// TestGetUserInput tests the GetUserInput function
func TestGetUserInput(t *testing.T) {
	testCases := []struct {
		input  string
		expect string
	}{
		{"123456789\n", "123456789"},
		{"XXX\n", ""},
		{"\n", ""},
	}

	for _, tc := range testCases {
		stdin := strings.NewReader(tc.input)
		result := GetUserInput(stdin)

		// Check the result
		if result != tc.expect {
			t.Errorf("Input: %s, Expected: %s, Got: %s", tc.input, tc.expect, result)
		}
	}
}
