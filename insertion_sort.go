package main

import (
	"fmt"
	"goritmos/randomize"
)

func main() {
	numbers := randomize.Generate(50, 700)

	fmt.Println("Before:", numbers)
	fmt.Println("After:", insertionSort(numbers))
}

func insertionSort(numbers []int) []int {
	qtd := len(numbers)
	for i := 1; i < qtd; i++ {
		position := i

		for position > 0 && numbers[position-1] > numbers[position] {
			value := numbers[position]
			numbers[position] = numbers[position-1]
			numbers[position-1] = value
			position--
		}
	}
	return numbers
}
