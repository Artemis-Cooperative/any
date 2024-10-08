package anytypes

import (
	"reflect"
)

// Return true if the value given is an array. Else, false.
func IsArray(v any) bool {
	return isOfReflectKind(v, reflect.Array)
}

// Return true if the value given is a basic type. Else, false.
func IsBasic(v any) bool {
	switch v.(type) {
	case bool, int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64, uintptr,
		float32, float64,
		complex64, complex128,
		string:
		return true
	default:
		return false
	}
}

// Return true if the value given is comparable. Else, false.
func IsComparable(v any) bool {
	return reflect.ValueOf(v).Comparable()
}

// Return true if the value given is a map. Else, false.
func IsMap[V any](v V) bool {
	return isOfReflectKind(v, reflect.Map)
}

// Return true if the value given is a valid numeric key. Else, false.
func IsNumKey(v any) bool {
	switch v := v.(type) {
	case int:
		return v > 0
	case int8:
		return v > 0
	case int16:
		return v > 0
	case int32:
		return v > 0
	case int64:
		return v > 0
	case uint:
		return v > 0
	case uint8:
		return v > 0
	case uint16:
		return v > 0
	case uint32:
		return v > 0
	case uint64:
		return v > 0
	default:
		return false
	}
}

// Return true if the value given adheres to cmp.Ordered. Else, false.
func IsOrdered(v any) bool {
	switch v.(type) {
	case int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64, uintptr,
		float32, float64,
		string:
		return true
	default:
		return false
	}
}

// Return true if the value given is a slice. Else, false.
func IsSlice(v any) bool {
	return isOfReflectKind(v, reflect.Slice)
}

// Return true if the value given is a struct. Else, false.
func IsStruct(v any) bool {
	return isOfReflectKind(v, reflect.Struct)
}

// Return true if the value given has a reflect type kind that matches the kind given. Else, false.
func isOfReflectKind(v any, kind reflect.Kind) bool {
	return reflect.TypeOf(v).Kind() == kind
}
