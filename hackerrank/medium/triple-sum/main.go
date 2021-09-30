package main

import (
	"fmt"
	"sort"
)

// Complete the triplets function below.
func triplets(a []int32, b []int32, c []int32) int64 {
	aa := []int{}
	bb := []int{}
	cc := []int{}

	am := make(map[int32]bool)
	bm := make(map[int32]bool)
	cm := make(map[int32]bool)

	for _, item := range a {
		if _, found := am[item]; !found {
			am[item] = true
		}
	}
	for _, item := range b {
		if _, found := bm[item]; !found {
			bm[item] = true
		}
	}
	for _, item := range c {
		if _, found := cm[item]; !found {
			cm[item] = true
		}
	}

	for key := range am {
		aa = append(aa, int(key))
	}
	for key := range bm {
		bb = append(bb, int(key))
	}
	for key := range cm {
		cc = append(cc, int(key))
	}

	sort.Ints(aa)
	sort.Ints(bb)
	sort.Ints(cc)

	var count int64 = 0

	for k := 0; k < len(bb); k++ {
		var count_aa int64 = 0
		var count_cc int64 = 0
		for i := 0; i < len(aa); i++ {
			if aa[i] <= bb[k] {
				count_aa++
			} else {
				break
			}
		}

		for i := 0; i < len(cc); i++ {
			if cc[i] <= bb[k] {
				count_cc++
			} else {
				break
			}
		}

		count += count_aa * count_cc
	}

	return count

}

func main() {

	var arra = []int32{1, 4, 5}
	var arrb = []int32{2, 3, 3}
	var arrc = []int32{1, 2, 3}

	ans := triplets(arra, arrb, arrc)
	fmt.Println(ans)

}
