package project_x

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	current_l1 := l1
	current_l2 := l2

	N := 0
	head := &ListNode{
		Val:  0,
		Next: nil,
	}
	result_current := head

	for (current_l1 != nil && current_l2 != nil) {
		sum := current_l1.Val + current_l2.Val + N/10
		m := sum % 10

		newNode := &ListNode{
			Val:  m,
			Next: nil,
		}

		result_current.Next = newNode
		result_current = newNode

		N = sum - m

		current_l1 = current_l1.Next
		current_l2 = current_l2.Next
	}

	for (current_l1 != nil) {
		sum := current_l1.Val + N/10
		m := sum % 10

		newNode := &ListNode{
			Val:  m,
			Next: nil,
		}

		result_current.Next = newNode
		result_current = newNode

		N = sum - m
		current_l1 = current_l1.Next
	}

	for (current_l2 != nil) {
		sum := current_l2.Val + N/10
		m := sum % 10

		newNode := &ListNode{
			Val:  m,
			Next: nil,
		}

		result_current.Next = newNode
		result_current = newNode

		N = sum - m
		current_l2 = current_l2.Next
	}

	if N != 0 {
		newNode := &ListNode{
			Val:  N / 10,
			Next: nil,
		}

		result_current.Next = newNode
		result_current = newNode

	}

	return head.Next
}
