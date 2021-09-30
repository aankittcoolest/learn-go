package main

import (
	"fmt"
)

/*
 * Complete the 'caesarCipher' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts following parameters:
 *  1. STRING s
 *  2. INTEGER k
 */

func caesarCipher(s string, k int32) string {
	// Write your code here
	var d []rune
	for _, b := range s {
		c := b
		if b >= 97 && b <= 122 {
			c = b + (k % 26)
			if c > 122 {
				c = 97 + (c - 122 - 1)
			}
		}
		if b >= 65 && b <= 90 {
			c = b + (k % 26)
			if c > 90 {
				c = 65 + (c - 90 - 1)
			}
		}
		d = append(d, c)
	}
	return string(d)
}

func main() {

	s := "middle-Outz"
	var k int32 = 2

	result := caesarCipher(s, k)

	fmt.Println(result)
}
