package leet_code

/*ReverseString
https://leetcode.com/explore/interview/card/top-interview-questions-easy/127/strings/879/
*/
func ReverseString(s []byte) {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
	}
}
