package basic

import (
	"sync"
)

type (
	SortFun[T Any] func(a []T, less SortSliceLessFunc[T])

	Tuple[V Equatable] struct {
		m    map[V]bool
		sort SortFun[V]

		lock sync.RWMutex
	}

	TupleOption[V Equatable] func(*Tuple[V])
)

// WithSort sets the sort function for the Tuple.
func WithSort[V Equatable](sort SortFun[V]) TupleOption[V] {
	return func(t *Tuple[V]) {
		t.sort = sort
	}
}

// NewTuple returns a new Tuple.
func NewTuple[V Equatable](opts ...TupleOption[V]) *Tuple[V] {
	tu := &Tuple[V]{
		m:    make(map[V]bool),
		sort: QuickSort[V],
	}
	for _, opt := range opts {
		opt(tu)
	}
	return tu
}

// Add adds a value to the Tuple.
func (t *Tuple[V]) Add(val ...V) {
	t.lock.Lock()
	defer t.lock.Unlock()

	for _, v := range val {
		t.m[v] = true
	}
}

// Remove removes a value from the Tuple.
func (t *Tuple[V]) Remove(val ...V) {
	t.lock.Lock()
	defer t.lock.Unlock()

	for _, v := range val {
		delete(t.m, v)
	}
}

// Has returns true if the Tuple contains the value.
func (t *Tuple[V]) Has(v V) bool {
	t.lock.RLock()
	defer t.lock.RUnlock()

	_, ok := t.m[v]
	return ok
}

// Len returns the number of values in the Tuple.
func (t *Tuple[V]) Len() int {
	t.lock.RLock()
	defer t.lock.RUnlock()

	return len(t.m)
}

// Clear removes all values from the Tuple.
func (t *Tuple[V]) Clear() {
	t.lock.Lock()
	defer t.lock.Unlock()

	t.m = make(map[V]bool)
}

// Values returns a slice of all values in the Tuple.
func (t *Tuple[V]) Values() []V {
	t.lock.RLock()
	defer t.lock.RUnlock()

	values := make([]V, 0, len(t.m))
	for v := range t.m {
		values = append(values, v)
	}
	return values
}

// SortValues sorts the values in the Tuple using the provided less function.
func (t *Tuple[V]) SortValues(less SortSliceLessFunc[V]) []V {
	res := t.Values()
	t.sort(res, less)
	return res
}

// Range calls f sequentially for each value present in the Tuple.
func (t *Tuple[V]) Range(f func(v V)) {
	t.lock.RLock()
	defer t.lock.RUnlock()

	for v := range t.m {
		f(v)
	}
}

// Copy returns a copy of the Tuple.
func (t *Tuple[V]) Copy() *Tuple[V] {
	t.lock.RLock()
	defer t.lock.RUnlock()

	c := NewTuple[V]()
	for v := range t.m {
		c.Add(v)
	}
	return c
}

// Equal returns true if the Tuple is equal to other.
func (t *Tuple[V]) Equal(other *Tuple[V]) bool {
	t.lock.RLock()
	defer t.lock.RUnlock()

	if t.Len() != other.Len() {
		return false
	}

	for v := range t.m {
		if !other.Has(v) {
			return false
		}
	}
	return true
}

// Union returns a new Tuple with the values from t and other.
func (t *Tuple[V]) Union(other *Tuple[V]) *Tuple[V] {
	t.lock.RLock()
	defer t.lock.RUnlock()

	c := t.Copy()
	other.Range(func(v V) {
		c.Add(v)
	})
	return c
}

// Intersect returns a new Tuple with the values that are in both t and other.
func (t *Tuple[V]) Intersect(other *Tuple[V]) *Tuple[V] {
	t.lock.RLock()
	defer t.lock.RUnlock()

	c := NewTuple[V]()
	t.Range(func(v V) {
		if other.Has(v) {
			c.Add(v)
		}
	})
	return c
}

// Difference returns a new Tuple with the values that are in t but not in other.
func (t *Tuple[V]) Difference(other *Tuple[V]) *Tuple[V] {
	t.lock.RLock()
	defer t.lock.RUnlock()

	c := NewTuple[V]()
	t.Range(func(v V) {
		if !other.Has(v) {
			c.Add(v)
		}
	})
	return c
}

// IsSubset returns true if other is a subset of t.
func (t *Tuple[V]) IsSubset(other *Tuple[V]) bool {
	t.lock.RLock()
	defer t.lock.RUnlock()

	if t.Len() > other.Len() {
		return false
	}

	for v := range t.m {
		if !other.Has(v) {
			return false
		}
	}
	return true
}

// IsSuperset returns true if other is a superset of t.
func (t *Tuple[V]) IsSuperset(other *Tuple[V]) bool {
	t.lock.RLock()
	defer t.lock.RUnlock()

	if t.Len() < other.Len() {
		return false
	}

	for v := range other.m {
		if !t.Has(v) {
			return false
		}
	}
	return true
}
