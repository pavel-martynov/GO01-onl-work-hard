package main

import (
	"fmt"
)

func main() {
	min := 48

	if min < 15 {
		fmt.Println("This minute is in the first quarter")
	} else if min < 30 {
		fmt.Println("This minute is in the second quarter")
	} else if min < 45 {
		fmt.Println("This minute is in the third quarter")
	} else {
		fmt.Println("This minute is in the fourth quarter")
	}
}
