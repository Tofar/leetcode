type intss [][]int

func (a intss) Len() int           { return len(a) }
func (a intss) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a intss) Less(i, j int) bool { 
    iL := len(a[i])
    if iL == 0 {
        return false
    }
    jL := len(a[j])
    if jL == 0 {
        return true
    }
    return a[i][iL-1] < a[j][jL-1]
}

func findLongestChain(pairs [][]int) int {
    if len(pairs) == 0 {
        return 0
    }
    sort.Sort(intss(pairs))
    right := pairs[0][len(pairs[0])-1]
    cnt := 1
    for i := 1; i < len(pairs); i++ {
        if pairs[i][0] > right {
            right = pairs[i][len(pairs[i])-1]
            cnt++
        } 
    }
    return cnt
}
