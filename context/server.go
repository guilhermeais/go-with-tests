package context

import (
	"context"
	"fmt"
	"net/http"
)

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, _ := store.Fetch(r.Context())

		if len(data) > 0 {
			fmt.Fprint(w, data)
		}
	}
}

type Store interface {
	Fetch(ctx context.Context) (string, error)
}
