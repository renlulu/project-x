package project_x

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var previous *ListNode
	current := head

	for current != nil {
		next := current.Next

		current.Next = previous

		previous = current
		current = next
	}

	return previous

}
