package basic

import (
	"sort"
	"testing"
)

func TestTuple_Add(t1 *testing.T) {
	t := NewTuple[int]()
	t.Add(1)
	t.Add(2)
	t.Add(3)
	t.Add(4)
	if t.Len() != 4 {
		t1.Error("TestTuple_Add failed")
	}

	if !t.Has(1) && !t.Has(2) && !t.Has(3) && !t.Has(4) {
		t1.Error("TestTuple_Add failed")
	}
}

func TestTuple_Remove(t1 *testing.T) {
	t := NewTuple[int]()
	t.Add(1)
	t.Add(2)
	t.Add(3)
	t.Add(4)
	t.Remove(1)
	t.Remove(2)
	t.Remove(3)
	t.Remove(4)
	if t.Len() != 0 {
		t1.Error("TestTuple_Remove failed")
	}

	if t.Has(1) && t.Has(2) && t.Has(3) && t.Has(4) {
		t1.Error("TestTuple_Remove failed")
	}
}

func TestTuple_Clear(t1 *testing.T) {
	t := NewTuple[int]()
	t.Add(1)
	t.Add(2)
	t.Add(3)
	t.Add(4)
	t.Clear()
	if t.Len() != 0 {
		t1.Error("TestTuple_Clear failed")
	}

	if t.Has(1) && t.Has(2) && t.Has(3) && t.Has(4) {
		t1.Error("TestTuple_Clear failed")
	}
}

func TestTuple_Values(t1 *testing.T) {
	t := NewTuple[int]()
	t.Add(1)
	t.Add(2)
	t.Add(3)
	t.Add(4)
	values := t.Values()
	if len(values) != 4 {
		t1.Error("TestTuple_Values failed")
	}

	sort.Ints(values)

	if values[0] != 1 && values[1] != 2 && values[2] != 3 && values[3] != 4 {
		t1.Error("TestTuple_Values failed")
	}
}

func TestTuple_Range(t1 *testing.T) {
	t := NewTuple[int]()
	t.Add(1)
	t.Add(2)
	t.Add(3)
	t.Add(4)
	count := 0
	t.Range(func(v int) {
		count++
	})
	if count != 4 {
		t1.Error("TestTuple_Range failed")
	}
}
