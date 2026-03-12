package main

import (
	"fmt"
	"net/http"
	"time"
)

func backgroundWorker() {
	for {
		fmt.Println("Фоновая обработка...")
		time.Sleep(5 * time.Second)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from Go server")
}

func main() {

	go backgroundWorker()

	http.HandleFunc("/", handler)

	fmt.Println("HTTP сервер запущен на порту 8080")

	http.ListenAndServe(":8080", nil)
}