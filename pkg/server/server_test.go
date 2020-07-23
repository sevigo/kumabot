package server

import (
	"context"
	"net/http"
	"sync"
	"testing"

	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
)

func TestServer_ListenAndServe(t *testing.T) {
	var wg sync.WaitGroup

	r := chi.NewRouter()
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(201)
	})

	s := Server{
		Addr:    ":18088",
		Handler: r,
	}
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go func() {
		s.ListenAndServe(ctx)
		wg.Done()
	}()
	resp, err := http.Get("http://localhost:18088/ping")
	assert.NoError(t, err)
	assert.Equal(t, 201, resp.StatusCode)
	cancel()
	wg.Wait()
}
