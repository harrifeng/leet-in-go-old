package leetcode

import (
	"math"
	"strings"
)

func longestPalindrome(s string) string {
	ns := "^#" + strings.Join(strings.Split(s, ""), "#") + "#$"
	cnt := [3000]int{}

	tmpi := 0
	tmpiCenter := 0
	maxv := 0
	maxvIndex := 0

	for i := 1; i < len(ns)-1; i++ {
		if i < tmpi {
			cnt[i] = int(math.Min(float64(cnt[2*tmpiCenter-i]),
				float64(tmpi-i)))
		} else {
			cnt[i] = 1
		}

		for ns[i+cnt[i]] == ns[i-cnt[i]] {
			cnt[i] += 1
		}

		if i+cnt[i] > tmpi {
			tmpi = i + cnt[i]
			tmpiCenter = i
		}

		if maxv < cnt[i] {
			maxv = cnt[i]
			maxvIndex = i
		}
	}

	return s[maxvIndex/2-maxv/2 : maxvIndex/2-maxv/2+maxv-1]
}
