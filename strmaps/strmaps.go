package strmaps

import (
	"fmt"
	"sort"

	"github.com/Artemis-Cooperative/types/maps"
)

// Return a string representation of a string map
func Pretty(m map[string]string) string {
	orderedKeys := maps.Keys(m)
	out := ""

	sort.Strings(orderedKeys)

	for _, key := range orderedKeys {
		out += key + ": " + m[key] + "\n"
	}

	return out
}

// Print a string representation of a string map
func Print(m map[string]string) {
	fmt.Printf("%s", Pretty(m))
}
