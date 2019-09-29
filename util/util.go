package util

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func EnterTextPrompt() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter text: ")
	text, _ := reader.ReadString('\n')

	return text
}

func HandleError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func GetSanitizedNumber(enteredNumber string) (int, error) {
	enteredNumberWithoutSuffix := strings.TrimSuffix(enteredNumber, "\n")
	integerEnteredNumber, err := strconv.Atoi(enteredNumberWithoutSuffix)

	return integerEnteredNumber, err
}
