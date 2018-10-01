package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net"

	"github.com/jonathanhaposan/gotask/tcp/server"
)

// only needed below for sample processing

func main() {

	fmt.Println("Launching server...")

	// listen on all interfaces
	ln, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Println(err)
	}

	defer func() {
		log.Println("Closed")
		ln.Close()
	}()
	// run loop forever (or until ctrl-c)
	for {
		// accept connection on port
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
		}

		requestTCP := server.TCPRequest{}
		decoder := gob.NewDecoder(conn)
		decoder.Decode(&requestTCP)

		log.Printf("%+v\n", requestTCP)

		requestTCP.RequestType = 123

		encoder := gob.NewEncoder(conn)
		encoder.Encode(requestTCP)

		// go server.HandleRequest(conn)
	}
}
