package reflection

import "reflect"

func Walk(x interface{}, fn func(string)) {
	val := getValue(x)

	walkValue := func(value reflect.Value) {
		Walk(value.Interface(), fn)
	}

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walkValue(val.Field(i))
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			walkValue(val.Index(i))
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walkValue(val.MapIndex(key))
		}
	}
	// numberOfValues := 0
	// var getField func(int) reflect.Value

	// switch val.Kind() {
	// case reflect.Struct:
	// 	numberOfValues = val.NumField()
	// 	getField = val.Field
	// case reflect.Slice, reflect.Array:
	// 	numberOfValues = val.Len()
	// 	getField = val.Index
	// case reflect.Map:
	// 	for _, key := range val.MapKeys() {
	// 		Walk(val.MapIndex(key).Interface(), fn)
	// 	}
	// case reflect.String:
	// 	fn(val.String())
	// }

	// for i := 0; i < numberOfValues; i++ {
	// 	Walk(getField(i).Interface(), fn)
	// }
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}
	return val
}
