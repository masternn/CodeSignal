package main

import (
	"fmt"
)

func main() {
	fares := []float64{
		0.7, 1, 1.3, 1.5, 1.7,
	}
	fmt.Println(solution(15, fares))
}

func solution(l int, fares []float64) string {
	ubers := []string{"UberX", "UberXL", "UberPlus", "UberBlack", "UberSUV"}

	for i := (len(fares) - 1); i >= 0; i-- {
		if (fares[i] * float64(l)) <= 20 {
			return ubers[i]
		}
	}

	return ubers[0]
}

// For l = 30 and fares = [0.3, 0.5, 0.7, 1, 1.3], the output should be
// solution(l, fares) = "UberXL".
