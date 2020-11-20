package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprint(w, "Hello, World")
	})
	http.ListenAndServe(":8080", nil)
}
