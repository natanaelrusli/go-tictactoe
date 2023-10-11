package tictactoe

import (
	"fmt"
	"strings"
)

func GetUserInput() string {
	var userInput string

	for userInput == "" {
		fmt.Println("Sample Input: ")

		fmt.Scanln(&userInput)
		strArr := strings.Split(userInput, "")

		if len(strArr) != 9 {
			userInput = ""
			fmt.Println("please input exactly 9 characters")
		}
	}

	return userInput
}
