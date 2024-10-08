package strslices

import (
	"regexp"
	"strings"

	intslices "github.com/Artemis-Cooperative/types/intslices"
)

func Explode(text string, delimiter string) []string {
	if len(delimiter) > len(text) {
		return strings.Split(delimiter, text)
	} else {
		return strings.Split(text, delimiter)
	}
}

func ToCSV(slice []string) string {
	return strings.Join(slice, ",")
}

func CSVToSlice(csv string) []string {
	delimiter := ","

	if regexp.MustCompile(", ").Match([]byte(csv)) {
		delimiter = ", "
	}

	return Explode(csv, delimiter)
}

func Pretty(a []string) string {
	out := "[\n    "

	out += strings.Join(a, "\n    ")

	out += "\n]\n"

	return out
}

func PrettyToSlice(s string) []string {
	// Remove brackets
	s = strings.TrimLeft(s, "[\n"+strings.Repeat(" ", 4))
	s = strings.TrimRight(s, "\n]")

	// Convert to array and return
	return strings.Split(s, "\n    ")
}

func Equals(slice1, slice2 []string) bool {
	if len(slice1) != len(slice2) {
		return false
	}

	for i := range slice1 {
		if slice1[i] != slice2[i] {
			return false
		}
	}

	return true
}

func MaxLength(slice []string) int {
	max := 0

	for _, value := range slice {
		if len(value) > max {
			max = len(value)
		}
	}

	return max
}

func Repeat(slice []string, count int) []string {
	repeated := []string{}

	for range intslices.Seq(1, count) {
		repeated = append(repeated, slice...)
	}

	return repeated
}
