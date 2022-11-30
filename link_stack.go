package basic

import (
	"bytes"
	"fmt"
)

type linkStack struct {
	head   *linkedNode
	size   int
	length int
}

func (s *linkStack) Values() []Any {
	data := make([]Any, 0, s.size)
	for node := s.head; node != nil; node = node.next {
		data = append(data, node.value)
	}
	return data
}

// Top returns the top element of the stack.
func (s *linkStack) Top() (element Any) {
	if s.head == nil || s.size == 0 {
		return
	}
	return s.head.value
}

// IsFull returns true if the stack is full.
func (s *linkStack) IsFull() bool {
	return s.size == s.length
}

// Clear clears the stack.
func (s *linkStack) Clear() {
	s.head = nil
	s.size = 0
}

// NewLinkStack returns a new linkStack.
func NewLinkStack(length int) *linkStack {
	return &linkStack{
		length: length,
	}
}

// Push pushes an element into the stack.
func (s *linkStack) Push(element Any) bool {
	if s.length == 0 || s.length == s.size {
		return false
	}
	node := &linkedNode{
		value: element,
	}
	node.next = s.head
	s.head = node
	s.size++
	return true
}

// Pop pops an element from the stack.
func (s *linkStack) Pop() (element Any) {
	if s.head == nil || s.size == 0 {
		return
	}
	element = s.head.value
	s.head = s.head.next
	s.size--
	return
}

// IsEmpty returns true if the stack is empty.
func (s *linkStack) IsEmpty() bool {
	return s.head == nil
}

// Size returns the size of the stack.
func (s *linkStack) Size() int {
	return s.size
}

// String returns the string of the stack.
func (s *linkStack) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("Stack: top ")
	node := s.head
	for node != nil {
		buffer.WriteString(fmt.Sprint(node.value))
		buffer.WriteString("->")
		node = node.next
	}
	buffer.WriteString("nil")
	return buffer.String()
}
