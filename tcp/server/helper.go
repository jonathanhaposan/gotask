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
		log.Printf("[server][SendTCPData]Error encoding send data. %+v\n", err)
		return
	}
	return
}

func ReadTCPData(conn net.Conn) (data TCPRequest, err error) {
	data = TCPRequest{}
	decoder := gob.NewDecoder(conn)
	err = decoder.Decode(&data)
	if err != nil {
		log.Printf("[server][ReadTCPData]Error encoding read data. %+v\n", err)
		return
	}
	return
}
