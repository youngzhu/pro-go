package main

import (
	"fmt"
)

func main() {
	pointerValues()

	/*
		names := [3]string{"Alice", "Charlie", "Bob"}

		secondPosition := &names[1]

		fmt.Println(*secondPosition)

		sort.Strings(names[:])

		fmt.Println(*secondPosition)
	*/
}

func pointerValues() {
	first := 100
	var second *int = &first

	first++

	var third int = *second
	*second++

	// 指向指针的指针
	forth := &second

	fmt.Println(first)
	fmt.Println(second)
	fmt.Println(*second)
	fmt.Println(third)
	fmt.Println("forth:", forth)
	fmt.Println("*forth:", *forth)
	fmt.Println("**forth:", **forth)
}
