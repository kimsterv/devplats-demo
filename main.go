package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi DevPlats All Hands! Nice gitops, you've requested.")
	})
	fmt.Println("Listening on port 8000")
	http.ListenAndServe(":8000", nil)
}
