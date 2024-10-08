package bytes

import (
	"bytes"
)

// Does not leave empty slices like bytes.Split()
func Explode(data []byte, delimiter []byte) [][]byte {
	if len(delimiter) > len(data) {
		return detonate(delimiter, data)
	} else {
		return detonate(data, delimiter)
	}
}

func detonate(data []byte, delimiter []byte) [][]byte {
	shrapnels := [][]byte{}
	shrapnel := []byte{}
	trash := []byte{}
	delimiterIndex := 0

	for _, dataByte := range data {
		if dataByte == delimiter[delimiterIndex] {
			trash = append(trash, dataByte)

			if delimiterIndex == (len(delimiter) - 1) {
				trash = []byte{}
				delimiterIndex = 0

				if len(shrapnel) > 0 {
					shrapnels = append(shrapnels, shrapnel)
				}

				shrapnel = []byte{}
			} else {
				delimiterIndex++
			}
		} else {
			shrapnel = append(shrapnel, trash...)
			shrapnel = append(shrapnel, dataByte)
			delimiterIndex = 0
		}
	}

	if len(shrapnel) > 0 {
		shrapnels = append(shrapnels, shrapnel)
	}

	return shrapnels
}

func SlicesToPrintString(byteSlices [][]byte) string {
	out := "[\n"

	for _, byteSlice := range byteSlices {
		out += "    [ " + OneLine(byteSlice) + " ]\n"
	}

	out += "]\n"

	return out
}

func OneLine(byteSlice []byte) string {
	out := ""

	for i, singleByte := range byteSlice {
		out += "'" + string(singleByte) + "'"

		if i < (len(byteSlice) - 1) {
			out += ", "
		}
	}

	return out
}

func AppendRepeat(slice []byte, symbol byte, count int) []byte {
	return append(slice, bytes.Repeat([]byte{symbol}, count)...)
}

func Strings(slices [][]byte) []string {
	out := []string{}

	for _, slice := range slices {
		out = append(out, string(slice))
	}

	return out
}
