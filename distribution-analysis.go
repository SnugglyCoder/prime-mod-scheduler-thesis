package main

import (
	"fmt"
)

const (
	trials = 10_000_000_000
)

func main() {
	twoCounter := 0
	threeCounter := 0
	fiveCounter := 0
	sevenCounter := 0
	for i := 0; i < trials; i++ {
		if i % 2 == 0 {
			twoCounter++
			continue
		}
		if i % 3 == 0 {
			threeCounter++
			continue
		}
		if i % 5 == 0 {
			fiveCounter++
			continue
		}
		if i % 7 == 0 {
			sevenCounter++
			continue
		}
		twoCounter++
	}
	fmt.Printf("Results:\n2: %f\n3: %f\n5: %f\n7: %f\n", float64(twoCounter)/trials, float64(threeCounter)/trials, float64(fiveCounter)/trials, float64(sevenCounter)/trials)
}