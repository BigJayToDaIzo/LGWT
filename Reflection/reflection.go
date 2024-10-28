package main

import "reflect"

func walk(x any, fn func(in string)) {
	// detect pointers before we get to
	// the NumField() call in for loop
	val := getVal(x)

	// now slices giving us attitude
	// walk them and short circuit return when done
	if val.Kind() == reflect.Slice {
		for i := 0; i < val.Len(); i++ {
			walk(val.Index(i).Interface(), fn)
		}
		return
	}

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		switch field.Kind() {
		case reflect.Struct:
			// recurse like a BOSS
			walk(field.Interface(), fn)
		case reflect.String:
			fn(field.String())
		}
	}
}

func getVal(x any) reflect.Value {
	val := reflect.ValueOf(x)
	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}
	return val
}
