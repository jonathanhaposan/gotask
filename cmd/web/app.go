package main

import (
	"log"
	"net/http"

	"github.com/jonathanhaposan/gotask/webserver"
)

func main() {

	router := webserver.InitRouter()

	s := &http.Server{
		Addr:    "localhost:9000",
		Handler: router,
	}

	log.Println("Start server on :9000")
	if err := s.ListenAndServe(); err != nil {
		log.Println("Server error:", err)
	}
	log.Println("Server stoped")
}
