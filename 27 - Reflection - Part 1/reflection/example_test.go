package main

import (
	"reflect"
	"strings"
)

type Payment struct {
	Currency string
	Amount   float64
}

var (
	product = Product{
		Name:     "Kayak",
		Category: "Watersprots",
		Price:    279,
	}
	customer = Customer{
		Name: "Alice",
		City: "New York",
	}
	payment = Payment{
		Currency: "USD",
		Amount:   100.5,
	}
)

var printValues = []interface{}{
	product, customer, payment, 10, true, "Hello",
}

func Example_TypeAndValue() {

	printDetails(printValues...)

	// Output:
	//Product: {Name: Kayak, Category: Watersprots, Price: 279}
	//Customer: {Name: Alice, City: New York}
	//Payment: {Currency: USD, Amount: 100.5}
	//int: 10
	//bool: true
	//string: Hello
}

func Example_PkgPath() {
	printDetailsWithPkgPath(printValues...)

	// Output:
	//Name: Product, PkgPath: reflection, Kind: struct
	//Name: Customer, PkgPath: reflection, Kind: struct
	//Name: Payment, PkgPath: reflection, Kind: struct
	//Name: int, PkgPath: (built-in), Kind: int
	//Name: bool, PkgPath: (built-in), Kind: bool
	//Name: string, PkgPath: (built-in), Kind: string
}

func Example_Value() {
	n := 101
	printValueDetails(true, 10, 11.1, "Hello", n, &n, product)

	// Output:
	//Bool: true
	//Int: 10
	//Float: 11.1
	//String: Hello
	//Int: 101
	//Pointer to Int: 101
	//Kind: struct, Value: <main.Product Value>
}

func Example_Value_identifyingTypes() {
	n := 101
	slice := []byte("Alice")

	printValueDetailsIdentifyingTypes(true, 10, 11.1, "Hello", n, &n, product, slice)

	// Output:
	//Bool: true
	//Int: 10
	//Float: 11.1
	//String: Hello
	//Int: 101
	//Pointer to Int: 101
	//Kind: struct, Value: <main.Product Value>
	//Pointer to Byte Slice: [65 108 105 99 101]
}

func incrementOrUpper(values ...interface{}) {
	for _, e := range values {
		eVal := reflect.ValueOf(e)
		if eVal.Kind() == reflect.Ptr {
			eVal = eVal.Elem()
		}
		if eVal.CanSet() {
			switch eVal.Kind() {
			case reflect.Int:
				eVal.SetInt(eVal.Int() + 1)
			case reflect.String:
				eVal.SetString(strings.ToUpper(eVal.String()))
			}
			Printfln("Modified Value: %v", eVal)
		} else {
			Printfln("Cannot set %v: %v", eVal.Kind(), eVal)
		}
	}
}
func Example_setValue() {
	name := "Alice"
	price := 279
	city := "London"

	// 必须用指针
	//incrementOrUpper(name, price, city)
	incrementOrUpper(&name, &price, &city)

	for _, val := range []interface{}{name, price, city} {
		Printfln("Value: %v", val)
	}

	// Output:
	//Modified Value: ALICE
	//Modified Value: 280
	//Modified Value: LONDON
	//Value: ALICE
	//Value: 280
	//Value: LONDON
}

func setAll(src interface{}, targets ...interface{}) {
	srcVal := reflect.ValueOf(src)
	for _, target := range targets {
		tVal := reflect.ValueOf(target)
		if tVal.Kind() == reflect.Ptr &&
			tVal.Elem().Type() == srcVal.Type() &&
			tVal.Elem().CanSet() {
			tVal.Elem().Set(srcVal)
		}
	}
}
func Example_setValue_another() {
	name := "Alice"
	price := 279
	city := "London"

	setAll("New string", &name, &price, &city)
	setAll(101, &name, &price, &city)

	for _, val := range []interface{}{name, price, city} {
		Printfln("Value: %v", val)
	}

	// Output:
	//Value: New string
	//Value: 101
	//Value: New string
}

// 将src转换成target的类型
func convertType(src, target interface{}) (result interface{}, ok bool) {
	srcVal := reflect.ValueOf(src)
	targetVal := reflect.ValueOf(target)

	if srcVal.Type().ConvertibleTo(targetVal.Type()) {
		result = srcVal.Convert(targetVal.Type()).Interface()
		ok = true
	} else {
		result = src
	}

	return
}
func Example_convert() {
	name := "Alice"
	price := 279

	newName, ok := convertType(name, 100)
	Printfln("Convert %v(%T): %v(%T), %v", name, name, newName, newName, ok)

	newPrice, ok := convertType(price, 100.00)
	Printfln("Convert %v(%T): %v(%T), %v", price, price, newPrice, newPrice, ok)

	// Output:
	//Convert Alice(string): Alice(string), false
	//Convert 279(int): 279(float64), true
}
