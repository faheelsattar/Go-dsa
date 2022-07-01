package main

import "fmt"

func matchingStrings(strings []string, queries []string) []int32 {
	// Write your code here
	dataObject := make(map[string]int32)
	var occurrence []int32

	for i := 0; i < len(queries); i++ {
		dataObject[queries[i]] = 0
	}

	for i := 0; i < len(strings); i++ {
		if _, ok := dataObject[strings[i]]; ok {
			dataObject[strings[i]] += 1
		}
	}

	for k := range dataObject {
		occurrence = append(occurrence, dataObject[k])
	}
	return occurrence
}
func main() {

	var strings []string = []string{"abcde", "sdaklfj", "asdjf", "na", "basdn", "sdaklfj", "asdjf", "na", "asdjf", "na", "basdn", "sdaklfj", "asdjf"}
	var queries []string = []string{"abcde", "sdaklfj", "asdjf", "na", "basdn"}

	fmt.Println(matchingStrings(strings, queries))
}
