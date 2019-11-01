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
		Age      int
		Name     Name
		Height   float32
		Children []Person
	}

	n0 := Name{FirstName: "Bill", LastName: "Junior"}
	p0 := Person{Age: 9, Name: n0, Height: 1.1}

	n1 := Name{FirstName: "Melinda", LastName: "Junior"}
	p1 := Person{Age: 12, Name: n1, Height: 1.3}

	n := Name{FirstName: "Bill", LastName: "Gates"}
	p := Person{Age: 55, Name: n, Children: []Person{p0, p1}, Height: 1.7}

	err := printer.PrintStruct(p)
	if err != nil {
		log.Fatal(err)
	}

	type Degree struct {
		Level int
		Field string
	}

	type Professor struct {
		Name   Name
		Degree Degree
	}

	type Course struct {
		Title  string
		Number int
		Prof   Professor
	}

	d := Degree{Level: 2, Field: "Computer Science"}
	prof0 := Professor{Name: Name{"Paul", "Carter"}, Degree: d}
	course0 := Course{Title: "Intro to Software Engineering", Number: 210, Prof: prof0}

	err = printer.PrintStruct(course0)
	if err != nil {
		log.Fatal(err)
	}

	type University struct {
		Name    string
		Courses map[int]Course
	}

	prof1 := Professor{Name: Name{"Don", "Acton"}, Degree: d}
	course1 := Course{Title: "Advanced Operating Systems", Number: 415, Prof: prof1}
	u := University{Name: "The University of British Columbia", Courses: map[int]Course{210: course0, 415: course1}}

	err = printer.PrintStruct(u)
	if err != nil {
		log.Fatal(err)
	}
}
