package main

import "fmt"

func main() {
	//  arr := []int32{1,4,3,2}
	//  arr := []int32{7,1,3,2,4,5,6}
	//arr := []int32{1, 2, 3, 5, 4, 6, 7, 8}
	//arr := []int32{4, 1, 2, 3}
	//arr := []int32{2, 1, 5, 3, 4}
	arr := []int32{1, 2, 5, 3, 7, 8, 6, 4}
	minimumBribes(arr)
}

// Traverse the array from back and try to match index with the element within range 2
func minimumBribes(q []int32) {
	var bribeCounts int32
	var i int32
	var b []int32
	for i = int32(len(q) - 1); i >= 1; i-- {
		if i+1 != q[i] {
			if i-2 >= 0 {
				for j := i; j >= i-2; j-- {
					if q[j] != i+1 {
						bribeCounts++
						b = append(b, q[j])
					} else {
						q[i] = q[j]
						k := i - 1
						for _, val := range b {
							q[k] = val
							k--
						}
						b = nil
						break
					}
				}
				if b != nil {
					fmt.Println("Too chaotic")
					return
				}
			} else if i == 1 && q[i] != q[i-1] {
				bribeCounts++
			}
		}
	}

	fmt.Println(bribeCounts)
}

//---------------------------
//1 2 5 3 7 8 6 4
//1 2 5 3 7 6 4 8 (2)
//1 2 5 3 6 4 7 8 (2)
//1 2 5 3 4 6 7 8 (1)
//1 2 3 4 5 6 7 8 (2)
//---------------------------
//
//--------------
//2 1 5 3 4
//2 1 3 4 5 (2)
//1 2 3 4 5 (1)
//--------------
