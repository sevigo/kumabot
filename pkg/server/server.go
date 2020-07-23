package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"
)

type Server struct {
	Addr    string
	Cert    string
	Host    string
	Handler http.Handler
}

func (s Server) ListenAndServe(ctx context.Context) error {
	srv := &http.Server{
		Addr:    s.Addr,
		Handler: s.Handler,
	}

	idleConnsClosed := make(chan struct{})
	go func() {
		sigChan := make(chan os.Signal, 1)
		// Setup our Ctrl+C handler
		signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

		// we block here until we got one of the signals from above
		select {
		case <-ctx.Done():
			log.Info("received context done signal, shutting down")
		case <-sigChan:
			log.Info("received an interrupt signal, shutting down")
		}
		if err := srv.Shutdown(ctx); err != nil {
			// Error from closing listeners, or context timeout:
			log.Infof("HTTP server Shutdown: %v", err)
		}
		close(idleConnsClosed)
	}()
	log.Info("ListenAndServe")
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		// Error starting or closing listener:
		log.Fatalf("HTTP server ListenAndServe: %v", err)
		return err
	}

	// block here until Shutdown in the go-routine is executed and idleConnsClosed is closed
	<-idleConnsClosed
	log.Info("ListenAndServe: service is stopped")
	return nil
}
