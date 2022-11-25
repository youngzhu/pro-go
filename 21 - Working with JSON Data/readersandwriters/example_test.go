package main

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

type MyDiscountedProduct struct {
	*Product `json:"product,omitempty"` // 空的话就不展示这个字段
	Discount float64                    `json:"-"` // 忽略这个字段
}

func Example_omittingField() {
	var builder strings.Builder
	encoder := json.NewEncoder(&builder)

	dp := MyDiscountedProduct{
		&Kayak, 10.50,
	}
	encoder.Encode(dp)

	dp2 := MyDiscountedProduct{Discount: 99.9}
	encoder.Encode(dp2)

	fmt.Println(builder.String())

	// Output:
	//{"product":{"Name":"Kayak","Category":"Watersports","Price":279}}
	//{}
}

type MyDiscountedProduct2 struct {
	*Product
	Discount float64 `json:",string"` //强制用string类型编码
}

// 强制使用string编码
func Example_encodedAsString() {
	var builder strings.Builder
	encoder := json.NewEncoder(&builder)

	dp := MyDiscountedProduct2{
		&Kayak, 10.50,
	}
	encoder.Encode(dp)

	fmt.Println(builder.String())

	// Output:
	//{"Name":"Kayak","Category":"Watersports","Price":279,"Discount":"10.5"}
}

func Example_Decode() {
	reader := strings.NewReader(`true "Hello" 99.99 200`)

	var vals []interface{}

	decoder := json.NewDecoder(reader)

	for {
		var decodedVal interface{}
		err := decoder.Decode(&decodedVal)
		if err != nil {
			if err != io.EOF {
				Printfln("Error: %v", err.Error())
			}
			break
		}
		vals = append(vals, decodedVal)
	}

	for _, v := range vals {
		Printfln("Decoded (%T): %v", v, v)
	}

	//Output:
	//Decoded (bool): true
	//Decoded (string): Hello
	//Decoded (float64): 99.99
	//Decoded (float64): 200
}

func Example_Decode_number() {
	reader := strings.NewReader(`true "Hello" 99.99 200`)

	var vals []interface{}

	decoder := json.NewDecoder(reader)
	decoder.UseNumber() // 将数字转为Number类型

	for {
		var decodedVal interface{}
		err := decoder.Decode(&decodedVal)
		if err != nil {
			if err != io.EOF {
				Printfln("Error: %v", err.Error())
			}
			break
		}
		vals = append(vals, decodedVal)
	}

	for _, v := range vals {
		if num, ok := v.(json.Number); ok {
			if i, err := num.Int64(); err == nil {
				Printfln("Decoded Int64: %v", i)
			} else if f, err := num.Float64(); err == nil {
				Printfln("Decoded Float64: %v", f)
			} else {
				Printfln("Decoded String: %v", num.String())
			}
		} else {
			Printfln("Decoded (%T): %v", v, v)
		}
	}

	//Output:
	//Decoded (bool): true
	//Decoded (string): Hello
	//Decoded Float64: 99.99
	//Decoded Int64: 200
}

func Example_Decode_specifiedType() {
	reader := strings.NewReader(`true "Hello" 99.99 200`)

	var bVal bool
	var sVal string
	var fVal float64
	var iVal int

	vals := []interface{}{
		&bVal,
		&sVal,
		&fVal,
		&iVal,
	}

	decoder := json.NewDecoder(reader)

	for i := 0; i < len(vals); i++ {
		err := decoder.Decode(vals[i])
		if err != nil {
			Printfln("Error: %v", err.Error())
			break
		}
	}

	Printfln("Decoded (%T): %v", bVal, bVal)
	Printfln("Decoded (%T): %v", sVal, sVal)
	Printfln("Decoded (%T): %v", fVal, fVal)
	Printfln("Decoded (%T): %v", iVal, iVal)

	//Output:
	//Decoded (bool): true
	//Decoded (string): Hello
	//Decoded (float64): 99.99
	//Decoded (int): 200
}

func Example_Decode_array() {
	reader := strings.NewReader(`[10,20,30]["Kayak","Lifejacket",9.99]`)

	var vals []interface{}

	decoder := json.NewDecoder(reader)

	for {
		var decodedVal interface{}
		err := decoder.Decode(&decodedVal)
		if err != nil {
			if err != io.EOF {
				Printfln("Error: %v", err.Error())
			}
			break
		}
		vals = append(vals, decodedVal)
	}

	for _, v := range vals {
		Printfln("Decoded (%T): %v", v, v)
	}

	//Output:
	//Decoded ([]interface {}): [10 20 30]
	//Decoded ([]interface {}): [Kayak Lifejacket 9.99]
}

func Example_Decode_array_specifiedType() {
	reader := strings.NewReader(`[10,20,30]["Kayak","Lifejacket",9.99]`)

	var ints []int
	var mixed []interface{}

	vals := []interface{}{
		&ints,
		&mixed,
	}

	decoder := json.NewDecoder(reader)

	for i := 0; i < len(vals); i++ {
		err := decoder.Decode(vals[i])
		if err != nil {
			Printfln("Error: %v", err.Error())
			break
		}
	}

	Printfln("Decoded (%T): %v", ints, ints)
	Printfln("Decoded (%T): %v", mixed, mixed)

	//Output:
	//Decoded ([]int): [10 20 30]
	//Decoded ([]interface {}): [Kayak Lifejacket 9.99]
}

func Example_Decode_map() {
	reader := strings.NewReader(`{"Kayak": 279,"Lifejacket":49.95}`)

	//m := make(map[string]interface{}) // 这个更安全，在不知道类型的情况下
	m := make(map[string]float64)

	decoder := json.NewDecoder(reader)

	err := decoder.Decode(&m)
	if err != nil {
		Printfln("Error: %v", err.Error())
	} else {
		//Printfln("Map(%T):  %v", m, m)
		for k, v := range m {
			Printfln("Key: %v, Value: %v", k, v)
		}
	}

	//Output:
	//Key: Kayak, Value: 279
	//Key: Lifejacket, Value: 49.95
}
