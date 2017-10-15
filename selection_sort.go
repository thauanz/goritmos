package main

import (
	"fmt"
	"goritmos/randomize"
)

func main() {
	numbers := randomize.Generate(100, 500)
	fmt.Println("Before:", numbers)
	fmt.Println("After:", selectionSort(numbers))
}

func selectionSort(numbers []int) []int {
	qtd := len(numbers)
	for x := 0; x < qtd; x++ {
		indexMin := x

		for y := x; y < qtd; y++ {
			if numbers[y] < numbers[indexMin] {
				indexMin = y
			}
		}
		// change the position of the smallest number
		number := numbers[x]
		numbers[x] = numbers[indexMin]
		numbers[indexMin] = number
	}
	return numbers
}
