package main

import "fmt"

func validTicTacToe(board []string) bool {
	xNum, oNum := getNumOfChar(board)
	if xNum < oNum || xNum-oNum > 1 {
		return false
	}
	ch := verifyOver(board)
	if len(ch) == 0 {
		return true
	}

	if len(ch) == 1 {
		if ch[0] == 'X' && xNum == oNum+1 {
			return true
		}
		if ch[0] == 'O' && xNum == oNum {
			return true
		}
	}

	return false
}

func getNumOfChar(board []string) (x int, o int) {
	for _, s := range board {
		for i := 0; i < len(s); i++ {
			if s[i] == 'X' {
				x++
			} else if s[i] == 'O' {
				o++
			}
		}
	}
	return x, o
}

func verifyOver(board []string) []byte {
	tempRes := []byte{}
	for i := 0; i < 3; i++ {
		if board[i][0] != ' ' && board[i][0] == board[i][1] && board[i][1] == board[i][2] {
			tempRes = append(tempRes, board[i][0])
		}
		if board[0][i] != ' ' && board[0][i] == board[1][i] && board[1][i] == board[2][i] {
			tempRes = append(tempRes, board[0][i])
		}
	}

	if board[0][0] != ' ' && board[0][0] == board[1][1] && board[1][1] == board[2][2] {
		tempRes = append(tempRes, board[0][0])
	}
	if board[1][1] != ' ' && board[0][2] == board[1][1] && board[1][1] == board[2][0] {
		tempRes = append(tempRes, board[1][1])
	}

	res := []byte{}
	for _, b := range tempRes {
		if len(res) == 0 {
			res = append(res, b)
		} else {
			if res[0] != b {
				res = append(res, b)
				return res
			}
		}
	}

	return res
}

func main() {
	fmt.Println(validTicTacToe([]string{"XXX", "OOX", "OOX"}))
}

