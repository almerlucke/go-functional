package functional

import (
	"errors"
	"reflect"
)

// GenericPipe creates a generic pipe function calling all passed in functions
// in sequence and passing the result through the chain.
func GenericPipe(in []reflect.Value) []reflect.Value {
	totalArgs := len(in)
	if totalArgs != 1 {
		panic(errors.New("Pipe expects a slice of functions"))
	}

	funcs := in[0]
	if funcs.Kind() != reflect.Slice {
		panic(errors.New("Pipe expects a slice of functions"))
	}

	numFuncs := funcs.Len()

	if numFuncs < 1 {
		panic(errors.New("Pipe expects at least one function as argument"))
	}

	// Unpack slice value to real slice
	funcsSlice := UnpackSliceValue(funcs)
	for i := 0; i < numFuncs; i++ {
		fun := funcsSlice[i]
		kind := fun.Type().Kind()

		// Check if value is func
		if kind != reflect.Func {
			if kind == reflect.Interface {
				// Change interface value to underlying function value if needed
				fun = reflect.ValueOf(fun.Interface())

				// Panic if not a func
				if fun.Type().Kind() != reflect.Func {
					panic(errors.New("Pipe expects only functions as arguments"))
				}

				// Store actual function value
				funcsSlice[i] = fun
			} else {
				panic(errors.New("Pipe expects only functions as arguments"))
			}
		}
	}

	// Return function signature copies input args from first function passed to pipe
	// and output values from last function passed to pipe
	firstFn := funcsSlice[0].Type()
	lastFn := funcsSlice[numFuncs-1].Type()
	inputTypes := CopyFunctionInputTypes(firstFn)
	outputTypes := CopyFunctionOutputTypes(lastFn)
	funcType := reflect.FuncOf(inputTypes, outputTypes, firstFn.IsVariadic())
	resultFun := reflect.MakeFunc(funcType, func(args []reflect.Value) []reflect.Value {
		inputArgs := args
		for _, f := range funcsSlice {
			inputArgs = f.Call(inputArgs)
		}
		return inputArgs
	})

	return []reflect.Value{resultFun}
}
