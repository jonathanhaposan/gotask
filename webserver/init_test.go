package webserver

import (
	"fmt"
	"io/ioutil"
	"net"
	"testing"
)

func TestInitTCPClient(t *testing.T) {
}

func TestOpenConn(t *testing.T) {
	message := "Hi there!\n"

	targetTCP = ":3000"

	go func() {
		conn := OpenConn()
		defer conn.Close()

		if _, err := fmt.Fprintf(conn, message); err != nil {
			t.Fatal(err)
		}
	}()

	l, err := net.Listen("tcp", ":3000")
	if err != nil {
		t.Fatal(err)
	}
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			return
		}
		defer conn.Close()

		buf, err := ioutil.ReadAll(conn)
		if err != nil {
			t.Fatal(err)
		}

		fmt.Println(string(buf[:]))
		if msg := string(buf[:]); msg != message {
			t.Fatalf("Unexpected message:\nGot:\t\t%s\nExpected:\t%s\n", msg, message)
		}
		return // Done
	}

}

func TestInitRouter(t *testing.T) {
	r := InitRouter()

	if r == nil {
		t.Fatal("Fail to init router")
	}
}
