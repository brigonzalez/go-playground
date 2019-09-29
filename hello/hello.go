package hello

import (
	"fmt"

	"github.com/brigonzalez/go-playground/util"
)

func HelloFromStandardInput() {
	text := util.EnterTextPrompt()
	fmt.Println("Hello " + text)
}
