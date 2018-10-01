package webserver

import (
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/jonathanhaposan/gotask/tcp/server"
	"github.com/julienschmidt/httprouter"
)

func handlerGetLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// _, err := r.Cookie("session_cookie")
	// if err == nil {
	// 	http.Redirect(w, r, "/profile", http.StatusSeeOther)
	// 	return
	// }

	tmpl := template.Must(template.ParseFiles(templateDirectory + "/login.html"))
	tmpl.Execute(w, nil)
}

func handlerPostLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	requestTCP := server.TCPRequest{}

	cookie, err := r.Cookie("session_cookie")
	if err == nil {
		requestTCP.Cookie = cookie.Value
		requestTCP.RequestType = server.RequestCheckCookie

		err = server.SendTCPData(tcpClient, requestTCP)
		log.Println("asdasdasda")
		if err != nil {
			log.Println("Error Send data to server")
			return
		}

		response, err := server.ReadTCPData(tcpClient)
		log.Println("xxxxxxxxxxxxxxxxxxxx")
		if err != nil {
			log.Println("Error Read data from server")
			return
		}

		if response.HasActiveSession {
			http.Redirect(w, r, "/profile", http.StatusSeeOther)
			return
		}
	}

	requestTCP, err = validateLogin(r.FormValue("username"), r.FormValue("password"), server.RequestLogin)
	if err != nil {
		log.Println("Error validate data")
		return
	}

	err = server.SendTCPData(tcpClient, requestTCP)
	if err != nil {
		log.Println("Error Send data to server")
		return
	}

	response, err := server.ReadTCPData(tcpClient)
	if err != nil {
		log.Println("Error Read data from server")
		return
	}

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
