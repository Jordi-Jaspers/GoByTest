package main

import (
	"reflect"
)

// golang challenge:
// Write a function walk(x interface{}, fn func(string))
// which takes a struct x and calls fn for all strings fields found inside.
// difficulty level: recursively.

// Solution: Reflection
// Reflection in computing is the ability of a program to examine its own structure,
// particularly through types; it's a form of metaprogramming. It's also a great source of confusion.

//val.Recv receives and returns a value from the channel v.
// for v, ok := val.Recv(); ok; v, ok = val.Recv() {
// 	walk(v.Interface(), fn)
// }
// A for loop over a channel. whenever it returns ok keep looping else stop.
// also possible with maps.

func walk(x interface{}, fn func(string)) {
	val := getValue(x)

	switch val.Kind() {
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walk(val.Field(i).Interface(), fn)
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			walk(val.Index(i).Interface(), fn)
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walk(val.MapIndex(key).Interface(), fn)
		}
	case reflect.Chan:
		for v, ok := val.Recv(); ok; v, ok = val.Recv() {
			walk(v.Interface(), fn)
		}
	case reflect.Func:
		for _, v := range val.Call(nil) {
			walk(v.Interface(), fn)
		}
	case reflect.String:
		fn(val.String())
	}
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)
	if val.Kind() == reflect.Ptr {
		return val.Elem()
	}
	return val
}
