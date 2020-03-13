package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4}
	for _, num := range arr {
		fmt.Println(num)
	}

	chinese := "go语言"
	chineseB := make([]rune, len(chinese))
	//迭代中文
	for l, c := range chinese{
		chineseB[l] = c
	}

	fmt.Println(string(chineseB))
}
