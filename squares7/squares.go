package main

import "fmt"

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

func main() {
	retCh := generate([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	retCh = squares(retCh)
	for result := range retCh {
		fmt.Println(result)
	}
}
