package main

import (
	"fmt"
)

// Complete the maxXor function below.
func maxXor(arr []int32, queries []int32) []int32 {
	// solve here
	for i := range arr {
		fmt.Printf("%04b", arr[i])
		fmt.Println()
	}
	for i := range queries {
		fmt.Printf("%04b", queries[i])
		fmt.Println()
	}

	return []int32{3, 7, 15, 10}
}

func main() {

	arr := []int32{3, 7, 10, 15}
	queries := []int32{3, 6}

	result := maxXor(arr, queries)
	fmt.Println(result)

}
