package config

import (
	"log"
	"net/http"
	"os"
	"time"
)

func StartServer(router http.Handler) {
	port := os.Getenv("PORT")
	srv := http.Server{
		Addr:         ":" + port,
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
	log.Println("Starting server on port " + port)
	log.Fatal(srv.ListenAndServe())
}
