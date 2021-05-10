package leet_code

/*MoveZeroes
https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/567/
*/
func MoveZeroes(nums []int) {
	indicator := 0
	for i := 0; i < len(nums)-indicator; i++ {
		if nums[i] == 0 {
			nums = append(nums[:i], nums[i+1:]...)
			nums = append(nums, 0)
			indicator++
			i--
		}
	}
}
