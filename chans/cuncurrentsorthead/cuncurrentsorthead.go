// Напишите функцию ConcurrentSortHead, которая из files ридеров, которые содержат
// упорядоченные по возрастанию строки, вернет m первых строк. Чтение из ридеров
// files должно быть конкурентным.
//
// Например:
// f1 := "aaa\nddd"
// f2 := "bbb\neee"
// f3 := "ccc\nfff"
// ConcurrentSortHead(4, files...) =>  []string{"aaa", "bbb", "ccc", "ddd"}

package cuncurrentsorthead

import (
	"bufio"
	"errors"
	"io"
	"slices"
)

func ConcurrencySortHead(m int, files ...io.Reader) ([]string, error) {
	if m <= 0 {
		return nil, errors.New("count strings must positive integer")
	}
	if len(files) == 0 {
		return nil, errors.New("no files provided")
	}

	chans := make(map[chan string]struct{}, len(files))

	for _, file := range files {
		file := file
		ch := make(chan string)

		go func() {
			scanner := bufio.NewScanner(bufio.NewReader(file))
			for scanner.Scan() {
				if scanner.Err() != nil {
					panic(scanner.Err())
				}
				ch <- scanner.Text()
			}

			close(ch)
		}()

		chans[ch] = struct{}{}
	}

	var res []string
	var counter int

	for len(chans) > 0 {
		r := make([]string, 0, len(chans))

		for ch := range chans {
			v, ok := <-ch
			if !ok {
				delete(chans, ch)
				continue
			}
			r = append(r, v)
		}

		slices.Sort(r)

		for _, i := range r {
			if i == "" {
				continue
			}

			counter++
			if counter > m {
				break
			}

			res = append(res, i)
		}
	}

	if len(res) < m {
		return res, errors.New("not enough lines")
	}

	return res, nil
}

func SortHead(m int, files ...io.Reader) ([]string, error) {
	if m <= 0 {
		return nil, errors.New("count strings must positive integer")
	}
	if len(files) == 0 {
		return nil, errors.New("no files provided")
	}

	var scans []*bufio.Scanner
	for _, file := range files {
		scans = append(scans, bufio.NewScanner(file))
	}

	var res []string
	linesCounter := 0
	for {
		ready := 0
		r := make([]string, 0, m)
		for _, scan := range scans {
			if scan.Scan() {
				if scan.Err() != nil {
					return nil, scan.Err()
				}

				ready++

				if scan.Text() == "" {
					continue
				}

				linesCounter++

				r = append(r, scan.Text())
				if linesCounter == m {
					slices.Sort(r)
					res = append(res, r...)
					return res, nil
				}
			}
		}

		slices.Sort(r)
		res = append(res, r...)

		if ready == 0 {
			break
		}
	}

	if linesCounter < m {
		return res, errors.New("not enough lines")
	}

	return res, nil
}
