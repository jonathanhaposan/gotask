package webserver

import (
	"log"
	"net"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var (
	assetDirectory    = "file/asset"
	templateDirectory = "file/asset/html"
	imageDirectory    = "file/image"
	targetTCP         = ":8081"
	retries           = 5
)

func InitRouter() (router *httprouter.Router) {
	router = httprouter.New()

	router.GET("/login", handlerGetLogin)
	router.POST("/login", handlerPostLogin)

	router.GET("/profile", handlerGetProfile)
	router.POST("/profile", handlerPostProfile)

	router.ServeFiles("/assets/*filepath", http.Dir(assetDirectory))
	router.ServeFiles("/image/*filepath", http.Dir(imageDirectory))
	return
}

func OpenConn() (conn *net.TCPConn) {
	raddr, err := net.ResolveTCPAddr("tcp", targetTCP)
	if err != nil {
		log.Printf("[webserver][OpenConn]Failed resolving TCP Addr. %+v\n", err)
		return
	}

	conn, err = net.DialTCP("tcp", nil, raddr)
	if err != nil {
		log.Printf("[webserver][OpenConn]Failed dialing to TCP Server. %+v\n", err)
		return
	}
	return
}
