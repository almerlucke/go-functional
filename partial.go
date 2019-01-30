package functional

import (
	"errors"
	"reflect"
)

// GenericPartial creates a generic partial function call
// expects a function and a number of partial arguments to that function
func GenericPartial(in []reflect.Value) []reflect.Value {
	totalArgs := len(in)
	if totalArgs < 2 {
		panic(errors.New("Partial expects at least two arguments"))
	}

	// Get number of partial arguments
	numArgs := totalArgs - 1

	f := in[0]
	ft := f.Type()
	if ft.Kind() != reflect.Func {
		panic(errors.New("Partial expects a function as first argument"))
	}

	numInputArgs := ft.NumIn()
	numLeftOverArgs := numInputArgs - numArgs

	if numLeftOverArgs <= 0 {
		panic(errors.New("Partial number of arguments exceed number of function input arguments"))
	}

	// Copy left over input argument types
	inTypes := make([]reflect.Type, numLeftOverArgs, numLeftOverArgs)
	j := 0
	for i := numArgs; i < numInputArgs; i++ {
		inTypes[j] = ft.In(i)
		j++
	}

	// Copy output types
	outTypes := make([]reflect.Type, ft.NumOut(), ft.NumOut())
	for i := 0; i < ft.NumOut(); i++ {
		outTypes[i] = ft.Out(i)
	}

	// Create a new function encapsulating the partial arguments with as
	// input arguments the remaining arguments for the original function
	resultFunType := reflect.FuncOf(inTypes, outTypes, ft.IsVariadic())
	fun := reflect.MakeFunc(resultFunType, func(args []reflect.Value) []reflect.Value {
		return f.Call(append(in[1:], args...))
	})

	// Return partial function
	return []reflect.Value{fun}
}
