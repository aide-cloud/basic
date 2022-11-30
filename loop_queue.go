package basic

import (
	"bytes"
	"fmt"
)

type loopQueue struct {
	head, tail int
	elements   []Any
	length     int
}

func (q *loopQueue) Values() []Any {
	return q.elements[q.head:q.tail]
}

func (q *loopQueue) IsFull() bool {
	return (q.tail+1)%len(q.elements) == q.head
}

func NewLoopQueue(length int) *loopQueue {
	return &loopQueue{
		elements: make([]Any, length+1),
		length:   length + 1,
	}
}

func (q *loopQueue) Push(element Any) bool {
	if (q.tail+1)%q.length == q.head {
		return false
	}
	q.elements[q.tail] = element
	q.tail = (q.tail + 1) % q.length
	return true
}

func (q *loopQueue) Pop() (element Any) {
	if q.head == q.tail {
		return
	}
	element = q.elements[q.head]
	q.head = (q.head + 1) % q.length
	return
}

func (q *loopQueue) IsEmpty() bool {
	return q.head == q.tail
}

func (q *loopQueue) Size() int {
	return (q.tail - q.head + q.length) % q.length
}

func (q *loopQueue) Clear() {
	q.elements = make([]Any, q.length)
	q.head, q.tail = 0, 0
}

func (q *loopQueue) String() string {
	if q.IsEmpty() {
		return "empty queue"
	}
	var buffer bytes.Buffer
	buffer.WriteString("head [")
	for i := q.head; i != q.tail; i = (i + 1) % q.length {
		buffer.WriteString(fmt.Sprintf("%v", q.elements[i]))
		if (i+1)%q.length != q.tail {
			buffer.WriteString(", ")
		}
	}
	buffer.WriteString("] tail")
	return buffer.String()
}
