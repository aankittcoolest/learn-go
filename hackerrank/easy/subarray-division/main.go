package main

import "fmt"

func main() {
	s := []int32{1, 2, 1, 3, 2}
	var d int32 = 3
	var m int32 = 2

	// s := []int32{1, 1, 1, 1, 1, 1}
	// var d int32 = 3
	// var m int32 = 2

	output := birthday(s, d, m)
	fmt.Println(output)

}

func birthday(s []int32, d int32, m int32) int32 {
	var count int32 = 0
	for i := range s {
		if int32(len(s)) >= int32(i)+m {
			a := s[i : m+int32(i)]

			var sum int32 = 0
			for j := range a {
				sum += a[j]
			}
			if sum == d {
				count++
			}
		}
	}
	return count

}
