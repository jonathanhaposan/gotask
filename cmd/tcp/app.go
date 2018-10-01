package main

import (
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
		log.Println("TCP Server Stop...")
		ln.Close()
	}()
	// run loop forever (or until ctrl-c)
	for {
		// accept connection on port
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
		}

		go server.HandleRequest(conn)
	}
}
