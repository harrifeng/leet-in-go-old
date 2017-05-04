package leetcode

import "math"

func reverse(x int) int {
	sign := 1
	if x < 0 {
		sign = -1
	}

	y := int(math.Abs(float64(x)))
	ret := 0
	for y > 0 {
		ret = ret*10 + y%10
		y = y / 10
	}

	if ret > int(math.Pow(2, 31))-1 {
		return 0
	}
	return ret * sign
}
