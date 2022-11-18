package main

import (
	"fmt"
	"time"
)

// 将一个普通方法，通过封装/适配器来支持chan
func Example_chanWrapper() {
	wrapper := func(price float64, c chan float64) {
		c <- calcTax(price)
	}

	resultChan := make(chan float64)
	go wrapper(275, resultChan)
	result := <-resultChan
	fmt.Println("Result:", result)

	// output:
	//Result: 330
}

func calcTax(price float64) float64 {
	return price + price*0.2
}

func Example_chanWrapper_nonVariable() {
	resultChan := make(chan float64)
	go func(price float64, c chan float64) {
		c <- calcTax(price)
	}(275, resultChan)

	result := <-resultChan
	fmt.Println("Result:", result)

	// output:
	//Result: 330
}

func Example_chan_unknownNumberOfValues() {
	dispatchChan := make(chan DispatchNotification, 100)
	go DispatchOrders(dispatchChan)
	for {
		// 校验chan是否已关闭
		if v, open := <-dispatchChan; open {
			fmt.Println("Dispatch to", v.Customer, ":", v.Quantity,
				"x", v.Product.Name)
		} else {
			fmt.Println("Channel has been closed")
			break
		}
	}

	// Output:
	//.
}

func Example_chan_enumerating() {
	dispatchChan := make(chan DispatchNotification, 100)
	go DispatchOrders(dispatchChan)
	// for-range 当chan关闭时结束
	for v := range dispatchChan {
		fmt.Println("Dispatch to", v.Customer, ":", v.Quantity,
			"x", v.Product.Name)
	}
	fmt.Println("Channel has been closed")

	// Output:
	//.
}

// 非阻塞
func Example_select_basic() {
	dispatchChan := make(chan DispatchNotification, 100)
	go DispatchOrders(dispatchChan)

	for {
		select {
		case v, ok := <-dispatchChan:
			if ok {
				fmt.Println("Dispatch to", v.Customer, ":", v.Quantity,
					"x", v.Product.Name)
			} else {
				fmt.Println("Channel has been closed")
				goto alldone
			}
		default:
			fmt.Println("wait...")
			time.Sleep(time.Millisecond * 500)
		}
	}
alldone:
	fmt.Println("All done!")

	// Output:
	//.
}

func receiveProducts(productChan chan<- *Product) {
	for _, p := range ProductList[:3] {
		productChan <- p
		time.Sleep(time.Millisecond * 500)
	}
	close(productChan)
}

func Example_select_multipleChannels() {
	dispatchChan := make(chan DispatchNotification, 100)
	go DispatchOrders(dispatchChan)

	productChan := make(chan *Product)
	go receiveProducts(productChan)

	openChannels := 2

	for {
		select {
		case dispatch, ok := <-dispatchChan:
			if ok {
				fmt.Println("Dispatch to", dispatch.Customer, ":", dispatch.Quantity,
					"x", dispatch.Product.Name)
			} else {
				fmt.Println("Dispatch Channel has been closed")
				dispatchChan = nil
				openChannels--
			}
		case product, ok := <-productChan:
			if ok {
				fmt.Println("Product:", product.Name)
			} else {
				fmt.Println("Product chan has been closed")
				productChan = nil
				openChannels--
			}
		default:
			if openChannels == 0 {
				goto alldone
			}
			fmt.Println("wait...")
			time.Sleep(time.Millisecond * 300)
		}
	}

alldone:
	fmt.Println("All done!")

	// Output:
}
