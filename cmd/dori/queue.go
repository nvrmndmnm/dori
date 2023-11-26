package main

type QNode struct {
	Value any
	next  *QNode
}

type Queue struct {
	Length int
	head   *QNode
	tail   *QNode
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Enqueue(item any) {
	q.Length++
	node := &QNode{Value: item}
	
	if q.tail == nil {
		q.tail = node
		q.head = q.tail
		return
	}

	q.tail.next = node
	q.tail = node 

}

func (q *Queue) Deque() any {
	if q.head == nil {
		return nil
	}

	q.Length--
	head := q.head
	q.head = q.head.next
	head.next = nil

	if q.Length == 0 {
		q.tail = nil
	}

	return head.Value
}

func (q *Queue) Peek() any {
	if q.head == nil {
		return nil
	}

	return q.head.Value
}
