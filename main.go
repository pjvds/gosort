package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	slack = 5
)

func main() {
	rand.Seed(time.Now().UnixNano())
	input := getRandomList(1000)

	for _, n := range input {
		fmt.Println(n)
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
