package main

func hourglassSum(arr [][]int32) int32 {
	var maxSum int32
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			var sum int32 = 0
			sum += arr[i][j] + arr[i][j+1] + arr[i][j+2]
			sum += arr[i+2][j] + arr[i+2][j+1] + arr[i+2][j+2]
			sum += arr[i+1][j+1]
			if i == 0 && j == 0 {
				maxSum = sum
			}
			if sum > maxSum {
				maxSum = sum
			}
		}
	}
	return maxSum
}

func main() {
	// var a = [][]int32{
	// 	{1, 1, 1, 0, 0, 0},
	// 	{0, 1, 0, 0, 0, 0},
	// 	{1, 1, 1, 0, 0, 0},
	// 	{0, 0, 2, 4, 4, 0},
	// 	{0, 0, 0, 2, 0, 0},
	// 	{0, 0, 1, 2, 4, 0},
	// }
	var a = [][]int32{
		{0, -4, -6, 0, -7, -6},
		{-1, -2, -6, -8, -3, -1},
		{-8, -4, -2, -8, -8, -6},
		{-3, -1, -2, -5, -7, -4},
		{-3, -5, -3, -6, -6, -6},
		{-3, -6, 0, -8, -6, -7},
	}
	hourglassSum(a)
}
