package main

import (
	easy "LeetCode-Go/interview-easy"
	"flag"
	"fmt"
)

func main() {
	problemNumber := flag.Int("problem", 1, "problem number")
	flag.Parse()
	data1 := []int{0, 0, 1}
	data2 := []int{0, 9, 4, 0, 8, 4, 0}
	switch *problemNumber {
	case 1:
		fmt.Println(easy.RemoveDuplicates(data2), data2)
	case 2:
		fmt.Println(easy.MaxProfit(data2))
	case 3:
		easy.Rotate(data2, 3)
	case 4:
		fmt.Println(easy.ContainsDuplicate(data2))
	case 5:
		fmt.Println(easy.SingleNumber(data2))
	case 6:
		fmt.Println(easy.Intersect(data1, data2))
	case 7:
		fmt.Println(easy.PlusOne(data1))
	case 8:
		easy.MoveZeroes(data1)
	}
}
