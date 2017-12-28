package squares

import (
	"context"
	"testing"
	"time"
)

var inputs = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

func TestGenerateCanceled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	retCh, errCh := Generate(ctx, inputs)
	// Assert we get a context canceled error back on the error channel
	select {
	case err, ok := <-errCh:
		if ok {
			if err != ctx.Err() {
				t.Fatalf(`expected error "%s", but got "%s"`, ctx.Err(), err)
			}
		} else {
			t.Fatalf(`expected error "%s", but error channel was closed`, ctx.Err())
		}
	case <-time.After(time.Second):
		t.Fatal("did not receive an error on the error channel, but should have")
	}
	// Assert that the return channel was closed
	select {
	case result, ok := <-retCh:
		if ok {
			t.Fatalf(
				"expected return channel to be closed, but received result %d",
				result,
			)
		}
	case <-time.After(time.Second):
		t.Fatalf("expected return channel to be closed, but it appears open")
	}
	// Assert that the error channel was closed
	select {
	case err, ok := <-errCh:
		if ok {
			t.Fatalf(
				`expected error channel to be closed, but received error "%s"`,
				err,
			)
		}
	case <-time.After(time.Second):
		t.Fatalf("expected error channel to be closed, but it appears open")
	}
}

func TestGenerateRunsToCompletion(t *testing.T) {
	resultCount := 0
	retCh, errCh := Generate(context.Background(), inputs)
loop:
	for {
		select {
		case result, ok := <-retCh:
			if ok {
				resultCount++
				if resultCount <= len(inputs) {
					if result != inputs[resultCount-1] {
						t.Fatalf(
							"expected result %d, received %d",
							inputs[resultCount-1],
							resultCount,
						)
					}
				} else {
					t.Fatalf(
						"expected only %d results, but have already received %d",
						len(inputs),
						resultCount,
					)
				}
			} else {
				// Channel was closed. Break out of the loop.
				break loop
			}
		case err, ok := <-errCh:
			if ok {
				t.Fatalf(`expected no errors, but received "%s"`, err)
			}
		case <-time.After(time.Second):
			t.Fatalf(
				"not receiving any result on the return channel, nor any errors on " +
					"the error channel",
			)
		}
	}
	if resultCount != len(inputs) {
		t.Fatalf(
			"expected exactly %d results, but only received %d",
			len(inputs),
			resultCount,
		)
	}
	// The only way to have broken out of the loop without a fatal event is
	// if the return channel was closed, so we know it is. Check also that
	// the error channel is closed.
	select {
	case err, ok := <-errCh:
		if ok {
			t.Fatalf(
				`expected error channel to be closed, but received error "%s"`,
				err,
			)
		}
	case <-time.After(time.Second):
		t.Fatalf("expected error channel to be closed, but it appears open")
	}
}

func TestSquare(t *testing.T) {
	testCases := []struct {
		num      int //input
		expected int // expected output
	}{
		{-3, 9},
		{-2, 4},
		{-1, 1},
		{0, 0},
		{1, 1},
		{2, 4},
		{3, 9},
	}
	for _, testCase := range testCases {
		result := square(testCase.num)
		if result != testCase.expected {
			t.Errorf(
				"doSquare(%d): expected %d, actual %d",
				testCase.num,
				testCase.expected,
				result,
			)
		}
	}
}

func TestSquaresCanceled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	retCh, errCh := Squares(ctx, make(chan int))
	// Assert we get a context canceled error back on the error channel
	select {
	case err, ok := <-errCh:
		if ok {
			if err != ctx.Err() {
				t.Fatalf(`expected error "%s", but got "%s"`, ctx.Err(), err)
			}
		} else {
			t.Fatalf(`expected error "%s", but error channel was closed`, ctx.Err())
		}
	case <-time.After(time.Second):
		t.Fatal("did not receive an error on the error channel, but should have")
	}
	// Assert that the return channel was closed
	select {
	case result, ok := <-retCh:
		if ok {
			t.Fatalf(
				"expected return channel to be closed, but received result %d",
				result,
			)
		}
	case <-time.After(time.Second):
		t.Fatalf("expected return channel to be closed, but it appears open")
	}
	// Assert that the error channel was closed
	select {
	case err, ok := <-errCh:
		if ok {
			t.Fatalf(
				`expected error channel to be closed, but received error "%s"`,
				err,
			)
		}
	case <-time.After(time.Second):
		t.Fatalf("expected error channel to be closed, but it appears open")
	}
}

func TestSquaresRunsToCompletion(t *testing.T) {
	inputCh := make(chan int)
	go func() {
		defer close(inputCh)
		for num := range inputs {
			inputCh <- num
		}
	}()
	retCh, errCh := Squares(context.Background(), inputCh)
	resultCount := 0
loop:
	for {
		select {
		case _, ok := <-retCh:
			if ok {
				resultCount++
				if resultCount > len(inputs) {
					t.Fatalf(
						"expected only %d results, but have already received %d",
						len(inputs),
						resultCount,
					)
				}
			} else {
				// Channel was closed. Break out of the loop.
				break loop
			}
		case err, ok := <-errCh:
			if ok {
				t.Fatalf(`expected no errors, but received "%s"`, err)
			}
		case <-time.After(time.Second):
			t.Fatalf(
				"not receiving any result on the return channel, nor any errors on " +
					"the error channel",
			)
		}
	}
	if resultCount != len(inputs) {
		t.Fatalf(
			"expected exactly %d results, but only received %d",
			len(inputs),
			resultCount,
		)
	}
	// The only way to have broken out of the loop without a fatal event is
	// if the return channel was closed, so we know it is. Check that the error
	// channel is also closed.
	select {
	case err, ok := <-errCh:
		if ok {
			t.Fatalf(
				`expected error channel to be closed, but received error "%s"`,
				err,
			)
		}
	case <-time.After(time.Second):
		t.Fatalf("expected error channel to be closed, but it appears open")
	}
}
