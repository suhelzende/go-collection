package main

import (
	"fmt"
	"go-collection/heap"
)

// Example: Given an array return return Kth smaller element
//
// number = [4,8,2,7,-1,23] k = 4
//
// ans = 7
func main() {
	list := []int{4, 8, 2, 7, -1, 23}

	var k int = 4
	comparator := func(t1, t2 int) bool { return t1 <= t2 }
	mh := heap.CreateHeapFromList[int](list, comparator)

	for i := 1; i < k && i < len(list); i++ {
		mh.Get()
	}
	ele, _ := mh.Get()
	fmt.Println(ele)

	// mh.Print()
	fmt.Print()
	fmt.Println("Hello World")
}
