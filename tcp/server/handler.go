package server

import (
	"bufio"
	"encoding/json"
	"log"
	"net"
)

func HandleRequest(conn net.Conn) {
	var requestTCP TCPRequest

	reader := bufio.NewReader(conn)
	buf, _ := reader.ReadBytes('\n')

	json.Unmarshal(buf, requestTCP)

	log.Printf("%+v\n", requestTCP)
	conn.Write([]byte("ok\n"))
}
