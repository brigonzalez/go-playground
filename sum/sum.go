package sum

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/brigonzalez/go-playground/util"
)

func SumNumbersForLoop() {
	text := util.EnterTextPrompt()
	arrayOfEnteredNumbers := strings.Split(text, " ")

	totalSum := 0
	for _, enteredNumber := range arrayOfEnteredNumbers {
		integerEnteredNumber, err := util.GetSanitizedNumber(enteredNumber)
		util.HandleError(err)
		totalSum += integerEnteredNumber
	}

	fmt.Println("total sum: " + strconv.Itoa(totalSum))
}

func SumNumbersRecursively() {
	text := util.EnterTextPrompt()
	arrayOfEnteredNumbers := strings.Split(text, " ")
	totalSum := recursionSumFunction(arrayOfEnteredNumbers)

	fmt.Println(totalSum)
}

func recursionSumFunction(arrayOfEnteredNumbers []string) int {
	integerEnteredNumber, err := util.GetSanitizedNumber(arrayOfEnteredNumbers[0])
	util.HandleError(err)

	if len(arrayOfEnteredNumbers) == 1 {
		return integerEnteredNumber
	}

	return integerEnteredNumber + recursionSumFunction(arrayOfEnteredNumbers[1:])
}
