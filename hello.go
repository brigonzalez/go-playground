package main

import (
	"fmt"

	"github.com/brigonzalez/stringutil"
)

func main() {
	fmt.Println("Hello, world.")
	fmt.Println(stringutil.Reverse("!oG ,olleH"))

	a := make([]int, 1)
	a[0] = 234
	fmt.Println(a);
}
