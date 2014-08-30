package gosort

import "testing"

func TestParallelQuickSort(t *testing.T) {
	VerifySort(ParallelQuickSort, t)
}

func BenchmarkParallelQuickSort(b *testing.B) {
	BenchmarkSort(ParallelQuickSort, b)
}
