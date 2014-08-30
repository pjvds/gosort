package gosort

import (
	"math/rand"
	"testing"
)

var (
	smallSet = getRandomList(10)
	largeSet = getRandomList(10 * 1000)
)

func VerifySort(sort func([]int) []int, t *testing.T) {
	output := sort(smallSet)

	var previous = output[0]
	for i := 1; i < len(output); i++ {
		current := output[i]

		if previous > current {
			t.Fatalf("number %v at %v is smaller than previous number %v", current, i, previous)
		}
	}

}

func getRandomList(length int) []int {
	result := make([]int, length, length)

	for i := 0; i < length; i++ {
		result[i] = i + 1
	}

	for i := 0; i < length; i++ {
		randomIndex := rand.Intn(length - 1)
		left := result[i]
		right := result[randomIndex]

		result[i] = right
		result[randomIndex] = left
	}

	return result
}
