package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/atyagi012/student-api/internal/config"
)

func main() {
	//load config
	cfg := config.MustLoad()

	//database setup

	//setup routes
	router := http.NewServeMux()
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	//setup server
	server := http.Server{
		Addr:    cfg.HTTPServer.Addr,
		Handler: router,
	}

	fmt.Printf("Server started at %s\n", cfg.HTTPServer.Addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Failed to start server")
	}

}
