package tabby

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

// WrappedFunction defines a test case for a given function
type WrappedFunction struct {
	originalFunction interface{}
	value            reflect.Value
	scenarios        map[string]scenario
}

// Func wraps a function
func Func(function interface{}) *WrappedFunction {
	f := reflect.ValueOf(function)
	if f.Kind() != reflect.Func {
		panic(fmt.Sprintf("expected reflect.Func type but input was %v", f.Kind()))
	}

	return &WrappedFunction{
		originalFunction: function,
		value:            f,
		scenarios:        make(map[string]scenario),
	}
}

// Run performs all scenarios as subtests
func (w *WrappedFunction) Run(t *testing.T) {
	for name, test := range w.scenarios {
		t.Run(name, test.t())
	}
}

// On creates a scenario based on the inputs.
// This call should be immediately followed by an Expect() call.
func (w *WrappedFunction) On(inputs ...interface{}) *OnEvent {
	o := &OnEvent{
		_w:          w,
		subtestName: strings.Replace(fmt.Sprint(inputs...), " ", "_", -1),
	}

	ns := scenario{
		testFunc: w.value,
	}

	for _, in := range inputs {
		ns.inputs = append(ns.inputs, reflect.ValueOf(in))
	}

	w.scenarios[o.subtestName] = ns
	return o
}
