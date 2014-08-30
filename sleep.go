package gosort

import "time"

const (
	sleepSlack = 5
)

func SleepSort(input []int) []int {
	length := len(input)
	sorting := make(chan int)

	for i := 0; i < length; i++ {
		go func(n int) {
			time.Sleep(time.Duration(n) * time.Millisecond * sleepSlack)
			sorting <- n
		}(input[i])
	}

	output := make([]int, length, length)
	for i := 0; i < length; i++ {
		output[i] = <-sorting
	}

	return output
}
