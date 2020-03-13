package main

import "fmt"

type pintInt int

func main() {
	var p pintInt
	p = 1
	p.dump()
}

func (p pintInt) dump() {
	fmt.Println(p)
}
