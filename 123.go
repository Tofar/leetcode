func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	b1 := -prices[0]
	s1 := 0
	b2 := -prices[0]
	s2 := 0
	for i := 1; i < len(prices); i++ {
		if b1 < -prices[i] {
			b1 = -prices[i]
		}
		if s1 < b1+prices[i] {
			s1 = b1 + prices[i]
		}
		if b2 < s1-prices[i] {
			b2 = s1 - prices[i]
		}
		if s2 < b2+prices[i] {
			s2 = b2 + prices[i]
		}
	}
	return s2
}
