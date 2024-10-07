package main

import (
	"fmt"
	"math"
)

func euclideanDistance(x1, y1, x2, y2 float64) float64 {
	dx := x2 - x1
	dy := y2 - y1
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	distance := euclideanDistance(1, 2, 4, 6)
	fmt.Printf("2点間の距離は %.2f です\n", distance)
}
