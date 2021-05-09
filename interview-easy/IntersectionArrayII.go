package leet_code

/*Intersect
https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/674/
*/
func Intersect(nums1 []int, nums2 []int) []int {
	tmp := make(map[int]int)
	ret := make([]int, 0)
	for i := range nums1 {
		tmp[nums1[i]] += 1
	}
	for i := range nums2 {
		if tmp[nums2[i]] > 0 {
			ret = append(ret, nums2[i])
			tmp[nums2[i]]--
		}
	}
	return ret
}
