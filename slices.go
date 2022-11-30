package basic

import (
	"fmt"
)

// SliceToNumber type a slice to b slice
func SliceToNumber[T, I Number](i []I) []T {
	out := make([]T, len(i))
	for k, v := range i {
		out[k] = T(v)
	}

	return out
}

// SliceToString type a slice to b slice
func SliceToString[T any](i []T) []string {
	out := make([]string, len(i))
	for k, v := range i {
		out[k] = fmt.Sprintf("%v", v)
	}

	return out
}

// SliceToBool type a slice to b slice
func SliceToBool[T any](i []T) []bool {
	out := make([]bool, len(i))
	for k, v := range i {
		out[k] = !IsZero(v)
	}

	return out
}

// IsSubset check if b is a subset of a
func IsSubset[T Comparable](a []T, b ...T) bool {
	if len(a) == 0 {
		return false
	}

	for _, v := range b {
		if !InSlice(a, v, func(i, j T) bool {
			return i < j
		}) {
			return false
		}
	}

	return true
}

// DiffSlice get the difference of two slices
func DiffSlice[T Comparable](a []T, b []T) []T {
	maxLen := len(a)
	if len(b) > maxLen {
		maxLen = len(b)
	}
	out := make([]T, 0, maxLen)
	for _, v := range a {
		if !InSlice(b, v, func(i, j T) bool {
			return i < j
		}) {
			out = append(out, v)
		}
	}

	return out
}

type RemoveSliceFunc[T any] func(i, v T) bool

// RemoveAnySlice remove a from b
func RemoveAnySlice[T any](b []T, a T, opts ...RemoveSliceFunc[T]) []T {
	out := make([]T, 0, len(b))

	for index := range b {
		// 判断地址
		if &a == &b[index] {
			continue
		}

		if len(opts) > 0 {
			if opts[0](a, b[index]) {
				continue
			}
		}

		out = append(out, b[index])
	}

	return out
}

// RemoveComparableSlice remove a from b
func RemoveComparableSlice[T Comparable](b []T, a T) []T {
	out := make([]T, 0, len(b))
	for _, v := range b {
		if a != v {
			out = append(out, v)
		}
	}

	return out
}

// InSlice check if a in b
func InSlice[T Comparable](b []T, a T, less func(i, j T) bool) bool {
	if len(b) == 0 {
		return false
	}

	QuickSort(b, less)

	return BinarySearch(b, a) != -1
}

// QuickSort sort a slice
func QuickSort[T any](a []T, less func(i, j T) bool) {
	if len(a) <= 1 {
		return
	}

	mid, i := a[0], 1
	head, tail := 0, len(a)-1

	for head < tail {
		if less(mid, a[i]) {
			a[i], a[tail] = a[tail], a[i]
			tail--
		} else {
			a[i], a[head] = a[head], a[i]
			head++
			i++
		}
	}

	QuickSort(a[:head], less)
	QuickSort(a[head+1:], less)
}

// MergeSlice merge two slice
func MergeSlice[T any](a []T, b []T) []T {
	out := make([]T, len(a)+len(b))
	copy(out, a)
	copy(out[len(a):], b)

	return out
}

// BinarySearch 二分查找
func BinarySearch[T Comparable](a []T, x T) int {
	less := func(i, j T) bool {
		return i < j
	}
	low, high := 0, len(a)-1
	for low <= high {
		mid := low + (high-low)>>1
		if less(x, a[mid]) {
			high = mid - 1
		} else if less(a[mid], x) {
			low = mid + 1
		} else {
			return mid
		}
	}

	return -1
}

// InsertSort 插入排序
func InsertSort[T Comparable](a []T, less func(i, j T) bool) {
	for i := 1; i < len(a); i++ {
		for j := i; j > 0 && less(a[j], a[j-1]); j-- {
			a[j], a[j-1] = a[j-1], a[j]
		}
	}
}
