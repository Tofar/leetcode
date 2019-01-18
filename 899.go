package main

import (
	"fmt"
	"sort"
	"strings"
)

type byteSlice []byte

func (this byteSlice) Len() int           { return len(this) }
func (this byteSlice) Swap(i, j int)      { this[i], this[j] = this[j], this[i] }
func (this byteSlice) Less(i, j int) bool { return this[i] < this[j] }

func orderlyQueue(S string, K int) string {
	if K > 1 {
		tmp := byteSlice(S)
		sort.Sort(tmp)
		return string(tmp)
	}

	min := S[0]
	minIs := []int{}
	for i := 0; i < len(S); i++ {
		if S[i] < min {
			min = S[i]
			minIs = []int{i}
		} else if S[i] == min {
			minIs = append(minIs, i)
		}
	}

	result := S

	for _, i := range minIs {
		tmp := S[i:] + S[:i]
		if strings.Compare(result, tmp) == 1 {
			result = tmp
		}
	}
	return result
}

func main() {
	fmt.Println(orderlyQueue("baaca", 3))
}

