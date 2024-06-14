package context1

import (
	"fmt"
	"net/http"
)

// Store fetches data.
type Store interface {
	Fetch() string
	Cancel()
}

// Server returns a handler for calling Store.
func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		store.Cancel()
		fmt.Fprint(w, store.Fetch())
	}
}
