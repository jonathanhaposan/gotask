package main

import (
	"log"
	"net"

	"github.com/jonathanhaposan/gotask/tcp/server"
)

func main() {

	server.InitRedisConn()

	log.Println("Start TCP server on :8081")

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

/**
tcpAddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:9999")
    if err != nil {
        return
    }
    listener, err := net.ListenTCP("tcp", tcpAddr)
    if err != nil {
        return
	}

	https://stackoverflow.com/questions/38646224/golang-tcp-client-does-not-receive-data-from-server-hangs-blocks-on-conn-read/38650064
**/
