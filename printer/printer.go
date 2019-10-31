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

	//fmt.Println(t.Name())


	return nil
}

