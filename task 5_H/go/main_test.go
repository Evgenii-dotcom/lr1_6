package main

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestCalculateSumOfSquares(t *testing.T) {
	result := calculateSumOfSquares([]int{1, 2, 3})
	if result != 14 {
		t.Errorf("expected 14, got %d", result)
	}
}

func TestProcess(t *testing.T) {
	input := Input{Numbers: []int{1, 2, 3}}

	inputBytes, _ := json.Marshal(input)

	in := bytes.NewBuffer(inputBytes)
	out := &bytes.Buffer{}

	err := process(in, out)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var result Output
	json.Unmarshal(out.Bytes(), &result)

	if result.Sum != 14 {
		t.Errorf("expected 14, got %d", result.Sum)
	}
}