package leetcode

func twoSum(nums []int, target int) []int {
	m := map[int]int{}

	for i, c := range nums {
		if val, ok := m[target-c]; ok {
			return []int{val, i}
		}
		m[c] = i

	}
	return []int{}
}
