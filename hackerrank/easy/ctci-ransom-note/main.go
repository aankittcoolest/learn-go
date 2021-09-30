package main

import (
	"fmt"
	"strings"
)

/*
 * Complete the 'checkMagazine' function below.
 *
 * The function accepts following parameters:
 *  1. STRING_ARRAY magazine
 *  2. STRING_ARRAY note
 */

func checkMagazine(magazine []string, note []string) {

	m := make(map[string]int32)

	for _, item := range magazine {
		if val, ok := m[item]; !ok {
			m[item] = 1
		} else {
			m[item] = val + 1
		}
	}

	for _, item := range note {
		if val, ok := m[item]; !ok {
			fmt.Println("No")
			return
		} else {
			if val == 0 {
				fmt.Println("No")
				return
			} else {
				m[item] = val - 1
			}
		}
	}
	fmt.Println("Yes")

}

func main() {

	magazine := "give me one grand today night"
	note := "give one grand today"
	checkMagazine(strings.Split(magazine, " "), strings.Split(note, " "))
}
