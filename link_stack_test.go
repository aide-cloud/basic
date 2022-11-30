package basic

import (
	"testing"
)

func Test_linkStack_Push(t *testing.T) {
	st := NewLinkStack(10)
	for i := 0; i < 10; i++ {
		if !st.Push(i) {
			t.Error("push error")
		}
	}

	if st.Size() != 10 {
		t.Error("size error")
	}

	if st.Push(10) {
		t.Error("push error")
	}

	for i := 9; i >= 0; i-- {
		if st.Top() != i {
			t.Error("top error")
		}

		if st.Size() != i+1 {
			t.Error("size error")
		}

		if st.Pop() != i {
			t.Error("pop error")
		}
	}
}

func Test_linkStack_Values(t *testing.T) {
	st := NewLinkStack(10)
	for i := 9; i >= 0; i-- {
		if !st.Push(i) {
			t.Error("push error")
		}
	}

	data := st.Values()
	for i := 9; i >= 0; i-- {
		if data[i] != i {
			t.Error("values error")
		}
	}
}
