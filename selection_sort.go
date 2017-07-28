package main

import "fmt"

func main() {
	numbers := []int{14, 33, 27, 10, 35, 19, 42, 44, 7, 16, 36, 100}

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
		fmt.Println(numbers)
	}
	return numbers
}
