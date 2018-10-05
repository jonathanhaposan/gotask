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
		log.Printf("[web server][main]Error start web server. %+v\n", err)
		return
	}
	log.Println("Server stoped")
}
