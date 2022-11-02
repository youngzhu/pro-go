package main

import (
	"fmt"
	"reflect"
)

func Example_appendFromArray() {
	names := [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}

	// [low:high]
	// from: low
	// len: high-low
	some := names[1:3]
	all := names[:]

	some = append(some, "Gloves")

	fmt.Println("names:", names)
	fmt.Printf("names len:%d, cap:%d\n", len(names), cap(names))
	fmt.Println("some:", some)
	fmt.Printf("some len:%d, cap:%d\n", len(some), cap(some))
	fmt.Println("all:", all)
	fmt.Printf("all len:%d, cap:%d\n", len(all), cap(all))

	// Output:
	// names: [Kayak Lifejacket Paddle Gloves]
	// names len:4, cap:4
	// some: [Lifejacket Paddle Gloves]
	// some len:3, cap:3
	// all: [Kayak Lifejacket Paddle Gloves]
	// all len:4, cap:4
}

func Example_appendFromArray_resize() {
	names := [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}

	// [low:high]
	// from: low
	// len: high-low
	some := names[1:3]
	all := names[:]

	some = append(some, "Gloves")
	// 容量不足，需要扩容，产生了新的数组
	some = append(some, "Jim")

	fmt.Println("names:", names)
	fmt.Printf("names len:%d, cap:%d\n", len(names), cap(names))
	fmt.Println("some:", some)
	fmt.Printf("some len:%d, cap:%d\n", len(some), cap(some))
	fmt.Println("all:", all)
	fmt.Printf("all len:%d, cap:%d\n", len(all), cap(all))

	// Output:
	// names: [Kayak Lifejacket Paddle Gloves]
	// names len:4, cap:4
	// some: [Lifejacket Paddle Gloves Jim]
	// some len:4, cap:6
	// all: [Kayak Lifejacket Paddle Gloves]
	// all len:4, cap:4
}

func Example_createFromArray_withSpecificCapacity() {
	names := [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}

	// [low:high:max]
	// from: low
	// len: high-low
	// cap: max-low
	some := names[1:3:3]
	all := names[:]

	fmt.Println("some:", some)
	fmt.Printf("some len:%d, cap:%d\n", len(some), cap(some))
	fmt.Println("===")

	some = append(some, "Gloves")

	fmt.Println("names:", names)
	fmt.Printf("names len:%d, cap:%d\n", len(names), cap(names))
	fmt.Println("some:", some)
	fmt.Printf("some len:%d, cap:%d\n", len(some), cap(some))
	fmt.Println("all:", all)
	fmt.Printf("all len:%d, cap:%d\n", len(all), cap(all))

	// Output:
	// some: [Lifejacket Paddle]
	//some len:2, cap:2
	//===
	//names: [Kayak Lifejacket Paddle Hat]
	//names len:4, cap:4
	//some: [Lifejacket Paddle Gloves]
	//some len:3, cap:4
	//all: [Kayak Lifejacket Paddle Hat]
	//all len:4, cap:4
}

func Example_createFromSlice() {
	names := [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}

	all := names[1:]
	// 虽然从all衍生，但底层数组指向的仍然是names
	// 不受all的改变影响
	theSlice := all[1:3]

	all = append(all, "Gloves")
	all[1] = "Canoe"

	fmt.Println("names:", names)
	fmt.Printf("names len:%d, cap:%d\n", len(names), cap(names))
	fmt.Println("theSlice:", theSlice)
	fmt.Printf("theSlice len:%d, cap:%d\n", len(theSlice), cap(theSlice))
	fmt.Println("all:", all)
	fmt.Printf("all len:%d, cap:%d\n", len(all), cap(all))

	// Output:
	// names: [Kayak Lifejacket Paddle Hat]
	//names len:4, cap:4
	//theSlice: [Paddle Hat]
	//theSlice len:2, cap:2
	//all: [Lifejacket Canoe Hat Gloves]
	//all len:4, cap:6
}

// 1. 会产生新的底层数组
// 2. 基于目标数组的len复制数据，而不是cap
func Example_copy() {
	names := [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}

	all := names[1:]
	theSlice := make([]string, 2, 3)
	copy(theSlice, all)

	all = append(all, "Gloves")
	all[1] = "Canoe"

	fmt.Println("names:", names)
	fmt.Printf("names len:%d, cap:%d\n", len(names), cap(names))
	fmt.Println("theSlice:", theSlice)
	fmt.Printf("theSlice len:%d, cap:%d\n", len(theSlice), cap(theSlice))
	fmt.Println("all:", all)
	fmt.Printf("all len:%d, cap:%d\n", len(all), cap(all))

	// Output:
	// names: [Kayak Lifejacket Paddle Hat]
	//names len:4, cap:4
	//theSlice: [Lifejacket Paddle]
	//theSlice len:2, cap:3
	//all: [Lifejacket Canoe Hat Gloves]
	//all len:4, cap:6
}

// 复制时不会对目标切片进行扩展
func Example_copy_uninitialized() {
	names := [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}

	all := names[1:]
	var theSlice []string
	copy(theSlice, all)

	fmt.Println("all:", all)
	fmt.Printf("all len:%d, cap:%d\n", len(all), cap(all))
	fmt.Println("theSlice:", theSlice)
	fmt.Printf("theSlice len:%d, cap:%d\n", len(theSlice), cap(theSlice))

	// Output:
	// all: [Lifejacket Paddle Hat]
	//all len:3, cap:3
	//theSlice: []
	//theSlice len:0, cap:0
}

func Example_copy_withSpecificRanges() {
	names := [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}

	all := names[1:]
	theSlice := []string{"Boots", "Canoe", "Loan"}
	fmt.Println("before:", theSlice)
	fmt.Printf("before len:%d, cap:%d\n", len(theSlice), cap(theSlice))

	copy(theSlice[1:], all[2:3])

	fmt.Println("after:", theSlice)
	fmt.Printf("after len:%d, cap:%d\n", len(theSlice), cap(theSlice))

	// Output:
	// all: [Lifejacket Paddle Hat]
	//all len:3, cap:3
	//theSlice: []
	//theSlice len:0, cap:0
}

func Example_copy_destinationLarger() {
	names := []string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	replacement := []string{"Boots", "Canoe", "Loan"}

	fmt.Println("names [before]:", names)
	fmt.Printf("names [before] len:%d, cap:%d\n", len(names), cap(names))

	copy(names, replacement)

	fmt.Println("names [after]:", names)
	fmt.Printf("names [after] len:%d, cap:%d\n", len(names), cap(names))

	// Output:
	// names [before]: [Kayak Lifejacket Paddle Hat]
	//names [before] len:4, cap:4
	//names [after]: [Boots Canoe Loan Hat]
	//names [after] len:4, cap:4
}

func Example_copy_destinationSmaller() {
	names := []string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	replacement := []string{"Boots", "Canoe", "Loan"}

	fmt.Println("names [before]:", names)
	fmt.Printf("names [before] len:%d, cap:%d\n", len(names), cap(names))

	copy(names[0:2], replacement)

	fmt.Println("names [after]:", names)
	fmt.Printf("names [after] len:%d, cap:%d\n", len(names), cap(names))

	// Output:
	// names [before]: [Kayak Lifejacket Paddle Hat]
	//names [before] len:4, cap:4
	//names [after]: [Boots Canoe Paddle Hat]
	//names [after] len:4, cap:4
}

// 没有内建的delete函数
// 只能通过 range和append组合
func Example_delete() {
	names := []string{"Kayak", "Lifejacket", "Paddle", "Hat"}

	deleted := append(names[:2], names[3:]...)

	fmt.Println("names:", names)
	fmt.Printf("names len:%d, cap:%d\n", len(names), cap(names))
	fmt.Println("deleted:", deleted)
	fmt.Printf("deleted len:%d, cap:%d\n", len(deleted), cap(deleted))

	// 居然改变了names，这是没想到的。。。

	// Output:
	// names: [Kayak Lifejacket Hat Hat]
	//names len:4, cap:4
	//deleted: [Kayak Lifejacket Hat]
	//deleted len:3, cap:4
}

func Example_equal() {
	s1 := []string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	s2 := []string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	s3 := []string{"Kayak", "Lifejacket", "Paddle"}

	//fmt.Println(s1 == s2) // 编译错误
	fmt.Println("s1 equals s2:", reflect.DeepEqual(s1, s2))
	fmt.Println("s1 equals s3:", reflect.DeepEqual(s1, s3))

	// Output:
	// s1 equals s2: true
	//s1 equals s3: false
}

func Example_getArray() {
	names := []string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	arrayPointer := (*[3]string)(names) // 长度不能大于slice的长度
	array := *arrayPointer

	fmt.Println("names:", names)
	fmt.Printf("names len:%d, cap:%d\n", len(names), cap(names))
	fmt.Println("array:", array)
	fmt.Printf("array len:%d, cap:%d\n", len(array), cap(array))

	// Output:
	// names: [Kayak Lifejacket Paddle Hat]
	//names len:4, cap:4
	//array: [Kayak Lifejacket Paddle]
	//array len:3, cap:3
}
