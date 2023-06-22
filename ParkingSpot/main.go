package main

import (
	"fmt"
)

func main() {
	car := []int{2, 1}
	luckySpot := []int{2, 2, 3, 2}
	parkingLot := [][]int{
		{0, 1, 1},
		{0, 0, 1},
		{1, 0, 0},
		{1, 0, 0}}

	fmt.Println("solution", solution(car, parkingLot, luckySpot))
}

func solution(carDimensions []int, parkingLot [][]int, luckySpot []int) bool {
	var x, y int
	x1 := luckySpot[0]
	y1 := luckySpot[1]
	x2 := luckySpot[2]
	y2 := luckySpot[3]

	beginX := x1
	endX := x2

	beginY := 0
	endY := len(parkingLot[0]) - 1

	//Vertical
	if x2 >= y2 {
		beginX = 0
		beginY = y1
		endY = y2
	}

	entrance := false
	inside := false
	space := carDimensions[1]

	for y = beginY; y <= endY; y++ {
		for x = beginX; x <= endX; x++ {
			//Are we inside luckySpot and it's occupied ?
			if (x >= x1 && x <= x2) && (y >= y1 && y <= y2) {
				if parkingLot[x][y] == 1 {
					return false
				}

				if !inside && !entrance && space == carDimensions[1] {
					entrance = true
				}

				if space != carDimensions[1] {
					space = carDimensions[1]
				}

				inside = true
			} else {
				if entrance {
					continue
				}
				space -= parkingLot[x][y]
			}
		}
	}

	return entrance || (space == carDimensions[1])
}
