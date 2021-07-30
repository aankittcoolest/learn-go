package main

import "fmt"

func main() {
	arr := []int64{1, 5, 5, 5, 6}
	result := countTriplets(arr, 1)
	fmt.Println(result)
}

func countTriplets(arr []int64, r int64) int64 {
	a := make(map[int64]int64)

	for _, item := range arr {
		if val, ok := a[item]; ok {
			a[item] = val + 1
		} else {
			a[item] = 1
		}
	}

	var total int64 = 0
	for item := range a {
		if r == 1 {
			if val1,ok := a[item]; ok {
				if val == 3 {}
			}

		}
		if val1, ok := a[item]; ok {
			second_item := item * r
			if val2, ok := a[second_item]; ok {
				third_item := second_item * r
				if val3, ok := a[third_item]; ok {
					fmt.Println("First item", item, val1)
					fmt.Println("Second item", second_item, val2)
					fmt.Println("Third item", third_item, val3)
					total = total + (val1 * val2 * val3)
				}
			}

		}
	}

	return total
}
