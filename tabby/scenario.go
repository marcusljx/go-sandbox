package tabby

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"testing"
)

type scenario struct {
	testFunc reflect.Value
	inputs   []reflect.Value
	outputs  []reflect.Value
}

func (s *scenario) t() func(t *testing.T) {
	return func(t *testing.T) {
		x := s.testFunc.Call(s.inputs)
		for i, a := range x {
			e := s.outputs[i]
			actual := a.Interface()
			expected := e.Interface()

			if expected != actual {
				t.Error(s.errorMsg(expected, actual))
			}
		}
	}
}

func (s *scenario) errorMsg(expected, actual interface{}) string {
	testName := runtime.FuncForPC(s.testFunc.Pointer()).Name()
	argList := make([]string, len(s.inputs))
	for i, in := range s.inputs {
		argList[i] = fmt.Sprint(in)
	}

	funcString := "\n" + formatKeyValueOutput("function", fmt.Sprintf("%s(%v)", testName, strings.Join(argList, ", ")))
	expectedLine := formatKeyValueOutput("expected", expected)
	actualLine := formatKeyValueOutput("actual", actual)
	return funcString + "\n" + expectedLine + "\n" + actualLine

}
