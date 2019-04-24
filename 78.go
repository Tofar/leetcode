func subsets(nums []int) [][]int {
	n := len(nums)
	res := make([][]int, 1<<uint(n))
	resLastI := make([]int, 1<<uint(n))

	res[0] = []int{}
	resLastI[0] = -1
	preL := 0
	preR := 1 // Left closed right away

	for i := 1; i <= n; i++ {
		m := preR
		for j := preL; j < preR; j++ {
			for x := resLastI[j] + 1; x < n; x++ {
				res[m] = append(clone(res[j]), nums[x])
				resLastI[m] = x
				m++
			}
		}
		preL, preR = preR, m
	}
	return res
}

func clone(a []int) []int {
	b := make([]int, len(a))
	for i, v := range a {
		b[i] = v
	}
	return b
}
