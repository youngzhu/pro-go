package main

import "fmt"

func Example_variadicParameters() {
	printSuppliers("Kayak", "Acme Kayaks", "Bob's Boats")
	printSuppliers("Lifejacket", "Sail Safe Co.")

	// Output:
	// Product:Kayak, Supplier:Acme Kayaks
	//Product:Kayak, Supplier:Bob's Boats
	//Product:Lifejacket, Supplier:Sail Safe Co.
}

func printSuppliers(product string, suppliers ...string) {
	for _, supplier := range suppliers {
		fmt.Printf("Product:%s, Supplier:%s\n", product, supplier)
	}
}

func Example_variadicParameters_noArgs() {
	printSuppliersIfNone("Kayak", "Acme Kayaks", "Bob's Boats")
	printSuppliersIfNone("Lifejacket", "Sail Safe Co.")
	printSuppliersIfNone("Soccer Ball")

	// Output:
	// Product:Kayak, Supplier:Acme Kayaks
	//Product:Kayak, Supplier:Bob's Boats
	//Product:Lifejacket, Supplier:Sail Safe Co.
	//Product:Soccer Ball, Supplier: (none)
}

func printSuppliersIfNone(product string, suppliers ...string) {
	if len(suppliers) == 0 {
		fmt.Printf("Product:%s, Supplier: (none)\n", product)
	} else {
		for _, supplier := range suppliers {
			fmt.Printf("Product:%s, Supplier:%s\n", product, supplier)
		}
	}
}

func Example_variadicParameters_usingSlice() {
	suppliers := []string{"Acme Kayaks", "Bob's Boats"}
	printSuppliersIfNone("Kayak", suppliers...)
	printSuppliersIfNone("Lifejacket", "Sail Safe Co.")
	printSuppliersIfNone("Soccer Ball")

	// Output:
	// Product:Kayak, Supplier:Acme Kayaks
	//Product:Kayak, Supplier:Bob's Boats
	//Product:Lifejacket, Supplier:Sail Safe Co.
	//Product:Soccer Ball, Supplier: (none)
}

func Example_swapValues() {
	v1, v2 := 10, 20
	fmt.Println("Before calling:", v1, v2)
	swapValues(v1, v2)
	fmt.Println("After calling:", v1, v2)

	// Output:
	//Before calling: 10 20
	//Before swap: 10 20
	//After swap: 20 10
	//After calling: 10 20
}

func swapValues(first int, second int) {
	fmt.Println("Before swap:", first, second)
	temp := first
	first = second
	second = temp
	fmt.Println("After swap:", first, second)
}

func Example_swapPointers() {
	v1, v2 := 10, 20
	fmt.Println("Before calling:", v1, v2)
	swapPointers(&v1, &v2)
	fmt.Println("After calling:", v1, v2)

	// Output:
	//Before calling: 10 20
	//Before swap: 10 20
	//After swap: 20 10
	//After calling: 20 10
}

func swapPointers(first, second *int) {
	fmt.Println("Before swap:", *first, *second)
	temp := *first
	*first = *second
	*second = temp
	fmt.Println("After swap:", *first, *second)
}
