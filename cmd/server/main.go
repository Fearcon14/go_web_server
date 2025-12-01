package main

import (
	"fmt"
	"net/http"

	"github.com/Fearcon14/go_web_server/cmd/internal/handlers"
)

func main() {
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/healthz", handlers.ReadinessHandler)
	serveMux.Handle("/app/", http.StripPrefix("/app", http.FileServer(http.Dir("../../web"))))

	server := &http.Server{
		Handler: serveMux,
		Addr:    ":8080",
	}

	fmt.Println("Server is running on port 8080")
	server.ListenAndServe()
}
