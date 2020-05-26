package main

import (
	"fmt"
	"github.com/google/btree"
)

type testTree struct {
	num int
}

func (t *testTree) Less(item btree.Item) bool {
	if i, ok := item.(*testTree); ok {
		return t.num < i.num
	}
	return false
}

func main() {
	t := btree.New(1000)

	num := 1000

	for i := 0; i < num; i++ {
		t.ReplaceOrInsert(&testTree{num: i})
	}

	desNum := 0

	t.Descend(func(i btree.Item) bool {
		defer func() {
			desNum++
		}()

		if desNum >= 10 {
			return false
		}

		fmt.Println(i)

		return true
	})
}
