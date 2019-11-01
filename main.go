package main

import (
	"log"

	"github.com/madelche/pendo-app/printer"
)

func main() {

	type Name struct {
		FirstName string
		LastName  string
	}

	type Person struct {
		Age  int
		Name Name
	}

	n := Name{FirstName: "Bill", LastName: "Gates"}
	p := Person{Age: 55, Name: n}

	err := printer.PrintStruct(p)
	if err != nil {
		log.Fatal(err)
	}
}
