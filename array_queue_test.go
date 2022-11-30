package basic

import "testing"

func Test_arrayQueue_Push(t *testing.T) {
	qu := NewArrayQueue(10)
	for i := 0; i < 10; i++ {
		if !qu.Push(i) {
			t.Error("push error")
		}
	}

	if qu.Size() != 10 {
		t.Error("push error")
	}

	for i := 0; i < 10; i++ {
		if qu.Pop() != i {
			t.Error("error")
		}
	}

	if !qu.IsEmpty() {
		t.Error("error")
	}
}

func Test_arrayQueue_String(t *testing.T) {
	qu := NewArrayQueue(10)
	for i := 0; i < 10; i++ {
		qu.Push(i)
	}

	if qu.String() != "head [0, 1, 2, 3, 4, 5, 6, 7, 8, 9] tail" {
		t.Error("string error, want: head [0, 1, 2, 3, 4, 5, 6, 7, 8, 9] tail, got: ", qu.String())
	}
}
