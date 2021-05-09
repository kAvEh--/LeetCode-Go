package main

/*ContainsDuplicate
https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/578/
*/
func ContainsDuplicate(nums []int) bool {
	tmp := make(map[int]int)
	for i := range nums {
		if tmp[nums[i]] > 0 {
			return true
		}
		tmp[nums[i]] = 1
	}
	return false
}
