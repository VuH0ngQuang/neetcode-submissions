/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
    mp := make(map[*ListNode]struct{})

	for head != nil {
		if _,exist := mp[head]; exist {
			return true
		}
		mp[head] = struct{}{}
		head = head.Next
	}

	return false
}
