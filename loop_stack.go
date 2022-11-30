package basic

import (
	"bytes"
	"fmt"
)

// loopStack is a stack based on loop array.
type loopStack struct {
	data   []Any
	top    int
	length int
}

func (s *loopStack) Top() (element Any) {
	return s.data[s.top]
}

func (s *loopStack) String() string {
	if s.IsEmpty() {
		return "empty stack"
	}

	var buffer bytes.Buffer
	buffer.WriteString("top [")
	for i := s.top; i >= 0; i-- {
		buffer.WriteString(fmt.Sprintf("%v", s.data[i]))
		if i != 0 {
			buffer.WriteString(", ")
		}
	}
	buffer.WriteString("] bottom")
	return buffer.String()
}

// NewLoopStack returns a new loopStack.
func NewLoopStack(length int) *loopStack {
	return &loopStack{
		data:   make([]Any, length),
		top:    -1,
		length: length,
	}
}

// Push pushes an element into the stack.
func (s *loopStack) Push(v Any) bool {
	if s.IsFull() {
		return false
	}

	s.top = (s.top + 1) % s.length
	s.data[s.top] = v
	return true
}

// Pop pops an element from the stack.
func (s *loopStack) Pop() (v Any) {
	if s.IsEmpty() {
		return
	}

	v = s.data[s.top]
	s.top = (s.top - 1 + s.length) % s.length
	return
}

// IsEmpty returns true if the stack is empty.
func (s *loopStack) IsEmpty() bool {
	return s.top == -1
}

// IsFull returns true if the stack is full.
func (s *loopStack) IsFull() bool {
	return !s.IsEmpty() && (s.top+1)%s.length == 0
}

// Size returns the size of the stack.
func (s *loopStack) Size() int {
	if s.IsEmpty() {
		return 0
	}

	if s.IsFull() {
		return s.length
	}

	return s.top + 1
}

// Clear clears the stack.
func (s *loopStack) Clear() {
	s.top = -1
	s.data = make([]Any, s.length)
}

// Values returns all elements in the stack.
func (s *loopStack) Values() []Any {
	if s.IsEmpty() {
		return nil
	}

	if s.IsFull() {
		return s.data
	}

	return s.data[:s.top+1]
}
