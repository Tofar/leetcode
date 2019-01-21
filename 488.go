package main

import "fmt"

func findMinStep(board string, hand string) int {
	if board == "" {
		return 0
	}
	if hand == "" {
		return -1
	}

	b := []byte(board + "#")

	handCount := [26]int{}
	for _, h := range hand {
		handCount[h-'A']++
	}

	rs := recursion(b, handCount)
	if rs == 6 {
		return -1
	}
	return rs
}

func recursion(board []byte, handCount [26]int) int {
	board = removeConsecutiveBall(board)
	if board[0] == '#' {
		return 0
	}
	rs := 6
	for i, j := 0, 1; j < len(board); j++ {
		if board[i] == board[j] {
			continue
		}
		count := 3 - (j - i)
		if handCount[board[i]-'A'] >= count {
			handCount[board[i]-'A'] -= count
			tmpRs := count + recursion(append(append([]byte{}, board[:i]...), board[j:]...), handCount)
			if tmpRs < rs {
				rs = tmpRs
			}
			handCount[board[i]-'A'] += count
		}
		i = j
	}

	return rs
}

func removeConsecutiveBall(board []byte) []byte {
	for i, j := 0, 1; j < len(board); j++ {
		if board[i] == board[j] {
			continue
		}
		if j-i >= 3 {
			return removeConsecutiveBall(append(append([]byte{}, board[:i]...), board[j:]...))
		}
		i = j
	}
	return board
}

func main() {
	fmt.Println(findMinStep("RBYYBBRRB", "YRBGB"))
}

