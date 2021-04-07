package main

import "fmt"

func season(num int) string {
	if num == 1 {
		return "winter"
	} else if num == 2 {
		return "spring"
	} else if num == 3 {
		return "summer"
	} else if num == 4 {
		return "autumn"
	}

	return "unknown season"
}

func main() {
	for i := 1; i < 6; i++ {
		fmt.Println(i, season(i))
	}
}
