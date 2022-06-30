package main

import "fmt"

func reverseArray(a []int32) []int32 {
	// Write your code here
	var reverse []int32

	for i := len(a) - 1; i >= 0; i-- {
		reverse = append(reverse, a[i])
	}

	return reverse
}

func main() {
	newArray := []int32{1, 4, 3, 2}
	fmt.Println(reverseArray(newArray))
}
