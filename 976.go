func largestPerimeter(A []int) int {
    if len(A) < 3 {
        return 0
    }
    sort.Sort(sort.Reverse(sort.IntSlice(A)))
    sum := 0
    for i := 0; i < len(A); i++ {
        for j := i+1; j < len(A); j++ {
            k := j+1
            for ; k < len(A); k++ {
                tmpSum := A[i] + A[j] + A[k]
                if tmpSum < sum {
                    break
                }
                if A[i] < A[j] + A[k] && A[i] - A[k] < A[j] {
                    sum = tmpSum
                }
            }
            if k == j+1 {
                break
            }
        }
    }
    return sum
}
