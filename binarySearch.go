package main

import "fmt"

func binarySearch(value int, array []int) bool {
	var i int = 0
	var j int = len(array)

	for i < j {
		mid := (i + j) / 2

		if array[mid] == value {
			return true
		}

		if array[mid] > value {
			j = mid - 1
		}

		if array[mid] < value {
			i = mid + 1
		}
	}
	return false
}

func main() {
	value := 20
	array := []int{1, 5, 6, 10, 13, 15, 16, 20}

	found := binarySearch(value, array)
	fmt.Println("founded: ", found)
}

// [1,5,6,10,13,15,16,20]
