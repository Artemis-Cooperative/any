package maps

import (
	"strconv"
	"testing"
)

func TestConcat(t *testing.T) {
	m1 := map[string]string{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	}
	m2 := map[string]string{
		"key4": "value4",
		"key5": "value5",
		"key6": "value6",
	}
	m3 := map[string]string{
		"key7": "value7",
		"key8": "value8",
		"key9": "value9",
	}
	expected := map[string]string{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
		"key4": "value4",
		"key5": "value5",
		"key6": "value6",
		"key7": "value7",
		"key8": "value8",
		"key9": "value9",
	}
	actual := Concat(m1, m2, m3)

	if !Equals(actual, expected) {
		t.Fatalf("\nExpected:\t%v\nActual:\t\t%v\n", expected, actual)
	}
}

func TestMapEquality(t *testing.T) {
	m1 := map[string]string{"key1": "value1", "key2": "value2", "key3": "value3"}
	m2 := map[string]string{"key1": "value1", "key2": "value2", "key3": "value3"}
	expected := true
	actual := Equals(m1, m2)

	if !actual {
		t.Fatalf("\nExpected: '" + strconv.FormatBool(expected) + "', they are equal\nActual: '" + strconv.FormatBool(actual) + "', they are not equal")
	}
}

func TestMapInequalityByLength(t *testing.T) {
	m1 := map[string]string{"key1": "value1", "key2": "value2", "key3": "value3"}
	m2 := map[string]string{"key1": "value1", "key2": "value2"}
	expected := false
	actual := Equals(m1, m2)

	if actual {
		t.Fatalf("\nExpected: '" + strconv.FormatBool(expected) + "', they are inequal\nActual: '" + strconv.FormatBool(actual) + "', they are equal")
	}
}

func TestMapInequalityByKeys(t *testing.T) {
	m1 := map[string]string{"key1": "value1", "key2": "value2", "key3": "value3"}
	m2 := map[string]string{"key1": "value1", "key2": "value2", "key": "value3"}
	expected := false
	actual := Equals(m1, m2)

	if actual {
		t.Fatalf("\nExpected: '" + strconv.FormatBool(expected) + "', they are inequal\nActual: '" + strconv.FormatBool(actual) + "', they are equal")
	}
}

func TestMapInequalityByValues(t *testing.T) {
	m1 := map[string]string{"key1": "value1", "key2": "value2", "key3": "value3"}
	m2 := map[string]string{"key1": "value1", "key2": "value2", "key3": "value"}
	expected := false
	actual := Equals(m1, m2)

	if actual {
		t.Fatalf("\nExpected: '" + strconv.FormatBool(expected) + "', they are inequal\nActual: '" + strconv.FormatBool(actual) + "', they are equal")
	}
}

func TestHasKey(t *testing.T) {
	mapp := map[string]string{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	}
	expected := true
	actual := HasKey[string](mapp, "key2")

	if expected != actual {
		t.Fatalf("\nExpected:\t%t, key was found\nActual:\t\t%t, key was NOT found\n", expected, actual)
	}
}

func TestNotHasKey(t *testing.T) {
	mapp := map[string]string{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	}
	expected := false
	actual := HasKey(mapp, "key4")

	if expected != actual {
		t.Fatalf("\nExpected:\t%t, key was NOT found\nActual:\t\t%t, key was found\n", expected, actual)
	}
}

func TestHasOneKeyType(t *testing.T) {
	mapp := map[string]string{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	}
	expected := true
	actual := HasOneKeyType(mapp)

	if expected != actual {
		t.Fatalf("\nExpected:\t%t, keys had one type\nActual:\t\t%t, keys did NOT have one type\n", expected, actual)
	}
}

func TestNotHasOneKeyType(t *testing.T) {
	mapp := map[any]string{
		"key1": "value1",
		2:      "value2",
		"key3": "value3",
	}
	expected := false
	actual := HasOneKeyType(mapp)

	if expected != actual {
		t.Fatalf("\nExpected:\t%t, keys had one type\nActual:\t\t%t, keys did NOT have one type\n", expected, actual)
	}
}

func TestKeys(t *testing.T) {
	mapp := map[any]string{
		1:      "value1",
		"key2": "value2",
	}
	expected := []any{1, "key2"}
	actual := Keys(mapp)
	notFound := []any{}

	for actualKey := range actual {
		found := false

		for expectedKey := range expected {
			if actualKey == expectedKey {
				found = true
				break
			}
		}

		if !found {
			notFound = append(notFound, actualKey)
		}
	}

	if len(notFound) > 0 {
		for notFoundKey := range notFound {
			t.Fatalf("\nExpected:\t%v to be found\nActual:\t\t%v\n", notFoundKey, actual)
		}
	}
}

func TestSortByKeys(t *testing.T) {
	unsorted := map[int]string{
		3: "value",
		2: "value",
		5: "value",
		1: "value",
		4: "value",
	}
	sorted := map[int]string{
		1: "value",
		2: "value",
		3: "value",
		4: "value",
		5: "value",
	}
	actual := SortByKeys(unsorted)

	if !Equals(actual, sorted) {
		t.Fatalf("\nExpected:\t%v\nActual:\t\t%v\n", sorted, actual)
	}
}

func TestValues(t *testing.T) {
	mapp := map[string]any{
		"key1": 1,
		"key2": "value2",
	}
	expected := []any{1, "value2"}
	actual := Values(mapp)
	notFound := []any{}

	for _, actualValue := range actual {
		found := false

		for _, expectedValue := range expected {
			if actualValue == expectedValue {
				found = true
				break
			}
		}

		if !found {
			notFound = append(notFound, actualValue)
		}
	}

	if len(notFound) > 0 {
		for notFoundValue := range notFound {
			t.Fatalf("\nExpected:\t%v to be found\nActual:\t\t%v\n", notFoundValue, actual)
		}
	}
}
