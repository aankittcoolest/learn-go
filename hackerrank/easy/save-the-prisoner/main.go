package main

import (
	"fmt"
)

/*
 * Complete the 'saveThePrisoner' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. INTEGER m
 *  3. INTEGER s
 */

func saveThePrisoner(n int32, m int32, s int32) int32 {
	return ((m-1)+(s-1))%n + 1

}

func main() {

	var n int32 = 7
	var m int32 = 19
	var s int32 = 2

	result := saveThePrisoner(n, m, s)
	fmt.Println(result)

}
