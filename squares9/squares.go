package main

import (
	"fmt"
	"time"
)

func printSquares(nums []int, timeout time.Duration) chan struct{} {
	timer := time.NewTimer(timeout)
	doneCh := make(chan struct{})
	go func() {
		defer close(doneCh)
		for _, num := range nums {
			select {
			case <-time.After(time.Second):
				fmt.Println(num * num)
			case <-timer.C:
				return
			}
		}
	}()
	return doneCh
}

func main() {
	doneCh := printSquares([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, time.Second*5)
	<-doneCh
}
