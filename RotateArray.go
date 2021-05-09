package main

/*Rotate
https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/646/
*/
func Rotate(nums []int, k int) {
	r := len(nums) - k%len(nums)
	temp := append(nums[r:], nums[:r]...)
	copy(nums, temp)
}
