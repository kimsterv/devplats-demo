package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>Happy Sunday! Nice gitops, you've requested.</h1>")
	})
	fmt.Println("Listening on port 8000")
	http.ListenAndServe(":8000", nil)
}
