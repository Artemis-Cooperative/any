package ints

import "strconv"

func IsNum(str string) bool {
	_, err := strconv.Atoi(str)

	if err == nil {
		return true
	} else {
		return false
	}
}
