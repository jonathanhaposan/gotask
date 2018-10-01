package server

import (
	"log"
	"net"

	"github.com/gomodule/redigo/redis"
	uuid "github.com/satori/go.uuid"
)

func HandleRequest(conn net.Conn) {
	var response TCPRequest

	requestTCP, err := ReadTCPData(conn)
	if err != nil {
		log.Println("Error read TCP data")
		return
	}

	log.Printf("%+v\n", requestTCP)

	switch requestTCP.RequestType {
	case RequestLogin:
		response = handleLogin(requestTCP)
	case RequestEdit:
		response = handleUpload(requestTCP)
	case RequestCheckCookie:
		response = getUserCookie(requestTCP)
	}

	err = SendTCPData(conn, response)
	if err != nil {
		log.Println("Error read TCP data")
		return
	}
}

func handleLogin(data TCPRequest) (resp TCPRequest) {
	resp = data
	cookie, err := setUserCookie(data.User.Username)
	if err != nil {
		log.Println("err login", err)
		return
	}
	resp.Cookie = cookie
	return
}

func handleUpload(data TCPRequest) (resp TCPRequest) {

	return
}

func getUserCookie(data TCPRequest) (resp TCPRequest) {
	conn := redisPool.Get()
	defer conn.Close()

	result, err := redis.Bytes(conn.Do("GET", data.Cookie))
	if err != nil {
		log.Println("Error get cookie from redis", err)
		return
	}

	if result != nil {
		resp.HasActiveSession = true
		log.Printf("This %s cookie is found", data.Cookie)
	}

	log.Println("get", result)
	return
}

func setUserCookie(username string) (cookie string, err error) {
	sessionCookie, _ := uuid.NewV4()

	conn := redisPool.Get()
	defer conn.Close()

	_, err = conn.Do("SETEX", sessionCookie.String(), 120, username)
	if err != nil {
		log.Println("Error set cookie from redis:", err)
		return
	}

	cookie = sessionCookie.String()
	return
}
