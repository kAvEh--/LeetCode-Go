package interview_medium

/*IncreasingTriplet
https://leetcode.com/explore/interview/card/top-interview-questions-medium/103/array-and-strings/781/
*/
func IncreasingTriplet(nums []int) bool {
	max := nums[len(nums)-1]
	if len(nums) < 3 {
		return false
	}
	tmp := make([]bool, len(nums))
	for i := len(nums) - 2; i > 0; i-- {
		if nums[i] < max {
			tmp[i] = true
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	min := nums[0]
	for i := 1; i < len(nums)-1; i++ {
		if min < nums[i] && tmp[i] {
			return true
		}
		if min > nums[i] {
			min = nums[i]
		}
	}
	return false
}
