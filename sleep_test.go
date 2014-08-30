package gosort

import "testing"

func TestSleepSort(t *testing.T) {
	VerifySort(SleepSort, t)
}
