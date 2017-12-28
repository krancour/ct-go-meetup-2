package main

import (
	"fmt"
	"time"
)

func printSquares(nums []int, stopCh chan struct{}) {
	for _, num := range nums {
		select {
		case <-time.After(time.Second):
			fmt.Println(num * num)
		case <-stopCh:
			return
		}
	}
}

func main() {
	stopCh := make(chan struct{})
	go printSquares([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, stopCh)
	time.Sleep(time.Second * 5)
	close(stopCh)
}
