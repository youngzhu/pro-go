package main

import (
	"math/rand"
	"sort"
)

func Example_random() {
	// 不设置种子的话，每次结果都一样
	//rand.Seed(time.Now().UnixNano())

	for i := 0; i < 5; i++ {
		Printfln("Value %v: %v", i, rand.Int())
	}

	// Output:
	//Value 0: 5577006791947779410
	//Value 1: 8674665223082153551
	//Value 2: 6129484611666145821
	//Value 3: 4037200794235010051
	//Value 4: 3916589616287113937
}

func intRange(min, max int) int {
	return rand.Intn(max-min) + min
}
func Example_random_range() {
	for i := 0; i < 5; i++ {
		Printfln("Value %v: %v", i, intRange(10, 20))
	}

	// Output:
	//Value 0: 11
	//Value 1: 17
	//Value 2: 17
	//Value 3: 19
	//Value 4: 11
}

func Example_shuffling() {
	var names = []string{"Alice", "Bob", "Charlie", "Dora", "Edith"}
	Printfln("Before: %v", names)

	rand.Shuffle(len(names), func(i, j int) {
		names[i], names[j] = names[j], names[i]
	})
	Printfln("After:  %v", names)

	// Output:
	//Before: [Alice Bob Charlie Dora Edith]
	//After:  [Charlie Alice Bob Edith Dora]
}

// 1. 从排过序的数组里搜索
// 2. 对于不存在的元素，不是返回-1，而是“如果存在，它应该在的位置
func search(ints []int, x int) {
	idx := sort.SearchInts(ints, x)
	Printfln("Index of %v: %v (present: %v)", x, idx, ints[idx] == x)
}
func Example_searching() {
	ints := []int{9, 4, 2, -1, 10}
	sort.Ints(ints)

	search(ints, 4)
	search(ints, 3)

	// output:
	//Index of 4: 2 (present: true)
	//Index of 3: 2 (present: false)
}
