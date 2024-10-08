package intslices

import (
	"slices"
	"testing"
)

func TestSeq(t *testing.T) {
	expected := []int{1, 2, 3, 4}
	actual := Seq(1, 4)

	if !slices.Equal[[]int](expected, actual) {
		t.Fatalf("\nExpected:\t%v\nActual:\t\t%v\n", expected, actual)
	}
}
