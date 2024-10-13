package slices

import (
	"reflect"
	"testing"

	testkit "github.com/Artemis-Cooperative/testkit/types"
	strslices "github.com/Artemis-Cooperative/types/strslices"
)

type MyStruct struct {
	A string
	B int
}

var data = 0

//region Concat

func TestConcat(t *testing.T) {
	slices := [][]string{{"a", "b"}, {"c", "d"}, {"e", "f"}}
	expected := []string{"a", "b", "c", "d", "e", "f"}
	actual := Concat(slices...)

	if !strslices.Equals(expected, actual) {
		t.Fatalf("\nExpected:\t%#v\nActual:\t\t%#v", expected, actual)
	}
}

//endregion

//region ExpectSlice

func TestExpectSlicePanicsWithoutSlice(t *testing.T) {
	typedValues := testkit.AllTypeInstances([]string{"slice"})
	expected := testkit.AllKinds([]string{"slice"})

	// Prepend message to expected values
	for i := 0; i < len(expected); i++ {
		expected[i] = "expected a slice, but received a(n) " + expected[i]
	}

	// Test actual values against expected
	for i := 0; i < len(typedValues); i++ {
		defer func() {
			if actual := recover(); actual != nil {
				if actual != expected[i] {
					t.Errorf("\nExpected:\t%#v\nActual:\t\t%#v", expected, actual)
				}
			} else {
				t.Errorf("expected panic, but none occurred")
			}
		}()

		ExpectSlice(typedValues[i])
	}
}

func TestExpectSliceDoesNotPanicWithSlice(t *testing.T) {
	typedValues := []any{
		[]any{},         // slice
		[]bool{},        // bool slice
		[]int{},         // int slice
		[]int8{},        // int8 slice
		[]int16{},       // int16 slice
		[]int32{},       // int32 slice
		[]int64{},       // int64 slice
		[]uint{},        // uint slice
		[]uint8{},       // uint8 slice
		[]uint16{},      // uint16 slice
		[]uint32{},      // uint32 slice
		[]uint64{},      // uint64 slice
		[]float32{},     // float32 slice
		[]float64{},     // float64 slice
		[]complex64{},   // complex64 slice
		[]complex128{},  // complex128 slice
		[]MyStruct{},    // struct slice
		[]map[any]any{}, // map slice
		[][1]any{},      // array slice
		[]func(){},      // function slice
		[]chan int{},    // channel slice
	}

	// Verify no panic is thrown
	for i := 0; i < len(typedValues); i++ {
		defer func() {
			if actual := recover(); actual != nil {
				t.Errorf("\nExpected:\tno panic\nActual:\t\t%#v", actual)
			}
		}()

		ExpectSlice(typedValues[i])
	}
}

//endregion

//region ItemType

func TestItemType(t *testing.T) {
	slices := testkit.SliceInstances()
	expected := testkit.ElementKinds()

	for i := 0; i < len(slices); i++ {
		actual := ItemType(slices[i]).String()

		if actual != expected[i] {
			t.Fatalf("\nExpected:\t%#v\nActual:\t\t%#v", expected[i], actual)
		}
	}
}

func TestItemTypePanicsWithoutASlice(t *testing.T) {
	expected := "expected a slice, but received a(n) int"
	defer func() {
		if actual := recover(); actual != nil {
			if actual != expected {
				t.Errorf("\nExpected:\t%#v\nActual:\t\t%#v", expected, actual)
			}
		} else {
			t.Errorf("expected panic, but none occurred")
		}
	}()

	ItemType(0)
}

//endregion

//region Indexes

func TestIndexes(t *testing.T) {
	slice := []string{"item1", "item2", "item3", "item4", "item5"}
	expected := []int{0, 1, 2, 3, 4}
	actual := Indexes(slice)

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("\nExpected:\t%#v\nActual:\t\t%#v\n", expected, actual)
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
		&data,                    // pointer
		MyStruct{},               // struct
		map[any]any{},            // map
		[1]any{"array"},          // array
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

//region Prepend

func TestPrepend(t *testing.T) {
	slice := []string{"b", "c"}
	item := "a"
	expected := []string{"a", "b", "c"}
	actual := Prepend(slice, item)

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("\nExpected:\t%#v\nActual:\t\t%#v", expected, actual)
	}
}

//endregion

//region Remove

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

//endregion

//region RemoveEmpty

func TestRemoveEmpty(t *testing.T) {
	slices := [][]string{{}, {"a", "b"}, {}, {"c"}, {}}
	expected := [][]string{{"a", "b"}, {"c"}}
	actual := RemoveEmpty(slices)

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("\nExpected:\t%#v\nActual:\t\t%#v", expected, actual)
	}
}

//endregion

//region Shift

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

//endregion
