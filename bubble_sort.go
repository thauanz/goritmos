package main

import (
	"fmt"
	"goritmos/randomize"
)

func main() {
	numbers := randomize.Generate(300, 300)

	fmt.Println("Before:", numbers)
	fmt.Println("After:", bubbleSort(numbers))
}

func bubbleSort(numbers []int) []int {
	flag := true
	qtd := len(numbers)

	for flag {
		flag = false
		for i := 0; i < qtd-1; i++ {
			if numbers[i+1] < numbers[i] {
				aux := numbers[i+1]
				numbers[i+1] = numbers[i]
				numbers[i] = aux
				flag = true
			}
		}
		qtd--
	}
	return numbers
}
