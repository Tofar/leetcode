func numDistinct(s, t string) int {
	dp := make([][]int, len(s)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(t)+1)
		for j := 0; j < len(t)+1; j++ {
			dp[i][j] = -1
		}
	}
	return helpFunc(s, t, dp, 0, 0)
}

func helpFunc(s, t string, dp [][]int, sI, tI int) int {
	if t == "" {
		return 1
	}

	if s == "" {
		return 0
	}

	for i := 0; i < len(s); i++ {
		if s[i] == t[0] {
			if dp[sI+i+1][tI+1] < 0 {
				dp[sI+i+1][tI+1] = helpFunc(s[i+1:], t[1:], dp, sI+i+1, tI+1)
			}
			if dp[sI+i+1][tI] < 0 {
				dp[sI+i+1][tI] = helpFunc(s[i+1:], t, dp, sI+i+1, tI)
			}

			return dp[sI+i+1][tI+1] + dp[sI+i+1][tI]
		}
	}
	return 0
}
