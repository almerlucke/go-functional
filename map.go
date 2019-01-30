package functional

import (
	"errors"
	"reflect"
)

// GenericMap maps a generic function over a generic array
// expects a map function as first argument and a slice as second
// argument
func GenericMap(in []reflect.Value) []reflect.Value {
	// Check number of input arguments
	if len(in) != 2 {
		panic(errors.New("Map expects two arguments"))
	}

	// The first arg needs to be a function of the form func (T1) T2
	f := in[0]
	ft := f.Type()

	if ft.Kind() != reflect.Func {
		panic(errors.New("Map expects a function as the first argument"))
	}

	if ft.NumIn() != 1 {
		panic(errors.New("Map expects the function to have one input argument"))
	}

	if ft.NumOut() != 1 {
		panic(errors.New("Map expects the function to have one output value"))
	}

	// The second arg needs to be a slice of T1 objects
	a := in[1]
	if a.Type().Kind() != reflect.Slice {
		panic(errors.New("Map expects a slice as the first argument"))
	}

	// Check if the slice elem type is the same as the input argument
	// of the mapped function
	if a.Type().Elem() != ft.In(0) {
		panic(errors.New("Map expects the function to have the same input type as the slice to map"))
	}

	// Get the length of the array
	l := a.Len()

	// We expect the output slice type of the map to be the same type
	// as the output type of the function passed to map
	outputType := f.Type().Out(0)

	// The result will be a slice of T2
	result := reflect.MakeSlice(reflect.SliceOf(outputType), l, l)

	// Call the function to map and assign the first output value to
	// the result slice at index
	for i := 0; i < l; i++ {
		result.Index(i).Set(f.Call([]reflect.Value{a.Index(i)})[0])
	}

	// Return the resulting slice as the only return value
	return []reflect.Value{result}
}
