package printer

import (
	"testing"
)

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
