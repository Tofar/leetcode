func numRabbits(answers []int) int {
    if len(answers) == 0 {
        return 0
    }
    sort.Ints(answers)
    sum := 0
    for i := 0; i < len(answers); i++ {
        cur := answers[i]
        sum += cur+1
        j := i+cur
        if i+cur >= len(answers) {
            j = len(answers)-1
        }
        for ; j > i; j-- {
            if answers[j] == cur {
                break
            }
        }
        i = j
    }
    
    return sum
}
