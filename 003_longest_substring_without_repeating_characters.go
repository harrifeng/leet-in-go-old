package leetcode

import "math"

func lengthOfLongestSubstring(s string) int {
	start := 0
	maxv := 0
	m := map[rune]int{}

	for i, r := range s {
		if v, ok := m[r]; ok {
			start = int(math.Max(float64(start), float64(v+1)))
		}
		m[r] = i
		maxv = int(math.Max(float64(maxv), float64(i-start+1)))
	}
	return maxv
}
