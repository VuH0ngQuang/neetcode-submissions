/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func sum(a int, b int, carry bool) (int, bool) {
	tmp := a + b
	if carry {
		tmp++
	}

	if tmp >= 10 {
		return tmp % 10, true
	}
	return tmp, false
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	ans := dummy
	var carry bool

	for l1 != nil && l2 != nil {
		tmp := 0
		tmp, carry = sum(l1.Val, l2.Val, carry)

		dummy.Next = &ListNode{Val: tmp}
		dummy = dummy.Next

		l1 = l1.Next
		l2 = l2.Next
	}

	for l1 != nil {
		tmp := 0
		tmp, carry = sum(l1.Val, 0, carry)

		dummy.Next = &ListNode{Val: tmp}
		dummy = dummy.Next

		l1 = l1.Next
	}

	for l2 != nil {
		tmp := 0
		tmp, carry = sum(l2.Val, 0, carry)

		dummy.Next = &ListNode{Val: tmp}
		dummy = dummy.Next

		l2 = l2.Next
	}

	if carry {
		dummy.Next = &ListNode{Val: 1}
	}

	return ans.Next
}
