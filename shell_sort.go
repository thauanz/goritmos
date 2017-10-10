package main

import (
	"fmt"
)

func main() {
	numbers := []int{34, 12, 23, 45, 55, 66, 89, 100, 96, 1000, 83, 88, 8, 9, 1}
	fmt.Println("Before:", numbers)
	shellSort(numbers)
	fmt.Println("After:", numbers)
}

func shellSort(numbers []int) {
	h := 1
	n := len(numbers)

	for h < n {
		h = h*3 + 1
	}

	h = h / 3
	var c, j int

	for h > 0 {
		for i := h; i < n; i++ {
			c = numbers[i]
			j = i
			for j >= h && numbers[j-h] > c {
				numbers[j] = numbers[j-h]
				j = j - h
			}
			numbers[j] = c
		}
		h = h / 2
	}
}
