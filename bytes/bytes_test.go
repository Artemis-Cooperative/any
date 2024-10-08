package bytes

import (
	"reflect"
	"testing"
)

// Explode a byte slice, where the delimiter is at the beginning, into byte slices
func TestExplodeWithDelimeterBeginning(t *testing.T) {
	byteSlice := []byte{'a', 'b', 'c'}
	delimiter := []byte{'a', 'b'}
	expected := [][]byte{{'c'}}
	actual := Explode(byteSlice, delimiter)

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("\nExpected:\t" + SlicesToPrintString(expected) + "\nActual:\t\t" + SlicesToPrintString(actual))
	}
}

// Explode a byte slice, where the delimiter is at the end, into byte slices
func TestExplodeWithDelimeterEnding(t *testing.T) {
	byteSlice := []byte{'a', 'b', 'c'}
	delimiter := []byte{'b', 'c'}
	expected := [][]byte{{'a'}}
	actual := Explode(byteSlice, delimiter)

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("\nExpected:\t" + SlicesToPrintString(expected) + "\nActual:\t\t" + SlicesToPrintString(actual))
	}
}

// Explode a byte slice, which is equal to the delimiter, into byte slices
func TestExplodeWithDelimeterOnly(t *testing.T) {
	byteSlice := []byte{'a', 'b'}
	delimiter := []byte{'a', 'b'}
	expected := [][]byte{}
	actual := Explode(byteSlice, delimiter)

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("\nExpected:\t" + SlicesToPrintString(expected) + "\nActual:\t\t" + SlicesToPrintString(actual))
	}
}

// Explode a byte slice into byte slices with a single byte delimiter
func TestExplodeWithSingleByteDelimiter(t *testing.T) {
	byteSlice := []byte{'a', 'b', ',', ' ', 'c'}
	delimiter := []byte{','}
	expected := [][]byte{{'a', 'b'}, {' ', 'c'}}
	actual := Explode(byteSlice, delimiter)

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("\nExpected:\t" + SlicesToPrintString(expected) + "\nActual:\t\t" + SlicesToPrintString(actual))
	}
}

// Explode a byte slice into byte slices with a multi-byte delimiter
func TestExplodeWithMultipleByteDelimiter(t *testing.T) {
	byteSlice := []byte{'a', 'b', ',', ' ', 'c'}
	delimiter := []byte{',', ' '}
	expected := [][]byte{{'a', 'b'}, {'c'}}
	actual := Explode(byteSlice, delimiter)

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("\nExpected:\t" + SlicesToPrintString(expected) + "\nActual:\t\t" + SlicesToPrintString(actual))
	}
}

func TestStrings(t *testing.T) {
	slices := [][]byte{{'a', 'b'}, {' '}, {'c', 'd'}}
	expected := []string{"ab", " ", "cd"}
	actual := Strings(slices)

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("\nExpected:\t%#v\nActual:\t\t%#v\n", expected, actual)
	}
}

func TestOneLine(t *testing.T) {
	slice := []byte{'a', 'b', 'c'}
	expected := "'a', 'b', 'c'"
	actual := OneLine(slice)

	if expected != actual {
		t.Fatalf("\nExpected:\t%#v\nActual:\t\t%#v\n", expected, actual)
	}
}
