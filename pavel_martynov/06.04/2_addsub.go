package main

import (
	"fmt"
)

func doSmth(a int, b int) int {
	if a <= 1 && b >= 3 {
		return a + b
	} else {
		return a - b
	}
}

func main() {
	a := 1
	b := 4

	fmt.Println(doSmth(a, b))

	a = 4
	b = 5

	fmt.Println(doSmth(a, b))
}
