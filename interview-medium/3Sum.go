package interview_medium

import (
	"sort"
)

/*ThreeSum
https://leetcode.com/explore/interview/card/top-interview-questions-medium/103/array-and-strings/776/
*/
func ThreeSum(nums []int) [][]int {
	var res [][]int
	if len(nums) < 3 {
		return res
	}
	for i := 0; i <= len(nums)-3; i++ {
		tmp := make(map[int]int)
		for j := i + 1; j < len(nums); j++ {
			k, ok := tmp[-nums[j]]
			if ok {
				tt := []int{nums[i], nums[k], nums[j]}
				sort.Ints(tt)
				for r := range res {
					if res[r][0] == tt[0] && res[r][1] == tt[1] && res[r][2] == tt[2] {
						goto skip
					}
				}
				res = append(res, tt)
			}
		skip:
			tmp[nums[i]+nums[j]] = j
		}
	}
	return res
}
