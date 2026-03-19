package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

type Input struct {
	Numbers []int `json:"numbers"`
}

type Output struct {
	Sum int `json:"sum"`
}

func calculateSumOfSquares(numbers []int) int {
	sum := 0
	for _, n := range numbers {
		sum += n * n
	}
	return sum
}

func process(r io.Reader, w io.Writer) error {
	var input Input

	if err := json.NewDecoder(r).Decode(&input); err != nil {
		return err
	}

	result := Output{
		Sum: calculateSumOfSquares(input.Numbers),
	}

	return json.NewEncoder(w).Encode(result)
}

func main() {
	if err := process(os.Stdin, os.Stdout); err != nil {
		log.Fatal(err)
	}
}