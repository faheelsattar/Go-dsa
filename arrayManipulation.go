package main

import (
	"fmt"
)

//----> NOT OPTIMIZED <-----
// func arrayManipulation(n int, queries [][]int) int {
// 	// Write your code here
// 	var i int
// 	var j int
// 	var max int
// 	data := make([][]int, n)
// 	for i = 0; i < n; i++ {
// 		data[0] = append(data[0], 0)
// 	}

// 	for i := 0; i < len(queries); i++ {
// 		row := queries[i]
// 		a := row[0]
// 		b := row[1]
// 		k := row[2]
// 		data[i+1] = append(data[i+1], data[i]...)
// 		for j = a - 1; j < b; j++ {
// 			data[i+1][j] = data[i+1][j] + k
// 			if data[i+1][j] > int(max) {
// 				max = int(data[i+1][j])
// 			}
// 		}
// 	}

// 	fmt.Println("data", data)
// 	return max
// }

func arrayManipulation(n int, queries [][]int) int {
	arr := make([]int, n)
	var max int = 0
	var curr int = 0

	for i := 0; i < len(queries); i++ {
		row := queries[i]
		a := row[0] - 1
		b := row[1] - 1
		k := row[2]
		arr[a] += k
		if b+1 < int(len(arr)) {
			arr[b+1] -= k
		}
	}

	for i := 0; i < int(n); i++ {
		curr += arr[i]
		if curr > int(max) {
			max = int(curr)
		}
	}
	return max
}

func main() {
	var a = [][]int{
		{1, 5, 3},
		{4, 8, 7},
		{6, 9, 1},
	}
	fmt.Println(arrayManipulation(10, a))
}
