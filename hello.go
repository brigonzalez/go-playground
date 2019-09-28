package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main() {
	argsWithoutProgramName := os.Args[1:]

	switch argsWithoutProgramName[0] {
	case "sayHello":
		helloFromStandardInput()
	case "sum":
		sumNumbersRecursively()
	default:
		fmt.Println("option entered not defined")
	}
}

func handleError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func enterTextPrompt() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter text: ")
	text, _ := reader.ReadString('\n')

	return text
}

func helloFromStandardInput() {
	text := enterTextPrompt()
	fmt.Println("Hello " + text)
}

func getSanitizedNumber(enteredNumber string) (int, error) {
	enteredNumberWithoutSuffix := strings.TrimSuffix(enteredNumber, "\n")
	integerEnteredNumber, err := strconv.Atoi(enteredNumberWithoutSuffix)

	return integerEnteredNumber, err
}

func sumNumbersForLoop() {
	text := enterTextPrompt()
	arrayOfEnteredNumbers := strings.Split(text, " ")

	totalSum := 0
	for _, enteredNumber := range arrayOfEnteredNumbers {
		integerEnteredNumber, err := getSanitizedNumber(enteredNumber)
		handleError(err)
		totalSum += integerEnteredNumber
	}

	fmt.Println("total sum: " + strconv.Itoa(totalSum))
}

func sumNumbersRecursively() {
	text := enterTextPrompt()
	arrayOfEnteredNumbers := strings.Split(text, " ")
	totalSum := recursionSumFunction(arrayOfEnteredNumbers)

	fmt.Println(totalSum)
}

func recursionSumFunction(arrayOfEnteredNumbers []string) int {
	integerEnteredNumber, err := getSanitizedNumber(arrayOfEnteredNumbers[0])
	handleError(err)

	if len(arrayOfEnteredNumbers) == 1 {
		return integerEnteredNumber
	}

	return integerEnteredNumber + recursionSumFunction(arrayOfEnteredNumbers[1:])
}
