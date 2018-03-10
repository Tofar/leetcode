package leetcode26_30

import (
	"testing"
)

func TestDivide(t *testing.T) {
	if num := divide(1, 1); num == 1 {
		t.Log("pass")
	} else {
		t.Error("faild")
	}
}
