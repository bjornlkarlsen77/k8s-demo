package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	rootContext := os.Getenv("ROOT_CONTEXT")
	if rootContext == "" {
		rootContext = "/" // Default to "/" if the variable is not set
	}
	http.HandleFunc(rootContext, Handler)
	http.HandleFunc(rootContext+"slow", SlowHandler)
	http.HandleFunc("/healthz", Healthz)
	http.HandleFunc("/ready", Ready)
	fmt.Println("Server starting on :8080")
	http.ListenAndServe(":8080", nil)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("api called")
	fmt.Fprintf(w, "Hello, from GoLang!")
}

func SlowHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("slow api called")
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
