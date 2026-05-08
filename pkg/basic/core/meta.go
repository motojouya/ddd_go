package core

import (
	"fmt"
	"reflect"
)

func GetType(i interface{}) (reflect.Type, error) {
	t := reflect.TypeOf(i)

	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if t.Kind() == reflect.Struct {
		return t, nil
	}

	if t.Kind() == reflect.Slice {
		return t.Elem(), nil
	}

	return nil, fmt.Errorf("type shoud be struct or slice. type: %v", reflect.TypeOf(i))
}
