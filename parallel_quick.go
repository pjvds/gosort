package gosort

import "sync"

func ParallelQuickSort(input []int) (output []int) {
	if len(input) == 0 {
		return input
	}

	var smaller, larger []int
	pivot := input[0]

	for _, n := range input[1:] {
		if n < pivot {
			smaller = append(smaller, n)
			continue
		}
		larger = append(larger, n)
	}

	var work sync.WaitGroup

	if len(smaller) > 0 {
		work.Add(1)
		go func() {
			smaller = ParallelQuickSort(smaller)
			work.Done()
		}()
	}

	if len(larger) > 0 {
		work.Add(1)
		go func() {
			larger = ParallelQuickSort(larger)
			work.Done()
		}()
	}

	work.Wait()

	return append(append(smaller, pivot), larger...)
}
