package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", serveMainPage)
	http.HandleFunc("/loginPage", serveLoginPage)
	http.HandleFunc("/login", handleLogin)
	http.ListenAndServe(":8080", nil)
}

func serveMainPage(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Serving request: %s\n, %s\n", r.Method, r.URL.Path)
	http.ServeFile(w, r, "Pages/index.html")
}

func serveLoginPage(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Serving request: %s\n, %s\n", r.Method, r.URL.Path)
	http.ServeFile(w, r, "Pages/login.html")
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Serving request: %s\n, %s\n", r.Method, r.URL.Path)
	myvar := map[string]interface{}{"email": r.FormValue("username"), "pass": r.FormValue("password")}
	outputHTML(w, "Pages/index.html", myvar)
}

func outputHTML(w http.ResponseWriter, filename string, data interface{}) {
	t, err := template.ParseFiles(filename)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	if err := t.Execute(w, data); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
