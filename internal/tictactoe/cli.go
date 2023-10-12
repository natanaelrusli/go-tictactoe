package tictactoe

import (
	"fmt"
	"io"
	"strings"
)

func GetUserInput(reader io.Reader) string {
	var userInput = ""
	var attempts = 0
	var maxAttempts = 3

	for userInput == "" {

		fmt.Println("Sample Input: ")

		fmt.Fscanln(reader, &userInput)
		strArr := strings.Split(userInput, "")

		if len(strArr) != 9 {
			userInput = ""
			attempts += 1

			if attempts == maxAttempts {
				fmt.Println("Exceeded maximum attempts. Exiting.")
				break
			}

			fmt.Println("please input exactly 9 characters")
		}
	}

	return userInput
}
