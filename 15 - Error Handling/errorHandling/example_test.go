package main

import "fmt"

type CategoryError struct {
	requestedCategory string
}

func (e *CategoryError) Error() string {
	return "Category [" + e.requestedCategory + "] does not exist"
}

//func (slice ProductSlice) TotalPrice(category string) (total float64, err *CategoryError) {
//	exists := false
//	for _, p := range slice {
//		if p.Category == category {
//			total += p.Price
//			exists = true
//		}
//	}
//	if !exists {
//		err = &CategoryError{
//			requestedCategory: category,
//		}
//	}
//	return
//}

// 如果没有异常，没有该类别导致总价为0和有该类别但总价为0，无法区分
func Example_error() {
	categories := []string{"Watersports", "Chess", "Running", "Toy"}

	for _, cat := range categories {
		total, err := Products.TotalPrice(cat)
		if err == nil {
			fmt.Println(cat, "Total:", ToCurrency(total))
		} else {
			fmt.Println(cat, "(no such category)")
		}
	}

	// Output:
	//Watersports Total: $328.95
	//Chess Total: $1291.00
	//Running (no such category)
	//Toy Total: $0.00
}

func Example_reportErrorsViaChannel() {
	categories := []string{"Watersports", "Chess", "Running", "Toy"}

	channel := make(chan ChannelMessage, 10)
	go Products.TotalPriceAsync(categories, channel)
	for msg := range channel {
		if msg.CategoryError == nil {
			fmt.Println(msg.Category, "Total:", ToCurrency(msg.Total))
		} else {
			fmt.Println(msg.Category, "(no such category)")
		}
	}

	// Output:
	//Watersports Total: $328.95
	//Chess Total: $1291.00
	//Running (no such category)
	//Toy Total: $0.00
}

// recover函数必须用defer修饰
func Example_recover() {
	recoverFunc := func() {
		if arg := recover(); arg != nil {
			if err, ok := arg.(error); ok {
				fmt.Println("Error:", err.Error())
			} else if str, ok := arg.(string); ok {
				fmt.Println("Message:", str)
			} else {
				fmt.Println("Panic recovered")
			}
		}
	}
	defer recoverFunc()

	//
	categories := []string{"Watersports", "Chess", "Running", "Toy"}

	channel := make(chan ChannelMessage, 10)
	go Products.TotalPriceAsync(categories, channel)
	for msg := range channel {
		if msg.CategoryError == nil {
			fmt.Println(msg.Category, "Total:", ToCurrency(msg.Total))
		} else {
			panic(msg.CategoryError)
		}
	}

	// Output:
	//Watersports Total: $328.95
	//Chess Total: $1291.00
	//Error: Cannot find category: Running
}

// 一般都用匿名函数
func Example_recover_anonymous() {
	defer func() {
		if arg := recover(); arg != nil {
			if err, ok := arg.(error); ok {
				fmt.Println("Error:", err.Error())
			} else if str, ok := arg.(string); ok {
				fmt.Println("Message:", str)
			} else {
				fmt.Println("Panic recovered")
			}
		}
	}()

	//
	categories := []string{"Watersports", "Chess", "Running", "Toy"}

	channel := make(chan ChannelMessage, 10)
	go Products.TotalPriceAsync(categories, channel)
	for msg := range channel {
		if msg.CategoryError == nil {
			fmt.Println(msg.Category, "Total:", ToCurrency(msg.Total))
		} else {
			panic(msg.CategoryError)
		}
	}

	// Output:
	//Watersports Total: $328.95
	//Chess Total: $1291.00
	//Error: Cannot find category: Running
}

func Example_panicAfterRecover() {
	defer func() {
		if arg := recover(); arg != nil {
			if err, ok := arg.(error); ok {
				fmt.Println("Error:", err.Error())
				panic(err)
			} else if str, ok := arg.(string); ok {
				fmt.Println("Message:", str)
			} else {
				fmt.Println("Panic recovered")
			}
		}
	}()

	//
	categories := []string{"Watersports", "Chess", "Running", "Toy"}

	channel := make(chan ChannelMessage, 10)
	go Products.TotalPriceAsync(categories, channel)
	for msg := range channel {
		if msg.CategoryError == nil {
			fmt.Println(msg.Category, "Total:", ToCurrency(msg.Total))
		} else {
			panic(msg.CategoryError)
		}
	}

}
