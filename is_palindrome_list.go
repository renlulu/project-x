package project_x

import "container/list"

func isPalindromeList(head *ListNode) bool {
	stack := list.New()
	if head == nil || head.Next == nil {
		return true
	}

	fast := head
	slow := head

	stack.PushFront(slow.Val)

	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		stack.PushFront(slow.Val)
	}

	if fast.Next == nil {
		front := stack.Front()
		if front != nil {
			stack.Remove(front)
		}
	}

	for slow.Next != nil {
		slow = slow.Next
		f := stack.Front()
		if f.Value.(int) != slow.Val {
			return false
		}
		stack.Remove(f)
	}

	return true
}
