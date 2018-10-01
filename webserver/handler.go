package webserver

import (
	"encoding/gob"
	"log"
	"net/http"
	"time"

	"github.com/jonathanhaposan/gotask/tcp/server"
	"github.com/julienschmidt/httprouter"
	"github.com/satori/go.uuid"
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

	user := server.User{r.FormValue("username"), r.FormValue("password")}
	requestTCP := server.TCPRequest{User: user}

	encoder := gob.NewEncoder(tcpClient)
	encoder.Encode(requestTCP)

	decoder := gob.NewDecoder(tcpClient)
	decoder.Decode(&requestTCP)

	log.Printf("%+v\n", requestTCP)

	sessionToken, _ := uuid.NewV4()

	http.SetCookie(w, &http.Cookie{
		Name:    "session_cookie",
		Value:   sessionToken.String(),
		Expires: time.Now().Add(120 * time.Second),
	})
	http.Redirect(w, r, "/profile", http.StatusSeeOther)
}

func handlerGetProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

func handlerPostProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}
