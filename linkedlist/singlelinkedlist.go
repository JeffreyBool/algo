/**
  * Author: JeffreyBool
  * Date: 2019/4/21
  * Time: 17:39
  * Software: GoLand
*/

package linkedlist

type ListNode struct {
	next   *ListNode //后续结点 next
	value interface{}
}

type LinkedList struct {
	head   *ListNode
	size int
}

//初始化 node
func NewListNode(v interface{}) *ListNode {
	return &ListNode{
		next:nil,
		value:v,
	}
}

//获取node next
func (node *ListNode) GetNext() *ListNode {
	return node.next
}

//获取node value
func (node *ListNode) GetValue() interface{} {
	return node.value
}

//初始化链表
func NewLinkedList() *LinkedList {
	return &LinkedList{head:NewListNode(nil)}
}

//判断单链表是否为空
func (linked *LinkedList) IsEmpty() bool {
	return linked.head == nil
}

//单链表长度
func (linked *LinkedList) Count () int {
	return linked.size
}

//获取头部结点
func (linked *LinkedList) GetHead() *ListNode {
	return linked.head
}

//从头部添加元素
func (linked *LinkedList) AddFirst(v interface{}) {
	node := NewListNode(v)
	node.next = linked.head
	linked.head = node
	linked.size ++
}

//从尾部添加元素
func (linked *LinkedList) AddLast(v interface{}) {
	node := NewListNode(v)
	if linked.IsEmpty() {
		linked.head = node
	}else{
		cur := linked.head
		for cur.next != nil{
			cur = cur.next
		}

		cur.next = node
	}
	linked.size ++
}

//根据索引添加元素
func (linked *LinkedList) Add(index uint, v interface{})  {
	if index < 0 {
		linked.AddFirst(v)
	}else if int(index) > linked.size {
		linked.AddLast(v)
	}else{
		prev := linked.head
		for i := 0; i < int(index) -1; i++ {
			prev = prev.next
		}

		node := NewListNode(v)
		node.next = prev.next
		prev.next = node
		linked.size ++
	}
}

//删除指定元素
func (linked *LinkedList) Remove(v interface{})  {
	prev := linked.head
	if prev.value == v {
		linked.head = prev.next
 	}else{
 		for prev.next != nil{
 			if prev.value == v {
 				prev.next = prev.next.next
 				break
			}else{
				prev = prev.next
			}
		}
	}
}
