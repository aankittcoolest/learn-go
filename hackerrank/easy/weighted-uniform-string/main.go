package main

import (
	"fmt"
)

/*
 * Complete the 'weightedUniformStrings' function below.
 *
 * The function is expected to return a STRING_ARRAY.
 * The function accepts following parameters:
 *  1. STRING s
 *  2. INTEGER_ARRAY queries
 */

func weightedUniformStrings(s string, queries []int32) []string {

	wm := make(map[int32]int32)

	var i int32
	for i = 1; i <= 26; i++ {
		wm[i] = 0
	}

	var initial_val int32
	var count int32 = 0
	for _, a := range s {
		w := int32(a) - 96

		if initial_val == w {
			count++
		} else {
			initial_val = w
			count = 1
		}

		_, ok := wm[w]
		if !ok {
			wm[w] = count
		} else {
			if count > wm[w] {
				wm[w] = count
			}
		}
	}

	items := []string{}
	for _, val := range queries {
		item := "No"
		for k, v := range wm {
			if val%k == 0 && val/k <= v {
				item = "Yes"
				break
			}
		}
		items = append(items, item)
	}

	// Write your code here
	return items

}

func main() {

	s := "abbcccdddd"
	var queries = []int32{1, 7, 5, 4, 15}

	result := weightedUniformStrings(s, queries)
	fmt.Println(result)

}
