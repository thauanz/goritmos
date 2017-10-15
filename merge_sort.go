package main

import (
	"fmt"
	"goritmos/randomize"
)

func main() {
	numbers := randomize.Generate(100, 500)
	fmt.Println("Before:", numbers)
	fmt.Println("After:", mergeSort(numbers))
}

func mergeSort(numbers []int) []int {
	qtd := len(numbers)

	if qtd == 1 {
		return numbers
	}
	mid := qtd / 2

	l := numbers[0:mid]
	r := numbers[mid:qtd]

	return merge(mergeSort(l), mergeSort(r))
}

func merge(l, r []int) []int {
	var result = make([]int, len(l)+len(r))

	x := 0
	for len(l) > 0 && len(r) > 0 {
		if l[0] < r[0] {
			result[x] = l[0]
			l = l[1:] // remove the list
		} else {
			result[x] = r[0]
			r = r[1:] // remove the list
		}
		x++
	}

	for i := 0; i < len(l); i++ {
		result[x] = l[i]
		x++
	}
	for i := 0; i < len(r); i++ {
		result[x] = r[i]
		x++
	}

	return result
}
