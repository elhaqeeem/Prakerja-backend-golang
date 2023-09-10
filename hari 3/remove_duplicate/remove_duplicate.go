// Golang program to remove duplicate
// values from Slice
package main

import (
	"fmt"
)

func removeDuplicateValues(stringSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}

	// If the key(values of the slice) is not equal
	// to the already present value in new slice (list)
	// then we append it. else we jump on another element.
	for _, entry := range stringSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func main() {

	// Assigning values to the slice
	stringSliceValues := []string{"as","ac","ad","as"}

	

	// Calling function where we
	// are removing the duplicates
	removeDuplicateValuesSlice := removeDuplicateValues(stringSliceValues)

	// Printing the filtered slice
	// without duplicates
	fmt.Println(removeDuplicateValuesSlice)
}
