package main

import (
	"fmt"
	"sort"
	"strings"
)

func areAnagrams(s1, s2 string) bool {
	// 長さが異なる場合はアナグラムではない
	if len(s1) != len(s2) {
		return false
	}
	// 文字列をソートして比較
	s1Slice := strings.Split(s1, "")
	s2Slice := strings.Split(s2, "")
	sort.Strings(s1Slice)
	sort.Strings(s2Slice)
	return strings.Join(s1Slice, "") == strings.Join(s2Slice, "")
}

func main() {
	fmt.Println(areAnagrams("listen", "silent")) // true
}
