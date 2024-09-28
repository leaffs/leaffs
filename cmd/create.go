package cmd

import (
	"log"
	"net/http"
)

func CreateServer() error {
	router := http.NewServeMux()
	server := http.Server{
		Addr:    "127.0.0.1:9999",
		Handler: router,
	}
	log.Printf("Server has been created on %s", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Error starting server: %v", err)
		return err
	}
	return nil
}
