package structs

import (
	"testing"
)

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

func TestEquals(t *testing.T) {
	testStruct := Original{
		Field1: 1,
		Field2: "2",
	}
	expected := true
	actual := Equals(testStruct, testStruct)

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
	actual := Equals(originalStruct, differentFieldTypesStruct)

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
	actual := Equals(originalStruct, differentFieldNamesStruct)

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
	actual := Equals(originalStruct, moreFieldssStruct)

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
	actual := Equals(originalStruct, lessFieldsStruct)

	if expected != actual {
		t.Fatalf("\nExpected:\t%v\nActual:\t\t%v\n", expected, actual)
	}
}
