package leetcode26_30

import (
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	if num := removeDuplicates2([]int{1, 1, 1, 2}); num == 2 {
		t.Log("pass")
	} else {
		t.Log("failed")
	}
}
