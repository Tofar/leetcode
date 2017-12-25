package leetcode22_25

/*
  Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

For example, given n = 3, a solution set is:

[
  "((()))",
  "(()())",
  "(())()",
  "()(())",
  "()()()"
]
*/

func GenerateParenthesis(n int) []string {
	var result []string
	var s string
	generate(n, n, s, &result)
	return result
}

func generate(leftNum, rightNum int, s string, result *[]string) {
	if leftNum == 0 && rightNum == 0 {
		*result = append(*result, s)
	}
	if leftNum > 0 {
		generate(leftNum-1, rightNum, s+"(", result)
	}
	if rightNum > 0 && leftNum < rightNum {
		generate(leftNum, rightNum-1, s+")", result)
	}
}
