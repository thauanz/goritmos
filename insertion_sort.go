package main

import "fmt"

func main() {
	numbers := []int{14, 33, 27, 10, 35, 19, 42, 44, 7}

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
