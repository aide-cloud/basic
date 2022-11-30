package basic

import "testing"

func TestArrayStack_Push(t *testing.T) {
	st := NewArrayStack(10)
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
