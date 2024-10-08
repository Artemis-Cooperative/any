package str

import (
	"testing"
)

func TestIsAlphaNum(t *testing.T) {
	alphaNum := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	expected := true
	actual := IsAlphaNum(alphaNum)

	if expected != actual {
		t.Fatalf("\nExpected:\t%t, the variable is alpha numberic\nActual:\t\t%t, the variable is not alpha numberic\n", expected, actual)
	}
}

func TestIsNotAlphaNum(t *testing.T) {
	alphaNum := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789*"
	expected := false
	actual := IsAlphaNum(alphaNum)

	if expected != actual {
		t.Fatalf("\nExpected:\t%t, the variable is not alpha numberic\nActual:\t\t%t, the variable is alpha numberic\n", expected, actual)
	}
}

func TestTrimSides(t *testing.T) {
	str := ">['value']"
	expected := "value"
	actual := TrimSides(str, 3, 2)

	if expected != actual {
		t.Fatalf("\nExpected:\t%s\nActual:\t\t%s", expected, actual)
	}
}

func TestCenter(t *testing.T) {
	str := "center me"
	width := 20
	expected := "     center me      " // Left padding is shorter than right
	actual := Center(str, width)

	if expected != actual {
		t.Fatalf("\nExpected:\t%s\nActual:\t\t%s", expected, actual)
	}
}

// Center text with custom multi-character left fill and right fill
func TestCenterCustom(t *testing.T) {
	str := "center me"
	width := 20
	expected := "[[[[[[center me]]]]]]" // Width is actually 21, fill is too large
	actual := CenterCustom(str, width, "[[", "]]")

	if expected != actual {
		t.Fatalf("\nExpected:\t%s\nActual:\t\t%s", expected, actual)
	}
}

func TestPadLeft(t *testing.T) {
	padCount := 5
	expected := "pad->     "
	actual := "pad->"

	PadLeft(&actual, &padCount)

	if expected != actual {
		t.Fatalf("\nExpected:\t'%s'\nActual:\t\t'%s'\n", expected, actual)
	}
}

func TestLeft(t *testing.T) {
	width := 10
	expected := "pad->     "
	actual := "pad->"

	Left(&actual, &width)

	if expected != actual {
		t.Fatalf("\nExpected:\t'%s'\nActual:\t\t'%s'\n", expected, actual)
	}
}
