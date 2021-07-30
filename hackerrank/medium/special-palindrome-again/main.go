package main

import (
	"fmt"
)

// Complete the substrCount function below.
func substrCount(n int32, s string) int64 {
	var count int64 = 0
	var spl_count int64 = 0

	for i := 0; i < len(s); i++ {
		count++
		spl_count = 0
		pivot := false
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] && pivot == false {
				count++
				spl_count++
			} else if s[i] != s[j] && pivot == true {
				break
			} else if s[i] != s[j] && pivot == false {
				spl_count++
				pivot = true
			} else if s[i] == s[j] && pivot == true {
				spl_count--
				if spl_count == 0 {
					count++
				}
			}
		}
	}
	return count
}

func main() {
	var n int32 = 4
	s := "asasd"

	result := substrCount(n, s)
	fmt.Println(result)

}
