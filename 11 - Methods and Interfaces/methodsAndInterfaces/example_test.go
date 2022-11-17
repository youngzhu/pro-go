package main

import "fmt"

/*
指针对象的比较，只有指向同一个内存地址时才返回true
值对象比较，只要类型和字段值一致，就返回true
但如果包含不可比较的字段，如切片，则会抛出运行时异常
*/
func Example_comparingInterface() {
	var e1 Expense = &Product{name: "Kayak"}
	var e2 Expense = &Product{name: "Kayak"}

	var e3 Expense = Service{description: "Boat Cover"}
	var e4 Expense = Service{description: "Boat Cover"}

	fmt.Println("e1 == e2", e1 == e2) // false
	fmt.Println("e3 == e4", e3 == e4)

	// Output:
	//
}

func Example_testBeforeTypeAssertion() {
	expenses := []Expense{
		Service{"Boat Cover", 12, 89.5, []string{}},
		Service{"Paddle Protect", 12, 8, []string{}},
		&Product{"Kayak", "Watersports", 275},
	}

	for _, expense := range expenses {
		if s, ok := expense.(Service); ok {
			fmt.Println("Service:", s.description, "Price:", s.monthlyFee*float64(s.durationMonths))
		} else {
			fmt.Println("Expense:", expense.getName(), "Cost:", expense.getCost(true))
		}
	}

	// Output:
	//Service: Boat Cover Price: 1074
	//Service: Paddle Protect Price: 96
	//Expense: Kayak Cost: 275
}

func Example_switchingOnDynamicTypes() {
	expenses := []Expense{
		Service{"Boat Cover", 12, 89.5, []string{}},
		Service{"Paddle Protect", 12, 8, []string{}},
		&Product{"Kayak", "Watersports", 275},
	}

	for _, expense := range expenses {
		switch value := expense.(type) {
		case Service:
			fmt.Println("Service:", value.description, "Price:", value.monthlyFee*float64(value.durationMonths))
		case *Product:
			fmt.Println("Product:", value.name, "Price:", value.price)
		default:
			fmt.Println("Expense:", expense.getName(), "Cost:", expense.getCost(true))
		}
	}

	// Output:
	//Service: Boat Cover Price: 1074
	//Service: Paddle Protect Price: 96
	//Product: Kayak Price: 275
}
