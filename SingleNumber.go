package main

import "sort"

/*SingleNumber
https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/549/
*/
func SingleNumber(nums []int) int {
	sort.Ints(nums)
	for i := 0; i < len(nums)/2; i++ {
		if nums[2*i] != nums[2*i+1] {
			return nums[2*i]
		}
	}
	return nums[len(nums)-1]
}
