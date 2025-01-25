package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", Handler)
	http.HandleFunc("/slow", SlowHandler)
	http.HandleFunc("/healthz", Healthz)
	http.HandleFunc("/ready", Ready)
	http.ListenAndServe(":8080", nil)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, from GoLang!")
}

func SlowHandler(w http.ResponseWriter, r *http.Request) {
	for i := 1; i <= 1000; i++ {
		fmt.Println(i)
	}
}

func Healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("ok"))
}

func Ready(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("ok"))
}
