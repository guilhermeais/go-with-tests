package sum

import (
	"reflect"
	"testing"
)

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func BenchmarkSumAll(b *testing.B) {
	numbersToSum := [][]int{
		{1, 2},
		{0, 9},
		{1, 2},
		{0, 9},
		{1, 2},
		{0, 9},
	}
	for i := 0; i < b.N; i++ {
		SumAll(numbersToSum...)
	}
}

func BenchmarkSumAllWithDynamicCap(b *testing.B) {
	numbersToSum := [][]int{
		{1, 2},
		{0, 9},
		{1, 2},
		{0, 9},
		{1, 2},
		{0, 9},
	}
	for i := 0; i < b.N; i++ {
		SumAllWithDynamicCap(numbersToSum...)
	}
}

// func a() {}
