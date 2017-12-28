package main

import "fmt"

func printSquares(nums []int) chan struct{} {
	doneCh := make(chan struct{})
	go func() {
		defer close(doneCh)
		for _, num := range nums {
			fmt.Println(num * num)
		}
	}()
	return doneCh
}

func main() {
	doneCh := printSquares([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	<-doneCh
}
