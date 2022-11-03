package main

import (
	"fmt"
	"sort"
)

func Example_map() {
	products := make(map[string]float64, 10)

	products["iPhone"] = 699.99
	products["Kindle"] = 69.99

	// cap() 对map不适用
	//fmt.Printf("size: %d, cap: %d", len(products), cap(products))
	fmt.Printf("size: %d\n", len(products))
	fmt.Println("iPhone:", products["iPhone"])
	fmt.Println("Apple:", products["Apple"])

	// Output:
	// size: 2
	//iPhone: 699.99
	//Apple: 0
}

func Example_map_literal() {
	products := map[string]float64{
		"iPhone": 699.99,
		"Kindle": 69.99,
	}

	// cap() 对map不适用
	//fmt.Printf("size: %d, cap: %d", len(products), cap(products))
	fmt.Printf("size: %d\n", len(products))
	fmt.Println("iPhone:", products["iPhone"])
	fmt.Println("Apple:", products["Apple"])

	// Output:
	// size: 2
	//iPhone: 699.99
	//Apple: 0
}

// 不存在的key，返回零值
// 但无法判断到底是没有key，还是key对应的就是零值
func Example_map_check() {
	products := map[string]float64{
		"iPhone":    699.99,
		"Kindle":    69.99,
		"HeadPhone": 0,
	}

	value, ok := products["Apple"]
	if ok {
		fmt.Println("Apple:", value)
	} else {
		fmt.Println("No Apple")
	}

	// 或者
	if val, exists := products["HeadPhone"]; exists {
		fmt.Println("HeadPhone:", val)
	} else {
		fmt.Println("No HeadPhone")
	}

	// Output:
	// No Apple
	//HeadPhone: 0
}

func Example_map_delete() {
	products := map[string]float64{
		"iPhone":    699.99,
		"Kindle":    69.99,
		"HeadPhone": 0,
	}

	if val, exists := products["HeadPhone"]; exists {
		fmt.Println("HeadPhone:", val)
	} else {
		fmt.Println("No HeadPhone")
	}

	delete(products, "HeadPhone")
	delete(products, "HeadPhone") // 删除一个不存在的key也不会报错

	if val, exists := products["HeadPhone"]; exists {
		fmt.Println("HeadPhone:", val)
	} else {
		fmt.Println("No HeadPhone")
	}

	// Output:
	// HeadPhone: 0
	//No HeadPhone
}

// map是无序的
// 要想对map进行有序遍历，可以先创建一个key的切片，在进行排序
// 最后，通过有序的key遍历map
func Example_map_enumerating() {
	products := map[string]float64{
		"iPhone":    699.99,
		"Kindle":    69.99,
		"HeadPhone": 0,
	}

	// 注意这里的初始大小是0，容量是map的大小
	keys := make([]string, 0, len(products))
	for key, _ := range products {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Println("Key:", key, "Value:", products[key])
	}

	// Output:
	// Key: HeadPhone Value: 0
	//Key: Kindle Value: 69.99
	//Key: iPhone Value: 699.99
}
