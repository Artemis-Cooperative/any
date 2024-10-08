package anytypes

import (
	"reflect"
	"testing"
)

type MyStruct struct {
	A int
	B string
}

var data = 0

//region IsArray

func TestIsArray(t *testing.T) {
	typedValues := []any{
		[1]any{"array"},             // any array
		[1]string{"array"},          // string array
		[1]int{1},                   // int array
		[1]int8{2},                  // int8 array
		[1]int16{3},                 // int16 array
		[1]int32{4},                 // int32 array
		[1]int64{5},                 // int64 array
		[1]uint{6},                  // uint array
		[1]uint8{7},                 // uint8 array
		[1]uint16{8},                // uint16 array
		[1]uint32{9},                // uint32 array
		[1]uint64{10},               // uint64 array
		[1]float32{11.1},            // float32 array
		[1]float64{12.2},            // float64 array
		[1]complex64{13.3 + 14.4i},  // complex64 array
		[1]complex128{15.5 + 16.6i}, // complex128 array
		[1]MyStruct{},               // struct array
		[2]any{"array", "array"},    // multi-element array
	}
	expected := []bool{}
	actual := []bool{}

	// Set expected and actual values
	for _, v := range typedValues {
		expected = append(expected, true)
		actual = append(actual, IsArray(v))
	}

	// Test actual values against expected
	for i := 0; i < len(typedValues); i++ {
		if actual[i] != expected[i] {
			valueType := reflect.TypeOf(typedValues[i])
			t.Fatalf("\nExpected:\t%t, %s is an array\nActual:\t\t%t, %s is NOT an array\n", expected[i], valueType, actual[i], valueType)
		}
	}
}

func TestIsNotArray(t *testing.T) {
	typedValues := []any{
		true,                     // bool
		int(1),                   // int
		int8(2),                  // int8
		int16(3),                 // int16
		int32(4),                 // int32
		int64(5),                 // int64
		uint(6),                  // uint
		uint8(7),                 // uint8
		uint16(8),                // uint16
		uint32(9),                // uint32
		uint64(10),               // uint64
		float32(11.1),            // float32
		float64(12.2),            // float64
		complex64(13.3 + 14.4i),  // complex64
		complex128(15.5 + 16.6i), // complex128
		"hello",                  // string
		MyStruct{},               // struct
		map[any]any{},            // map
		&data,                    // pointer
		func() {},                // function
		make(chan int),           // channel
	}
	expected := []bool{}
	actual := []bool{}

	// Set expected and actual values
	for _, v := range typedValues {
		expected = append(expected, false)
		actual = append(actual, IsArray(v))
	}

	// Test actual values against expected
	for i := 0; i < len(typedValues); i++ {
		if actual[i] != expected[i] {
			valueType := reflect.TypeOf(typedValues[i])
			t.Fatalf("\nExpected:\t%t, %s is not an array\nActual:\t\t%t, %s IS an array\n", expected[i], valueType, actual[i], valueType)
		}
	}
}

//endregion

//region IsBasic

func TestIsBasic(t *testing.T) {
	typedValues := []any{
		true,                     // bool
		int(1),                   // int
		int8(2),                  // int8
		int16(3),                 // int16
		int32(4),                 // int32
		int64(5),                 // int64
		uint(6),                  // uint
		uint8(7),                 // uint8
		uint16(8),                // uint16
		uint32(9),                // uint32
		uint64(10),               // uint64
		float32(11.1),            // float32
		float64(12.2),            // float64
		complex64(13.3 + 14.4i),  // complex64
		complex128(15.5 + 16.6i), // complex128
		"hello",                  // string
	}
	expected := []bool{}
	actual := []bool{}

	// Set expected and actual values
	for _, v := range typedValues {
		expected = append(expected, true)
		actual = append(actual, IsBasic(v))
	}

	// Test actual values against expected
	for i := 0; i < len(typedValues); i++ {
		if actual[i] != expected[i] {
			valueType := reflect.TypeOf(typedValues[i])
			t.Fatalf("\nExpected:\t%t, %s is a basic\nActual:\t\t%t, %s is NOT a basic\n", expected[i], valueType, actual[i], valueType)
		}
	}
}

func TestIsNotBasic(t *testing.T) {
	typedValues := []any{
		MyStruct{},     // struct
		[]any{},        // slice
		[1]any{},       // array
		map[any]any{},  //map
		&data,          // pointer
		func() {},      // function
		make(chan int), // channel
	}
	expected := []bool{}
	actual := []bool{}

	// Set expected and actual values
	for _, v := range typedValues {
		expected = append(expected, false)
		actual = append(actual, IsBasic(v))
	}

	// Test actual values against expected
	for i := 0; i < len(typedValues); i++ {
		if actual[i] != expected[i] {
			valueType := reflect.TypeOf(typedValues[i])
			t.Fatalf("\nExpected:\t%t, %s is a basic\nActual:\t\t%t, %s is NOT a basic\n", expected[i], valueType, actual[i], valueType)
		}
	}
}

//endregion

//region IsComparable

func TestIsComparable(t *testing.T) {
	typedValues := []any{
		true,                     // bool
		int(1),                   // int
		int8(2),                  // int8
		int16(3),                 // int16
		int32(4),                 // int32
		int64(5),                 // int64
		uint(6),                  // uint
		uint8(7),                 // uint8
		uint16(8),                // uint16
		uint32(9),                // uint32
		uint64(10),               // uint64
		float32(11.1),            // float32
		float64(12.2),            // float64
		complex64(13.3 + 14.4i),  // complex64
		complex128(15.5 + 16.6i), // complex128
		"hello",                  // string
		MyStruct{},               // struct
		[1]any{"array"},          // array
		&data,                    // pointer
		make(chan int),           // channel
	}
	expected := []bool{}
	actual := []bool{}

	// Set expected and actual values
	for _, v := range typedValues {
		expected = append(expected, true)
		actual = append(actual, IsComparable(v))
	}

	// Test actual values against expected
	for i := 0; i < len(typedValues); i++ {
		if actual[i] != expected[i] {
			valueType := reflect.TypeOf(typedValues[i])
			t.Fatalf("\nExpected:\t%t, %s is comparable\nActual:\t\t%t, %s is NOT comparable\n", expected[i], valueType, actual[i], valueType)
		}
	}
}

func TestIsNotComparable(t *testing.T) {
	typedValues := []any{
		[]any{"slice"}, // slice
		map[any]any{},  // map
		func() {},      // function
	}
	expected := []bool{}
	actual := []bool{}

	// Set expected and actual values
	for _, v := range typedValues {
		expected = append(expected, false)
		actual = append(actual, IsComparable(v))
	}

	// Test actual values against expected
	for i := 0; i < len(typedValues); i++ {
		if actual[i] != expected[i] {
			valueType := reflect.TypeOf(typedValues[i])
			t.Fatalf("\nExpected:\t%t, %s is not comparable\nActual:\t\t%t, %s IS comparable\n", expected[i], valueType, actual[i], valueType)
		}
	}
}

//endregion

//region IsMap

func TestIsMap(t *testing.T) {
	typedValues := []any{
		map[any]any{},       // any, any map
		map[int]any{},       // int, any map
		map[string]any{},    // string, any map
		map[any]int{},       // any, int map
		map[int]int{},       // int, int map
		map[string]int{},    // string, int map
		map[any]string{},    // any, string map
		map[int]string{},    // int, string map
		map[string]string{}, // string, string map
	}
	expected := []bool{}
	actual := []bool{}

	// Set expected and actual values
	for _, v := range typedValues {
		expected = append(expected, true)
		actual = append(actual, IsMap(v))
	}

	// Test actual values against expected
	for i := 0; i < len(typedValues); i++ {
		if actual[i] != expected[i] {
			valueType := reflect.TypeOf(typedValues[i])
			t.Fatalf("\nExpected:\t%t, %s is a map\nActual:\t\t%t, %s is NOT a map\n", expected[i], valueType, actual[i], valueType)
		}
	}
}

func TestIsNotMap(t *testing.T) {
	typedValues := []any{
		true,                     // bool
		int(1),                   // int
		int8(2),                  // int8
		int16(3),                 // int16
		int32(4),                 // int32
		int64(5),                 // int64
		uint(6),                  // uint
		uint8(7),                 // uint8
		uint16(8),                // uint16
		uint32(9),                // uint32
		uint64(10),               // uint64
		float32(11.1),            // float32
		float64(12.2),            // float64
		complex64(13.3 + 14.4i),  // complex64
		complex128(15.5 + 16.6i), // complex128
		"hello",                  // string
		MyStruct{},               // struct
		[]any{},                  // slice
		[1]any{"array"},          // array
		&data,                    // pointer
		func() {},                // function
		make(chan int),           // channel
	}
	expected := []bool{}
	actual := []bool{}

	// Set expected and actual values
	for _, v := range typedValues {
		expected = append(expected, false)
		actual = append(actual, IsMap(v))
	}

	// Test actual values against expected
	for i := 0; i < len(typedValues); i++ {
		if actual[i] != expected[i] {
			valueType := reflect.TypeOf(typedValues[i])
			t.Fatalf("\nExpected:\t%t, %s is not a map\nActual:\t\t%t, %s IS a map\n", expected[i], valueType, actual[i], valueType)
		}
	}
}

//endregion

//region IsNumKey

func TestIsNumKey(t *testing.T) {
	typedValues := []any{
		int(1),     // int
		int8(2),    // int8
		int16(3),   // int16
		int32(4),   // int32
		int64(5),   // int64
		uint(6),    // uint
		uint8(7),   // uint8
		uint16(8),  // uint16
		uint32(9),  // uint32
		uint64(10), // uint64
	}
	expected := []bool{}
	actual := []bool{}

	// Set expected and actual values
	for _, v := range typedValues {
		expected = append(expected, true)
		actual = append(actual, IsNumKey(v))
	}

	// Test actual values against expected
	for i := 0; i < len(typedValues); i++ {
		if actual[i] != expected[i] {
			valueType := reflect.TypeOf(typedValues[i])
			t.Fatalf("\nExpected:\t%t, %s is a number key\nActual:\t\t%t, %s is NOT a number key\n", expected[i], valueType, actual[i], valueType)
		}
	}
}

func TestIsNotNumKey(t *testing.T) {
	typedValues := []any{
		true,                     // bool
		float32(11.1),            // float32
		float64(12.2),            // float64
		complex64(13.3 + 14.4i),  // complex64
		complex128(15.5 + 16.6i), // complex128
		"hello",                  // string
		MyStruct{},               // struct
		[]any{},                  // slice
		[1]any{"array"},          // array
		&data,                    // pointer
		func() {},                // function
		make(chan int),           // channel
	}
	expected := []bool{}
	actual := []bool{}

	// Set expected and actual values
	for _, v := range typedValues {
		expected = append(expected, false)
		actual = append(actual, IsNumKey(v))
	}

	// Test actual values against expected
	for i := 0; i < len(typedValues); i++ {
		if actual[i] != expected[i] {
			valueType := reflect.TypeOf(typedValues[i])
			t.Fatalf("\nExpected:\t%t, %s is not a number key\nActual:\t\t%t, %s IS a number key\n", expected[i], valueType, actual[i], valueType)
		}
	}
}

//endregion

//region IsSlice

func TestIsSlice(t *testing.T) {
	typedValues := []any{
		[]any{},      // any slice
		[]int{},      // int slice
		[]string{},   // string slice
		[]MyStruct{}, // struct slice
	}
	expected := []bool{}
	actual := []bool{}

	// Set expected and actual values
	for _, v := range typedValues {
		expected = append(expected, true)
		actual = append(actual, IsSlice(v))
	}

	// Test actual values against expected
	for i := 0; i < len(typedValues); i++ {
		if actual[i] != expected[i] {
			valueType := reflect.TypeOf(typedValues[i])
			t.Fatalf("\nExpected:\t%t, %s is a slice\nActual:\t\t%t, %s is NOT a slice\n", expected[i], valueType, actual[i], valueType)
		}
	}
}

func TestIsNotSlice(t *testing.T) {
	typedValues := []any{
		true,                     // bool
		int(1),                   // int
		int8(2),                  // int8
		int16(3),                 // int16
		int32(4),                 // int32
		int64(5),                 // int64
		uint(6),                  // uint
		uint8(7),                 // uint8
		uint16(8),                // uint16
		uint32(9),                // uint32
		uint64(10),               // uint64
		float32(11.1),            // float32
		float64(12.2),            // float64
		complex64(13.3 + 14.4i),  // complex64
		complex128(15.5 + 16.6i), // complex128
		"hello",                  // string
		MyStruct{},               // struct
		map[any]any{},            // map
		[1]any{"array"},          // array
		&data,                    // pointer
		func() {},                // function
		make(chan int),           // channel
	}
	expected := []bool{}
	actual := []bool{}

	// Set expected and actual values
	for _, v := range typedValues {
		expected = append(expected, false)
		actual = append(actual, IsSlice(v))
	}

	// Test actual values against expected
	for i := 0; i < len(typedValues); i++ {
		if actual[i] != expected[i] {
			valueType := reflect.TypeOf(typedValues[i])
			t.Fatalf("\nExpected:\t%t, %s is not a slice\nActual:\t\t%t, %s IS a slice\n", expected[i], valueType, actual[i], valueType)
		}
	}
}

//endregion

//region IsStruct

func TestIsStruct(t *testing.T) {
	typedValues := []any{
		MyStruct{}, // struct
	}
	expected := []bool{}
	actual := []bool{}

	// Set expected and actual values
	for _, v := range typedValues {
		expected = append(expected, true)
		actual = append(actual, IsStruct(v))
	}

	// Test actual values against expected
	for i := 0; i < len(typedValues); i++ {
		if actual[i] != expected[i] {
			valueType := reflect.TypeOf(typedValues[i])
			t.Fatalf("\nExpected:\t%t, %s is a struct\nActual:\t\t%t, %s is NOT a struct\n", expected[i], valueType, actual[i], valueType)
		}
	}
}

func TestIsNotStruct(t *testing.T) {
	typedValues := []any{
		true,                     // bool
		int(1),                   // int
		int8(2),                  // int8
		int16(3),                 // int16
		int32(4),                 // int32
		int64(5),                 // int64
		uint(6),                  // uint
		uint8(7),                 // uint8
		uint16(8),                // uint16
		uint32(9),                // uint32
		uint64(10),               // uint64
		float32(11.1),            // float32
		float64(12.2),            // float64
		complex64(13.3 + 14.4i),  // complex64
		complex128(15.5 + 16.6i), // complex128
		"hello",                  // string
		[]any{},                  // slice
		[1]any{"array"},          // array
		map[any]any{},            // map
		&data,                    // pointer
		func() {},                // function
		make(chan int),           // channel
	}
	expected := []bool{}
	actual := []bool{}

	// Set expected and actual values
	for _, v := range typedValues {
		expected = append(expected, false)
		actual = append(actual, IsStruct(v))
	}

	// Test actual values against expected
	for i := 0; i < len(typedValues); i++ {
		if actual[i] != expected[i] {
			valueType := reflect.TypeOf(typedValues[i])
			t.Fatalf("\nExpected:\t%t, %s is not a struct\nActual:\t\t%t, %s IS a struct\n", expected[i], valueType, actual[i], valueType)
		}
	}
}

//endregion

//region IsOrdered

func TestIsOrdered(t *testing.T) {
	typedValues := []any{
		int(1),        // int
		int8(2),       // int8
		int16(3),      // int16
		int32(4),      // int32
		int64(5),      // int64
		uint(6),       // uint
		uint8(7),      // uint8
		uint16(8),     // uint16
		uint32(9),     // uint32
		uint64(10),    // uint64
		float32(11.1), // float32
		float64(12.2), // float64
		"hello",       // string
	}
	expected := []bool{}
	actual := []bool{}

	// Set expected and actual values
	for _, v := range typedValues {
		expected = append(expected, true)
		actual = append(actual, IsOrdered(v))
	}

	// Test actual values against expected
	for i := 0; i < len(typedValues); i++ {
		if actual[i] != expected[i] {
			valueType := reflect.TypeOf(typedValues[i])
			t.Fatalf("\nExpected:\t%t, %s is ordered\nActual:\t\t%t, %s is NOT ordered\n", expected[i], valueType, actual[i], valueType)
		}
	}
}

func TestIsNotOrdered(t *testing.T) {
	typedValues := []any{
		complex64(13.3 + 14.4i),  // complex64
		complex128(15.5 + 16.6i), // complex128
		MyStruct{},               // struct
		[]any{},                  // slice
		[1]any{"array"},          // array
		&data,                    // pointer
		func() {},                // function
		make(chan int),           // channel
	}
	expected := []bool{}
	actual := []bool{}

	// Set expected and actual values
	for _, v := range typedValues {
		expected = append(expected, false)
		actual = append(actual, IsOrdered(v))
	}

	// Test actual values against expected
	for i := 0; i < len(typedValues); i++ {
		if actual[i] != expected[i] {
			valueType := reflect.TypeOf(typedValues[i])
			t.Fatalf("\nExpected:\t%t, %s is not ordered\nActual:\t\t%t, %s IS ordered\n", expected[i], valueType, actual[i], valueType)
		}
	}
}

//endregion
