package ordone

import (
	"context"
)

func OrDone(ctx context.Context, in <-chan int) <-chan int {
	res := make(chan int)

	go func() {
		defer close(res)

		for {
			select {
			case <-ctx.Done():
				return
			case n, ok := <-in:
				if !ok {
					return
				}

				select {
				case <-ctx.Done():
					return
				case res <- n:
				}
			}
		}
	}()

	return res
}
