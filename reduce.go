package functional

import (
	"errors"
	"reflect"
)

// GenericReduce reduces a slice to a single value
// expects the first arg to be a function which accepts an accumulator and a
// next value like func(T1, T2) T1. The second argument is the initial
// accumulation value, the third arg is the slice to reduce
func GenericReduce(in []reflect.Value) []reflect.Value {
	if len(in) != 3 {
		panic(errors.New("Reduce expects three arguments"))
	}

	f := in[0]
	ft := f.Type()
	acc := in[1]
	a := in[2]

	if ft.Kind() != reflect.Func {
		panic(errors.New("Reflect expects a function as the first argument"))
	}

	if ft.NumIn() != 2 {
		panic(errors.New("Reflect expects the function to have two input arguments"))
	}

	if ft.NumOut() != 1 {
		panic(errors.New("Reflect expects the function to have one output value"))
	}

	if ft.In(0) != ft.Out(0) {
		panic(errors.New("Reflect expects the function accumulation argument to have the same type as the output value"))
	}

	if ft.In(0) != acc.Type() {
		panic(errors.New("Reflect expects the function accumulation argument to have the same type as the initial accumulation argument"))
	}

	if a.Kind() != reflect.Slice {
		panic(errors.New("Reflect expects a slice as the third argument"))
	}

	if ft.In(1) != a.Type().Elem() {
		panic(errors.New("Reflect expects the function next argument to have the same type as the slice elements"))
	}

	// Do the actual reduction
	l := a.Len()

	for i := 0; i < l; i++ {
		acc = f.Call([]reflect.Value{acc, a.Index(i)})[0]
	}

	return []reflect.Value{acc}
}
