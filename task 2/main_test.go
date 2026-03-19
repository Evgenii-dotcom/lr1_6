package main

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHelloHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	helloHandler(rec, req)

	res := rec.Result()
	body, _ := io.ReadAll(res.Body)

	expected := "Hello from Go server"

	if string(body) != expected {
		t.Errorf("expected %s, got %s", expected, string(body))
	}
}

func TestServer_StartAndShutdown(t *testing.T) {
	server := NewServer(":0") 

	go func() {
		_ = server.Start()
	}()

	time.Sleep(100 * time.Millisecond)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	err := server.Shutdown(ctx)
	if err != nil {
		t.Errorf("shutdown failed: %v", err)
	}
}

func TestWorker_StartStop(t *testing.T) {
	worker := NewWorker(10 * time.Millisecond)
	worker.Start()

	time.Sleep(50 * time.Millisecond)

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("worker stop caused panic: %v", r)
		}
	}()

	worker.Stop()
}