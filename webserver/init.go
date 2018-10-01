package webserver

import (
	"github.com/julienschmidt/httprouter"
)

func InitRouter() (router *httprouter.Router) {
	router = httprouter.New()

	router.GET("/login", handlerGetLogin)
	router.POST("/login", handlerPostLogin)

	router.GET("/profile", handlerGetProfile)
	router.POST("/profile", handlerPostProfile)

	return
}
