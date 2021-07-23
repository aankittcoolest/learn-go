package main

import "fmt"

func main() {
	s := make([]int, 11)

	for i := range s {
		s[i] = i
	}

	for i := range s {
		if s[i]%2 == 0 {
			fmt.Println(s[i], " is even")
		} else {
			fmt.Println(s[i], " is odd")
		}
	}
}
