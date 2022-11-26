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
		if !InSlice(a, v) {
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
		if !InSlice(b, v) {
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
func InSlice[T Comparable](b []T, a T) bool {
	for _, v := range b {
		if a == v {
			return true
		}
	}

	return false
}
