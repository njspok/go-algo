package runbeforefirsterror

import "sync"

func Run(fn ...func() error) error {
	errs := make(chan error, len(fn))

	wg := sync.WaitGroup{}
	for _, f := range fn {
		wg.Add(1)

		go func(f func() error) {
			defer wg.Done()
			errs <- f()
		}(f)
	}

	go func() {
		wg.Wait()
		close(errs)
	}()

	for err := range errs {
		if err != nil {
			go func() {
				for range errs {
				}
			}()
			return err
		}
	}

	return nil
}
