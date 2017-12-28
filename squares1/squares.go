package main

import (
	"fmt"
	"sync"
)

func printSquares(nums []int, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, num := range nums {
		fmt.Println(num * num)
	}
}

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go printSquares([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, wg)
	wg.Wait()
}
