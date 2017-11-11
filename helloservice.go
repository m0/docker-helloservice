package main

import (
	"os"
	"io"
	"net/http"
	"fmt"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("serving %s request from %s\n", r.Method, r.RemoteAddr)
	message := os.Getenv("MESSAGE")
	if len(message) < 1 {
		message = "hello docker!"
	}
	io.WriteString(w, fmt.Sprintf("%s", message))
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8080", nil)
}
