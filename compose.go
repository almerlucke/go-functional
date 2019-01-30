package functional

import (
	"errors"
	"reflect"
)

// GenericCompose composes two functions together, call second function first,
// the result is passed as arg(s) to the first function, the result is returned
func GenericCompose(in []reflect.Value) []reflect.Value {
	if len(in) != 2 {
		panic(errors.New("Compose expects two arguments"))
	}

	f := in[0]
	g := in[1]
	gt := g.Type()
	ft := f.Type()

	if gt.Kind() != reflect.Func || ft.Kind() != reflect.Func {
		panic(errors.New("Compose expects two functions as arguments"))
	}

	inputArgs := CopyFunctionInputTypes(gt)
	outputArgs := CopyFunctionOutputTypes(ft)

	rt := reflect.FuncOf(inputArgs, outputArgs, gt.IsVariadic())
	r := reflect.MakeFunc(rt, func(args []reflect.Value) []reflect.Value {
		return f.Call(g.Call(args))
	})

	return []reflect.Value{r}
}
