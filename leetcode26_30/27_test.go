package leetcode26_30

import (
	"testing"
)

func TestRemoveElement(t *testing.T) {
	if num := removeElement([]int{3, 2, 2, 3}, 3); num == 2 {
		t.Log("pass")
	} else {
		t.Log("failed")
	}
}
