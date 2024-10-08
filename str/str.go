package str

import (
	"fmt"
	"math"
	"regexp"
	"strings"
)

func IsAlphaNum(str string) bool {
	return regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString(str)
}

func TrimSides(str string, leftTrimCount int, rightTrimCount int) string {
	return (str[leftTrimCount:])[:(len(str) - (leftTrimCount + rightTrimCount))]
}

func Center(str string, width int) string {
	return CenterCustom(str, width, " ", " ")
}

func CenterCustom(str string, width int, leftFill string, rightFill string) string {
	padding := width - len(str)
	leftPadding := padding / 2
	rightPadding := divideAndCeil(padding, 2)
	leftFillCount := divideAndCeil(leftPadding, len(leftFill))
	rightFillCount := divideAndCeil(rightPadding, len(rightFill))

	return fmt.Sprintf(
		"%s%s%s",
		strings.Repeat(leftFill, leftFillCount),
		str,
		strings.Repeat(rightFill, rightFillCount),
	)
}

func divideAndCeil(num1 int, num2 int) int {
	return int(math.Ceil(float64(num1) / float64(num2)))
}

func PadLeft(str *string, padCount *int) {
	*str = *str + strings.Repeat(" ", *padCount)
}

func Left(str *string, width *int) {
	padCount := 0

	if *width > len(*str) {
		padCount = (*width)-len(*str)
	}

	*str = *str + strings.Repeat(" ", padCount)
}