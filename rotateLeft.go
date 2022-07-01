package main

import "fmt"

func rotateLeft(d int32, arr []int32) []int32 {
	// Write your code here
	var rotatedArray []int32
	var i int32
	for i = d; i < int32(len(arr)); i++ {
		rotatedArray = append(rotatedArray, arr[i])
	}

	for i = 0; i < d; i++ {
		rotatedArray = append(rotatedArray, arr[i])
	}
	return rotatedArray
}
func main() {
	var arr []int32 = []int32{5, 1, 2, 3, 4}
	fmt.Println(rotateLeft(4, arr))
}
