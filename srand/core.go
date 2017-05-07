package srand

import (
	"math/rand"
	"reflect"
	"time"
)

var (
	rng = rand.New(rand.NewSource(time.Now().UnixNano()))
)

// RandomOfSet returns a random value from a set.
// If the input is a map, the return value is a random key from the map.
// If the input is not a container type, the input is returned.
//
// This function operates on generics, and uses reflection in the background.
// The return value must be casted back into the original type by the caller, eg.:
//   arr := []int{1,2,3,4,5}
//   x := RandomOfSet(arr).(int)
func RandomOfSet(set interface{}) interface{} {
	v := reflect.ValueOf(set)

	switch v.Kind() {
	case reflect.Array, reflect.Slice:
		return _randList(v)
	case reflect.Map:
		return _randMap(v)
	}
	return set
}

// RandomOfSet returns a slice of random values from a set.
// If the input is a map, the return value is a slice of random keys from the map.
// If the input is not a container type, it is returned as a single element []interface{} containing only the input.
//
// This function operates on generics, and uses reflection in the background.
// The return value must be copy-casted back into the original container type by the caller, eg.:
//   n := 20
//   arr := []int{1,2,3,4,5}
//   result := make([]int, n)
//   for i, d := range RandomNOfSet(arr, n).([]interface{}) {
//       result[i] := d.(int)
//   }
func RandomNOfSet(set interface{}, n int) []interface{} {
	v := reflect.ValueOf(set)
	result := make([]interface{}, n)
	var f func(reflect.Value) interface{}

	switch v.Kind() {
	case reflect.Array, reflect.Slice:
		f = _randList
	case reflect.Map:
		f = _randMap
	default:
		return []interface{}{set}
	}

	for i := 0; i < n; i++ {
		result[i] = f(v)
	}
	return result
}

func _randList(v reflect.Value) interface{} {
	return v.Index(rng.Intn(v.Len())).Interface()
}

func _randMap(v reflect.Value) interface{} {
	return v.MapIndex(RandomOfSet(v.MapKeys()).(reflect.Value)).Interface()
}
