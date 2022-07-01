package main

import (
	"fmt"
	"math"
)

func maxSubArraySum(array []float64) float64 {
	var length int = len(array)
	var best float64 = 0
	var sum float64 = 0
	for i := 0; i < length; i++ {
		sum = math.Max(array[i], sum+array[i])
		best = math.Max(best, sum)
	}

	return best
}

func main() {

	var array []float64 = []float64{-1, 2, 4, -3, 5, 2, -5, 2}

	var sum float64 = maxSubArraySum(array)
	fmt.Println("sum", sum)
}

// [−1 2 4 −3 5 2 −5 2]
