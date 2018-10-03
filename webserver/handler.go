package webserver

import (
	"bytes"
	"html/template"
	"io"
	"log"
	"mime"
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
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
	}
	requestTCP, err := validateLogin(r.FormValue("username"), r.FormValue("password"), server.RequestLogin)
	if err != nil {
		log.Println("Error validate data:", err)
		errorJSONResponse(w, err.Error())
		return
	}

	conn := OpenConn()

	err = server.SendTCPData(conn, requestTCP)
	if err != nil {
		log.Println("Error Send data to server:", err)
		errorJSONResponse(w, err.Error())
		return
	}

	response, err := server.ReadTCPData(conn)
	if err != nil {
		log.Println("Error Read data from server:", err)
		errorJSONResponse(w, err.Error())
		return
	}
	conn.Close()

	if len(response.Error) != 0 {
		log.Println("Error from server:", response.Error)
		errorJSONResponse(w, response.Error)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "session_cookie",
		Value:   response.Cookie,
		Expires: time.Now().Add(120 * time.Second),
	})
}

func handlerGetProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	cookie, err := r.Cookie("session_cookie")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	conn := OpenConn()
	user, _ := checkUserSession(conn, cookie)
	conn.Close()

	tmpl := template.Must(template.ParseFiles(templateDirectory + "/profile.html"))
	tmpl.Execute(w, user)
}

func handlerPostProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var rawPicture server.UploadedPicture

	cookie, err := r.Cookie("session_cookie")
	if err != nil {
		http.Error(w, "Unauthorized - Session Expired", http.StatusUnauthorized)
		return
	}

	conn := OpenConn()
	userSession, hasSession := checkUserSession(conn, cookie)
	conn.Close()

	if !hasSession {
		http.Error(w, "Unauthorized - Session Expired", http.StatusUnauthorized)
		return
	}

	nickname := r.FormValue("nickname")
	if len(nickname) == 0 {
		errorJSONResponse(w, "Nickname cannot empty")
		return
	}

	file, head, err := r.FormFile("picture") // img is the key of the form-data
	if err != nil {
		if err.Error() != http.ErrMissingFile.Error() {
			errorJSONResponse(w, err.Error())
			return
		}
	}

	if file != nil {
		buffer := bytes.NewBuffer(nil)
		_, err := io.Copy(buffer, file)
		if err != nil {
			log.Println("Error parse to buffer:", err)
			errorJSONResponse(w, err.Error())
			return
		}
		defer file.Close()

		switch http.DetectContentType(buffer.Bytes()) {
		case "image/jpeg", "image/jpg", "image/png":
			rawPicture.FileType = http.DetectContentType(buffer.Bytes())
			rawPicture.File = buffer.Bytes()
			rawPicture.FileSize = head.Size
			raw, _ := mime.ExtensionsByType(rawPicture.FileType)
			rawPicture.FileExt = raw[0]
		default:
			errorJSONResponse(w, "unknown file type uploaded")
			return
		}
	}

	userSession.Nickname = nickname

	requestTCP := server.TCPRequest{
		RequestType:     server.RequestEdit,
		UploadedPicture: rawPicture,
		User:            userSession,
	}

	conn = OpenConn()
	err = server.SendTCPData(conn, requestTCP)
	if err != nil {
		log.Println("Error Send data to server:", err)
		errorJSONResponse(w, err.Error())
		return
	}

	response, err := server.ReadTCPData(conn)
	if err != nil {
		log.Println("Error Read data from server:", err)
		errorJSONResponse(w, err.Error())
		return
	}
	conn.Close()

	if len(response.Error) != 0 {
		log.Println("Error from server:", response.Error)
		errorJSONResponse(w, response.Error)
		return
	}
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

	userData = response.User
	isActive = response.HasActiveSession
	return
}
