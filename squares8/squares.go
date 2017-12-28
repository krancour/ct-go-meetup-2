package main

import (
	"fmt"
	"sync"
)

func generate(nums []int) chan int {
	retCh := make(chan int)
	go func() {
		defer close(retCh)
		for _, num := range nums {
			retCh <- num
		}
	}()
	return retCh
}

func squares(inputCh chan int) chan int {
	retCh := make(chan int)
	go func() {
		defer close(retCh)
		for num := range inputCh {
			retCh <- num * num
		}
	}()
	return retCh
}

func merge(inputChs ...chan int) chan int {
	retCh := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(len(inputChs))
	for _, inputCh := range inputChs {
		go func(inputCh chan int) {
			defer wg.Done()
			for num := range inputCh {
				retCh <- num
			}
		}(inputCh)
	}
	go func() {
		defer close(retCh)
		wg.Wait()
	}()
	return retCh
}

func main() {
	retCh := generate([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	retCh = merge(squares(retCh), squares(retCh))
	for result := range retCh {
		fmt.Println(result)
	}
}
