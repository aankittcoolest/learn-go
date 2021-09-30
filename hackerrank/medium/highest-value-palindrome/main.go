package main

import (
	"fmt"
	"io/ioutil"
)

func highestValuePalindrome(s string, n int32, k int32) string {
	// Write your code here

	chars := []rune(s)
	mid := (len(chars)) / 2

	diff_count := 0
	for i := 0; i < mid; i++ {
		if chars[i] != chars[len(chars)-1-i] {
			diff_count++
		}
	}

	for i := 0; i < mid; i++ {
		if diff_count > 0 {
			if k == 0 {
				return "-1"
			}

			start_element := chars[i]
			end_element := chars[len(chars)-1-i]

			if start_element != end_element && start_element != 57 && end_element != 57 && k > int32(diff_count) {
				chars[i] = 57
				chars[len(chars)-1-i] = 57
				k = k - 2
				diff_count--
			} else if start_element != 57 && end_element != 57 && k > int32(diff_count) {
				chars[i] = 57
				chars[len(chars)-1-i] = 57
				k = k - 2
			} else if start_element != end_element {
				if start_element > end_element {
					chars[len(chars)-1-i] = start_element
				}
				if end_element > start_element {
					chars[i] = end_element

				}
				k = k - 1
				diff_count--
			}
		}
	}

	if diff_count > 0 {
		return "-1"
	}

	for i := 0; i < mid; i++ {
		if k >= 2 {
			if chars[i] != 57 {
				chars[i] = 57
				chars[len(chars)-1-i] = 57
				k = k - 2
			}
		}
	}

	if k > 0 && len(chars)%2 != 0 {
		chars[mid] = 57
	}

	result := string(chars)
	return result

}

func main() {
	s := read_from_file()
	var n int32 = 77543
	var k int32 = 58343

	// s := "932239"
	// var n int32 = 6
	// var k int32 = 2

	result := highestValuePalindrome(s, n, k)

	fmt.Println(result)

	read_from_file()
}

func read_from_file() string {
	data, err := ioutil.ReadFile("test_case.txt")
	if err != nil {
		fmt.Println("something happened...")
	}
	return string(data)
}
