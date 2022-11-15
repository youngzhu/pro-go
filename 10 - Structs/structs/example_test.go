package main_test

import (
	"encoding/json"
	"fmt"
	"strings"
)

func Example_embedded() {
	type Product struct {
		name, category string
		price          float64
	}
	type StockLevel struct {
		Product
		count int
	}

	stock := StockLevel{
		Product{"Kayak", "Watersports", 275},
		100,
	}

	fmt.Println("Name:", stock.Product.name)
	fmt.Println("Count:", stock.count)

	// output:
	//Name: Kayak
	//Count: 100
}

// 多个相同的嵌套类，需要字段名
func Example_embedded2() {
	type Product struct {
		name, category string
		price          float64
	}
	type StockLevel struct {
		Product
		alternate Product
		count     int
	}

	stock := StockLevel{
		Product:   Product{"Kayak", "Watersports", 275},
		alternate: Product{"Lifejacket", "Watersports", 48.95},
		count:     100,
	}

	fmt.Println("Name:", stock.Product.name)
	fmt.Println("Alt Name:", stock.alternate.name)

	// output:
	//Name: Kayak
	//Alt Name: Lifejacket
}

// 当结构里的字段都是客比较时
// 可以直接使用==比较结构值
func Example_comparable() {
	type Product struct {
		name, category string
		price          float64
	}

	p1 := Product{"Kayak", "Watersports", 275}
	p2 := Product{"Kayak", "Watersports", 275}
	p3 := Product{"Kayak", "Boats", 275}

	fmt.Println(p1 == p2)
	fmt.Println(p1 == p3)

	// Output:
	// true
	//false
}

// 如果结构中包括不可比较的字段，如切片
// 则编译错误：invalid operation: p1 == p2 (the operator == is not defined on Product)
func Example_incomparable() {
	type Product struct {
		name, category string
		price          float64
		otherNames     []string
	}

	//p1 := Product{name: "Kayak", category: "Watersports", price: 275}
	//p2 := Product{name: "Kayak", category: "Watersports", price: 275}
	//p3 := Product{name: "Kayak", category: "Boats", price: 275}
	//
	//fmt.Println(p1 == p2)
	//fmt.Println(p1 == p3)

}

// 相同的（字段相同，顺序相同）结构可以互相转换
func Example_convert() {
	type Product struct {
		name, category string
		price          float64
	}
	type Item struct {
		name     string
		category string
		price    float64
	}
	type Other struct {
		name     string
		price    float64
		category string
	}

	prod := Product{name: "Kayak", category: "Watersports", price: 275}
	item := Item{name: "Kayak", category: "Watersports", price: 275}
	//other := Other{name: "Kayak", category: "Watersports", price: 275}

	//fmt.Println(prod == item) // 编译错误
	fmt.Println(prod == Product(item))
	//fmt.Println(prod == Product(other)) // 编译错误

	// Output:
	//true
}

func Example_anonymous() {
	type Product struct {
		name, category string
		price          float64
	}

	prod := Product{name: "Kayak", category: "Watersports", price: 275}

	var sb strings.Builder
	json.NewEncoder(&sb).Encode(struct {
		ProductName  string
		ProductPrice float64
	}{
		ProductName:  prod.name,
		ProductPrice: prod.price,
	})

	fmt.Println(sb.String())

	// Output:
	//{"ProductName":"Kayak","ProductPrice":275}
}

func Example_pointer() {
	type Product struct {
		name, category string
		price          float64
	}

	p1 := Product{name: "Kayak", category: "Watersports", price: 275}

	p2 := p1
	p3 := &p1

	p1.name = "Original Kayak"

	fmt.Println(p1.name)
	fmt.Println(p2.name)
	fmt.Println(p3.name)
	fmt.Println((*p3).name)

	// Output:
	//Original Kayak
	//Kayak
	//Original Kayak
	//Original Kayak

}

type Supplier struct {
	name, city string
}
type Product struct {
	name, category string
	price          float64
	*Supplier
}

// 通过工厂方法定制化创建
func newProduct(name, category string, price float64, supplier *Supplier) *Product {
	return &Product{name: name,
		category: category,
		price:    price - 10,
		Supplier: supplier,
	}
}

func Example_pointerFieldCopy() {
	acme := &Supplier{
		name: "Acme Co",
		city: "New York",
	}

	p1 := newProduct("Kayak", "Watersports", 275, acme)
	// 值复制
	// 产品名字没变，供应商的名字跟着变了
	// 结论：指针属性，即使是值复制，也是指向相同的指针地址
	p2 := *p1

	p1.name = "Original Kayak"
	p1.Supplier.name = "Boat Co."

	for _, p := range []Product{*p1, p2} {
		fmt.Println("Name:", p.name, "Supplier:", p.Supplier.name, p.Supplier.city)
	}

	// Output:
	//Name: Original Kayak Supplier: Boat Co. New York
	//Name: Kayak Supplier: Boat Co. New York
}

func Example_pointerFieldCopy_deep() {
	acme := &Supplier{
		name: "Acme Co",
		city: "New York",
	}

	p1 := newProduct("Kayak", "Watersports", 275, acme)
	p2 := copyProduct(p1)

	p1.name = "Original Kayak"
	p1.Supplier.name = "Boat Co."

	for _, p := range []Product{*p1, p2} {
		fmt.Println("Name:", p.name, "Supplier:", p.Supplier.name, p.Supplier.city)
	}

	// Output:
	//Name: Original Kayak Supplier: Boat Co. New York
	//Name: Kayak Supplier: Acme Co New York
}

func copyProduct(product *Product) Product {
	p := *product
	// 创建了一个新的变量，然后取址，与源地址实现切割
	s := *product.Supplier
	p.Supplier = &s
	return p
}
