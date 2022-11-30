package basic

import (
	"reflect"
	"testing"
)

func TestToNumber_IntToInt32(t *testing.T) {
	type args struct {
		i []int
	}
	tests := []struct {
		name string
		args args
		want []int32
	}{
		{
			name: "test int to int32",
			args: args{
				i: []int{1, 2, 3},
			},
			want: []int32{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SliceToNumber[int32](tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SliceToNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToNumber_IntToInt64(t *testing.T) {
	type args struct {
		i []int
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{
			name: "test int to int64",
			args: args{
				i: []int{1, 2, 3},
			},
			want: []int64{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SliceToNumber[int64](tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SliceToNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToNumber_IntToFloat32(t *testing.T) {
	type args struct {
		i []int
	}
	tests := []struct {
		name string
		args args
		want []float32
	}{
		{
			name: "test int to float32",
			args: args{
				i: []int{1, 2, 3},
			},
			want: []float32{1.0, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SliceToNumber[float32](tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SliceToNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToNumber_IntToFloat64(t *testing.T) {
	type args struct {
		i []int
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{
			name: "test int to float64",
			args: args{
				i: []int{1, 2, 3},
			},
			want: []float64{1.0, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SliceToNumber[float64](tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SliceToNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToString(t *testing.T) {
	type args struct {
		i []int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test int to string",
			args: args{
				i: []int{1, 2, 3},
			},
			want: []string{"1", "2", "3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SliceToString(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SliceToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSliceToBool(t *testing.T) {
	type args struct {
		i []int
	}
	tests := []struct {
		name string
		args args
		want []bool
	}{
		{
			name: "test int to bool",
			args: args{
				i: []int{1, 2, 3, 0},
			},
			want: []bool{true, true, true, false},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SliceToBool(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SliceToBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInSlice(t *testing.T) {
	type args struct {
		slice []int
		value int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test in slice",
			args: args{
				slice: []int{1, 2, 3},
				value: 1,
			},
			want: true,
		},
		{
			name: "test in slice",
			args: args{
				slice: []int{1, 2, 3},
				value: 4,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InSlice(tt.args.slice, tt.args.value, func(i, j int) bool {
				return i < j
			}); got != tt.want {
				t.Errorf("InSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsSubset(t *testing.T) {
	type args struct {
		slice1 []int
		slice2 []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test is subset",
			args: args{
				slice1: []int{1, 2, 3},
				slice2: []int{1, 2, 3, 4},
			},
			want: false,
		},
		{
			name: "test is subset",
			args: args{
				slice1: []int{1, 2, 3},
				slice2: []int{1, 2},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsSubset(tt.args.slice1, tt.args.slice2...); got != tt.want {
				t.Errorf("IsSubset() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveSlice(t *testing.T) {
	type args struct {
		slice []int
		value int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test remove slice",
			args: args{
				slice: []int{1, 2, 3},
				value: 1,
			},
			want: []int{2, 3},
		},
		{
			name: "test remove slice",
			args: args{
				slice: []int{1, 2, 3},
				value: 4,
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveComparableSlice(tt.args.slice, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDiffSlice(t *testing.T) {
	type args struct {
		slice1 []int
		slice2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test diff slice",
			args: args{
				slice1: []int{1, 2, 3},
				slice2: []int{1, 2, 3, 4},
			},
			want: []int{},
		},
		{
			name: "test diff slice",
			args: args{
				slice1: []int{1, 2, 3},
				slice2: []int{1, 2},
			},
			want: []int{3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DiffSlice(tt.args.slice1, tt.args.slice2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DiffSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveAnySlice(t *testing.T) {
	type args struct {
		slice []int
		value int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test remove any slice",
			args: args{
				slice: []int{1, 2, 3},
				value: 1,
			},
			want: []int{2, 3},
		},
		{
			name: "test remove any slice",
			args: args{
				slice: []int{1, 2, 3},
				value: 4,
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveAnySlice(tt.args.slice, tt.args.value, func(i, v int) bool {
				return i == v
			}); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveAnySlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuickSort(t *testing.T) {
	type args struct {
		slice []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test quick sort",
			args: args{
				slice: []int{1, 2, 3},
			},
			want: []int{1, 2, 3},
		},
		{
			name: "test quick sort",
			args: args{
				slice: []int{3, 2, 1},
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if QuickSort[int](tt.args.slice, func(i, t int) bool {
				return i < t
			}); !reflect.DeepEqual(tt.args.slice, tt.want) {
				t.Errorf("QuickSort() = %v, want %v", tt.args.slice, tt.want)
			}
		})
	}
}

func TestMergeSlice(t *testing.T) {
	type args struct {
		slice1 []int
		slice2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test merge slice",
			args: args{
				slice1: []int{1, 2, 3},
				slice2: []int{1, 2, 3, 4},
			},
			want: []int{1, 2, 3, 1, 2, 3, 4},
		},
		{
			name: "test merge slice",
			args: args{
				slice1: []int{1, 2, 3},
				slice2: []int{1, 2},
			},
			want: []int{1, 2, 3, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeSlice(tt.args.slice1, tt.args.slice2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinarySearch(t *testing.T) {
	type args struct {
		slice []int
		value int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test binary search",
			args: args{
				slice: []int{1, 2, 3},
				value: 1,
			},
			want: 0,
		},
		{
			name: "test binary search",
			args: args{
				slice: []int{1, 2, 3},
				value: 4,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearch[int](tt.args.slice, tt.args.value); got != tt.want {
				t.Errorf("BinarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInsertSort(t *testing.T) {
	type args struct {
		slice []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test insert sort",
			args: args{
				slice: []int{1, 2, 3},
			},
			want: []int{1, 2, 3},
		},
		{
			name: "test insert sort",
			args: args{
				slice: []int{3, 2, 1},
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InsertSort[int](tt.args.slice, func(i, t int) bool {
				return i < t
			})
			if !reflect.DeepEqual(tt.args.slice, tt.want) {
				t.Errorf("InsertSort() = %v, want %v", tt.args.slice, tt.want)
			}
		})
	}
}
