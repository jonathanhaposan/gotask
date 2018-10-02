package webserver

import (
	"html/template"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/jonathanhaposan/gotask/tcp/server"
	"github.com/julienschmidt/httprouter"
)

func handlerGetLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	cookie, err := r.Cookie("session_cookie")
	if err == nil {
		conn := OpenConn()
		_, hasSession := checkUserSession(conn, cookie)
		conn.Close()

		if hasSession {
			http.Redirect(w, r, "/profile", http.StatusSeeOther)
			return
		}
	}

	tmpl := template.Must(template.ParseFiles(templateDirectory + "/login.html"))
	tmpl.Execute(w, nil)
}

func handlerPostLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	cookie, err := r.Cookie("session_cookie")
	if err == nil {
		conn := OpenConn()
		_, hasSession := checkUserSession(conn, cookie)
		conn.Close()

		if hasSession {
			http.Redirect(w, r, "/profile", http.StatusSeeOther)
			return
		}
	}
	requestTCP, err := validateLogin(r.FormValue("username"), r.FormValue("password"), server.RequestLogin)
	if err != nil {
		log.Println("Error validate data")
		return
	}

	conn := OpenConn()

	err = server.SendTCPData(conn, requestTCP)
	if err != nil {
		log.Println("Error Send data to server")
		return
	}

	response, err := server.ReadTCPData(conn)
	if err != nil {
		log.Println("Error Read data from server")
		return
	}

	conn.Close()

	log.Printf("a %+v\n", response)

	http.SetCookie(w, &http.Cookie{
		Name:    "session_cookie",
		Value:   response.Cookie,
		Expires: time.Now().Add(120 * time.Second),
	})
}

func handlerGetProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	cookie, err := r.Cookie("session_cookie")
	if err != nil || err.Error() == http.ErrNoCookie.Error() {
		http.Redirect(w, r, "/profile", http.StatusSeeOther)
		return
	}

	conn := OpenConn()
	user, _ := checkUserSession(conn, cookie)
	conn.Close()

	log.Println(user)

	// w.Write([]byte("profile"))
}

func handlerPostProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

func checkUserSession(conn net.Conn, cookie *http.Cookie) (userData server.User, isActive bool) {
	requestTCP := server.TCPRequest{}
	requestTCP.Cookie = cookie.Value
	requestTCP.RequestType = server.RequestCheckCookie

	err := server.SendTCPData(conn, requestTCP)
	if err != nil {
		log.Println("Error Send data to server")
		return
	}

	response, err := server.ReadTCPData(conn)
	if err != nil {
		log.Println("Error Read data from server")
		return
	}

	isActive = response.HasActiveSession
	return
}
