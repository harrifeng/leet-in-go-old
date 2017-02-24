package leetcode

import "math"

func findNth(nums1 []int, nums2 []int, nth int) int {

	if len(nums1) > len(nums2) {
		return findNth(nums2, nums1, nth)
	}
	if len(nums1) == 0 {
		return nums2[nth-1]
	}
	if len(nums2) == 0 {
		return nums1[nth-1]
	}
	if nth == 1 {
		return int(math.Min(float64(nums1[0]), float64(nums2[0])))
	}

	partA := int(math.Min(float64(len(nums1)), float64(nth/2)))
	partB := nth - partA
	if nums1[partA-1] < nums2[partB-1] {
		return findNth(nums1[partA:], nums2, nth-partA)
	} else {
		return findNth(nums1, nums2[partB:], nth-partB)
	}
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	sum := len(nums1) + len(nums2)

	if sum%2 == 1 {
		return float64(findNth(nums1, nums2, sum/2+1))
	} else {
		return float64(findNth(nums1, nums2, sum/2+1)+findNth(nums1, nums2, sum/2)) / 2.0
	}
	return 0
}
