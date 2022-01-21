package decorator

import (
	"context"
	"log"
)

type Service interface {
	Sum(ctx context.Context, a, b int) (int, error)
	Concat(ctx context.Context, a, b string) (string, error)
}

type Middleware func(service Service) Service

type LoggingMiddleware struct {
	logger *log.Logger
	next   Service
}

func NewLoggingMiddleware(logger *log.Logger) Middleware {
	return func(next Service) Service {
		return LoggingMiddleware{logger, next}
	}
}

func (mw LoggingMiddleware) Sum(ctx context.Context, a, b int) (v int, err error) {
	defer func() {
		mw.logger.Printf("Method: Sum a: %v b: %v result: %v err: %v", a, b, v, err)
	}()
	return mw.next.Sum(ctx, a, b)
}

func (mw LoggingMiddleware) Concat(ctx context.Context, a, b string) (v string, err error) {
	defer func() {
		mw.logger.Printf("Method: Concat a: %v b: %v result: %v err: %v", a, b, v, err)
	}()
	return mw.next.Concat(ctx, a, b)
}
