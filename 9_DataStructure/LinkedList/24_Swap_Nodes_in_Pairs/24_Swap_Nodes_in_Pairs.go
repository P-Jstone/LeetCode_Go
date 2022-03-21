package DS

import (
	"github.com/P-Jstone/LeetCode_Go/Structure"
)

type ListNode = Structure.ListNode

// 递归
//func swapPairs(head *ListNode) *ListNode {
//	if head == nil || head.Next == nil {
//		return head
//	}
//	newHead := head.Next
//	head.Next = swapPairs(newHead.Next)
//	newHead.Next = head
//	return newHead
//}

// 迭代
func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{
		Next: head,
	}
	temp := dummy
	for temp.Next != nil && temp.Next.Next != nil {
		node1 := temp.Next
		node2 := temp.Next.Next
		temp.Next = node2
		node1.Next = node2.Next
		node2.Next = node1
		temp = node1
	}
	return dummy.Next
}
