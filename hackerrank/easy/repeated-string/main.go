package main

import (
	"fmt"
	"regexp"
)

/*
 * Complete the 'repeatedString' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts following parameters:
 *  1. STRING s
 *  2. LONG_INTEGER n
 */

func repeatedString(s string, n int64) int64 {
	r, _ := regexp.Compile("a")
	times := len(r.FindAllString(s, -1))
	if times == 0 {
		return 0
	}
	quo, rem := n/int64(len(s)), n%int64(len(s))

	substr := s[:rem]
	remain_times := len(r.FindAllString(substr, -1))

	return quo*int64(times) + int64(remain_times)

}

func main() {
	s := "a"
	var n int64 = 1000000000000
	result := repeatedString(s, n)
	fmt.Println(result)
}
