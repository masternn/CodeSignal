package main

import (
	"fmt"
	"math"
)

func main() {
	departure := []float64{0, 0.4}
	destination := []float64{1, 0.6}

	fmt.Println(solution(departure, destination))
}

func solution(departure []float64, destination []float64) float64 {
	var axis, total float64

	for idx, val := range departure {
		if math.Floor(val) == math.Floor(destination[idx]) {
			axis = (math.Ceil(val) - val) + math.Ceil(destination[idx]) - destination[idx]
			if sum := (val + destination[idx]); sum < axis {
				axis = sum
			}
		} else {
			axis = math.Abs(destination[idx] - val)
		}
		total += axis
	}

	return total
}
