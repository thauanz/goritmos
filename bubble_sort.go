package main

import "fmt"

func main() {
	numbers := []int{5, 3, 2, 6, 10, 8, 20, 44, 33, 12, 16, 50, 37, 1, 24}
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
	fmt.Println(numbers)
}
