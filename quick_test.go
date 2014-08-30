package gosort

import "testing"

func TestQuickSort(t *testing.T) {
	VerifySort(QuickSort, t)
}

func BenchmarkQuickSort(b *testing.B) {
	BenchmarkSort(QuickSort, b)
}
