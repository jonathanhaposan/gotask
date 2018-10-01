package webserver

import (
	"net/http"
	"strconv"
	"time"

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

	sessionToken := strconv.FormatInt(time.Now().Unix(), 10)

	http.SetCookie(w, &http.Cookie{
		Name:    "session_cookie",
		Value:   sessionToken,
		Expires: time.Now().Add(120 * time.Second),
	})
	http.Redirect(w, r, "/profile", http.StatusSeeOther)
}

func handlerGetProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

func handlerPostProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}
