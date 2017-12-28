package main

import (
	"context"
	"fmt"
	"log"

	"github.com/krancour/ct-go-meetup-2/squares12/pkg/squares"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	genRetCh, genErrCh := squares.Generate(ctx, []int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	sqrRetCh, sqrErrCh := squares.Squares(ctx, genRetCh)
	for {
		select {
		case result, ok := <-sqrRetCh:
			if !ok {
				return
			}
			fmt.Println(result)
		case err, ok := <-genErrCh:
			if ok {
				log.Fatal(err)
			}
		case err, ok := <-sqrErrCh:
			if ok {
				log.Fatal(err)
			}
		}
	}
}
