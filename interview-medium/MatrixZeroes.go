package interview_medium

/*SetZeroes
https://leetcode.com/explore/interview/card/top-interview-questions-medium/103/array-and-strings/777/
*/
func SetZeroes(matrix [][]int) {
	f1, f2 := false, false
	for i := 0; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			f2 = true
			continue
		}
	}
	for i := 0; i < len(matrix[0]); i++ {
		if matrix[0][i] == 0 {
			f1 = true
			continue
		}
	}
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[i]); j++ {
			if matrix[0][j] == 0 || matrix[i][0] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	if f2 {
		for i := 0; i < len(matrix); i++ {
			matrix[i][0] = 0
		}
	}
	if f1 {
		for i := 0; i < len(matrix[0]); i++ {
			matrix[0][i] = 0
		}
	}
}
