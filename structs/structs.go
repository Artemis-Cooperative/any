package structs

import (
	"reflect"

	anytypes "github.com/Artemis-Cooperative/types/anytypes"
)

type Structs struct{}

// Return true if the two structs given are equal. Else, false.
func (s Structs) Equals(s1 interface{}, s2 interface{}) bool {
	r1 := reflect.ValueOf(s1)
	r2 := reflect.ValueOf(s2)
	t1 := reflect.TypeOf(s1)
	t2 := reflect.TypeOf(s2)

	// Verify are structs and have equal field count
	if !s.IsStruct(s1) || !s.IsStruct(s2) {
		return false
	} else if r1.NumField() != r2.NumField() {
		return false
	}

	// Verify field names values and types
	for i := 0; i < r1.NumField(); i++ {
		i1 := r1.Field(i).Interface()
		i2 := r2.Field(i).Interface()
		n1 := t1.Field(i).Name
		n2 := t2.Field(i).Name

		if i1 != i2 || n1 != n2 {
			return false
		}
	}

	return true
}

// Return true if the value given is a struct. Else, false.
func (s Structs) IsStruct(v any) bool {
	return anytypes.HasKind(v, "struct")
}
