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

	for _, n := range sleepSort(input) {
		fmt.Println(n)
	}
}

func sleepSort(input []int) []int {
	length := len(input)
	sorting := make(chan int)

	for i := 0; i < length; i++ {
		go func(n int) {
			time.Sleep(time.Duration(n) * time.Millisecond * slack)
			sorting <- n
		}(input[i])
	}

	output := make([]int, length, length)
	for i := 0; i < length; i++ {
		output[i] = <-sorting
	}

	return output
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
