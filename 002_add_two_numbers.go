package leetcode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	ret := &ListNode{}
	tmp := ret
	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		ret.Next = &ListNode{sum % 10, nil}
		ret = ret.Next
		carry = sum / 10
	}
	return tmp.Next
}
