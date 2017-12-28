package main

import "fmt"

func printSquares(nums []int) {
	for _, num := range nums {
		fmt.Println(num * num)
	}
}

func main() {
	go printSquares([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
}
