package main

import (
	"flag"
	"fmt"
)

func main() {
	problemNumber := flag.Int("problem", 1, "problem number")
	flag.Parse()
	//data1 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	data2 := []int{1, 2, 3, 4, 5}
	switch *problemNumber {
	case 1:
		fmt.Println(RemoveDuplicates(data2), data2)
	case 2:
		fmt.Println(MaxProfit(data2))
	}
}
