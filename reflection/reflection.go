package reflection

import (
	"reflect"
)

func walk(x interface{}, fn func(string)) {
	value := getValue(x)

	switch value.Kind() {
	case reflect.Struct:
		for i := 0; i < value.NumField(); i++ {
			walk(value.Field(i).Interface(), fn)
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < value.Len(); i++ {
			walk(value.Index(i).Interface(), fn)
		}
	case reflect.Map:
		for _, key := range value.MapKeys() {
			walk(value.MapIndex(key).Interface(), fn)
		}
	case reflect.Chan:
		for {
			if v, ok := value.Recv(); ok {
				walk(v.Interface(), fn)
			} else {
				break
			}
		}
	case reflect.Func:
		result := value.Call(nil)

		for _, v := range result {
			walk(v.Interface(), fn)
		}
	case reflect.String:
		fn(value.String())
	}
}

func getValue(x interface{}) reflect.Value {
	value := reflect.ValueOf(x)

	// handle x ptr
	if value.Kind() == reflect.Pointer {
		value = value.Elem()
	}

	return value
}
