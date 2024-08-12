package main

import (
	"reflect"
	"testing"
)

func TestConvertStringToInterfaceSlice(t *testing.T) {
	type args struct {
		strings []string
	}
	tests := []struct {
		name string
		args args
		want []interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertStringToInterfaceSlice(tt.args.strings); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertStringToInterfaceSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSortRows_Len(t *testing.T) {
	tests := []struct {
		name string
		r    SortRows
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.Len(); got != tt.want {
				t.Errorf("Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSortRows_Less(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		r    SortRows
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSortRows_Swap(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		r    SortRows
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.r.Swap(tt.args.i, tt.args.j)
		})
	}
}

func TestConvertStringToInterfaceSlice1(t *testing.T) {
	type args struct {
		strings []string
	}
	tests := []struct {
		name string
		args args
		want []interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertStringToInterfaceSlice(tt.args.strings); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertStringToInterfaceSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSortRows_Len1(t *testing.T) {
	tests := []struct {
		name string
		r    SortRows
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.Len(); got != tt.want {
				t.Errorf("Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSortRows_Less1(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		r    SortRows
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSortRows_Swap1(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		r    SortRows
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.r.Swap(tt.args.i, tt.args.j)
		})
	}
}
