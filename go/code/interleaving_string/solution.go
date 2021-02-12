package interleaving_string

func IsInterleave(lhs, rhs, target string) bool {
	lhsLen, rhsLen := len(lhs), len(rhs)
	if lhsLen+rhsLen != len(target) {
		return false
	}

	dp := make([][]bool, lhsLen+1)
	for i := range dp {
		dp[i] = make([]bool, rhsLen+1)
	}
	dp[0][0] = true

	for i := 0; i <= lhsLen; i++ {
		for j := 0; j <= rhsLen; j++ {
			if i > 0 && dp[i-1][j] && target[i+j-1] == lhs[i-1] ||
				j > 0 && dp[i][j-1] && target[i+j-1] == rhs[j-1] {
				dp[i][j] = true
			}
		}
	}
	return dp[lhsLen][rhsLen]
}
