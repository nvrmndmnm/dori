package main

type SNode struct {
	Value any
	prev  *SNode
}

type Stack struct {
	Length int
	head   *SNode
}

func NewStack() *Stack {
	return &Stack{
		Length: 0,
		head:   nil,
	}
}

func (s *Stack) Push(item any) {
	node := &SNode{
		Value: item,
	}

	s.Length ++
	if s.head == nil {
		s.head = node
	}

	node.prev = s.head
	s.head = node
}

func (s *Stack) Pop() any{
	if s.Length == 0 {
		return nil
	}
	
	s.Length--
	head := s.head
	s.head = s.head.prev

	return head.Value
}

func (s *Stack) Peek() any{
	return s.head.Value
}