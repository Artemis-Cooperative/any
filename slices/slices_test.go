package slices

import (
	"reflect"
	"testing"

	strslices "github.com/Artemis-Cooperative/types/strslices"
)

func TestConcat(t *testing.T) {
	slices := [][]string{{"a", "b"}, {"c", "d"}, {"e", "f"}}
	expected := []string{"a", "b", "c", "d", "e", "f"}
	actual := Concat(slices...)

	if !strslices.Equals(expected, actual) {
		t.Fatalf("\nExpected:\t%#v\nActual:\t\t%#v", expected, actual)
	}
}

func TestRemoveBeginning(t *testing.T) {
	slice := []string{"a", "b", "c", "d"}
	expected := []string{"b", "c", "d"}
	actual := Remove(slice, 0)

	if !strslices.Equals(expected, actual) {
		t.Fatalf("\nExpected:\t%#v\nActual:\t\t%#v", expected, actual)
	}
}

func TestRemoveMiddle(t *testing.T) {
	slice := []string{"a", "b", "c", "d"}
	expected := []string{"a", "b", "d"}
	actual := Remove(slice, 2)

	if !strslices.Equals(expected, actual) {
		t.Fatalf("\nExpected:\t%#v\nActual:\t\t%#v", expected, actual)
	}
}

func TestRemoveEnd(t *testing.T) {
	slice := []string{"a", "b", "c", "d"}
	expected := []string{"a", "b", "c"}
	actual := Remove(slice, (len(slice) - 1))

	if !strslices.Equals(expected, actual) {
		t.Fatalf("\nExpected:\t%#v\nActual:\t\t%#v", expected, actual)
	}
}

func TestRemoveEmpty(t *testing.T) {
	slices := [][]string{{}, {"a", "b"}, {}, {"c"}, {}}
	expected := [][]string{{"a", "b"}, {"c"}}
	actual := RemoveEmpty(slices)

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("\nExpected:\t%#v\nActual:\t\t%#v", expected, actual)
	}
}

func TestPrepend(t *testing.T) {
	slice := []string{"b", "c"}
	item := "a"
	expected := []string{"a", "b", "c"}
	actual := Prepend(slice, item)

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("\nExpected:\t%#v\nActual:\t\t%#v", expected, actual)
	}
}

func TestShiftWithEmpty(t *testing.T) {
	slice := []string{}
	expected := []string{}
	actual := Shift[string](slice)

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("\nExpected:\t%#v\nActual:\t\t%#v", expected, actual)
	}
}

func TestShiftWithOne(t *testing.T) {
	slice := []string{"item"}
	expected := []string{}
	actual := Shift[string](slice)

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("\nExpected:\t%#v\nActual:\t\t%#v", expected, actual)
	}
}

func TestShiftWithMany(t *testing.T) {
	slice := []string{"item1", "item2"}
	expected := []string{"item2"}
	actual := Shift[string](slice)

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("\nExpected:\t%#v\nActual:\t\t%#v", expected, actual)
	}
}

func TestIndexes(t *testing.T) {
	slice := []string{"item1", "item2", "item3", "item4", "item5"}
	expected := []int{0, 1, 2, 3, 4}
	actual := Indexes(slice)

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("\nExpected:\t%#v\nActual:\t\t%#v\n", expected, actual)
	}
}
