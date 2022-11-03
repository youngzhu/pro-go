package main

import (
	"fmt"
	"strconv"
)

func Example_strings() {
	var price string = "$48.95"

	var currencyByte byte = price[0]
	var currencyStr string = string(price[0])
	var amountStr string = price[1:]
	amount, err := strconv.ParseFloat(amountStr, 64)

	fmt.Println("Currency(byte):", currencyByte)
	fmt.Println("Currency(string):", currencyStr)
	if err == nil {
		fmt.Println("Amount:", amount)
	} else {
		fmt.Println("Parse error:", err)
	}

	// Output:
	// Currency(byte): 36
	//Currency(string): $
	//Amount: 48.95
}

// 但并不是所有的字符都是ASCII
// 所以可以用rune，它保证每个字符只占一位
func Example_strings_rune() {
	var priceStr string = "€48.95"

	var price []rune = []rune(priceStr)

	var currencyStr string = string(price[0])
	var amountStr string = string(price[1:])
	amount, err := strconv.ParseFloat(amountStr, 64)

	fmt.Println("priceStr len:", len(priceStr))
	fmt.Println("price len:", len(price))
	fmt.Println("Currency:", currencyStr)
	if err == nil {
		fmt.Println("Amount:", amount)
	} else {
		fmt.Println("Parse error:", err)
	}

	// Output:
	// priceStr len: 8
	//price len: 6
	//Currency: €
	//Amount: 48.95
}

// 在使用for-range时，Go将string当做rune处理
// 注意这里的索引不连续
func Example_strings_enumerating() {
	var price string = "€48.95"

	for i, val := range price {
		fmt.Println(i, val, string(val))
	}

	// output:
	// 0 8364 €
	//3 52 4
	//4 56 8
	//5 46 .
	//6 57 9
	//7 53 5
}
