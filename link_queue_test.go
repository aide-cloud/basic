package basic

import "testing"

func Test_linkedQueue_Push(t *testing.T) {
	qu := NewLinkedQueue(10)
	for i := 0; i < 10; i++ {
		qu.Push(i)
	}

	if qu.Size() != 10 {
		t.Errorf("size error, want: %d, got: %d", 10, qu.Size())
	}

	for i := 0; i < 10; i++ {
		if qu.Pop() != i {
			t.Errorf("pop error, want: %d, got: %d", i, qu.Pop())
		}
	}

	if qu.Size() != 0 {
		t.Errorf("size error, want: %d, got: %d", 0, qu.Size())
	}
}

func Test_linkedQueue_String(t *testing.T) {
	qu := NewLinkedQueue(10)
	for i := 0; i < 10; i++ {
		qu.Push(i)
	}

	if qu.String() != "head [0, 1, 2, 3, 4, 5, 6, 7, 8, 9] tail" {
		t.Errorf("string error, want: %s, got: %s", "0 1 2 3 4 5 6 7 8 9 ", qu.String())
	}
}
