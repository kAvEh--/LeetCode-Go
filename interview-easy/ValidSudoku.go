package leet_code

import (
	"strconv"
)

/*IsValidSudoku
https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/769/
*/
func IsValidSudoku(board [][]byte) bool {
	for i := 0; i < len(board); i++ {
		if !isValid(board[i]...) {
			return false
		}
		var tmp []byte
		for j := 0; j < len(board); j++ {
			tmp = append(tmp, board[j][i])
		}
		if !isValid(tmp...) {
			return false
		}
	}
	if !isValid(getSquare(board, 0, 0)...) {
		return false
	}
	if !isValid(getSquare(board, 0, 3)...) {return false}
	if !isValid(getSquare(board, 0, 6)...) {return false}

	if !isValid(getSquare(board, 3, 0)...) {return false}
	if !isValid(getSquare(board, 3, 3)...) {return false}
	if !isValid(getSquare(board, 3, 6)...) {return false}

	if !isValid(getSquare(board, 6, 0)...) {return false}
	if !isValid(getSquare(board, 6, 3)...) {return false}
	if !isValid(getSquare(board, 6, 6)...) {return false}

	return true
}

func getSquare(board [][]byte, i, j int) []byte {
	var tt []byte
	tt = append(tt, board[i][j:j+3]...)
	tt = append(tt, board[i+1][j:j+3]...)
	tt = append(tt, board[i+2][j:j+3]...)
	return tt
}

func isValid(nums ...byte) bool {
	tmp := make(map[int]bool)
	for _, i := range nums {
		if string(i) != "." {
			tt, _ := strconv.Atoi(string(i))
			if tmp[tt] {
				return false
			}
			tmp[tt] = true
		}
	}
	return true
}
