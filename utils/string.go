package utils

import "strconv"

//ConvertStringToInt ...
func ConvertStringToInt(str string) (int, error) {
	res, err := strconv.Atoi(str)
	return res, err
}

//ConvertUInt64ToString ...
func ConvertUInt64ToString(number uint64) string {
	return strconv.FormatUint(number, 10)
}
