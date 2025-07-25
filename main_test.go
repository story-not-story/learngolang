package main

import (
	"testing"
)

type arg struct {
	a, b int
}

func Test_min1(t *testing.T) {
	test := []struct {
		id       string
		argvalue arg
		expected int
	}{
		{id: "0", argvalue: arg{a: 1, b: 2}, expected: 1},
		{id: "1", argvalue: arg{a: -1, b: 0}, expected: -1},
		{id: "2", argvalue: arg{a: -5, b: -1}, expected: -5},
		{id: "3", argvalue: arg{a: 15, b: 2}, expected: 2},
	}
	for _, item := range test {
		t.Run(item.id, func(t *testing.T) {
			if expected := min1(item.argvalue.a, item.argvalue.b); expected != item.expected {
				t.Errorf("input: %s, expected: %d, got: %d \n", item.id, item.expected, expected)
			}

		})
	}

}

/*
func Test_min3(t *testing.T) {
	type args[MyNumber3 cmp.Ordered] struct {
		x MyNumber3
		y MyNumber3
	}
	type testCase[MyNumber3 cmp.Ordered] struct {
		name string
		args args[MyNumber3]
		want MyNumber3
	}
	tests := []testCase[MyNumber3 cmp.Ordered]{
		{name: "0", args: args{x: 1, y: 2}, want: 1},
		{name: "1", args: args{x: -1, y: 0}, want: -1},
		{name: "2", args: args{x: -5, y: -1}, want: -5},
		{name: "3", args: args{x: 15, y: 2}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := min3(tt.args.x, tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("min3() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_min2(t *testing.T) {
	type args[MyNumber2 interface {
		~int | ~int8 | ~int16 | ~int32 | ~int64 |
			~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
			~float32 | ~float64 |
			~string
	}] struct {
		x MyNumber2
		y MyNumber2
	}
	type testCase[MyNumber2 interface {
		~int | ~int8 | ~int16 | ~int32 | ~int64 |
			~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
			~float32 | ~float64 |
			~string
	}] struct {
		name string
		args args[MyNumber2]
		want MyNumber2
	}
	tests := []testCase[interface {
		~int | ~int8 | ~int16 | ~int32 | ~int64 |
			~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
			~float32 | ~float64 |
			~string
	}]{
		{name: "0", args: args{x: 1, y: 2}, want: 1},
		{name: "1", args: args{x: -1, y: 0}, want: -1},
		{name: "2", args: args{x: -5, y: -1}, want: -5},
		{name: "3", args: args{x: 15, y: 2}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := min2(tt.args.x, tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("min2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_min11(t *testing.T) {
	type args[T MyNumber] struct {
		x T
		y T
	}
	type testCase[T MyNumber] struct {
		name string
		args args[T]
		want T
	}
	tests := []testCase[T MyNumber]{
	{name: "0", args: args{x: 1, y: 2}, want: 1},
	{name: "1", args: args{x: -1, y: 0}, want: -1},
	{name: "2", args: args{x: -5, y: -1}, want: -5},
	{name: "3", args: args{x: 15, y: 2}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := min1(tt.args.x, tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("min1() = %v, want %v", got, tt.want)
			}
		})
	}
}*/
