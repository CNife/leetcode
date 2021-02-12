package binary_prefix_divisible_by_5

import "math/big"

func PrefixDividedBy5(array []int) []bool {
	if len(array) <= 64 {
		return uint64Path(array)
	}
	return bigIntPath(array)
}

func uint64Path(array []int) []bool {
	var num uint64 = 0
	result := make([]bool, len(array))
	for i, bit := range array {
		num *= 2
		if bit == 1 {
			num++
		}
		result[i] = num%5 == 0
	}
	return result
}

func bigIntPath(array []int) []bool {
	num, remainder := big.NewInt(0), big.NewInt(0)
	result := make([]bool, len(array))

	for i, bit := range array {
		num.Mul(num, two)
		if bit == 1 {
			num.Add(num, one)
		}

		result[i] = remainder.Rem(num, five).Cmp(zero) == 0
	}
	return result
}

var (
	zero = big.NewInt(0)
	one  = big.NewInt(1)
	two  = big.NewInt(2)
	five = big.NewInt(5)
)
