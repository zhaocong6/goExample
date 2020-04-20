package main

import (
	"fmt"
	"unsafe"
)

type Num struct {
	i int64
	j int64
}

func main() {
	n := Num{
		i: 1,
		j: 1,
	}

	nPointer := unsafe.Pointer(&n)
	niPointer := (*int)(nPointer)
	*niPointer = 2

	njPointer := (*int)(unsafe.Pointer(uintptr(nPointer) + unsafe.Offsetof(n.j)))
	*njPointer = 2
	fmt.Println(nPointer, n)
}
