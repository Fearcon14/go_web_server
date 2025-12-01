package main

import (
	"fmt"
	"net/http"
)

func main() {
	serveMux := http.NewServeMux()
	serveMux.Handle("/", http.FileServer(http.Dir("../../web")))

	server := &http.Server{
		Handler: serveMux,
		Addr:    ":8080",
	}

	fmt.Println("Server is running on port 8080")
	server.ListenAndServe()
}
