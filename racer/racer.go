package racer

import (
	"net/http"
)

func Racer(urls ...string) (winner string) {
	ch := make(chan string)
	for _, url := range urls {
		go func(_url string) {
			ping(url)
			ch <- _url
		}(url)
	}

	return <-ch
}

func ping(url string) {
	resp, err := http.Head(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
}
