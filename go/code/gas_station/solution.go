package gas_station

func CanCompleteCircuit(gas, cost []int) int {
	total, current, start := 0, 0, -1
	for i, g := range gas {
		c := cost[i]
		current += g - c
		total += g - c
		if current < 0 {
			current, start = 0, -1
		} else if start < 0 {
			start = i
		}
	}
	if total >= 0 {
		return start
	} else {
		return -1
	}
}
