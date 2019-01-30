package functional

import (
	"errors"
	"reflect"
)

// GenericCurry creates a wrapper around a given function so the
// original function can be called a parameter at the time.
func GenericCurry(in []reflect.Value) []reflect.Value {
	// Check number of input arguments
	if len(in) != 1 {
		panic(errors.New("Curry expects one argument"))
	}

	f := in[0]
	ft := f.Type()

	if ft.Kind() == reflect.Interface {
		ft = reflect.TypeOf(f.Interface())
	}

	if ft.Kind() != reflect.Func {
		panic(errors.New("Curry expects a function as the first argument"))
	}

	numIn := ft.NumIn()

	// If number of inputs is 1 or 0 return the original function
	if numIn <= 1 {
		return []reflect.Value{f}
	}

	// Copy function input types
	inputTypes := CopyFunctionInputTypes(ft)

	// Copy function output types
	outputTypes := CopyFunctionOutputTypes(ft)

	// Create last output function type
	outputFuncTypes := make([]reflect.Type, numIn, numIn)
	outputFuncTypes[numIn-1] = reflect.FuncOf(
		[]reflect.Type{inputTypes[numIn-1]},
		outputTypes,
		ft.IsVariadic(),
	)

	// Create rest of output function types
	for i := numIn - 2; i >= 0; i-- {
		outputFuncTypes[i] = reflect.FuncOf(
			[]reflect.Type{inputTypes[i]},
			[]reflect.Type{outputFuncTypes[i+1]},
			false,
		)
	}

	// Forward actual curry declaration
	var curryFun func([]reflect.Value) []reflect.Value

	// Create input values cache and initial count and encapsulate them in
	// the curry func implementation
	inputValues := make([]reflect.Value, numIn, numIn)
	inputCount := 0
	curryFun = func(args []reflect.Value) []reflect.Value {
		inputValues[inputCount] = args[0]
		inputCount++

		// If input count is the same as number of input arguments return
		// original function call
		if inputCount == numIn {
			return f.Call(inputValues)
		}

		// Return the next function step of the curry wrapper
		returnFunc := reflect.MakeFunc(outputFuncTypes[inputCount], curryFun)
		return []reflect.Value{returnFunc}
	}

	// Return the first curry wrapper
	returnFunc := reflect.MakeFunc(outputFuncTypes[0], curryFun)
	return []reflect.Value{returnFunc}
}
