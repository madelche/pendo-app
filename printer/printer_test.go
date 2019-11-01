package printer

import (
	"fmt"
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

type PersonChan struct {
	Age  int
	Chan chan int
}

func TestPrintStruct_InvalidArgNil(t *testing.T) {
	result := PrintStruct(nil)
	if result == nil {
		t.Error("PrintStruct(nil) failed, expected error")
	}
}

func TestPrintStruct_InvalidArgString(t *testing.T) {
	result := PrintStruct("invalid arg")
	if result == nil {
		t.Error("PrintStruct(string) failed, expected error")
	}
}

func TestPrintStruct_InvalidArgFunc(t *testing.T) {
	result := PrintStruct(func() { fmt.Println("A test function") })
	if result == nil {
			t.Error("PrintStruct(func) failed, expected error")
	}
}

func TestPrintStruct_UnsupportedStructFieldPtr(t *testing.T) {
	n := Name{FirstName: "Bill", LastName: "Gates"}
	p := PersonPtr{Age: 55, Name: &n}

	result := PrintStruct(p)
	if result == nil {
		t.Error("PrintStruct(ptr) failed, expected error")
	}
}

func TestPrintStruct_UnsupportedStructFieldChan(t *testing.T) {
	p := PersonChan{Age: 55, Chan: make(chan int)}

	result := PrintStruct(p)
	if result == nil {
		t.Error("PrintStruct(ptr) failed, expected error")
	}
}
