package main

func Rotate(nums []int, k int) {
	r := len(nums) - k%len(nums)
	temp := append(nums[r:], nums[:r]...)
	copy(nums, temp)
}
