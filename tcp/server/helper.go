package server

import (
	"encoding/gob"
	"net"
)

func SendTCPData(conn net.Conn, data TCPRequest) {
	encoder := gob.NewEncoder(conn)
	encoder.Encode(data)
}

func ReadTCPData(conn net.Conn) (data TCPRequest) {
	data = TCPRequest{}
	decoder := gob.NewDecoder(conn)
	decoder.Decode(&data)
	return
}
