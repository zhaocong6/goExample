package main

import (
	"fmt"
	"github.com/shopspring/decimal"
)

func main() {
	//非负数表示
	var _ uint64
	//高精度计算
	s, _ := decimal.NewFromString("1.0000000000000000002")
	s2, _ := decimal.NewFromString("0.0000000000000000002")

	mul := s.Mul(s2).String()
	fmt.Println(mul)
}
