package downloads

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// timeoutLimit - вероятность, с которой не будет возвращаться ошибка от fakeDownload():
// timeoutLimit = 100 - ошибок не будет;
// timeoutLImit = 0 - всегда будет возвращаться ошибка.
// Можете "поиграть" с этим параметром, для проверки случаев с возвращением ошибки.
const timeoutLimit = 50

type Result struct {
	msg string
	err error
}

// fakeDownload - имитирует разное время скачивания для разных адресов
func fakeDownload(url string) Result {
	r := rand.Intn(100)
	time.Sleep(time.Duration(r) * time.Millisecond)
	if r > timeoutLimit {
		return Result{
			err: errors.New(fmt.Sprintf("failed to download data from %s: timeout", url)),
		}
	}
	return Result{
		msg: fmt.Sprintf("downloaded data from %s", url),
	}
}

// download - параллельно скачивает данные из urls
func download(urls []string) ([]string, error) {
	ch := make(chan Result, len(urls))

	for _, url := range urls {
		// url := url // no need
		go func() {
			ch <- fakeDownload(url)
		}()
	}

	var res []string
	var errs []error
	//var err error

	for i := 0; i < len(urls); i++ {
		r := <-ch
		if r.err != nil {

			// good, but not work interface{Unwrap() []errors} !!!
			//err = errors.Join(err, r.err)

			errs = append(errs, r.err)
		} else {
			res = append(res, r.msg)
		}
	}

	close(ch)

	//return res, err

	return res, errors.Join(errs...)
}
