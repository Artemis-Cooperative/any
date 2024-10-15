package structs

import (
	"reflect"
	"testing"

	anytypes "github.com/Artemis-Cooperative/types/anytypes"
)

var s Structs

type MyStruct struct {
	A string
	B int
}

var data = 0

type Original struct {
	Field1 int
	Field2 string
}

type DifferentFieldTypes struct {
	Field1 string
	Field2 int
}

type DifferentFieldNames struct {
	Field3 int
	Field4 string
}

type MoreFields struct {
	Field1 int
	Field2 string
	Field3 string
}

type LessFields struct {
	Field1 int
}

//region Equals

func TestEquals(t *testing.T) {
	testStruct := Original{
		Field1: 1,
		Field2: "2",
	}
	expected := true
	actual := s.Equals(testStruct, testStruct)

	if expected != actual {
		t.Fatalf("\nExpected:\t%v\nActual:\t\t%v\n", expected, actual)
	}
}

func TestNotEqualsDifferentTypes(t *testing.T) {
	originalStruct := Original{
		Field1: 1,
		Field2: "2",
	}
	differentFieldTypesStruct := DifferentFieldTypes{
		Field1: "1",
		Field2: 2,
	}
	expected := false
	actual := s.Equals(originalStruct, differentFieldTypesStruct)

	if expected != actual {
		t.Fatalf("\nExpected:\t%v\nActual:\t\t%v\n", expected, actual)
	}
}

func TestNotEqualsDifferentNames(t *testing.T) {
	originalStruct := Original{
		Field1: 1,
		Field2: "2",
	}
	differentFieldNamesStruct := DifferentFieldNames{
		Field3: 1,
		Field4: "2",
	}
	expected := false
	actual := s.Equals(originalStruct, differentFieldNamesStruct)

	if expected != actual {
		t.Fatalf("\nExpected:\t%v\nActual:\t\t%v\n", expected, actual)
	}
}

func TestNotEqualsMoreFields(t *testing.T) {
	originalStruct := Original{
		Field1: 1,
		Field2: "2",
	}
	moreFieldssStruct := MoreFields{
		Field1: 1,
		Field2: "2",
		Field3: "3",
	}
	expected := false
	actual := s.Equals(originalStruct, moreFieldssStruct)

	if expected != actual {
		t.Fatalf("\nExpected:\t%v\nActual:\t\t%v\n", expected, actual)
	}
}

func TestNotEqualsLessFields(t *testing.T) {
	originalStruct := Original{
		Field1: 1,
		Field2: "2",
	}
	lessFieldsStruct := LessFields{
		Field1: 1,
	}
	expected := false
	actual := s.Equals(originalStruct, lessFieldsStruct)

	if expected != actual {
		t.Fatalf("\nExpected:\t%v\nActual:\t\t%v\n", expected, actual)
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
		actual = append(actual, s.IsStruct(v))
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
		actual = append(actual, s.IsStruct(v))
	}

	// Test actual values against expected
	for i := 0; i < len(typedValues); i++ {
		if actual[i] != expected[i] {
			valueType := anytypes.TypeString(typedValues[i])
			t.Fatalf("\nExpected:\t%t, %s is not a struct\nActual:\t\t%t, %s IS a struct\n", expected[i], valueType, actual[i], valueType)
		}
	}
}

//endregion
