package main

import (
	"log"
	"net"
)

var conn net.Conn

func main() {

	// connect to this socket
	conn, err := net.Dial("tcp", "127.0.0.1:8081")
	if err != nil {
		log.Println("Error connecting to tcp")
	}
	log.Println(conn)
}
