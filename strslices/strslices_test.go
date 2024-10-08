package strslices

import (
	"testing"
)

// Convert a slice to a csv string
func TestToCSV(t *testing.T) {
	slice := []string{"value1", "value2", "value3"}
	expected := "value1,value2,value3"
	actual := ToCSV(slice)

	if expected != actual {
		t.Fatalf("\nExpected: " + expected + "\nActual: " + actual)
	}
}

// Convert a csv string to a slice
func TestCSVToSlice(t *testing.T) {
	csv := "value1,value2,value3"
	expected := []string{"value1", "value2", "value3"}
	actual := CSVToSlice(csv)

	if !Equals(expected, actual) {
		t.Fatalf("\nExpected: " + Pretty(expected) + "\nActual: " + Pretty(actual))
	}
}

func TestCSVWithSpacesToSlice(t *testing.T) {
	csv := "value1.0 value1.1, value2.0 value2.1, value3.0 value3.1"
	expected := []string{"value1.0 value1.1", "value2.0 value2.1", "value3.0 value3.1"}
	actual := CSVToSlice(csv)

	if !Equals(expected, actual) {
		t.Fatalf("\nExpected: " + Pretty(expected) + "\nActual: " + Pretty(actual))
	}
}

// Convert string slice to a print string
func TestSliceToPrintString(t *testing.T) {
	a := []string{"value1", "value2", "value3"}
	expected := "[\n    value1\n    value2\n    value3\n]\n"
	actual := Pretty(a)

	if expected != actual {
		t.Fatalf("\nExpected: " + expected + "\nActual: " + actual)
	}
}

// Convert a print string to a string slice
func TestPrintStringToSlice(t *testing.T) {
	s := "[\n    value1\n    value2\n    value3\n]\n"
	expected := []string{"value1", "value2", "value3"}
	actual := PrettyToSlice(s)

	if !Equals(expected, actual) {
		t.Fatalf("\nExpected: " + Pretty(expected) + "\nActual: " + Pretty(actual))
	}
}

func TestMaxLength(t *testing.T) {
	slice := []string{
		"1",
		"2 ",
		"3  ",
		"2 ",
	}
	expected := 3
	actual := MaxLength(slice)

	if expected != actual {
		t.Fatalf("\nExpected:\t%d\nActual:\t\t%d", expected, actual)
	}
}

func TestRepeat(t *testing.T) {
	slice := []string{"value1", "value2", "value3"}
	expected := []string{"value1", "value2", "value3",
		"value1", "value2", "value3",
		"value1", "value2", "value3"}
	actual := Repeat(slice, 3)

	if !Equals(expected, actual) {
		t.Fatalf("\nExpected:\t%#v\nActual:\t\t%#v\n", expected, actual)
	}
}

func TestRepeatZero(t *testing.T) {
	slice := []string{"value1", "value2", "value3"}
	expected := []string{}
	actual := Repeat(slice, 0)

	if !Equals(expected, actual) {
		t.Fatalf("\nExpected:\t%#v\nActual:\t\t%#v\n", expected, actual)
	}
}
