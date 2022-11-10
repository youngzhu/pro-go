package main_test

import (
	"fmt"
	"sort"
)

func Example_funcType() {
	products := map[string]float64{
		"Kayak":      275,
		"Lifejacket": 48.95,
	}

	for product, price := range products {
		var calcFunc func(float64) float64
		if price > 100 {
			calcFunc = calcWithTax
		} else {
			calcFunc = calcWithoutTax
		}
		totalPrice := calcFunc(price)
		fmt.Println("Product:", product, "Price:", totalPrice)
	}

	// Output:
	//Product: Kayak Price: 330
	//Product: Lifejacket Price: 48.95
}

func calcWithTax(price float64) float64 {
	return price + price*0.2
}
func calcWithoutTax(price float64) float64 {
	return price
}

func Example_comparison() {
	products := map[string]float64{
		"Kayak":      275,
		"Lifejacket": 48.95,
	}

	for _, price := range products {
		var calcFunc func(float64) float64
		fmt.Println("Function assigned:", calcFunc == nil)
		if price > 100 {
			calcFunc = calcWithTax
		} else {
			calcFunc = calcWithoutTax
		}
		fmt.Println("Function assigned:", calcFunc == nil)
		calcFunc(price)
	}

	// Output:
	//Function assigned: true
	//Function assigned: false
	//Function assigned: true
	//Function assigned: false
}

func Example_asArgument() {
	products := map[string]float64{
		"Kayak":      275,
		"Lifejacket": 48.95,
	}

	for product, price := range products {
		if price > 100 {
			printPrice(product, price, calcWithTax)
		} else {
			printPrice(product, price, calcWithoutTax)
		}
	}

	// Output:
	//Product: Kayak Price: 330
	//Product: Lifejacket Price: 48.95
}

func printPrice(product string, price float64, calculator func(float642 float64) float64) {
	fmt.Println("Product:", product, "Price:", calculator(price))
}

func Example_asResult() {
	products := map[string]float64{
		"Kayak":      275,
		"Lifejacket": 48.95,
	}

	for product, price := range products {
		printPrice(product, price, selectCalculator(price))
	}

	// Output:
	//Product: Kayak Price: 330
	//Product: Lifejacket Price: 48.95
}

func selectCalculator(price float64) func(float64) float64 {
	if price > 100 {
		return calcWithTax
	}
	return calcWithoutTax
}

// aliases
type calcFunc func(float64) float64

func Example_literal() {
	withTax := func(price float64) float64 {
		return price + price*0.2
	}

	_ = withTax(99.9)
}

func Example_closure_before() {
	watersportsProducts := map[string]float64{
		"Kayak":      275,
		"Lifejacket": 48.95,
	}
	soccerProducts := map[string]float64{
		"Soccer Ball": 19.5,
		"Stadium":     79500,
	}
	calc := func(price float64) float64 {
		if price > 100 {
			return price + price*0.2
		}
		return price
	}
	printPriceOrderByProduct(watersportsProducts, calc)

	calc = func(price float64) float64 {
		if price > 50 {
			return price + price*0.1
		}
		return price
	}
	printPriceOrderByProduct(soccerProducts, calc)

	// Output:
	//Product: Kayak Price: 330
	//Product: Lifejacket Price: 48.95
	//Product: Soccer Ball Price: 19.5
	//Product: Stadium Price: 87450
}

func printPriceOrderByProduct(products map[string]float64, calculator calcFunc) {
	keys := make([]string, 0, len(products))
	for key, _ := range products {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Println("Product:", key, "Price:", calculator(products[key]))
	}
}

func Example_closure_after() {
	watersportsProducts := map[string]float64{
		"Kayak":      275,
		"Lifejacket": 48.95,
	}
	soccerProducts := map[string]float64{
		"Soccer Ball": 19.5,
		"Stadium":     79500,
	}

	waterCalc := priceCalcFactory(100, 0.2)
	soccerCalc := priceCalcFactory(50, 0.1)

	printPriceOrderByProduct(watersportsProducts, waterCalc)
	printPriceOrderByProduct(soccerProducts, soccerCalc)

	// Output:
	//Product: Kayak Price: 330
	//Product: Lifejacket Price: 48.95
	//Product: Soccer Ball Price: 19.5
	//Product: Stadium Price: 87450
}

func priceCalcFactory(threshold, rate float64) calcFunc {
	return func(price float64) float64 {
		if prizeGiveaway {
			return 0
		} else if price > threshold {
			return price + price*rate
		} else {
			return price
		}
	}
}

var prizeGiveaway = false

// 在函数执行的时候，闭包才执行
// 所以，值都是0
func Example_closure_evaluation() {
	watersportsProducts := map[string]float64{
		"Kayak":      275,
		"Lifejacket": 48.95,
	}
	soccerProducts := map[string]float64{
		"Soccer Ball": 19.5,
		"Stadium":     79500,
	}

	prizeGiveaway = false
	waterCalc := priceCalcFactory(100, 0.2)
	prizeGiveaway = true
	soccerCalc := priceCalcFactory(50, 0.1)

	printPriceOrderByProduct(watersportsProducts, waterCalc)
	printPriceOrderByProduct(soccerProducts, soccerCalc)

	// Output:
	//Product: Kayak Price: 0
	//Product: Lifejacket Price: 0
	//Product: Soccer Ball Price: 0
	//Product: Stadium Price: 0
}

func Example_closure_forcingEarlyEvaluation() {
	watersportsProducts := map[string]float64{
		"Kayak":      275,
		"Lifejacket": 48.95,
	}
	soccerProducts := map[string]float64{
		"Soccer Ball": 19.5,
		"Stadium":     79500,
	}

	prizeGiveaway = false
	waterCalc := priceCalcFactoryForcingEarlyEvaluation(100, 0.2)
	prizeGiveaway = true
	soccerCalc := priceCalcFactoryForcingEarlyEvaluation(50, 0.1)

	printPriceOrderByProduct(watersportsProducts, waterCalc)
	printPriceOrderByProduct(soccerProducts, soccerCalc)

	// Output:
	//Product: Kayak Price: 330
	//Product: Lifejacket Price: 48.95
	//Product: Soccer Ball Price: 0
	//Product: Stadium Price: 0
}

// 创建一个副本，会使用factory被调用时的值
func priceCalcFactoryForcingEarlyEvaluation(threshold, rate float64) calcFunc {
	fixedPrizeGiveaway := prizeGiveaway
	return func(price float64) float64 {
		if fixedPrizeGiveaway {
			return 0
		} else if price > threshold {
			return price + price*rate
		} else {
			return price
		}
	}
}

func Example_closure_forcingEarlyEvaluation_param() {
	watersportsProducts := map[string]float64{
		"Kayak":      275,
		"Lifejacket": 48.95,
	}
	soccerProducts := map[string]float64{
		"Soccer Ball": 19.5,
		"Stadium":     79500,
	}

	waterCalc := priceCalcFactoryForcingEarlyEvaluationWithParam(100, 0.2, false)
	soccerCalc := priceCalcFactoryForcingEarlyEvaluationWithParam(50, 0.1, true)

	printPriceOrderByProduct(watersportsProducts, waterCalc)
	printPriceOrderByProduct(soccerProducts, soccerCalc)

	// Output:
	//Product: Kayak Price: 330
	//Product: Lifejacket Price: 48.95
	//Product: Soccer Ball Price: 0
	//Product: Stadium Price: 0
}

// 创建一个副本，会使用factory被调用时的值
func priceCalcFactoryForcingEarlyEvaluationWithParam(threshold, rate float64, prizeGiveaway bool) calcFunc {
	return func(price float64) float64 {
		if prizeGiveaway {
			return 0
		} else if price > threshold {
			return price + price*rate
		} else {
			return price
		}
	}
}
