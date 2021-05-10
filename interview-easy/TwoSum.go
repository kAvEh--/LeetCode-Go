package leet_code

/*TwoSum
https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/546/
*/
func TwoSum(nums []int, target int) []int {
	hash := make(map[int]int)
	ret := make([]int, 2)
	for i := 0; i < len(nums); i++ {
		if _, ok := hash[target-nums[i]]; ok {
			ret[0] = hash[target-nums[i]]
			ret[1] = i
			return ret
		}
		hash[nums[i]] = i
	}
	return ret
}
