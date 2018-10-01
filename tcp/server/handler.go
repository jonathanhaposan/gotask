package server

import (
	"log"
	"net"

	uuid "github.com/satori/go.uuid"
)

func HandleRequest(conn net.Conn) {
	requestTCP := ReadTCPData(conn)

	log.Printf("%+v\n", requestTCP)

	response := TCPRequest{}

	switch requestTCP.RequestType {
	case 1:
		response = handleLogin(requestTCP)
	case 2:
		response = handleUpload(requestTCP)
	case 3:
		response = getUserCookie(requestTCP)
	}

	SendTCPData(conn, response)
}

func handleLogin(data TCPRequest) (resp TCPRequest) {
	resp = data
	asd := setUserCookie()
	resp.Cookie = asd
	log.Println(asd)
	return
}

func handleUpload(data TCPRequest) (resp TCPRequest) {

	return
}

func getUserCookie(data TCPRequest) (resp TCPRequest) {

	return
}

func setUserCookie() (cookie string) {
	sessionCookie, _ := uuid.NewV4()
	cookie = sessionCookie.String()
	log.Println(cookie)
	return
}
