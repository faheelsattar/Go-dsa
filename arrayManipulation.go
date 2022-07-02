package main

import "fmt"

//----> NOT OPTIMIZED <-----
func arrayManipulation(n int32, queries [][]int32) int64 {
	// Write your code here
	var i int32
	var j int32
	var max int64
	data := make([][]int32, n)
	for i = 0; i < n; i++ {
		data[0] = append(data[0], 0)
	}

	for i := 0; i < len(queries); i++ {
		row := queries[i]
		a := row[0]
		b := row[1]
		k := row[2]
		data[i+1] = append(data[i+1], data[i]...)
		for j = a - 1; j < b; j++ {
			data[i+1][j] = data[i+1][j] + k
			if data[i+1][j] > int32(max) {
				max = int64(data[i+1][j])
			}
		}
	}

	fmt.Println("data", data)
	return max
}

func main() {
	var a = [][]int32{
		{1, 5, 3},
		{4, 8, 7},
		{6, 9, 1},
	}
	fmt.Println(arrayManipulation(10, a))
}
