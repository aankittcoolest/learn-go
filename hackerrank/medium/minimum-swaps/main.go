package main

func main() {
	//  arr := []int32{1,4,3,2}
	//  arr := []int32{7,1,3,2,4,5,6}
	arr := []int32{4, 3, 1, 2}
	minimumSwaps(arr)
}

// Complete the minimumSwaps function below.
func minimumSwaps(arr []int32) int32 {

	var a = make(map[int32]int32)

	for key, val := range arr {
		a[int32(key)+1] = val
	}

	var swapCounts int32 = 0
	for key, val := range a {
		if key == val {
			continue
		}
		for {
			if key != a[key] {
				swapCounts++
				temp := a[key]
				a[key] = a[val]
				a[val] = temp
				val = a[key]
				continue
			}
			break
		}
	}
	return swapCounts
}
