package main

import (
	"fmt"
)

// Complete the freqQuery function below.
func freqQuery(queries [][]int32) []int32 {

	f := make(map[int32]int32)
	c := make(map[int32]int32)

	var out = []int32{}
	for _, v := range queries {
		if v[0] == 1 {
			c[f[v[1]]] -= 1
			f[v[1]] += 1
			c[f[v[1]]] += 1
		} else if v[0] == 2 {
			if f[v[1]] > 0 {
				c[f[v[1]]] -= 1
				f[v[1]] -= 1
				c[f[v[1]]] += 1

			}
		} else {
			if c[v[1]] > 0 {
				out = append(out, 1)
			} else {
				out = append(out, 0)
			}
		}
	}

	return out

}

func main() {
	test_case1()

}

func test_case1() {
	var queries = make([][]int32, 10)

	queries[0] = []int32{1, 3}
	queries[1] = []int32{2, 3}
	queries[2] = []int32{3, 2}
	queries[3] = []int32{1, 4}
	queries[4] = []int32{1, 5}
	queries[5] = []int32{1, 5}
	queries[6] = []int32{1, 4}
	queries[7] = []int32{3, 2}
	queries[8] = []int32{2, 4}
	queries[9] = []int32{3, 2}

	ans := freqQuery(queries)

	fmt.Println(ans)

}

func test_case2() {
	var queries = make([][]int32, 4)

	queries[0] = []int32{1, 2}
	queries[1] = []int32{3, 0}
	queries[2] = []int32{2, 2}
	queries[3] = []int32{3, 0}
	ans := freqQuery(queries)

	fmt.Println(ans)

}

func test_case3() {

	var queries = make([][]int32, 81)

	queries[0] = []int32{3, 5}
	queries[1] = []int32{3, 3}
	queries[2] = []int32{1, 10000004}
	queries[3] = []int32{1, 10000016}
	queries[4] = []int32{1, 10000011}
	queries[5] = []int32{3, 10}
	queries[6] = []int32{1, 10000006}
	queries[7] = []int32{3, 5}
	queries[8] = []int32{2, 4}
	queries[9] = []int32{2, 3}
	queries[10] = []int32{2, 6}
	queries[11] = []int32{1, 10000037}
	queries[12] = []int32{3, 10}
	queries[13] = []int32{3, 3}
	queries[14] = []int32{1, 10000013}
	queries[15] = []int32{1, 10000013}
	queries[16] = []int32{3, 10}
	queries[17] = []int32{3, 10}
	queries[18] = []int32{1, 10000025}
	queries[19] = []int32{1, 10000021}
	queries[20] = []int32{2, 7}
	queries[21] = []int32{1, 10000002}
	queries[22] = []int32{3, 7}
	queries[23] = []int32{3, 9}
	queries[24] = []int32{2, 9}
	queries[25] = []int32{2, 8}
	queries[26] = []int32{3, 4}
	queries[27] = []int32{3, 4}
	queries[28] = []int32{1, 10000025}
	queries[29] = []int32{3, 6}
	queries[30] = []int32{1, 10000037}
	queries[31] = []int32{2, 9}
	queries[32] = []int32{2, 8}
	queries[33] = []int32{1, 10000042}
	queries[34] = []int32{2, 7}
	queries[35] = []int32{2, 10}
	queries[36] = []int32{1, 10000002}
	queries[37] = []int32{2, 2}
	queries[38] = []int32{2, 4}
	queries[39] = []int32{2, 5}
	queries[40] = []int32{1, 10000005}
	queries[41] = []int32{1, 10000021}
	queries[42] = []int32{1, 10000031}
	queries[43] = []int32{3, 4}
	queries[44] = []int32{1, 10000013}
	queries[45] = []int32{3, 8}
	queries[46] = []int32{3, 2}
	queries[47] = []int32{3, 4}
	queries[48] = []int32{1, 10000024}
	queries[49] = []int32{3, 5}
	queries[50] = []int32{2, 2}
	queries[51] = []int32{2, 5}
	queries[52] = []int32{3, 3}
	queries[53] = []int32{2, 1}
	queries[54] = []int32{3, 6}
	queries[55] = []int32{1, 10000021}
	queries[56] = []int32{2, 4}
	queries[57] = []int32{3, 1}
	queries[58] = []int32{3, 5}
	queries[59] = []int32{1, 10000049}
	queries[60] = []int32{1, 10000010}
	queries[61] = []int32{1, 10000036}
	queries[62] = []int32{2, 8}
	queries[63] = []int32{1, 10000001}
	queries[64] = []int32{3, 2}
	queries[65] = []int32{1, 10000017}
	queries[66] = []int32{1, 10000002}
	queries[67] = []int32{1, 10000003}
	queries[68] = []int32{3, 2}
	queries[69] = []int32{1, 10000048}
	queries[70] = []int32{1, 10000009}
	queries[71] = []int32{3, 3}
	queries[72] = []int32{1, 10000031}
	queries[73] = []int32{1, 10000044}
	queries[74] = []int32{3, 7}
	queries[75] = []int32{1, 10000025}
	queries[76] = []int32{1, 10000041}
	queries[77] = []int32{1, 10000031}
	queries[78] = []int32{1, 10000049}
	queries[79] = []int32{2, 6}
	queries[80] = []int32{1, 10000015}

	ans := freqQuery(queries)

	fmt.Println(ans)
}
