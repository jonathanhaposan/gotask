package webserver

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var (
	assetDirectory    = "../../file/asset"
	templateDirectory = "../../file/asset/html"
	tcpClient         *net.TCPConn
)

func InitRouter() (router *httprouter.Router) {
	router = httprouter.New()

	router.GET("/login", handlerGetLogin)
	router.POST("/login", handlerPostLogin)

	router.GET("/profile", handlerGetProfile)
	router.POST("/profile", handlerPostProfile)

	router.ServeFiles("/assets/*filepath", http.Dir(assetDirectory))

	return
}

func InitTCPClient() {
	target := "localhost:8081"

	raddr, err := net.ResolveTCPAddr("tcp", target)
	if err != nil {
		fmt.Println(err)
		return
	}

	conn, err := net.DialTCP("tcp", nil, raddr)
	if err != nil {
		fmt.Println(err)
		return
	}

	log.Println("Start TCP Client on :8081")
	tcpClient = conn
	tcpClient.SetKeepAlive(true)
}

func OpenConn() (conn *net.TCPConn) {
	target := "localhost:8081"

	raddr, err := net.ResolveTCPAddr("tcp", target)
	if err != nil {
		fmt.Println(err)
		return
	}

	conn, err = net.DialTCP("tcp", nil, raddr)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}
