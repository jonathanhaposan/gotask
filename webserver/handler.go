package webserver

import (
	"log"
	"net/http"
	"time"

	"github.com/jonathanhaposan/gotask/tcp/server"
	"github.com/julienschmidt/httprouter"
)

func handlerGetLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	_, err := r.Cookie("session_cookie")
	if err == nil {
		http.Redirect(w, r, "/profile", http.StatusSeeOther)
		return
	}

	w.Write([]byte("asdasdasd"))
}

func handlerPostLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	_, err := r.Cookie("session_cookie")
	if err == nil {
		http.Redirect(w, r, "/profile", http.StatusSeeOther)
		return
	}

	user := server.User{Username: r.FormValue("username"), Password: r.FormValue("password")}
	requestTCP := server.TCPRequest{RequestType: 1, User: user}

	server.SendTCPData(tcpClient, requestTCP)

	response := server.ReadTCPData(tcpClient)

	log.Printf("a %+v\n", response)

	http.SetCookie(w, &http.Cookie{
		Name:    "session_cookie",
		Value:   response.Cookie,
		Expires: time.Now().Add(120 * time.Second),
	})
	http.Redirect(w, r, "/profile", http.StatusSeeOther)
}

func handlerGetProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

func handlerPostProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}
