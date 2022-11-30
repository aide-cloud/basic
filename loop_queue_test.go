package basic

import "testing"

func TestLoopQueue_Push(t *testing.T) {
	q := NewLoopQueue(10)
	for i := 0; i < 10; i++ {
		q.Push(i)
	}

	if q.Size() != 10 {
		t.Error("size error")
	}

	for i := 0; i < 10; i++ {
		if q.Pop() != i {
			t.Error("pop error")
		}
	}

	if !q.IsEmpty() {
		t.Error("empty error")
	}
}

func TestLoopQueue_String(t *testing.T) {
	q := NewLoopQueue(10)
	for i := 0; i < 10; i++ {
		q.Push(i)
	}

	if q.String() != "head [0, 1, 2, 3, 4, 5, 6, 7, 8, 9] tail" {
		t.Error("string error: want head [0, 1, 2, 3, 4, 5, 6, 7, 8, 9] tail, got", q.String())
	}
}
