package server

import (
	"net"
	"testing"
)

func TestReadSendTCPData(t *testing.T) {
	client, server := net.Pipe()

	go func() {
		conn := client
		defer conn.Close()

		if err := SendTCPData(conn, TCPRequest{}); err != nil {
			t.Fatal(err)
		}
	}()

	for {
		conn := server
		defer conn.Close()

		_, err := ReadTCPData(conn)
		if err != nil {
			t.Fatal(err)
		}

		return // Done
	}
}
