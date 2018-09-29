package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<center><h1>Welcome to my page<h1></center>")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server is running...")
	http.ListenAndServe("0.0.0.0:3000", nil)
}
