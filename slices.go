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
func SliceToString[T Any](i []T) []string {
	out := make([]string, len(i))
	for k, v := range i {
		out[k] = fmt.Sprintf("%v", v)
	}

	return out
}

// SliceToBool type a slice to b slice
func SliceToBool[T Any](i []T) []bool {
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

type RemoveSliceFunc[T Any] func(i, v T) bool

// RemoveAnySlice remove a from b
func RemoveAnySlice[T Any](b []T, a T, opts ...RemoveSliceFunc[T]) []T {
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

// MergeSlice merge two slice
func MergeSlice[T Any](a []T, b []T) []T {
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

type SortSliceLessFunc[T Any] func(i, j T) bool

// QuickSort sort a slice
func QuickSort[T Any](a []T, less SortSliceLessFunc[T]) {
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

// InsertSort 插入排序
func InsertSort[T Any](a []T, less SortSliceLessFunc[T]) {
	for i := 1; i < len(a); i++ {
		for j := i; j > 0 && less(a[j], a[j-1]); j-- {
			a[j], a[j-1] = a[j-1], a[j]
		}
	}
}

// BubbleSort 冒泡排序
func BubbleSort[T Any](a []T, less SortSliceLessFunc[T]) {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a)-i-1; j++ {
			if less(a[j+1], a[j]) {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}

// SelectSort 选择排序
func SelectSort[T Any](a []T, less SortSliceLessFunc[T]) {
	for i := 0; i < len(a); i++ {
		min := i
		for j := i + 1; j < len(a); j++ {
			if less(a[j], a[min]) {
				min = j
			}
		}
		a[i], a[min] = a[min], a[i]
	}
}

// ShellSort 希尔排序
func ShellSort[T Any](a []T, less SortSliceLessFunc[T]) {
	for gap := len(a) / 2; gap > 0; gap /= 2 {
		for i := gap; i < len(a); i++ {
			for j := i; j >= gap && less(a[j], a[j-gap]); j -= gap {
				a[j], a[j-gap] = a[j-gap], a[j]
			}
		}
	}
}

// HeapSort 堆排序
func HeapSort[T Any](a []T, less SortSliceLessFunc[T]) {
	n := len(a)
	for i := n/2 - 1; i >= 0; i-- {
		heapify(a, n, i, less)
	}

	for i := n - 1; i >= 0; i-- {
		a[0], a[i] = a[i], a[0]
		heapify(a, i, 0, less)
	}
}

func heapify[T Any](a []T, n, i int, less SortSliceLessFunc[T]) {
	largest := i
	l, r := 2*i+1, 2*i+2

	if l < n && less(a[largest], a[l]) {
		largest = l
	}

	if r < n && less(a[largest], a[r]) {
		largest = r
	}

	if largest != i {
		a[i], a[largest] = a[largest], a[i]
		heapify(a, n, largest, less)
	}
}

// MergeSort 归并排序
func MergeSort[T Any](a []T, less SortSliceLessFunc[T]) {
	if len(a) <= 1 {
		return
	}

	mid := len(a) / 2
	MergeSort(a[:mid], less)
	MergeSort(a[mid:], less)
	merge(a, less)
}

func merge[T Any](a []T, less SortSliceLessFunc[T]) {
	mid := len(a) / 2
	tmp := make([]T, len(a))
	copy(tmp, a)

	i, j, k := 0, mid, 0
	for i < mid && j < len(a) {
		if less(tmp[i], tmp[j]) {
			a[k] = tmp[i]
			i++
		} else {
			a[k] = tmp[j]
			j++
		}
		k++
	}

	for i < mid {
		a[k] = tmp[i]
		i++
		k++
	}

	for j < len(a) {
		a[k] = tmp[j]
		j++
		k++
	}
}
