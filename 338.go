func countBits(num int) []int {
    if num == 0 {
        return []int{0}
    }
    bits := []int{0, 1}
    left := 1
    step := 2
    for i := 1; i <= num; step<<=1 {
        t := num-i
        if t > step {
            t = step
        }
        var tempBits []int
        if t <= step>>1 {
            tempBits = bits[left:left+t]
        } else {
            pre := bits[left:]
            tempBits = append([]int{}, pre...)
            r := t-1-step>>1
            for j, v := range pre {
                if j > r {
                    break
                }
                tempBits = append(tempBits, v+1)
            }
        }
        bits = append(bits, tempBits...)
        i += step
        left += step>>1
    }
    return bits
}
