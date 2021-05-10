package main

import (
	easy "LeetCode-Go/interview-easy"
	"flag"
	"fmt"
)

func main() {
	problemNumber := flag.Int("problem", 1, "problem number")
	flag.Parse()
	switch *problemNumber {
	case 1:
		fmt.Println(easy.RemoveDuplicates([]int{2, 7, 11, 15}))
	case 2:
		fmt.Println(easy.MaxProfit([]int{2, 7, 11, 15}))
	case 3:
		easy.Rotate([]int{2, 7, 11, 15}, 3)
	case 4:
		fmt.Println(easy.ContainsDuplicate([]int{2, 7, 11, 15}))
	case 5:
		fmt.Println(easy.SingleNumber([]int{2, 7, 11, 15}))
	case 6:
		fmt.Println(easy.Intersect([]int{0, 0, 1}, []int{2, 7, 11, 15}))
	case 7:
		fmt.Println(easy.PlusOne([]int{0, 0, 1}))
	case 8:
		easy.MoveZeroes([]int{0, 0, 1})
	case 9:
		fmt.Println(easy.TwoSum([]int{-3, 4, 3, 90}, 0))
	case 10:
		fmt.Println(easy.IsValidSudoku([][]byte{{'8', '3', '.', '.', '7', '.', '.', '.', '.'}, {'6', '.', '.', '1', '9', '5', '.', '.', '.'}, {'.', '9', '8', '.', '.', '.', '.', '6', '.'}, {'8', '.', '.', '.', '6', '.', '.', '.', '3'}, {'4', '.', '.', '8', '.', '3', '.', '.', '1'}, {'7', '.', '.', '.', '2', '.', '.', '.', '6'}, {'.', '6', '.', '.', '.', '.', '2', '8', '.'}, {'.', '.', '.', '4', '1', '9', '.', '.', '5'}, {'.', '.', '.', '.', '8', '.', '.', '7', '9'}}))
	case 11:
		easy.RotateII([][]int{{1, 2}, {3, 4}})
	}
}
