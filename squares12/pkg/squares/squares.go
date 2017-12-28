package squares

import (
	"context"
)

// Generate is a squares pipeline stage that loads a given array of integers
// onto a return channel. Presumably, this return channel is used as the input
// channel to another stage.
func Generate(ctx context.Context, nums []int) (chan int, chan error) {
	retCh := make(chan int)
	errCh := make(chan error)
	go func() {
		defer close(retCh)
		defer close(errCh)
		for _, num := range nums {
			select {
			case retCh <- num:
			case <-ctx.Done():
				errCh <- ctx.Err()
				return
			}
		}
	}()
	return retCh, errCh
}

func square(num int) int {
	return num * num
}

// Squares receives integers over an input channel and places their squares on
// a return channel.
func Squares(
	ctx context.Context,
	inputCh chan int,
) (chan int, chan error) {
	retCh := make(chan int)
	errCh := make(chan error)
	go func() {
		defer close(retCh)
		defer close(errCh)
		for {
			select {
			case num, ok := <-inputCh:
				if ok {
					select {
					case retCh <- square(num):
					case <-ctx.Done():
						errCh <- ctx.Err()
						return
					}
				} else {
					// The input channel was closed
					return
				}
			case <-ctx.Done():
				errCh <- ctx.Err()
				return
			}
		}
	}()
	return retCh, errCh
}
