package main

import (
	"fmt"
	"log"
	"net/http"
)

var (
	port = 7777
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Recevied request with path: %s\n", r.URL.Path)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Hello world!!!!")
	})
	addr := fmt.Sprintf(":%d", port)
	log.Printf("Server is running on port %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
