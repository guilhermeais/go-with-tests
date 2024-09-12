package racer

import (
	"fmt"
	"net/http"
	"time"
)

func Racer(urls ...string) (winner string, err error) {
	ch := make(chan string)
	for _, url := range urls {
		go func(u string) {
			ping(u)
			ch <- u
		}(url)
	}

	select {
	case winner = <-ch:
		return winner, nil
	case <-time.After(10 * time.Second):
		return "", fmt.Errorf("timed out waiting for responses")
	}
}

func ping(url string) {
	resp, err := http.Head(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
}
