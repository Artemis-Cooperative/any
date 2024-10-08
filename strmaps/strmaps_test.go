package strmaps

import (
	"testing"
)

// Convert string map to string
func TestPretty(t *testing.T) {
	mapp := map[string]string{"key1": "value1", "key2": "value2", "key3": "value3"}
	expected := "key1: value1\nkey2: value2\nkey3: value3\n"
	actual := Pretty(mapp)

	if expected != actual {
		t.Fatalf("\nExpected: '" + expected + "'\nActual: '" + actual + "'")
	}
}
