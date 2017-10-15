package randomize

import (
	"math/rand"
)

func Generate(quantity, interval int) []int {
	//rand.Seed(int64(interval))
	if quantity > interval {
		panic("the quantity of number is bigger than interval")
	}

	var numbers []int
	generated := make(map[int]bool, quantity)

	var number int
	for i := 0; i < quantity; i++ {
		number = rand.Intn(interval)
		if !generated[number] {
			generated[number] = true
			numbers = append(numbers, number)
		}
	}
	return numbers
}
