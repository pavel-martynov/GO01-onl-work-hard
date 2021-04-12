package main

import (
	"errors"
	"fmt"
)

func fibonacci(n int) []uint64 {
	if n == 0 {
		return []uint64{}
	}

	if n == 1 {
		return []uint64{0}
	}

	if n > 93 {
		fmt.Println("94th and further fibonacci numbers are not fitting into uint64")
		n = 93
	}

	fseq := []uint64{0, 1}
	for i := 2; i < n; i++ {
		fseq = append(fseq, fseq[i-2] + fseq[i-1])
	}

	return fseq
}

func sequenceOfNaturalNumbers(n int) bool {
	sum := 0
	for i := 0; i <= n; i++ {
		sum += i
	}

	fmt.Println("Summary of the sequence: ", sum)

	equationResult := (n * (n + 1)) / 2

	fmt.Println("Equation result: ", equationResult)

	return sum == equationResult
}

func factorial(n int) (int, error) {
	if n < 0 {
		return n, errors.New("cannot calculate factorial of negative number")
	}
	if n <= 1 {
		return 1, nil
	}

	previousValue, _ := factorial(n-1)

	return n * previousValue, nil
}

func main() {
	fmt.Println(fibonacci(100))
	fmt.Println(sequenceOfNaturalNumbers(100))
	fmt.Println(factorial(4))
}
