package helper

import "strconv"

func StrToUint(str string) uint {
	num, _ := strconv.ParseUint(str, 10, 0)
	return uint(num)
}
