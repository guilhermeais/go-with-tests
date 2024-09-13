package racer

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

const tenSecondTimeout = 10 * time.Second

func Racer(urls ...string) (winner string, err error) {
	return ConfigurableRacer(urls, tenSecondTimeout)
}

func ConfigurableRacer(urls []string, timeout time.Duration) (winner string, err error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ch := make(chan string)
	for _, url := range urls {
		go func(u string) {
			if err := ping(ctx, u); err != nil {
				return
			}
			ch <- u
		}(url)
	}

	select {
	case winner = <-ch:
		cancel()
		return winner, nil
	case <-time.After(timeout):
		cancel()
		fmt.Println("Deu timeout")
		return "", fmt.Errorf("timed out waiting for responses")
	}
}

func ping(ctx context.Context, url string) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
