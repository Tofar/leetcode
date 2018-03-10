package leetcode26_30

import (
	"testing"
)

func TestStrStr(t *testing.T) {
	if num := strStr("hello", "ll"); num == 2 {
		t.Log("pass")
	} else {
		t.Error("faild")
	}
}
