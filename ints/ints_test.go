package ints

import (
	"testing"
)

func TestIsNumWithNumber(t *testing.T) {
	expected := true
	actual := IsNum("0")

	if expected != actual {
		t.Fatalf("\nExpected\t%v\nActual:\t\t%v\n", expected, actual)
	}
}

func TestIsNumWithLetter(t *testing.T) {
	expected := false
	actual := IsNum("a")

	if expected != actual {
		t.Fatalf("\nExpected\t%v\nActual:\t\t%v\n", expected, actual)
	}
}

func TestIsNumWithBlank(t *testing.T) {
	expected := false
	actual := IsNum("")

	if expected != actual {
		t.Fatalf("\nExpected\t%v\nActual:\t\t%v\n", expected, actual)
	}
}

func TestIsNumWithSpace(t *testing.T) {
	expected := false
	actual := IsNum(" ")

	if expected != actual {
		t.Fatalf("\nExpected\t%v\nActual:\t\t%v\n", expected, actual)
	}
}

func TestIsNumWithNumbersAndSpace(t *testing.T) {
	expected := false
	actual := IsNum("2 2")

	if expected != actual {
		t.Fatalf("\nExpected\t%v\nActual:\t\t%v\n", expected, actual)
	}
}

func TestIsNumWithCommaSeparatedNumbers(t *testing.T) {
	expected := false
	actual := IsNum("2,200") // TODO: Maybe this should be convertable

	if expected != actual {
		t.Fatalf("\nExpected\t%v\nActual:\t\t%v\n", expected, actual)
	}
}
