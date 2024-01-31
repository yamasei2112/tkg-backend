package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", hanlder)
	http.ListenAndServe(":8080", nil)
}

func hanlder(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}
