package main

import "fmt"

func main() {
	fmt.Println(solution(30, 7, []float64{0.2, 0.35, 0.4, 0.45}, []float64{1.1, 1.8, 2.3, 3.5}))
}

func solution(ride_time int, ride_distance int, cost_per_minute []float64, cost_per_mile []float64) []float64 {
	var costPerMile []float64
	var fareEstimator float64

	for idx, v := range cost_per_minute {
		fareEstimator = (v * float64(ride_time)) + (cost_per_mile[idx] * float64(ride_distance))
		//(Cost per minute) * (ride time) + (Cost per mile) * (ride distance)
		costPerMile = append(costPerMile, fareEstimator)
	}

	return costPerMile
}
