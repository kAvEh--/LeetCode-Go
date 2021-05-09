package main

func Rotate(nums []int, k int) {
	r := len(nums) - k%len(nums)
	nums = append(nums[r:], nums[:r]...)
}
