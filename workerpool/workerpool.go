package workerpool

import "sync"

type Result struct {
	WorkerID int
	Value    int
	Result   int
}

func Worker(
	id int,
	in <-chan int,
	task func(x int) int,
	results chan<- Result,
) {
	for val := range in {
		results <- Result{
			WorkerID: id,
			Value:    val,
			Result:   task(val),
		}
	}
}

func NewWorkerPool(
	n int,
	in <-chan int,
	task func(x int) int,
) (results <-chan Result) {
	res := make(chan Result)

	go func() {
		wg := sync.WaitGroup{}
		for i := 0; i < n; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				Worker(i, in, task, res)
			}()
		}
		wg.Wait()
		close(res)
	}()

	return res
}
