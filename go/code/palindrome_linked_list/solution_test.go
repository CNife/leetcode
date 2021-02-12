package palindrome_linked_list

import (
	"testing"

	. "github.com/CNife/leetcode/go/types"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		head *ListNode
		want bool
	}{
		{NewList(1, 2), false},
		{NewList(1, 2, 2, 1), true},
	}
	for _, tt := range tests {
		if got := IsPalindrome(tt.head); got != tt.want {
			t.Errorf("IsPalindrome(%v) = %v, want %v", tt.head, got, tt.want)
		}
	}
}

func BenchmarkIsPalindrome(b *testing.B) {
	list := makePalindromeList(1_000_000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if !IsPalindrome(list) {
			b.Error("BenchmarkIsPalindrome failed")
		}
	}
}

func makePalindromeList(size int) *ListNode {
	values := make([]int, size)
	for i := 0; i < size/2; i++ {
		values[i] = i
	}
	for i := size / 2; i < size; i++ {
		values[i] = size - 1 - i
	}
	return NewList(values...)
}
