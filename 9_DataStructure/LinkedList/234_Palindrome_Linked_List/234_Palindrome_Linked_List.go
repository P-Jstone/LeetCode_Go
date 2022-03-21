package DS

import "github.com/P-Jstone/LeetCode_Go/Structure"

type ListNode = Structure.ListNode

func isPalindrome(head *ListNode) bool {
	vals := make([]int, 0)
	for ; head != nil; head = head.Next {
		vals = append(vals, head.Val)
	}
	n := len(vals)
	for l, r := 0, n-1; l < r; l, r = l+1, r-1 {
		if vals[l] != vals[r] {
			return false
		}
	}
	return true
}
