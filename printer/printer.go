package printer

import (
	"fmt"
	"reflect"
)

// PrintStruct recurses over a struct to print its name, field names,
// and values in a readable format.
// Supports public attributes of the following types: booleans, strings,
// numeric types, structs, arrays, slices, and maps.
func PrintStruct(i interface{}) (err error) {
	t := reflect.TypeOf(i)
	if t == nil {
		err = fmt.Errorf("error: expected type struct but received type nil")
		return err
	}
	if t.Kind() != reflect.Struct {
		err = fmt.Errorf("error: expected type struct but received type %s", t.Kind().String())
		return err
	}

	fmt.Println()

	err = printStructRec(i, "")
	if err != nil {
		return err
	}

	fmt.Println()

	return nil
}

func printStructRec(i interface{}, tab string) (err error) {
	_type := reflect.TypeOf(i)
	structName(_type.Name(), tab)

	value := reflect.ValueOf(i)
	for j := 0; j < value.NumField(); j++ {
		field := value.Field(j)
		name := _type.Field(j).Name

		err := printField(name, tab, field)
		if err != nil {
			return err
		}
	}

	return nil
}

func printField(name, tab string, value reflect.Value) (err error) {
	kind := value.Kind()
	switch {
	case isPrimitive(kind):
		primitive(name, tab, value.Interface())
	case kind == reflect.Struct:
		header(kind, name, tab)
		err = printStructRec(value.Interface(), incTab(tab))
		if err != nil {
			return err
		}
	case kind == reflect.Array || kind == reflect.Slice:
		header(kind, name, tab)
		for i := 0; i < value.Len(); i++ {
			v := value.Index(i)
			listElem(v.Interface(), incTab(tab))
		}
		footer(kind, tab)
	case kind == reflect.Map:
		header(kind, name, tab)
		for _, k := range value.MapKeys() {
			v := value.MapIndex(k)
			mapPair(k.Interface(), v.Interface(), incTab(tab))
		}
		footer(kind, tab)
	default:
		err = fmt.Errorf("error: field type %s not supported", kind)
		return err
	}

	return nil
}

//
// Helpers for printing struct elements and formatting
//

func structName(name, tab string) {
	fmt.Printf("%sObject of Class \"%s\"\n", tab, name)
	fmt.Printf("%s--------------------------------\n", tab)
}

func header(kind reflect.Kind, name, tab string) {
	switch kind {
	case reflect.Struct:
		fmt.Printf("%s%s =\n", tab, name)
	case reflect.Array, reflect.Slice:
		fmt.Printf("%s%s = [\n", tab, name)
	case reflect.Map:
		fmt.Printf("%s%s = map[\n", tab, name)
	}
}

func footer(kind reflect.Kind, tab string) {
	switch kind {
	case reflect.Array, reflect.Slice, reflect.Map:
		fmt.Printf("%s]\n", tab)
	}
}

func primitive(n, tab string, i interface{}) {
	fmt.Printf("%s%s = %v\n", tab, n, i)
}

func listElem(i interface{}, tab string) {
	fmt.Printf("%s%v\n", tab, i)
}

func mapPair(k, v interface{}, tab string) {
	fmt.Printf("%s%v : %v\n", tab, k, v)
}

func incTab(tab string) string {
	return tab + "\t"
}
