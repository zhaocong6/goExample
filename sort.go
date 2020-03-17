package main

import (
	"fmt"
	"sort"
)

type ByLength []string

func main() {
	strs := []string{"c", "b", "a"}
	sort.Strings(strs)
	fmt.Println(strs)

	ints := []int{2, 3, 1}
	sort.Ints(ints)
	fmt.Println(ints)

	//自定义排序
	var leng ByLength
	leng = []string{"apple", "kiwi", "banana"}
	sort.Sort(leng)
	fmt.Println(leng)
}

func (s ByLength) Len() int {
	return len(s)
}

func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}
