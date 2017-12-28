package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func printSquares(ctx context.Context, nums []int) (chan struct{}, chan error) {
	doneCh := make(chan struct{})
	errCh := make(chan error)
	go func() {
		defer close(doneCh)
		defer close(errCh)
		for _, num := range nums {
			select {
			case <-time.After(time.Second):
				fmt.Println(num * num)
			case <-ctx.Done():
				errCh <- ctx.Err()
				return
			}
		}
	}()
	return doneCh, errCh
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	doneCh, errCh := printSquares(ctx, []int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	select {
	case <-doneCh:
		log.Fatal("We finished processing the list, but shouldn't have!")
	case err := <-errCh:
		if err == ctx.Err() {
			log.Println(err)
		} else {
			log.Fatal(err)
		}
	}
}
