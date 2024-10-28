package main

import (
	"reflect"
)

func walk(x any, fn func(in string)) {
	// detect pointers before we get to
	// the NumField() call in for loop
	val := getVal(x)
	// very clever abstraction keeping the fn call global
	walkVal := func(value reflect.Value) {
		walk(value.Interface(), fn)
	}

	// recursion allows walk to pass nested things
	// recursively back through the switch statement
	// until they've all made it through the last case
	// very water fountainy in nature
	// unevaporated water is pumped right back out the top again
	// to tumble down each layer until it finds the one where
	// the sun deems it evaporation worthy (case doesnt recurse)
	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walkVal(val.Field(i))
		}
	case reflect.Slice, reflect.Array:
		// now slices giving us attitude
		for i := 0; i < val.Len(); i++ {
			walkVal(val.Index(i))
		}
	case reflect.Map:
		// and of course the heavy lifting map
		for _, key := range val.MapKeys() {
			walkVal(val.MapIndex(key))
		}
	case reflect.Chan:
		for {
			if v, ok := val.Recv(); ok {
				walkVal(v)
			} else {
				break
			}
		}
	case reflect.Func:
		valFnResult := val.Call(nil)
		for _, res := range valFnResult {
			walkVal(res)
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
