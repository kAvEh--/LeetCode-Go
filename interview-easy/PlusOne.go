package interview_easy

import "fmt"

/*PlusOne
https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/559/
*/
func PlusOne(digits []int) []int {
	if digits[len(digits)-1] == 9 {
		flag := true
		digits[len(digits)-1] = 0
		for i := len(digits) - 2; i >= 0; i-- {
			fmt.Println(digits, i)
			if digits[i] != 9 {
				flag = false
				digits[i] += 1
				break
			}
			digits[i] = 0
		}
		fmt.Println(digits)
		if flag {
			digits[0] = 1
			digits = append(digits, 0)
		}
		return digits
	}
	digits[len(digits)-1] += 1
	return digits
}
