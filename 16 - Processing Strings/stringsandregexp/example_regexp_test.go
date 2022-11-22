package main

import (
	"fmt"
	"regexp"
)

func Example_FindString() {
	pattern := regexp.MustCompile("K[a-z]{4}|[a-z]oat")

	s := "Kayak. A boat for one person."

	firstMatch := pattern.FindString(s)
	fmt.Println("First match:", firstMatch)

	allMatches := pattern.FindAllString(s, -1)
	for i, m := range allMatches {
		fmt.Println("Match", i, "=", m)
	}

	// Output:
	//First match: Kayak
	//Match 0 = Kayak
	//Match 1 = boat
}

func Example_Split() {
	pattern := regexp.MustCompile(" |boat|one")

	s := "Kayak. A boat for one person."

	split := pattern.Split(s, -1)
	for _, s := range split {
		if s != "" {
			fmt.Println(s)
		}
	}

	// Output:
	//Kayak.
	//A
	//for
	//person.
}

func Example_subexpressions() {
	pattern := regexp.MustCompile("A ([a-z]*) for ([a-z]*) person")

	s := "Kayak. A boat for one person."

	subs := pattern.FindStringSubmatch(s)

	for _, sub := range subs {
		fmt.Println("Match:", sub)
	}

	// Output:
	//Match: A boat for one person
	//Match: boat
	//Match: one
}

func Example_namedSubexpressions() {
	pattern := regexp.MustCompile(
		"A (?P<type>[a-z]*) for (?P<capacity>[a-z]*) person")

	s := "Kayak. A boat for one person."

	subs := pattern.FindStringSubmatch(s)

	for _, name := range []string{"type", "capacity"} {
		fmt.Println(name, "=", subs[pattern.SubexpIndex(name)])
	}

	// Output:
	//type = boat
	//capacity = one
}

func Example_namedSubexpressions_replace() {
	pattern := regexp.MustCompile(
		"A (?P<type>[a-z]*) for (?P<capacity>[a-z]*) person")

	s := "Kayak. A boat for one person."

	template := "(type: ${type}, capacity: ${capacity})"

	replaced := pattern.ReplaceAllString(s, template)
	fmt.Println("replaced:", replaced)

	replacedLiteral := pattern.ReplaceAllLiteralString(s, template)
	fmt.Println("replacedLiteral:", replacedLiteral)

	// Output:
	//replaced: Kayak. (type: boat, capacity: one).
	//replacedLiteral: Kayak. (type: ${type}, capacity: ${capacity}).
}

func Example_namedSubexpressions_replace_byNum() {
	pattern := regexp.MustCompile(
		"A (?P<type>[a-z]*) for (?P<capacity>[a-z]*) person")

	s := "Kayak. A boat for one person."

	template := "(type: ${1}, capacity: ${2})"

	replaced := pattern.ReplaceAllString(s, template)
	fmt.Println("replaced:", replaced)

	replacedLiteral := pattern.ReplaceAllLiteralString(s, template)
	fmt.Println("replacedLiteral:", replacedLiteral)

	// Output:
	//replaced: Kayak. (type: boat, capacity: one).
	//replacedLiteral: Kayak. (type: ${1}, capacity: ${2}).
}

func Example_() {
	pattern := regexp.MustCompile(
		"A (?P<type>[a-z]*) for (?P<capacity>[a-z]*) person")

	s := "Kayak. A boat for one person."

	f := func(s string) string {
		return "This is the replacement content"
	}

	replaced := pattern.ReplaceAllStringFunc(s, f)
	fmt.Println("replaced:", replaced)

	// Output:
	//replaced: Kayak. This is the replacement content.
}
