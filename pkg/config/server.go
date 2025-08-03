package config

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	*http.Server
}

func NewServer(addr string, handler http.Handler) *Server {
	return &Server{
		Server: &http.Server{
			Addr:    addr,
			Handler: handler,
		},
	}
}

func (s *Server) Start() {
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Printf("Server listening on %s", s.Addr)
		if err := s.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Could not listen on %s: %v", s.Addr, err)
		}
	}()

	<-stopChan
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown failed: %v", err)
	}

	log.Println("Server gracefully stopped")
}
