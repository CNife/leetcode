package ugly_number_2

import "testing"

func TestNthUglyNumber(t *testing.T) {
	uglyNumbers := []int{1, 2, 3, 4, 5, 6, 8, 9, 10, 12}
	for i := range uglyNumbers {
		if got := NthUglyNumber(i + 1); got != uglyNumbers[i] {
			t.Fatalf("NthUglyNumber(%v) = %v, want %v",
				i+1, got, uglyNumbers[i])
		}
	}
}
