package mergechans

import (
	"sync"
)

func MergeChans(chans ...<-chan int) <-chan int {
	res := make(chan int)

	go func() {
		wg := &sync.WaitGroup{}
		for _, ch := range chans {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for v := range ch {
					res <- v
				}
			}()
		}
		wg.Wait()
		close(res)
	}()

	return res
}
