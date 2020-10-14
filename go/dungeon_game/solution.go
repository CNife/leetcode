package dungeon_game

func CalculateMinimumHP(dungeon [][]int) int {
	m, n := len(dungeon), len(dungeon[0])
	memo := make([][]int, 0, m)
	for i := 0; i < m; i++ {
		memo = append(memo, make([]int, n))
	}

	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if i == m-1 && j == n-1 {
			return max(1-dungeon[i][j], 1)
		}
		if memo[i][j] > 0 {
			return memo[i][j]
		}

		var result int
		if i == m-1 {
			result = max(dfs(i, j+1)-dungeon[i][j], 1)
		} else if j == n-1 {
			result = max(dfs(i+1, j)-dungeon[i][j], 1)
		} else {
			result = max(min(dfs(i+1, j), dfs(i, j+1))-dungeon[i][j], 1)
		}
		memo[i][j] = result
		return result
	}
	return dfs(0, 0)
}

func max(lhs, rhs int) int {
	if lhs < rhs {
		return rhs
	}
	return lhs
}

func min(lhs, rhs int) int {
	if lhs < rhs {
		return lhs
	}
	return rhs
}
