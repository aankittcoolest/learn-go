package main

import (
	"fmt"
	"math"
)

/*
 * Complete the 'commonChild' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. STRING s1
 *  2. STRING s2
 */

func commonChild(s1 string, s2 string) int32 {
	s := make([][]int32, len(s1)+1)

	for i := range s {
		s[i] = make([]int32, len(s2)+1)
	}

	A := make([]string, len(s1)+1)
	B := make([]string, len(s2)+1)

	A[0] = ""
	B[0] = ""
	for i := range s1 {
		A[i+1] = fmt.Sprintf("%c", s1[i])
	}
	for i := range s2 {
		B[i+1] = fmt.Sprintf("%c", s2[i])
	}

	var max int32 = 0
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(B); j++ {
			if i == 0 || j == 0 {
				s[i][i] = 0
				continue
			}
			if A[i] == B[j] {
				s[i][j] = 1 + s[i-1][j-1]
			} else {
				s[i][j] = int32(math.Max(float64(s[i-1][j]), float64(s[i][j-1])))
			}
			if s[i][j] > max {
				max = s[i][j]
			}
		}
	}
	return max

}

func main() {
	s1 := "AAA"

	s2 := "AAA"

	result := commonChild(s1, s2)

	fmt.Println(result)
}
