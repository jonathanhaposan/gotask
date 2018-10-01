package server

import (
	"encoding/gob"
	"log"
	"net"
)

func SendTCPData(conn net.Conn, data TCPRequest) (err error) {
	encoder := gob.NewEncoder(conn)
	err = encoder.Encode(data)
	if err != nil {
		log.Println("Error encode TCPData", err)
		return
	}
	return
}

func ReadTCPData(conn net.Conn) (data TCPRequest, err error) {
	data = TCPRequest{}
	decoder := gob.NewDecoder(conn)
	err = decoder.Decode(&data)
	if err != nil {
		log.Println("Error decode TCPData", err)
		return
	}
	return
}
