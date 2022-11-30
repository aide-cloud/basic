package basic

import (
	"bytes"
	"fmt"
)

type linkedNode struct {
	value Any
	next  *linkedNode
}

type linkedQueue struct {
	head, tail   *linkedNode
	length, size int
}

func (q *linkedQueue) Values() []Any {
	data := make([]Any, 0, q.size)
	for node := q.head; node != nil; node = node.next {
		data = append(data, node)
	}

	return data
}

func NewLinkedQueue(length int) *linkedQueue {
	return &linkedQueue{
		length: length,
	}
}

func (q *linkedQueue) IsFull() bool {
	return q.size == q.length
}

func (q *linkedQueue) Push(element Any) bool {
	if q.IsEmpty() || q.IsFull() {
		return false
	}
	node := &linkedNode{
		value: element,
	}
	if q.tail == nil {
		q.head = node
	} else {
		q.tail.next = node
	}
	q.tail = node
	q.size++
	return true
}

func (q *linkedQueue) Pop() (element Any) {
	if q.head == nil || q.size == 0 {
		return
	}
	element = q.head.value
	q.head = q.head.next
	if q.head == nil {
		q.tail = nil
	}
	q.size--
	return
}

func (q *linkedQueue) IsEmpty() bool {
	return q.length == 0
}

func (q *linkedQueue) Size() int {
	return q.size
}

func (q *linkedQueue) Clear() {
	q.head, q.tail = nil, nil
	q.size = 0
}

func (q *linkedQueue) String() string {
	if q.IsEmpty() {
		return "empty queue"
	}
	var buffer bytes.Buffer
	buffer.WriteString("head [")
	for node := q.head; node != nil; node = node.next {
		buffer.WriteString(fmt.Sprintf("%v", node.value))
		if node.next != nil {
			buffer.WriteString(", ")
		}
	}
	buffer.WriteString("] tail")
	return buffer.String()
}
