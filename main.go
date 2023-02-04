package main

// From https://gobyexample.com/http-servers

import (
	"fmt"
	"net/http"
	"os"
)

func hi(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hi🐈\n")
	fmt.Println("hi called")
}

func healthz(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "ok\n")
	fmt.Println("healthz called")
}

func main() {
	http.HandleFunc("/hi", hi)
	http.HandleFunc("/healthz", healthz)

	fmt.Println("Listening on port", os.Getenv("PORT"))
	http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), nil)
}
