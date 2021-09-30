package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	testcase2()
}

// Complete the countTriplets function below.
func countTriplets(arr []int64, r int64) int64 {
	lm := make(map[int64]int64)
	rm := make(map[int64]int64)

	for _, a := range arr {
		if e, ok := rm[a]; ok {
			rm[a] = e + 1
		} else {
			rm[a] = 1
		}
	}

	var sum, pivot, left, right, val1, val2 int64

	for i := 0; i < len(arr); i++ {
		pivot = arr[i]
		rm[pivot] = rm[pivot] - 1

		left = pivot / r
		right = pivot * r

		val1 = lm[left]
		val2 = rm[right]
		sum += val1 * val2
		lm[pivot] = lm[pivot] + 1
	}

	return sum
}

func testcase1() {
	arr := []int64{1, 3, 9, 9, 27, 81}
	result := countTriplets(arr, 3)
	fmt.Println(result)
}

func testcase2() {
	data, err := ioutil.ReadFile("test-case.txt")
	if err != nil {
		fmt.Println(err)
	}
	a := strings.Split(string(data), " ")

	var b = []int64{}

	for _, val := range a {
		c, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			fmt.Println(err)
		}
		b = append(b, c)
	}

	result := countTriplets(b, 3)
	fmt.Println(result)

}
