package webserver

import (
	"log"
	"net"

	"github.com/julienschmidt/httprouter"
)

var tcpClient net.Conn

func InitRouter() (router *httprouter.Router) {
	router = httprouter.New()

	router.GET("/login", handlerGetLogin)
	router.POST("/login", handlerPostLogin)

	router.GET("/profile", handlerGetProfile)
	router.POST("/profile", handlerPostProfile)

	return
}

func InitTCPClient() {
	conn, err := net.Dial("tcp", ":8081")
	if err != nil {
		log.Println("Error connecting to tcp")
	}
	tcpClient = conn
}
