package main

import (
	"fmt"
)

/*
 * Complete the 'steadyGene' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts STRING gene as parameter.
 */

func steadyGene(gene string) int32 {
	// Write your code here
	m := make(map[string]int)
	n := make(map[string]int)

	c := len(gene)
	m["G"] = c / 4
	m["A"] = c / 4
	m["C"] = c / 4
	m["T"] = c / 4

	n["G"] = 0
	n["A"] = 0
	n["C"] = 0
	n["T"] = 0

	for _, val := range gene {
		n[string(val)] += 1
	}

	initial := true
	count := 0
	for _, val := range gene {
		if m[string(val)] > 0 {
			m[string(val)] -= 1
			if !initial {
				count++
			}
			continue
		}
		initial = false
		min_key, min_val := min_val(n)
		if min_val == c/4 {
			break
		}
		fmt.Println(m)
		fmt.Println(n)
		n[min_key] = min_val + 1
		n[string(val)] -= 1
		m[min_key] -= 1
		fmt.Println(m)
		fmt.Println(n)
		count++
	}

	return int32(count)

}

func min_val(m map[string]int) (string, int) {
	min_val := m["G"]
	min_key := "G"
	for key, val := range m {
		if val < min_val {
			min_val = val
			min_key = key
		}
	}
	return min_key, min_val
}

func main() {
	gene := "ACTGAAAG"
	result := steadyGene(gene)

	fmt.Println(result)
}
