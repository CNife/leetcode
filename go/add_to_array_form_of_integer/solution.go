package add_to_array_form_of_integer

func AddToArrayForm(array []int, k int) []int {
	reverse(array)
	kArray := toReversedArray(k)

	result := addReversedArray(array, kArray)
	reverse(result)
	return result
}

func reverse(array []int) {
	for i, j := 0, len(array)-1; i < j; i, j = i+1, j-1 {
		array[i], array[j] = array[j], array[i]
	}
}

func toReversedArray(num int) []int {
	if num == 0 {
		return []int{0}
	}

	result := make([]int, 0, 20)
	for num > 0 {
		digit := num % 10
		num /= 10
		result = append(result, digit)
	}
	return result
}

func addReversedArray(lhs, rhs []int) []int {
	result := make([]int, 0, max(len(lhs), len(rhs))+1)
	carry := false

	i, j := 0, 0
	for i < len(lhs) && j < len(rhs) {
		sum := lhs[i] + rhs[j]
		if carry {
			sum++
			carry = false
		}
		if sum > 9 {
			sum -= 10
			carry = true
		}
		result = append(result, sum)
		i, j = i+1, j+1
	}
	for i < len(lhs) {
		digit := lhs[i]
		if carry {
			digit++
			carry = false
		}
		if digit > 9 {
			digit -= 10
			carry = true
		}
		result = append(result, digit)
		i++
	}
	for j < len(rhs) {
		digit := rhs[j]
		if carry {
			digit++
			carry = false
		}
		if digit > 9 {
			digit -= 10
			carry = true
		}
		result = append(result, digit)
		j++
	}

	if carry {
		result = append(result, 1)
	}
	return result
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
