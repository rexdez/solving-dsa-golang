package main

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var carry int = 0
	var result *ListNode = &ListNode{}
	res_ptr := result
	for l1 != nil || l2 != nil || carry > 0 {
		var sum int
		sum = carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		result.Next = &ListNode{
			Val: sum%10,
		}
		result =  result.Next
		carry = sum/10
	}
	return res_ptr.Next
}