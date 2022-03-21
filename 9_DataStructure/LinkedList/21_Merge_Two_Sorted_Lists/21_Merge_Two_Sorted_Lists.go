package DS

import "github.com/P-Jstone/LeetCode_Go/Structure"

type ListNode = Structure.ListNode

// 递归
//func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
//	if list1 == nil {
//		return list2
//	} else if list2 == nil {
//		return list1
//	} else if list1.Val < list2.Val {
//		list1.Next = mergeTwoLists(list1.Next, list2)
//		return list1
//	} else {
//		list2.Next = mergeTwoLists(list1, list2.Next)
//		return list2
//	}
//}

// 迭代
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	prehead := &ListNode{}
	prev := prehead
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			prev.Next = list1
			list1 = list1.Next
		} else {
			prev.Next = list2
			list2 = list2.Next
		}
		prev = prev.Next
	}
	if list1 == nil {
		prev.Next = list2
	} else {
		prev.Next = list1
	}
	return prehead.Next
}
