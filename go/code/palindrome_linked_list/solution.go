package palindrome_linked_list

import . "github.com/CNife/leetcode/go/types"

func IsPalindrome(head *ListNode) bool {
	var values []int
	for node := head; node != nil; node = node.Next {
		values = append(values, node.Val)
	}
	for i := 0; i < len(values)/2; i++ {
		if values[i] != values[len(values)-1-i] {
			return false
		}
	}
	return true
}
