package gosort

import "testing"

func TestSleepSort(t *testing.T) {
	VerifySort(SleepSort, t)
}

func BenchmarkSleepSort(b *testing.B) {
	BenchmarkSort(SleepSort, b)
}
