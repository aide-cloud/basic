package basic

import (
	"bytes"
	"fmt"
)

// arrayStack is a stack implemented by array.
type arrayStack struct {
	data   []Any
	length int
}

func (s *arrayStack) Values() []Any {
	return s.data
}

// Top returns the top element of the stack.
func (s *arrayStack) Top() (element Any) {
	if s.IsEmpty() {
		return
	}
	return s.data[len(s.data)-1]
}

// IsFull returns true if the stack is full.
func (s *arrayStack) IsFull() bool {
	return s.length == len(s.data)
}

// String returns the string of the stack.
func (s *arrayStack) String() string {
	if s.IsEmpty() {
		return "empty stack"
	}

	var buffer bytes.Buffer
	buffer.WriteString("top [")
	for i := s.length - 1; i >= 0; i-- {
		buffer.WriteString(fmt.Sprintf("%v", s.data[i]))
		if i != 0 {
			buffer.WriteString(" ")
		}
	}
	buffer.WriteString("] bottom")
	return buffer.String()
}

// NewArrayStack returns a new arrayStack.
func NewArrayStack(length int) *arrayStack {
	return &arrayStack{
		data:   make([]Any, 0, length),
		length: length,
	}
}

// Push pushes an element into the stack.
func (s *arrayStack) Push(element Any) bool {
	if s.length == s.Size() {
		return false
	}
	s.data = append(s.data, element)
	return true
}

// Pop pops an element from the stack.
func (s *arrayStack) Pop() (v Any) {
	if s.IsEmpty() {
		return
	}
	v = s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return
}

// IsEmpty returns true if the stack is empty.
func (s *arrayStack) IsEmpty() bool {
	return len(s.data) == 0
}

// Size returns the size of the stack.
func (s *arrayStack) Size() int {
	return len(s.data)
}

// Clear clears the stack.
func (s *arrayStack) Clear() {
	s.data = make([]Any, 0, s.length)
}
