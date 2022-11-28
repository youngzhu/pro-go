package main

import (
	"html/template"
	"os"
)

func Example_ParseFiles() {
	t, err := template.ParseFiles("templates/template.html")
	if err == nil {
		t.Execute(os.Stdout, &Kayak)
	} else {
		Printfln("Error: %v", err.Error())
	}

	// output:

}

func Example_ParseFiles_multiple() {
	all, err := template.ParseFiles("templates/template.html",
		"templates/extras.html")
	if err == nil {
		all.ExecuteTemplate(os.Stdout, "template.html", &Kayak)
		os.Stdout.WriteString("\n\n")
		all.ExecuteTemplate(os.Stdout, "extras.html", &Kayak)
	} else {
		Printfln("Error: %v", err.Error())
	}

	// output:

}

func Example_ParseGlob() {
	all, err := template.ParseGlob("templates/*.html")
	if err == nil {
		for _, t := range all.Templates() {
			Printfln("Template name: %v", t.Name())
		}
	} else {
		Printfln("Error: %v", err.Error())
	}

	// Output:

}

func Example_ParseGlob_Lookup() {
	all, err := template.ParseGlob("templates/*.html")
	if err == nil {
		selected := all.Lookup("template.html")
		err = selected.Execute(os.Stdout, &Kayak)
	}

	if err != nil {
		Printfln("Error: %v", err.Error())
	}

	// Output:

}
