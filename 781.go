func numRabbits(answers []int) int {
    if len(answers) == 0 {
        return 0
    }
    answersMap := map[int]int{}
    for _, v := range answers {
        answersMap[v]++
    }
    sum := 0
    for k, v := range answersMap {
        sum += (k+1)*(int(math.Ceil(float64(v)/float64(k+1))))
    }
    
    return sum
}
