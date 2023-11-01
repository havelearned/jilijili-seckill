package utils

import "strconv"

func isStringEmpty(s string) bool {
	return s == ""
}

func StringToInt(str string) (int, error) {
	num, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}
	return num, nil
}
