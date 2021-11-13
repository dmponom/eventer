package stringtools

import (
	"strconv"
	"strings"
)

func JoinInts(nArr []int, sep string) string {
	if len(nArr) == 0 {
		return ""
	}
	strArr := make([]string, 0, len(nArr))
	for _, n := range nArr {
		strArr = append(strArr, strconv.Itoa(n))
	}
	return strings.Join(strArr, sep)
}
