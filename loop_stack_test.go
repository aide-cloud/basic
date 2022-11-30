package basic

import (
	"testing"
)

func Test_loopStack_Push(t *testing.T) {
	st := NewLoopStack(10)
	for i := 0; i < 10; i++ {
		if !st.Push(i) {
			t.Errorf("push %d failed", i)
		}
	}

	if st.Push(10) {
		t.Errorf("push 10 should failed")
	}

	if st.Size() != 10 {
		t.Errorf("size should be 10")
	}

	for i := 9; i >= 0; i-- {
		if st.Top() != i {
			t.Errorf("top should be %d", i)
		}

		if st.Size() != i+1 {
			t.Errorf("size should be %d", i+1)
		}

		if st.Pop() != i {
			t.Errorf("pop %d failed", i)
		}
	}
}

func Test_loopStack_Clear(t *testing.T) {
	st := NewLoopStack(10)
	for i := 0; i < 10; i++ {
		if !st.Push(i) {
			t.Errorf("push %d failed", i)
		}
	}

	st.Clear()

	if !st.IsEmpty() {
		t.Errorf("stack should be empty")
	}
}

func Test_loopStack_Values(t *testing.T) {
	st := NewLoopStack(10)
	for i := 0; i < 10; i++ {
		if !st.Push(i) {
			t.Errorf("push %d failed", i)
		}
	}

	values := st.Values()
	for i := 9; i >= 0; i-- {
		if values[i] != i {
			t.Errorf("values[%d] should be %d", i, i)
		}
	}
}

func Test_loopStack_String(t *testing.T) {
	st := NewLoopStack(10)
	for i := 0; i < 10; i++ {
		if !st.Push(i) {
			t.Errorf("push %d failed", i)
		}
	}

	if st.String() != "top [9, 8, 7, 6, 5, 4, 3, 2, 1, 0] bottom" {
		t.Errorf("string error: want [9, 8, 7, 6, 5, 4, 3, 2, 1, 0] bottom, got %s", st.String())
	}
}
