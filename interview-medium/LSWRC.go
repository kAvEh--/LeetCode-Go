package interview_medium

import "fmt"

/*LengthOfLongestSubstring
https://leetcode.com/explore/interview/card/top-interview-questions-medium/103/array-and-strings/779/
*/
func LengthOfLongestSubstring(s string) int {
	start, end, max := 0, 0, 0
	tmp := make(map[byte]int)
	if len(s) == 0 {
		return 0
	}
	for i := 0; i < len(s); i++ {
		fmt.Println(tmp, start, i)
		if tmp[s[i]] > 0 {
			for s[start] != s[i] {
				tmp[s[start]] -= 1
				start++
			}
			start++
			continue
		}
		tmp[s[i]] += 1
		end = i
		if end-start+1 > max {
			max = end - start + 1
		}
	}
	return max
}
