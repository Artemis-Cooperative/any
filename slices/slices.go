package slices

import (
	intslices "github.com/Artemis-Cooperative/types/intslices"
)

// Concatenate the slices given
func Concat[T any](slices ...[]T) []T {
	var concatenated []T

	for _, slice := range slices {
		concatenated = append(concatenated, slice...)
	}

	return concatenated
}

// Remove an element from the slice by its index
func Remove[T any](slice []T, i int) []T {
	return append(slice[:i], slice[(i+1):]...)
}

func RemoveEmpty[T any](slices [][]T) [][]T {
	emptyIndexes := []int{}

	// Determine which slices are empty
	for i := range slices {
		if len(slices[i]) == 0 {
			emptyIndexes = Prepend(emptyIndexes, i)
		}
	}

	// Remove those slices
	for _, index := range emptyIndexes {
		slices = Remove(slices, index)
	}

	return slices
}

func Reverse[T any](slice []T) []T {
	reversed := []T{}

	for i := (len(slice) - 1); i >= 0; i-- {
		reversed = append(reversed, slice[i])
	}

	return reversed
}

func Prepend[T any](slice []T, item T) []T {
	return append([]T{item}, slice...)
}

func Shift[T any](slice []T) []T {
	switch len(slice) {
	case 0:
		return slice
	case 1:
		return []T{}
	default:
		return slice[1:]
	}
}

func Indexes[T any](slice []T) []int {
	return intslices.Seq(0, len(slice)-1)
}
