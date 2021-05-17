package interview_medium

import (
	"sort"
	"strings"
)

/*GroupAnagrams
https://leetcode.com/explore/interview/card/top-interview-questions-medium/103/array-and-strings/778/
*/
func GroupAnagrams(strs []string) [][]string {
	anagram := make([][]string, 0)
	tmp := make(map[string]int)
	for i := 0; i < len(strs); i++ {
		s := strings.Split(strs[i], "")
		sort.Strings(s)
		s2 := strings.Join(s, "")
		num, ok := tmp[s2]
		if ok {
			anagram[num] = append(anagram[num], strs[i])
			continue
		}
		tmp[s2] = len(anagram)
		tt := make([]string, 0)
		tt = append(tt, strs[i])
		anagram = append(anagram, tt)
	}
	return anagram
}
