package main

// From https://gobyexample.com/http-servers

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strconv"
)

func pi(w http.ResponseWriter, req *http.Request) {
	fmt.Println("pi called")
	n, err := strconv.Atoi(req.URL.Query().Get("n"))
	if err != nil || n < 1 {
		http.NotFound(w, req)
		return
	}

	var total int = 0
	var totalIn int = 0
	for i := 0; i < n; i++ {
		var x = rand.Float64()
		var y = rand.Float64()
		if x*x+y*y < 1 {
			totalIn++
		}
		total++
	}
	fmt.Fprintf(w, "%f\n", float64(totalIn)*4.0/float64(total))
}

func hi(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hi🐈\n")
	fmt.Println("hi called")
}

func healthz(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "ok\n")
	fmt.Println("healthz called")
}

func healthzButBroken(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("500 - Something bad happened!"))
	fmt.Println("healthz-but-broken called")
}

func healthzButSometimesBroken(w http.ResponseWriter, req *http.Request) {
	if rand.Intn(10)%2 == 0 {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
	}
	fmt.Println("healthz-but-sometimes-broken called")
}

func main() {
	http.HandleFunc("/hi", hi)
	http.HandleFunc("/pi", pi)
	http.HandleFunc("/healthz", healthz)
	http.HandleFunc("/healthz-but-broken", healthzButBroken)
	http.HandleFunc("/healthz-but-sometimes-broken", healthzButSometimesBroken)

	fmt.Println("Listening on port", os.Getenv("PORT"))
	err := http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), nil)
	if err != nil {
		panic(err)
	}
}
