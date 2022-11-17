package store_test

import (
	"composition/store"
	"fmt"
)

/*
嵌套类型的字段和方法
可以通过嵌套类型访问，也可以直接访问（好像从父类继承了一样）
*/
func Example_embedType() {
	boats := []*store.Boat{
		store.NewBoat("Kayak", 275, 1, false),
		store.NewBoat("Canoe", 400, 3, false),
		store.NewBoat("Tender", 625.5, 2, true),
	}

	for _, b := range boats {
		fmt.Println("[Conventional Name]:", b.Product.Name, "[Direct Name]:", b.Name, "[Price]:", b.Price(0.2))
	}

	// Output:
	//[Conventional Name]: Kayak [Direct Name]: Kayak [Price]: 330
	//[Conventional Name]: Canoe [Direct Name]: Canoe [Price]: 480
	//[Conventional Name]: Tender [Direct Name]: Tender [Price]: 750.6
}

// 嵌套类的字段和方法都会被提升到顶层类型里
// RentalBoat -> Boat -> Product
func Example_chainOfNestedTypes() {
	rentals := []*store.RentalBoat{
		store.NewRentalBoatShort("Rubber Ring", 10, 1, false, false),
		store.NewRentalBoatShort("Yacht", 50000, 5, true, true),
		store.NewRentalBoatShort("Super Yacht", 150000, 15, true, true),
	}

	for _, r := range rentals {
		fmt.Println("[Rental Boat]:", r.Name, "[Rental Price]:", r.Price(0.2))
	}

	// Output:
	//[Rental Boat]: Rubber Ring [Rental Price]: 12
	//[Rental Boat]: Yacht [Rental Price]: 60000
	//[Rental Boat]: Super Yacht [Rental Price]: 180000
}

func Example_typeSwitchLimitation() {
	products := map[string]store.ItemForSale{
		"Kayak": store.NewBoat("Kayak", 279, 1, false),
		"Ball":  store.NewProduct("Soccer Ball", "Soccer", 19.5),
	}

	// 编译错误：v.Name 无法识别
	//for k, p := range products {
	//	switch v := p.(type) {
	//	case *store.Product, *store.Boat:
	//		fmt.Println("Name:", v.Name,)
	//	}
	//}

	for k, p := range products {
		switch v := p.(type) {
		case *store.Product:
			fmt.Println("Name:", v.Name, "Category:", v.Category,
				"Price:", v.Price(0.2))
		case *store.Boat:
			fmt.Println("Name:", v.Name, "Category:", v.Category,
				"Price:", v.Price(0.2))
		default:
			fmt.Println("Key:", k, "Price:", p.Price(0.2))
		}
	}

	// Output:
	//Name: Kayak Category: Watersports Price: 334.8
	//Name: Soccer Ball Category: Soccer Price: 23.4
}
