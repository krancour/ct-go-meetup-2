package main

import (
	"fmt"
	"time"
)

func printSquares(nums []int, stopCh chan struct{}) chan struct{} {
	doneCh := make(chan struct{})
	go func() {
		defer close(doneCh)
		for _, num := range nums {
			select {
			case <-time.After(time.Second):
				fmt.Println(num * num)
			case <-stopCh:
				return
			}
		}
	}()
	return doneCh
}

func main() {
	stopCh := make(chan struct{})
	doneCh := printSquares([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, stopCh)
	time.Sleep(time.Second * 5)
	close(stopCh)
	<-doneCh
}
