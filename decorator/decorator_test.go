package decorator

import (
	"context"
	"errors"
	"log"
	"math"
	"os"
)

var ErrTwoZeroes = errors.New("non zero number required")
var ErrIntOverflow = errors.New("integer overflow")
var ErrMaxSizeExceeded = errors.New("max size exceeded")

type basicService struct{}

func (s basicService) Sum(_ context.Context, a, b int) (int, error) {
	if a == 0 && b == 0 {
		return 0, ErrTwoZeroes
	}
	if b > 0 && a > (math.MaxInt32-b) || b < 0 && a < (math.MinInt32) {
		return 0, ErrIntOverflow
	}

	return a + b, nil
}

func (s basicService) Concat(_ context.Context, a, b string) (string, error) {
	if len(a)+len(b) > 2048 {
		return "", ErrMaxSizeExceeded
	}
	return a + b, nil
}

// ExampleMiddleware - based on https://www.youtube.com/watch?v=NX0sHF8ZZgw
func ExampleMiddleware() {
	loggingMiddleware := NewLoggingMiddleware(log.New(os.Stdout, "", 0))
	service := loggingMiddleware(basicService{})
	service.Sum(context.TODO(), 1, 2)
	// Output: Method: Sum a: 1 b: 2 result: 3 err: <nil>
}
