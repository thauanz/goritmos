package main

import (
	"fmt"
	"goritmos/randomize"
)

func main() {
	numbers := randomize.Generate(100, 2000)
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
