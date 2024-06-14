package context1

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type StubStore struct {
	response string
	cancled  bool
}

func (s *StubStore) Fetch() string {
	time.Sleep(100 * time.Millisecond)
	return s.response
}

func (s *StubStore) Cancel() {
	s.cancled = true
}

func TestServer(t *testing.T) {
	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
		data := "hello, world"
		store := &StubStore{response: data, cancled: false}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingCtx, cancle := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancle)

		request = request.WithContext(cancellingCtx)

		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if !store.cancled {
			t.Error("store was not told to cancel")
		}
	})
}
