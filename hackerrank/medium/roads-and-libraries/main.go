package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type node struct {
	id         int32
	neighbours []int32
	isVisited  bool
}

//func main() {
//	//a := [][]int32{{1, 2}, {3, 1}, {2, 3}}
//	//roadsAndLibraries(3, 2, 1, a)
//
//	a := [][]int32{{1, 3}, {3, 4}, {2, 4}, {1, 2}, {2, 3}, {5, 6}}
//	roadsAndLibraries(6, 2, 1, a)
//
//}

func main() {
	f, err := os.Open("data.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()

	os.Stdin = f

	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	qTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
		checkError(err)
		n := int32(nTemp)

		mTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
		checkError(err)
		m := int32(mTemp)

		c_libTemp, err := strconv.ParseInt(firstMultipleInput[2], 10, 64)
		checkError(err)
		c_lib := int32(c_libTemp)

		c_roadTemp, err := strconv.ParseInt(firstMultipleInput[3], 10, 64)
		checkError(err)
		c_road := int32(c_roadTemp)

		var cities [][]int32
		for i := 0; i < int(m); i++ {
			citiesRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

			var citiesRow []int32
			for _, citiesRowItem := range citiesRowTemp {
				citiesItemTemp, err := strconv.ParseInt(citiesRowItem, 10, 64)
				checkError(err)
				citiesItem := int32(citiesItemTemp)
				citiesRow = append(citiesRow, citiesItem)
			}

			if len(citiesRow) != 2 {
				panic("Bad input")
			}

			cities = append(cities, citiesRow)
		}

		result := roadsAndLibraries(n, c_lib, c_road, cities)
		fmt.Println(result)
	}
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

// DFS to find connected components
func roadsAndLibraries(n int32, c_lib int32, c_road int32, cities [][]int32) int64 {

	var cost int64
	if c_road >= c_lib {
		cost = int64(n) * int64(c_lib)
		return cost
	}

	var nodes = make(map[int32]*node)

	var i int32
	for i = 1; i <= n; i++ {
		n := node{i, []int32{}, false}
		nodes[i] = &n
	}

	for _, val := range cities {
		nodes[val[0]].neighbours = append(nodes[val[0]].neighbours, val[1])
		nodes[val[1]].neighbours = append(nodes[val[1]].neighbours, val[0])
	}

	for _, node := range nodes {
		if node.isVisited == false {
			var numberOfNodesInComponent int64 = 0
			numberOfNodesInComponent++
			numberOfNodesInComponent = dfs(node, nodes, numberOfNodesInComponent)
			cost += (numberOfNodesInComponent-1)*int64(c_road) + int64(c_lib)
		}
	}

	return cost
}

func dfs(n *node, nodes map[int32]*node, numberOfNodesInComponent int64) int64 {
	n.isVisited = true
	for _, neighbour := range n.neighbours {
		if nodes[neighbour].isVisited == false {
			numberOfNodesInComponent++
			numberOfNodesInComponent = dfs(nodes[neighbour], nodes, numberOfNodesInComponent)
		}
	}
	return numberOfNodesInComponent
}
