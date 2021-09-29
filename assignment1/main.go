package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request From:", r.RemoteAddr)
	for k, v := range r.Header {
		for _, vv := range v {
			fmt.Println(k, vv)
			w.Header().Add(k, vv)
		}
	}
	w.Header().Add("VERSION", os.Getenv("VERSION"))
	code := 200
	w.WriteHeader(code)
	fmt.Println("ResponseCode:", code)
}

func Healthhandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/healthz", Healthhandler)
	http.ListenAndServe("localhost:8898", nil)
}
