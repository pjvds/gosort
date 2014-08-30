package gosort

func QuickSort(input []int) []int {
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

	return append(append(QuickSort(smaller), pivot), QuickSort(larger)...)
}
