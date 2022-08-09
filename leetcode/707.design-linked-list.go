/*
 * @lc app=leetcode id=707 lang=golang
 *
 * [707] Design Linked List
 */

package leetcode

// @lc code=start
type MyLinkedList struct {
    val int
	next *MyLinkedList
}

func Constructor() MyLinkedList {
    return MyLinkedList{}
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 {
		return -1
	}
	cur := this.next
    for index > 0 && cur != nil {
        cur = cur.next
        index--
    }
    if index > 0 || cur == nil{
        return -1
    }
	return cur.val
}

func (this *MyLinkedList) AddAtHead(val int)  {
    newNode := &MyLinkedList{val,this.next}
    this.next = newNode
}

func (this *MyLinkedList) AddAtTail(val int)  {
    cur := this
	for cur.next != nil {
		cur = cur.next
	}
	newNode := &MyLinkedList{val, nil}
	cur.next = newNode
}

func (this *MyLinkedList) AddAtIndex(index int, val int)  {
    if index < 0 {
		return
	}

	cur := this
	for index > 0 && cur.next != nil{
		cur = cur.next
		index--
	}
	if index > 0 {
		return
	}
	newNode := &MyLinkedList{val, cur.next}
	cur.next = newNode
}

func (this *MyLinkedList) DeleteAtIndex(index int)  {
    if index < 0 {
		return
	}
	cur := this
	for index > 0 && cur.next != nil{
		cur = cur.next
		index--
	}
	if index > 0 || cur.next == nil {
        return
    }
    cur.next = cur.next.next
}


/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
// @lc code=end

