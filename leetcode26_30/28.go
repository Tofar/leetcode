package leetcode26_30

/*
Implement strStr().

Return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.

Example 1:

Input: haystack = "hello", needle = "ll"
Output: 2
Example 2:

Input: haystack = "aaaaa", needle = "bba"
Output: -1
*/

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	for k := 0; k < len(haystack); k++ {
		if haystack[k] == needle[0] {
			if (k + len(needle)) <= len(haystack) {
				i := 1
				for ; i < len(needle); i++ {
					if haystack[k+i] != needle[i] {
						break
					}
				}
				if i == len(needle) {
					return k
				}
			} else {
				break
			}
		}
	}
	return -1
}
