package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", serveMainPage)
	http.HandleFunc("/login", serveLoginPage)
	http.ListenAndServe(":8080", nil)
}

func serveMainPage(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Serving request: %s\n, %s\n", r.Method, r.URL.Path)
	http.ServeFile(w, r, "index.html")
}

func serveLoginPage(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Serving request: %s\n, %s\n", r.Method, r.URL.Path)
	http.ServeFile(w, r, "login.html")
}
