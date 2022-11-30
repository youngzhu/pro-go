package main

import (
	"fmt"
	"reflect"
	"strings"
)

func Printfln(template string, values ...interface{}) {
	fmt.Printf(template+"\n", values...)
}

func printDetails(values ...interface{}) {
	for _, elem := range values {
		fields := []string{}

		elemType := reflect.TypeOf(elem)
		elemValue := reflect.ValueOf(elem)
		if elemType.Kind() == reflect.Struct {
			for i := 0; i < elemType.NumField(); i++ {
				fieldName := elemType.Field(i).Name
				fieldValue := elemValue.Field(i)
				fields = append(fields,
					fmt.Sprintf("%v: %v", fieldName, fieldValue))
			}
			Printfln("%v: {%v}", elemType.Name(), strings.Join(fields, ", "))
		} else {
			Printfln("%v: %v", elemType.Name(), elemValue)
		}
	}
}

func getTypePath(t reflect.Type) (path string) {
	path = t.PkgPath()
	if path == "" {
		path = "(built-in)"
	}
	return
}

func printDetailsWithPkgPath(values ...interface{}) {
	for _, elem := range values {
		elemType := reflect.TypeOf(elem)
		Printfln("Name: %v, PkgPath: %v, Kind: %v",
			elemType.Name(), getTypePath(elemType), elemType.Kind())
	}
}

func printValueDetails(values ...interface{}) {
	for _, e := range values {
		eVal := reflect.ValueOf(e)
		switch eVal.Kind() {
		case reflect.Bool:
			var v bool = eVal.Bool()
			Printfln("Bool: %v", v)
		case reflect.Int:
			var v int64 = eVal.Int()
			Printfln("Int: %v", v)
		case reflect.Float64, reflect.Float32:
			var v float64 = eVal.Float()
			Printfln("Float: %v", v)
		case reflect.String:
			var v string = eVal.String()
			Printfln("String: %v", v)
		case reflect.Ptr:
			var v reflect.Value = eVal.Elem()
			if v.Kind() == reflect.Int {
				Printfln("Pointer to Int: %v", v.Int())
			}
		default:
			Printfln("Kind: %v, Value: %v", eVal.Kind(), eVal.String())
		}
	}
}

var (
	intPtrType    = reflect.TypeOf((*int)(nil))
	byteSliceType = reflect.TypeOf([]byte(nil))
)

func printValueDetailsIdentifyingTypes(values ...interface{}) {
	for _, e := range values {
		eVal := reflect.ValueOf(e)
		eType := reflect.TypeOf(e)

		if eType == intPtrType {
			Printfln("Pointer to Int: %v", eVal.Elem().Int())
		} else if eType == byteSliceType {
			Printfln("Pointer to Byte Slice: %v", eVal.Bytes())
		} else {
			switch eVal.Kind() {
			case reflect.Bool:
				var v bool = eVal.Bool()
				Printfln("Bool: %v", v)
			case reflect.Int:
				var v int64 = eVal.Int()
				Printfln("Int: %v", v)
			case reflect.Float64, reflect.Float32:
				var v float64 = eVal.Float()
				Printfln("Float: %v", v)
			case reflect.String:
				var v string = eVal.String()
				Printfln("String: %v", v)
			default:
				Printfln("Kind: %v, Value: %v", eVal.Kind(), eVal.String())
			}
		}
	}
}
