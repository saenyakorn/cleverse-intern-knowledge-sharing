package fibo

import "context"

func Fibonacci(ctx context.Context, n int) int {
	select {
	case <-ctx.Done():
		return 0
	default:
		if n < 2 {
			return n
		}
		return Fibonacci(ctx, n-1) + Fibonacci(ctx, n-2)
	}
}
