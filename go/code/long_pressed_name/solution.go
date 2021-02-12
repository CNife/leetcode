package long_pressed_name

func IsLongPressedName(name, typed string) bool {
	var prev byte
	i, j := 0, 0
	for i < len(name) && j < len(typed) {
		if name[i] == typed[j] {
			prev = name[i]
			i++
			j++
		} else if prev != 0 && typed[j] == prev {
			j++
		} else {
			return false
		}
	}
	for j < len(typed) {
		if typed[j] != prev {
			return false
		}
		j++
	}
	return i == len(name)
}
