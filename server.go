package main

import (
	"fmt"
	_ "encoding/json"
	"net/http"
)

func welcome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, "Welcome to our website")
	return
}

func main() {
	http.HandleFunc("/", welcome)
	
	http.ListenAndServe(":3000", nil)
}