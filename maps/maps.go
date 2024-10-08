package maps

import (
	"cmp"
	"reflect"
	"slices"
)

// Concatenate the maps given
func Concat[K comparable, V any](maps ...map[K]V) map[K]V {
	concatenated := map[K]V{}

	for _, m := range maps {
		for key, value := range m {
			concatenated[key] = value
		}
	}

	return concatenated
}

// Return true if the maps given are equal, else false. O(nlogn)
func Equals[K cmp.Ordered, V cmp.Ordered](m1 map[K]V, m2 map[K]V) bool {
	// Confirm equal lengths
	if len(m1) != len(m2) {
		return false
	}

	// Get sorted values for each map, O(nlogn)
	keys1 := SortedKeys(m1)
	keys2 := SortedKeys(m2)

	// Compare keys for equality, O(n)
	for i := 0; i < len(keys1); i++ {
		if keys1[i] != keys2[i] {
			return false
		}
	}

	// Get sorted values for each map, O(nlogn)
	values1 := SortedValues(m1)
	values2 := SortedValues(m2)

	// Compare values for equality, O(n)
	for i := 0; i < len(values1); i++ {
		if values1[i] != values2[i] {
			return false
		}
	}

	return true
}

// Return true if the key given is in the map given. Else, false.
func HasKey[K comparable, V any](m map[K]V, key K) bool {
	for mKey := range m {
		if key == mKey {
			return true
		}
	}

	return false
}

func HasOneKeyType[K comparable, V any](m map[K]V) bool {
	if len(m) == 0 {
		return true
	}
	var keyType reflect.Type

	for key := range m {
		if keyType == nil {
			keyType = reflect.TypeOf(key)
		} else if keyType != reflect.TypeOf(key) {
			return false
		}
	}

	return true
}

// Return keys of the map given
func Keys[K comparable, V any](m map[K]V) []K {
	var keys []K

	for key := range m {
		keys = append(keys, key)
	}

	return keys
}

// Sort the map given by its keys
func SortByKeys[K cmp.Ordered, V any](m map[K]V) map[K]V {
	keys := SortedKeys(m)
	sorted := map[K]V{}

	for _, key := range keys {
		sorted[key] = m[key]
	}

	return sorted
}

// Return a sorted array of keys from the map given
func SortedKeys[K cmp.Ordered, V any](m map[K]V) []K {
	keys := Keys(m)

	slices.Sort(keys)

	return keys
}

// Return a sorted array of values from the map given
func SortedValues[K comparable, V cmp.Ordered](m map[K]V) []V {
	values := Values(m)

	slices.Sort(values)

	return values
}

// Return values of the map given
func Values[K comparable, V any](m map[K]V) []V {
	var values []V

	for _, value := range m {
		values = append(values, value)
	}

	return values
}
