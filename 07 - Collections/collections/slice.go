package main

import "fmt"

// 会产生一个新的底层数组
func appendPure() {
	names := []string{"Kayak", "Lifejacket", "Paddle"}

	appendedNames := append(names, "Hat", "Gloves")

	names[0] = "Canoe"

	fmt.Println("names:", names)
	fmt.Println("appendedNames:", appendedNames)
}

// 只有有足够的容量，扩展后的切片仍然指向相同的底层数组
func appendWithEnoughCapacity() {
	names := make([]string, 3, 6)

	names[0] = "Kayak"
	names[1] = "Lifejacket"
	names[2] = "Paddle"

	appendedNames := append(names, "Hat", "Gloves")

	names[0] = "Canoe"

	fmt.Println("names:", names)
	fmt.Println("appendedNames:", appendedNames)
}

// 产生的切片共用底层数组
func createFromArray() {
	names := [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}

	// [low:high]
	// from: low
	// len: high-low
	some := names[1:3]
	all := names[:]

	some[0] = "Tom"

	fmt.Println("names:", names)
	fmt.Println("some:", some)
	fmt.Println("all:", all)
}
