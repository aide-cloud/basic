package basic

import (
	"bytes"
	"fmt"
)

type arrayQueue struct {
	elements []Any
	length   int
}

func (q *arrayQueue) Values() []Any {
	return q.elements
}

func (q *arrayQueue) IsFull() bool {
	return q.length == len(q.elements)
}

func NewArrayQueue(length int) *arrayQueue {
	return &arrayQueue{
		elements: make([]Any, 0, length),
		length:   length,
	}
}

func (q *arrayQueue) Push(element Any) bool {
	if q.IsFull() {
		return false
	}

	q.elements = append(q.elements, element)

	return true
}

func (q *arrayQueue) Pop() (element Any) {
	if q.IsEmpty() {
		return
	}

	element = q.elements[0]
	q.elements = q.elements[1:q.Size()]
	return
}

func (q *arrayQueue) IsEmpty() bool {
	return q.Size() == 0
}

func (q *arrayQueue) Size() int {
	return len(q.elements)
}

func (q *arrayQueue) Clear() {
	q.elements = make([]Any, 0, q.length)
}

func (q *arrayQueue) String() string {
	if q.IsEmpty() {
		return "empty queue"
	}
	var buffer bytes.Buffer
	buffer.WriteString("head [")
	for i := 0; i < q.length; i++ {
		buffer.WriteString(fmt.Sprintf("%v", q.elements[i]))
		if i != q.length-1 {
			buffer.WriteString(", ")
		}
	}
	buffer.WriteString("] tail")
	return buffer.String()
}
