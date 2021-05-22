package interview_medium

/*AddTwoNumbers
https://leetcode.com/explore/interview/card/top-interview-questions-medium/107/linked-list/783/
*/
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	head := new(ListNode)
	cur := head
	for l1 != nil || l2 != nil || carry != 0 {
		n1, n2 := 0, 0
		if l1 != nil {
			n1, l1 = l1.Val, l1.Next
		}
		if l2 != nil {
			n2, l2 = l2.Val, l2.Next
		}
		num := n1 + n2 + carry
		carry = num / 10
		cur.Next = &ListNode{num % 10, nil}
		cur = cur.Next
	}
	return head.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
