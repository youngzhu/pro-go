package main

import (
	"bytes"
	"fmt"
	"strings"
)

// bytes 和 strings 有一套对应的函数
func Example_strings_bytes() {
	price := "€100"

	fmt.Println("strings prefix:", strings.HasPrefix(price, "€"))
	fmt.Println("bytes prefix:", bytes.HasPrefix([]byte(price), []byte("€")))
	fmt.Println("bytes prefix:", bytes.HasPrefix([]byte(price), []byte{226, 130}))

	// output:
	//strings prefix: true
	//bytes prefix: true
	//bytes prefix: true
}

// Title: 所有首字母都大写（包括介词等）
// ToTitle: 所有字母都大写（主要是英文字母）
func Example_convertingStringCase() {
	strEn := "A boat for sailing"
	strSpecial := "中国"

	convertingStringCase(strEn)
	fmt.Println()
	convertingStringCase(strSpecial)

	// Output:
	//Original: A boat for sailing [65 32 98 111 97 116 32 102 111 114 32 115 97 105 108 105 110 103]
	//Upper: A BOAT FOR SAILING [65 32 66 79 65 84 32 70 79 82 32 83 65 73 76 73 78 71]
	//Title: A Boat For Sailing [65 32 66 111 97 116 32 70 111 114 32 83 97 105 108 105 110 103]
	//ToTitle: A BOAT FOR SAILING [65 32 66 79 65 84 32 70 79 82 32 83 65 73 76 73 78 71]
	//
	//Original: 中国 [228 184 173 229 155 189]
	//Upper: 中国 [228 184 173 229 155 189]
	//Title: 中国 [228 184 173 229 155 189]
	//ToTitle: 中国 [228 184 173 229 155 189]
}

func convertingStringCase(s string) {
	fmt.Println("Original:", s, []byte(s))

	upper := strings.ToUpper(s)
	fmt.Println("Upper:", upper, []byte(upper))

	title := strings.Title(s)
	fmt.Println("Title:", title, []byte(title))

	toTitle := strings.ToTitle(s)
	fmt.Println("ToTitle:", toTitle, []byte(toTitle))
}

func Example_IndexFunc() {
	s := "A boat for sailing"

	isLetterB := func(r rune) bool {
		return r == 'B' || r == 'b'
	}

	fmt.Println("IndexFunc:", strings.IndexFunc(s, isLetterB))

	// Output:
	//IndexFunc: 2
}

// After函数分割的部分会带有分割符
func Example_split() {
	s := "A boat for sailing"

	split := strings.Split(s, " ")
	for _, x := range split {
		fmt.Println("Split >>" + x + "<<")
	}

	splitAfter := strings.SplitAfter(s, " ")
	for _, x := range splitAfter {
		fmt.Println("SplitAfter >>" + x + "<<")
	}

	// Output:
	//Split >>A<<
	//Split >>boat<<
	//Split >>for<<
	//Split >>sailing<<
	//SplitAfter >>A <<
	//SplitAfter >>boat <<
	//SplitAfter >>for <<
	//SplitAfter >>sailing<<
}

// SplitN 限定返回的切片的长度
func Example_splitN() {
	s := "A boat for sailing"

	split := strings.SplitN(s, " ", 2)
	for _, x := range split {
		fmt.Println("Split >>" + x + "<<")
	}

	splitAfter := strings.SplitAfterN(s, " ", 5)
	for _, x := range splitAfter {
		fmt.Println("SplitAfter >>" + x + "<<")
	}

	// Output:
	//Split >>A<<
	//Split >>boat for sailing<<
	//SplitAfter >>A <<
	//SplitAfter >>boat <<
	//SplitAfter >>for <<
	//SplitAfter >>sailing<<
}

// Fields 按空白字符分割，不管是一个空格还是多个
func Example_fields() {
	s := "This  is  double  spaced"

	splits := strings.Split(s, " ")
	showWhiteSpace("splits", splits)

	fmt.Println()

	fields := strings.Fields(s)
	showWhiteSpace("fields", fields)

	// Output:
	//splits >>This<<
	//splits >><<
	//splits >>is<<
	//splits >><<
	//splits >>double<<
	//splits >><<
	//splits >>spaced<<
	//
	//fields >>This<<
	//fields >>is<<
	//fields >>double<<
	//fields >>spaced<<
}

func showWhiteSpace(prefix string, slice []string) {
	for _, s := range slice {
		fmt.Println(prefix + " >>" + s + "<<")
	}
}

func Example_map() {
	s := "It was a boat. A small boat."

	mapper := func(r rune) rune {
		if r == 'b' {
			return 'c'
		}
		return r
	}

	mapped := strings.Map(mapper, s)
	fmt.Println(mapped)

	// Output:
	//It was a coat. A small coat.
}

func Example_Replacer() {
	s := "It was a boat. A small boat."

	// 新老值成对出现
	replacer := strings.NewReplacer("boat", "kayak", "small", "huge")
	replaced := replacer.Replace(s)
	fmt.Println(replaced)

	// Output:
	//It was a kayak. A huge kayak.
}

func Example_Join() {
	s := "It was a boat. A small boat."

	elements := strings.Fields(s)
	joined := strings.Join(elements, "--")
	fmt.Println(joined)

	// Output:
	//It--was--a--boat.--A--small--boat.
}

func Example_Builder() {
	s := "It  was  a  boat. A  small  boat."

	var builder strings.Builder

	for _, sub := range strings.Fields(s) {
		if sub == "small" {
			builder.WriteString("very ")
		}
		builder.WriteString(sub)
		builder.WriteString(" ")
	}

	fmt.Println(builder.String())

	// Output:
	//It was a boat. A very small boat.
}
