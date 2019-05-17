package project_x

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode {
		Val: 0,
		Next: head,
	}

	length := 0
	first := head
	for first != nil {
		length++
		first = first.Next
	}

	length = length - n
	first = dummy

	for length > 0 {
		length--
		first = first.Next
	}

	if first.Next.Next == nil {
		first.Next = nil
	} else {
		first.Next = first.Next.Next
	}

	return dummy.Next
}