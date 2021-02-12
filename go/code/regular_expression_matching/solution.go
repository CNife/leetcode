package regular_expression_matching

func IsMatch(s string, p string) bool {
	sLen, pLen := len(s), len(p)
	dp := make([][]bool, sLen+1)
	for i := range dp {
		dp[i] = make([]bool, pLen+1)
	}

	dp[0][0] = true
	if pLen > 0 && p[0] == '*' {
		return false
	}
	for i := 2; i <= pLen; i++ {
		if p[i-1] == '*' {
			dp[0][i] = dp[0][i-2]
		}
	}

	for i := 1; i <= sLen; i++ {
		for j := 1; j <= pLen; j++ {
			if s[i-1] == p[j-1] || p[j-1] == '.' {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {
				if s[i-1] == p[j-2] || p[j-2] == '.' {
					dp[i][j] = dp[i][j-2] || dp[i-1][j-2] || dp[i-1][j]
				} else {
					dp[i][j] = dp[i][j-2]
				}
			}
		}
	}
	return dp[sLen][pLen]
}
