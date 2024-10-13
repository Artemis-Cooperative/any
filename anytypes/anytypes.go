package anytypes

import (
	"reflect"
)

//region Public

// Return true if the values given are equal. Else, false.
// TODO: accomodate structs, slices, arrays, maps and other structures
func Equals(v1 any, v2 any) bool {
	return reflect.DeepEqual(v1, v2)
}

// Panic if the value given is not of the kind given.
// TODO: Test after collections is rolled out
func ExpectKind(v any, expected string) {
	actual := KindString(v)

	if actual != expected {
		panic("expected a " + expected + ", but received a(n) " + actual)
	}
}

// Panic if the value given is not of the type given.
// TODO: Test after collections is rolled out
func ExpectType(v any, expected string) {
	actual := TypeString(v)

	if actual != expected {
		panic("expected a " + expected + ", but received a(n) " + actual)
	}
}

// Return true if the kind of the value given equals the expected kind given. Else, false.
func HasKind(v any, expected string) bool {
	return KindString(v) == expected
}

// Return true if the type of the value given equals the expected type given. Else, false.
func HasType(v any, expected string) bool {
	return TypeString(v) == expected
}

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

// Return the kind of the value given as a string.
func KindString(v any) string {
	return reflect.TypeOf(v).Kind().String()
}

// Return true if the values given are of the same type. Else, false.
func TypeEquals(v1 any, v2 any) bool {
	return reflect.TypeOf(v1) == reflect.TypeOf(v2)
}

func TypeString(v any) string {
	return reflect.TypeOf(v).String()
}

//endregion

//region Private

// Return true if the value given has a reflect type kind that matches the kind given. Else, false.
func isOfReflectKind(v any, kind reflect.Kind) bool {
	return reflect.TypeOf(v).Kind() == kind
}

//endregion
