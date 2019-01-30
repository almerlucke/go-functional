package functional

import (
	"container/list"
	"errors"
	"reflect"
)

// GenericFilter implements a generic functional filter
// expects a filter function as first argument and a slice
// as second value
func GenericFilter(in []reflect.Value) []reflect.Value {
	// Check number of input arguments
	if len(in) != 2 {
		panic(errors.New("Filter expects two arguments"))
	}

	// The first arg needs to be a function of the form func (T1) T2
	f := in[0]
	ft := f.Type()

	if ft.Kind() != reflect.Func {
		panic(errors.New("Filter expects a function as the first argument"))
	}

	if ft.NumIn() != 1 {
		panic(errors.New("Filter expects the function to have one input argument"))
	}

	if ft.NumOut() != 1 {
		panic(errors.New("Filter expects the function to have one output value"))
	}

	if ft.Out(0).Kind() != reflect.Bool {
		panic(errors.New("Filter expects the function to have a boolean output argument"))
	}

	// The second arg needs to be a slice of T1 objects
	a := in[1]
	if a.Type().Kind() != reflect.Slice {
		panic(errors.New("Filter expects a slice as the first argument"))
	}

	// Check if the slice elem type is the same as the input argument
	// of the filter function
	if a.Type().Elem() != ft.In(0) {
		panic(errors.New("Filter expects the function to have the same input type as the slice to filter"))
	}

	// Create a temp list
	lst := list.New()
	n := a.Len()

	// Filter input slice based on filter function boolean result
	for i := 0; i < n; i++ {
		elem := a.Index(i)
		result := f.Call([]reflect.Value{elem})[0]
		if result.Interface().(bool) {
			lst.PushBack(elem)
		}
	}

	// Copy list to slice, in this way the resulting slice has an exact
	// amount of elements
	len := lst.Len()

	// The result will be a filtered slice of T1
	filteredResult := reflect.MakeSlice(a.Type(), len, len)

	i := 0
	for el := lst.Front(); el != nil; el = el.Next() {
		filteredResult.Index(i).Set(el.Value.(reflect.Value))
		i++
	}

	// Return the resulting slice as the only return value
	return []reflect.Value{filteredResult}
}
