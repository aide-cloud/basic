package basic

import "testing"

func TestIsZero(t *testing.T) {
	type args struct {
		i interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "int 0 test is zero",
			args: args{
				i: 0,
			},
			want: true,
		},
		{
			name: "int 1 test is not zero",
			args: args{
				i: 1,
			},
			want: false,
		},
		{
			name: "float 0.0 test is zero",
			args: args{
				i: 0.0,
			},
			want: true,
		},
		{
			name: "float 1.0 test is not zero",
			args: args{
				i: 1.0,
			},
			want: false,
		},
		{
			name: "string \"\" test is zero",
			args: args{
				i: "",
			},
			want: true,
		},
		{
			name: "string \"1\" test is not zero",
			args: args{
				i: "1",
			},
			want: false,
		},
		{
			name: "bool false test is zero",
			args: args{
				i: false,
			},
			want: true,
		},
		{
			name: "bool true test is not zero",
			args: args{
				i: true,
			},
			want: false,
		},
		{
			name: "nil test is zero",
			args: args{
				i: nil,
			},
			want: true,
		},
		{
			name: "struct test is not zero",
			args: args{
				i: struct{}{},
			},
			want: false,
		},
		{
			name: "slice test is not zero",
			args: args{
				i: []int{},
			},
			want: false,
		},
		{
			name: "map test is not zero",
			args: args{
				i: map[string]int{},
			},
			want: false,
		},
		{
			name: "func test is not zero",
			args: args{
				i: func() {},
			},
			want: false,
		},
		{
			name: "chan test is not zero",
			args: args{
				i: make(chan int),
			},
			want: false,
		},
		{
			name: "complex test is not zero",
			args: args{
				i: complex(1, 1),
			},
			want: false,
		},
		{
			name: "complex test is zero",
			args: args{
				i: complex(0, 0),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsZero(tt.args.i); got != tt.want {
				t.Errorf("IsZero() = %v, want %v", got, tt.want)
			}
		})
	}
}
