package sort_integers_by_the_number_of_1_bits

import (
	"math/bits"
	"sort"
)

func SortByBits(array []int) []int {
	for i, num := range array {
		bitsCount := bits.OnesCount(uint(num))
		array[i] = num | (bitsCount << (bits.UintSize - 6))
	}
	sort.Ints(array)
	for i, num := range array {
		array[i] = num & int((^uint(0))>>6)
	}
	return array
}
