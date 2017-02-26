package leetcode

import (
	"math"
	"strings"
)

func longestPalindrome(s string) string {
	ns := "^" + strings.Join(strings.Split(s, ""), "#") + "$"
	cnt := [3000]int{}

	tmpi_c := 0
	_ = tmpi_c
	tmpi := 0
	maxi := 0
	maxv := 0

	for i := 1; i < len(ns)-1; i++ {
		if i < tmpi {
			cnt[i] = int(math.Min(float64(cnt[2*tmpi_c-i]),
				float64(tmpi-i)))
		} else {
			cnt[i] = 1
		}

		for ns[i+cnt[i]] == ns[i-cnt[i]] {
			cnt[i] += 1
		}

		if i+cnt[i] > tmpi {
			tmpi = i + cnt[i]
			tmpi_c = i
		}

		if maxv < cnt[i] {
			maxv = cnt[i]
			maxi = i
		}
	}

	return s[maxi/2-maxv/2 : maxi/2-maxv/2+1]
}
