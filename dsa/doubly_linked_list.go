package main

type DLLNode struct {
	Value any
	prev  *DLLNode
	next  *DLLNode
}

type DoublyLinkedList struct {
	Length int
	head   *DLLNode
	tail   *DLLNode
}

func NewDoublDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{
		Length: 0,
		head:   nil,
		tail:   nil,
	}
}

func (list *DoublyLinkedList) Prepend(item any) {
	list.Length++

	node := &DLLNode{Value: item}

	if list.head == nil {
		list.head = node
		list.tail = node

		return
	}

	node.next = list.head
	list.head.prev = node
	list.head = node
}

func (list *DoublyLinkedList) InsertAt(item any, idx int) error {
	var err error

	if idx > list.Length {
		return err
	}

	if idx == list.Length {
		list.Append(item)

		return nil
	} else if idx == 0 {
		list.Prepend(item)

		return nil
	}

	list.Length++

	curr := list.getAt(idx)
	node := &DLLNode{Value: item}

	node.next = curr
	node.prev = curr.prev
	curr.prev = node

	if node.prev != nil {
		node.prev.next = curr
	}

	return nil
}

func (list *DoublyLinkedList) Append(item any) {
	list.Length++

	node := &DLLNode{Value: item}

	if list.tail == nil {
		list.head = node
		list.tail = node

		return
	}

	node.prev = list.tail
	list.tail.next = node
	list.tail = node
}

func (list *DoublyLinkedList) Remove(item any) any {
	curr := list.head

	for i := 0; curr != nil && i < list.Length; i++ {
		if curr.Value == item {
			break
		}

		curr = curr.next
	}

	if curr == nil {
		return nil
	}

	return list.removeNode(curr)
}

func (list *DoublyLinkedList) Get(idx int) any {
	return list.getAt(idx).Value
}

func (list *DoublyLinkedList) RemoveAt(idx int) any {
	node := list.getAt(idx)

	if node == nil {
		return nil
	}

	return list.removeNode(node)
}

func (list *DoublyLinkedList) getAt(idx int) *DLLNode {
	curr := list.head

	for i := 0; curr != nil && i < idx; i++ {
		curr = curr.next
	}

	return curr
}

func (list *DoublyLinkedList) removeNode(node *DLLNode) any {
	list.Length--

	if list.Length == 0 {
		out := list.head.Value
		list.head = nil
		list.tail = nil

		return out
	}

	if node.prev != nil {
		node.prev.next = node.next
	}

	if node.next != nil {
		node.next.prev = node.prev
	}

	if node == list.head {
		list.head = node.next
	}

	if node == list.tail {
		list.tail = node.prev
	}

	node.prev = nil
	node.next = nil

	return node.Value
}
