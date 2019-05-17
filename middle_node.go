package project_x

func middleNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return head
	}

	if head.Next.Next == nil {
		return head.Next
	}

	slow := head
	fast := head

	for (fast.Next != nil && fast.Next.Next != nil) {
		fast = fast.Next.Next
		slow = slow.Next
	}

	if (fast.Next == nil) {
		return slow
	} else {
		return slow.Next
	}
}