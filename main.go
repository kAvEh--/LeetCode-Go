package main

import (
	"flag"
	"fmt"
)

func main() {
	problemNumber := flag.Int("problem", 1, "problem number")
	flag.Parse()
	data1 := []int{4, 9, 5}
	data2 := []int{9, 4, 9, 8, 4}
	switch *problemNumber {
	case 1:
		fmt.Println(RemoveDuplicates(data2), data2)
	case 2:
		fmt.Println(MaxProfit(data2))
	case 3:
		Rotate(data2, 3)
	case 4:
		fmt.Println(ContainsDuplicate(data2))
	case 5:
		fmt.Println(SingleNumber(data2))
	case 6:
		fmt.Println(Intersect(data1, data2))
	}
}
