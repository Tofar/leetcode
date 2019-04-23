func uniquePaths(m int, n int) int {
    if m <= 0 || n <= 0 {
        return 0
    }
    dp := make([][]int, n)
    for i := 0; i < n; i++ {
        dp[i] = make([]int, m)
        if i == 0 {
            dp[0][0] = 1
        }
        for j := 0; j < m; j++ {
            if j > 0 {
                dp[i][j] += dp[i][j-1]                 
            }
            if i > 0 {
                dp[i][j] += dp[i-1][j]
            }
        }
    }
    return dp[n-1][m-1]
}
