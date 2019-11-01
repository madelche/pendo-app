package printer

import (
	"testing"
)

type Name struct {
	FirstName string
	LastName  string
}

type PersonPtr struct {
	Age  int
	Name *Name
}

func TestPrintStruct_NilArg(t *testing.T) {
	result := PrintStruct(nil)
	if result == nil {
		t.Error("PrintStruct(nil) failed, expected error")
	}
}

func TestPrintStruct_InvalidArg(t *testing.T) {
	result := PrintStruct("invalid arg")
	if result == nil {
		t.Error("PrintStruct(string) failed, expected error")
	}
}

func TestPrintStruct_UnsupportedStructField(t *testing.T) {
	n := Name{FirstName: "Bill", LastName: "Gates"}
	p := PersonPtr{Age: 55, Name: &n}

	result := PrintStruct(p)
	if result == nil {
		t.Error("PrintStruct(ptr) failed, expected error")
	}
}
