package main

import (
	"log"
	"net"

	"github.com/jonathanhaposan/gotask/tcp/server"
)

func main() {

	server.InitRedisConn()
	server.InitDBConn()

	log.Println("Start TCP server on :8081")

	// listen on all interfaces
	// ln, err := net.Listen("tcp", "localhost:8081")
	// if err != nil {
	// 	log.Println(err)
	// }

	// defer func() {
	// 	log.Println("TCP Server Stop...")
	// 	ln.Close()
	// }()
	// // run loop forever (or until ctrl-c)
	// for {
	// 	// accept connection on port
	// 	conn, err := ln.Accept()
	// 	if err != nil {
	// 		log.Println(err)
	// 	}

	// 	go server.HandleRequest(conn)
	// }

	tcpAddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:8081")
	if err != nil {
		log.Printf("[tcp server][main]Error Resolve TCP Addr. %+v\n", err)
		return
	}

	listener, err := net.ListenTCP("tcp4", tcpAddr)
	if err != nil {
		log.Printf("[tcp server][main]Error initiate listener to address. %+v\n", err)
		return
	}
	defer listener.Close()

	for {
		// accept connection on port
		conn, err := listener.AcceptTCP()
		if err != nil {
			log.Printf("[tcp server][main]Error when accepting connection. %+v\n", err)
			return
		}

		go server.HandleRequest(conn)
	}
}
