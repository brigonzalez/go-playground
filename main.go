package main

import (
	"fmt"
	"os"

	"github.com/brigonzalez/go-playground/hello"
	"github.com/brigonzalez/go-playground/sum"
)

func main() {
	argsWithoutProgramName := os.Args[1:]

	if len(argsWithoutProgramName) == 0 {
		fmt.Println("enter option for program name: sum, sayHello")
	} else {
		switch argsWithoutProgramName[0] {
		case "sayHello":
			hello.HelloFromStandardInput()
		case "sum":
			sum.SumNumbersRecursively()
		default:
			fmt.Println("option entered not defined")
		}
	}
}
