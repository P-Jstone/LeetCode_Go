package Pointer

import ""

type ListNode = Structure.ListNode

func detectCycle(head *ListNode) *ListNode {
    fast, slow := head, head
    for fast != nil && fast.Next != nil {
    	fast = fast.Next.Next
    	slow = slow.Next
		if slow == fast {
			for slow != head {
				slow = slow.Next
				head = head.Next
			}
			return head
		}
	}
	return nil
}
//leetcode submit region end(Prohibit modification and deletion)
