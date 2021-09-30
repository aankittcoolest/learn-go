package main

import (
	"fmt"
	"sort"
	"strings"
)

func sortString(s string) string {
	w := strings.Split(s, "")
	sort.Strings(w)
	return strings.Join(w, "")
}

func sum(n int32) int32 {
	var sum int32 = 0
	var i int32 = 0
	for i = 1; i <= n; i++ {
		sum += i
	}
	return sum
}

func sherlockAndAnagrams(s string) int32 {

	m := make(map[string]int32)

	for i := 0; i < len(s); i++ {
		b := ""
		for j := i; j < len(s); j++ {
			b = b + fmt.Sprintf("%c", s[j])
			if len(b) > 1 {
				b = sortString(b)
			}
			if v, ok := m[b]; !ok {
				m[b] = 1
			} else {
				m[b] = v + 1
			}
		}
	}

	var t int32 = 0
	for _, v := range m {
		if v > 1 {
			t += sum(v - 1)
		}
	}

	return t

}

func main() {
	s := "kkkk"
	result := sherlockAndAnagrams(s)
	fmt.Println(result)

}
