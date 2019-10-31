package printer

import (
	"fmt"
	"reflect"
)

// Questions:
// Pointers to struct?

// TypeOf ValueOf
func PrintStruct(i interface{}) (err error) {
	t := reflect.TypeOf(i)
	if t == nil {
		err = fmt.Errorf("error: expected type struct but received type nil")
		return err
	}
	if t.Kind() != reflect.Struct {
		err = fmt.Errorf("error: expected type struct but received type %s", t.Name())
		return err
	}

	fmt.Println()

	printStructRec(i, "")

	fmt.Println()

	return nil
}

func printStructHeader(n, tab string) {
	fmt.Printf("%sObject of Class \"%s\"\n", tab, n)
	fmt.Printf("%s--------------------------------\n", tab)
}

func printStructRec(i interface{}, tab string) {
	t := reflect.TypeOf(i)
	printStructHeader(t.Name(), tab)

	v := reflect.ValueOf(i)
	for j := 0; j < v.NumField(); j++ {
		f := v.Field(j)
		n := t.Field(j).Name
		if isPrimitive(f.Type()) {
			printPrimitive(n, tab, f.Interface())
		}

		if f.Kind() == reflect.Struct {
			fmt.Printf("%s =\n", n)
			printStructRec(f.Interface(), incTab(tab))
		}
	}
}

func printPrimitive(n, tab string, i interface{}) {
	fmt.Printf("%s%s = %v\n", tab, n, i) // TODO: add quotation marks to string printout
}

func incTab(tab string) string {
	return tab + "\t"
}

// TODO: move into reflect utils file?
func isPrimitive(t reflect.Type) bool {
	switch t.Kind() {
	case reflect.String,
		reflect.Bool,
		reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64,
		reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64,
		reflect.Uintptr,
		reflect.Float32,
		reflect.Float64,
		reflect.Complex64,
		reflect.Complex128:
		return true
	}
	return false
}
