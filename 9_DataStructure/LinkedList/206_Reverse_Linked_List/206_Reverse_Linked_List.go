package DS

import "github.com/P-Jstone/LeetCode_Go/Structure"

type ListNode = Structure.ListNode

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	pre = nil
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}
