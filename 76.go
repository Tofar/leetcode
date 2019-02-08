func minWindow(s string, t string) string {
    tH := getSHash(t)
    wH := map[byte]int{} // window Hash
    l, r := -1, len(s)-1
    sum := 0
    for tmpL, i := 0, 0; i < len(s); i++ {
        wH[s[i]]++
        if wH[s[i]] <= tH[s[i]] {
            sum++
        }
        if sum == len(t) {
            for tmpL < i && wH[s[tmpL]] > tH[s[tmpL]] {
                wH[s[tmpL]]--
                tmpL++
            }
            if i-tmpL < r-l {
                r = i
                l = tmpL
            }
            wH[s[tmpL]]--
            sum--
            tmpL++
        }
    }
    
    if l == -1 {
        return ""
    }
    return s[l:r+1]
}

func getSHash(t string) map[byte]int {
    tH := map[byte]int{}
    for i := 0; i < len(t); i++ {
        tH[t[i]]++
    }   
    return tH
}
