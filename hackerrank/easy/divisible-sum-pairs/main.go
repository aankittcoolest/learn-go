package main

import "fmt"

func main() {

	var n int32 = 6
	var k int32 = 3
	ar := []int32{1, 3, 2, 6, 1, 2}

	fmt.Println(divisibleSumPairs(n, k, ar))

}

func divisibleSumPairs(n int32, k int32, ar []int32) int32 {

	var count int32 = 0
	for i := range ar {
		subAr := ar[i+1:]
		for j := range subAr {
			if (ar[i]+subAr[j])%k == 0 {
				count++
			}
		}
	}

	// Write your code here
	return count

}
