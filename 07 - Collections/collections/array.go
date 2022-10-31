package main

import "fmt"

// 数组赋值会产生一个新的数组
func arrayValues() {
	names := [3]string{"Kayak", "Lifejacket", "Paddle"}

	copyArr := names
	pointerArr := &names

	names[0] = "Canoe"

	fmt.Println("names:", names)
	fmt.Println("copyArr:", copyArr)
	fmt.Println("pointerArr:", *pointerArr)
}

// ==，类型一样，包含的元素一样（包括顺序），就返回true
// 类型：[3]string
func comparing() {
	names := [3]string{"Kayak", "Lifejacket", "Paddle"}

	names2 := [3]string{"Kayak", "Lifejacket", "Paddle"}
	//names3 := [4]string{"Kayak", "Lifejacket", "Paddle"}
	names4 := [...]string{"Kayak", "Lifejacket", "Paddle"}
	names5 := [3]string{"Kayak", "Paddle", "Lifejacket"}

	fmt.Println("names == names2:", names == names2)
	//fmt.Println("", names == names3) // 编译错误
	fmt.Println("names == names4:", names == names4)
	fmt.Println("names == names5:", names == names5)
}
