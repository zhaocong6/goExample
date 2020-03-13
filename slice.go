package main

import "fmt"

func main() {
	var slice []int
	slice = append(slice, []int{1,2,3}...)

	slice2 := make([]int, 4)

	slice3 := make([][]int, 3)
	slice3[0] = slice
	slice3[1] = slice2

	arr := [...]int{1,2,3,4}
	//通过切割array实现slice
	slice4 := arr[:3]

	slice3[2] = slice4

	fmt.Println(slice, slice2, slice3, slice4)
}
