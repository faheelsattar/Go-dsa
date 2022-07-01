package main

import "fmt"

func dynamicArray(n int32, queries [][]int32) []int32 {
	// Write your code here
	var lastAnswer int32 = 0
	arr := make([][]int32, n)
	var answers []int32
	for i := 0; i < len(queries); i++ {
		if queries[i][0] == 1 {
			idx := (queries[i][1] ^ lastAnswer) % n
			arr[idx] = append(arr[idx], queries[i][2])
		} else {
			idx := (queries[i][1] ^ lastAnswer) % n
			lastAnswer = arr[idx][queries[i][2]%int32(len(arr[idx]))]
			answers = append(answers, lastAnswer)
		}
	}
	return answers
}

func main() {
	var a = [][]int32{
		{1, 0, 5},
		{1, 1, 7},
		{1, 0, 3},
		{2, 1, 0},
		{2, 1, 1},
	}
	fmt.Println(dynamicArray(2, a))
}
