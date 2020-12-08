package split_array_into_fibonacci_sequence

import (
	"strconv"
	"testing"
)

func TestSplitIntoFibonacci(t *testing.T) {
	tests := []struct {
		s           string
		hasSolution bool
	}{
		{"0123", false},
		{"123456579", true},
		{"11235813", true},
		{"112358130", false},
		{"1101111", true},
	}
	for _, tt := range tests {
		got := SplitIntoFibonacci(tt.s)
		if !isValidSolution(tt.s, tt.hasSolution, got) {
			t.Errorf("s is \"%v\", got is %v", tt.s, got)
		}
	}
}

func isValidSolution(s string, hasSolution bool, got []int) bool {
	if !hasSolution {
		return len(got) == 0
	} else if len(got) == 0 {
		return false
	}
	for i := 0; i < len(got)-2; i++ {
		if got[i]+got[i+1] != got[i+2] {
			return false
		}
	}
	i := 0
	for _, num := range got {
		numString := strconv.Itoa(num)
		endI := i + len(numString)
		if endI > len(s) || s[i:endI] != numString {
			return false
		}
		i = endI
	}
	return true
}
