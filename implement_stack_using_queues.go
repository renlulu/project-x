package project_x
//
//import "container/list"
//
//type MyStack struct {
//	list *list.List
//}
//
///** Initialize your data structure here. */
//func Constructor() MyStack {
//	l := list.New()
//	return MyStack{
//		list: l,
//	}
//}
//
///** Push element x onto stack. */
//func (this *MyStack) Push(x int) {
//	this.list.PushBack(x)
//	c := this.list.Len() - 1
//	for c > 0 {
//		e := this.list.Front()
//		this.list.Remove(e)
//		this.list.PushBack(e.Value.(int))
//		c--
//	}
//}
//
///** Removes the element on top of the stack and returns that element. */
//func (this *MyStack) Pop() int {
//	e := this.list.Front()
//	res := e.Value.(int)
//	this.list.Remove(e)
//	return res
//}
//
///** Get the top element. */
//func (this *MyStack) Top() int {
//	return this.list.Front().Value.(int)
//}
//
///** Returns whether the stack is empty. */
//func (this *MyStack) Empty() bool {
//	return this.list.Len() == 0
//}
