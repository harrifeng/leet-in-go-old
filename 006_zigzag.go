package leetcode

import "bytes"

func convert(s string, numRows int) string {

	if numRows == 1 {
		return s
	}

	ret := [1000]bytes.Buffer{}
	size := 2*numRows - 2

	for i, c := range s {
		cur := i % size
		if cur < numRows {
			ret[cur].WriteRune(c)
		} else {
			ret[size-cur].WriteRune(c)
		}
	}

	for i := 1; i < numRows; i++ {
		ret[0].WriteString(ret[i].String())
	}

	return ret[0].String()
}
