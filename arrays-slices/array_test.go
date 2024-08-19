package arraysslices

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("test of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d, want %d given, %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("test of multiple slices", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d, want %d.", got, want)
		}
	})
}

func TestAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v.", got, want)
		}
	}
	t.Run("test with slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		checkSums(t, got, want)
	})
	t.Run("sum eith empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checkSums(t, got, want)
	})
}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
        Sum([]int{1, 2, 3, 4, 5})
    }
}

func BenchmarkSumAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
        SumAll([]int{1, 2}, []int{3, 4})
    }
}

func BenchmarkSumAllTails(b *testing.B) {
	for i := 0; i < b.N; i++ {
        SumAllTails([]int{1, 2}, []int{3, 4})
    }
}

func ExampleSum() {
	numbers := []int{10, 20, 30, 40, 50}
	sum := Sum(numbers)
	fmt.Println(sum)
	// Output: 150
}

func ExampleSumAll() {
	sums := SumAll([]int{1, 2}, []int{5, 6})
	fmt.Println(sums)
	// Output: [3 11]
}

func ExampleSumAllTails() {
	tails := SumAllTails([]int{1, 2, 3}, []int{4, 5, 6})
	fmt.Println(tails)
	// Output: [5 11]
}
