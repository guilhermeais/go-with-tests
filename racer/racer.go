package racer

import (
	"net/http"
	"time"
)

func Racer(urls ...string) (winner string) {
	var fastest time.Duration
	for i, url := range urls {
		start := time.Now()
		http.Get(url)
		duration := time.Since(start)

		isFirstUrl := i == 0

		if isFirstUrl || duration < fastest {
			fastest = duration
			winner = url
			continue
		}
	}

	return winner
}
