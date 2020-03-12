package main

import "fmt"

func main() {
	//[]int 是切片
	var arr [3]int

	//数组不能使用new和make创建
	arr2 := [...]int{1,2,3}

	//二位数组
	arrToArr := [2][3]int{arr, arr2}

	fmt.Println(arr, arr2, arrToArr)
}
