package functional

import "reflect"

// UnpackSliceValue unpacks a slice value to a slice
func UnpackSliceValue(v reflect.Value) []reflect.Value {
	l := v.Len()
	unpack := make([]reflect.Value, l, l)
	for i := 0; i < l; i++ {
		unpack[i] = v.Index(i)
	}

	return unpack
}

// CopyFunctionInputTypes copies function input types to a slice
func CopyFunctionInputTypes(t reflect.Type) []reflect.Type {
	numIn := t.NumIn()
	inputs := make([]reflect.Type, numIn, numIn)
	for i := 0; i < numIn; i++ {
		inputs[i] = t.In(i)
	}

	return inputs
}

// CopyFunctionOutputTypes copies function output types to a slice
func CopyFunctionOutputTypes(t reflect.Type) []reflect.Type {
	numOut := t.NumOut()
	outputs := make([]reflect.Type, numOut, numOut)
	for i := 0; i < numOut; i++ {
		outputs[i] = t.Out(i)
	}

	return outputs
}

// MakeSpecificFunc assigns a generic func implementation to a
// specific func definition
func MakeSpecificFunc(fptr interface{}, genericFunc func([]reflect.Value) []reflect.Value) {
	fn := reflect.ValueOf(fptr).Elem()
	fn.Set(reflect.MakeFunc(fn.Type(), genericFunc))
}
