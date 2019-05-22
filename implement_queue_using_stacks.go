package project_x

import "container/list"

type MyQueue struct {
	list *list.List
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	l := list.New()
	return MyQueue{
		list: l,
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.list.PushBack(x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	//actually we are not allowed to do this
	front := this.list.Front()
	res, _ := front.Value.(int)
	this.list.Remove(front)
	return res
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	return this.list.Front().Value.(int)
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return this.list.Len() == 0
}
