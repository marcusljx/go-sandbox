package tabby

import "reflect"

// GivenEvent is the intermediate object returned by Given(),
// which must be followed immediately by Expect()
type GivenEvent struct {
	subtestName string
	_w          *WrappedFunction
}

// Expect sets the expected output values for the given values in the preceding Given() call
func (o *GivenEvent) Expect(outputs ...interface{}) *WrappedFunction {
	s := o._w.scenarios[o.subtestName]
	for _, out := range outputs {
		s.outputs = append(s.outputs, reflect.ValueOf(out))
	}
	o._w.scenarios[o.subtestName] = s
	return o._w
}
