package leetcode

import (
	"strings"
)

func convert(s string, numRows int) string {

	if numRows == 1 {
		return s
	}

	ret := [10]string{}
	size := 2*numRows - 2

	for i, c := range s {
		cur := i % size
		if cur < numRows {
			ret[cur] += string(c)
		} else {
			ret[size-cur] += string(c)
		}
	}

	return strings.Join(ret[:numRows+1], "")
}
