package main

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

func TestBuildResponse(t *testing.T) {
	input := "hello"
	expected := "Echo: hello\n"

	result := buildResponse(input)

	if result != expected {
		t.Errorf("expected %s, got %s", expected, result)
	}
}

func TestProcessSingleMessage(t *testing.T) {
	input := "hello\n"
	expected := "Echo: hello\n"

	in := strings.NewReader(input)
	out := &bytes.Buffer{}

	rw := struct {
		io.Reader
		io.Writer
	}{
		Reader: in,
		Writer: out,
	}

	err := ProcessSingleMessage(rw)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if out.String() != expected {
		t.Errorf("expected %q, got %q", expected, out.String())
	}
}