package min_cost_climbing_stairs

func MinCostClimbingStairs(costs []int) int {
	f1, f2 := 0, 0
	for i := len(costs) - 1; i > -1; i-- {
		f0 := costs[i]
		if f1 < f2 {
			f0 += f1
		} else {
			f0 += f2
		}
		f1, f2 = f0, f1
	}
	if f1 < f2 {
		return f1
	} else {
		return f2
	}
}
