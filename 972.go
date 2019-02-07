package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)


func isRationalEqual(S string, T string) bool {
	Ss := strings.Split(S, "(")
	Ts := strings.Split(T, "(")
	var s, t string
	if len(Ss) > 1 {
		s = expandS(Ss[0], Ss[1][:len(Ss[1])-1], 18)
	} else {
		s = Ss[0]
	}

	if len(Ts) > 1 {
		t = expandS(Ts[0], Ts[1][:len(Ts[1])-1], 18)
	} else {
		t = Ts[0]
	}

	sN, _ := strconv.ParseFloat(s, 64)
	tN, _ := strconv.ParseFloat(t, 64)

	fmt.Println(sN, tN)
	if math.Abs(sN-tN) < 1e-14 {
		return true
	}
	return false
}

func expandS(or, ep string, limit int) string {
	orL := len(or)
	epL := len(ep)
	for i := orL; i <= limit; i += epL {
		or += ep
	}
	return or
}

func main() {
	fmt.Println(1e2 == 100)
	fmt.Println(isRationalEqual("0.(52)",
		"0.5(25)"))
}

