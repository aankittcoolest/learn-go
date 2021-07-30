package main

import "fmt"

var count int64 = 0

func main() {

	var a = []int32{7, 5, 3, 1}
	mergeSort(a)
	fmt.Println(count)
}

func mergeSort(arr []int32) {
	if len(arr) > 1 {
		mid := int32(len(arr)) / 2
		L := make([]int32, len(arr[:mid]))
		copy(L, arr[:mid])
		R := make([]int32, len(arr[mid:]))
		copy(R, arr[mid:])
		mergeSort(L)
		mergeSort(R)

		var i, j, k int64 = 0, 0, 0

		for {
			if i < int64(len(L)) && j < int64(len(R)) {
				if L[i] <= R[j] {
					arr[k] = L[i]
					i++
				} else {
					count += int64(len(L)) - i
					arr[k] = R[j]
					j++
				}
				k++
			} else {
				break
			}
		}

		for {
			if i < int64(len(L)) {
				arr[k] = L[i]
				i++
				k++
			} else {
				break
			}
		}

		for {
			if j < int64(len(R)) {
				arr[k] = R[j]
				j++
				k++
			} else {
				break
			}
		}

	}
}
