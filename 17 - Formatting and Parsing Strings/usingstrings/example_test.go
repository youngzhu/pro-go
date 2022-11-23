package main

import "fmt"

func printfln(format string, values ...interface{}) {
	fmt.Printf(format+"\n", values...)
}

func Example_formattingVerbs() {
	printfln("Value: %v", Kayak)
	// 对于没实现Stringer的类型还是有区别的
	//printfln("Value: %+v", Kayak) // 跟上面的没区别
	printfln("Go syntax: %#v", Kayak)
	printfln("Type: %T", Kayak)

	// Output:
	//Value: Product: Kayak, Price: $275.00
	//Go syntax: main.Product{Name:"Kayak", Category:"Watersports", Price:275}
	//Type: main.Product
}

// 需要输入，无法使用Example测试
func Example_scanningStrings() {
	var name string
	var category string
	var price float64

	fmt.Print("Enter text to scan:")
	n, err := fmt.Scan(&name, &category, &price)

	if err == nil {
		printfln("Scanned %v values", n)
		printfln("Name: %v, Category: %v, Price: %.2f", name, category, price)
	} else {
		printfln("Error: %v", err.Error())
	}

	// Output:
	//
}

func Example_scanningIntoSlice() {
	vals := make([]string, 3)
	fmt.Print("Enter text to scan: ")

	ivals := make([]interface{}, 3)
	for i := 0; i < len(vals); i++ {
		ivals[i] = &vals[i]
	}
	// 必须用[]interface{}
	//fmt.Scan(vals...)
	fmt.Scan(ivals...)
	printfln("Name: %v", vals)
}

func Example_scanningStrings_source() {
	var name string
	var category string
	var price float64

	source := "Lifejacket Watersports 48.95"
	n, err := fmt.Sscan(source, &name, &category, &price)

	if err == nil {
		printfln("Scanned %v values", n)
		printfln("Name: %v, Category: %v, Price: %.2f", name, category, price)
	} else {
		printfln("Error: %v", err.Error())
	}

	// Output:
	//Scanned 3 values
	//Name: Lifejacket, Category: Watersports, Price: 48.95
}

func Example_scanningStrings_source_template() {
	var name string
	var category string
	var price float64

	source := "Product Lifejacket Watersports 48.95"
	template := "Product %s %s %f"
	n, err := fmt.Sscanf(source, template, &name, &category, &price)

	if err == nil {
		printfln("Scanned %v values", n)
		printfln("Name: %v, Category: %v, Price: %.2f", name, category, price)
	} else {
		printfln("Error: %v", err.Error())
	}

	// Output:
	//Scanned 3 values
	//Name: Lifejacket, Category: Watersports, Price: 48.95
}
