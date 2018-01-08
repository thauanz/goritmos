package main

import (
	"fmt"
)

type Peg struct {
	Letter string
	Index  int
}

var pegs []Peg = []Peg{
	{"A", 1},
	{"B", 2},
	{"C", 3},
}

func main() {
	towersOfHanoi(5, "A", "B")
}

func towersOfHanoi(disks int, fromPeg, toPeg string) {
	if disks <= 0 {
		return
	}

	towersOfHanoi(disks-1, fromPeg, getSparePeg(fromPeg, toPeg))
	fmt.Println(disks, "-", fromPeg, "=>", toPeg)
	towersOfHanoi(disks-1, getSparePeg(fromPeg, toPeg), toPeg)
}

func getSparePeg(pegLetter1, pegLetter2 string) string {
	peg1 := convertLetterToPeg(pegLetter1)
	peg2 := convertLetterToPeg(pegLetter2)

	if peg1 == nil || peg2 == nil {
		return ""
	}
	sparePegIndex := len(pegs) - (peg1.Index - 1) - (peg2.Index - 1)
	return pegs[sparePegIndex].Letter
}

func convertLetterToPeg(pegLetter string) *Peg {
	for _, peg := range pegs {
		if pegLetter == peg.Letter {
			return &peg
		}
	}
	return nil
}
