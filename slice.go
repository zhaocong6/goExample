package main

import "fmt"

func main() {
	var slice []int
	slice = append(slice, []int{1,2,3}...)

	slice2 := make([]int, 4)

	slice3 := make([][]int, 3)
	slice3[0] = slice
	slice3[1] = slice2

	fmt.Println(slice, slice2, slice3)
}
