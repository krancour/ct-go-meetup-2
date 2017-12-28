package main

import "fmt"

func squares(nums []int) chan int {
	retCh := make(chan int)
	go func() {
		for _, num := range nums {
			retCh <- num * num
		}
	}()
	return retCh
}

func main() {
	retCh := squares([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	for result := range retCh {
		fmt.Println(result)
	}
}
