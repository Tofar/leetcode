func maximalSquare(matrixByte [][]byte) int {
    m := len(matrixByte)
    if m == 0 {
        return 0
    }
    n := len(matrixByte[0])
    if n == 0 {
        return 0
    }
    
    matrix := make([][]int, m)
    maxD := 0
    for i := m-1; i >= 0; i-- {
        matrix[i] = make([]int, n)
        for j := n-1; j >= 0; j-- {
            if matrixByte[i][j] == '1' {
                if maxD == 0 {
                    maxD = 1
                }
                matrix[i][j] = 1
                if j < n-1 && i < m-1 {
                    min := matrix[i+1][j+1]
                    if min > matrix[i+1][j] {
                        min = matrix[i+1][j]
                    }
                    if min > matrix[i][j+1] {
                        min = matrix[i][j+1]
                    }
                    matrix[i][j] = min + 1
                    if min+1 > maxD {
                        maxD = min + 1
                    }
                }
            }
        }
    }
    return maxD*maxD
}
