package main

/*RemoveDuplicates
https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/727/
*/
func RemoveDuplicates(nums []int) int {
	pivot := 0
	for i := 1; i < len(nums); i++ {
		if nums[pivot] != nums[i] {
			pivot++
			nums[pivot] = nums[i]
		}
	}
	return pivot + 1
}
