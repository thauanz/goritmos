package randomize

import (
	"math/rand"
)

func Generate(n ...int) []int {
	//rand.Seed(int64(interval))
	if len(n) <= 0 {
		panic("hey give one number, please!")
	}
	var interval, quantity int = n[0], n[0]

	if len(n) >= 2 {
		interval = n[1]
	}

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
