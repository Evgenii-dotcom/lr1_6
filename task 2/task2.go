package main

import (
	"context"
	"log"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(addr string) *Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/", helloHandler)

	return &Server{
		httpServer: &http.Server{
			Addr:    addr,
			Handler: mux,
		},
	}
}

func (s *Server) Start() error {
	log.Printf("HTTP сервер запущен на %s\n", s.httpServer.Addr)
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Hello from Go server"))
}

type Worker struct {
	interval time.Duration
	stop     chan struct{}
}

func NewWorker(interval time.Duration) *Worker {
	return &Worker{
		interval: interval,
		stop:     make(chan struct{}),
	}
}

func (w *Worker) Start() {
	go func() {
		ticker := time.NewTicker(w.interval)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				log.Println("Фоновая обработка...")
			case <-w.stop:
				return
			}
		}
	}()
}

func (w *Worker) Stop() {
	close(w.stop)
}

func main() {
	worker := NewWorker(5 * time.Second)
	worker.Start()

	server := NewServer(":8080")

	if err := server.Start(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}