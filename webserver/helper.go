package webserver

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/jonathanhaposan/gotask/tcp/server"
)

var (
	errorEmptyUsername = errors.New("username is empty")
	errorEmptyPassword = errors.New("password is empty")
)

func validateLogin(username, password string, reqType int) (data server.TCPRequest, err error) {
	if len(username) == 0 {
		err = errorEmptyUsername
		log.Printf("[webserver][validateLogin]Username should not empty. %+v\n", err)
		return
	}

	if len(password) == 0 {
		err = errorEmptyPassword
		log.Printf("[webserver][validateLogin]Password should not empty. %+v\n", err)
		return
	}

	data = server.TCPRequest{
		RequestType: reqType,
		User: server.User{
			Username: username,
			Password: password,
		},
	}

	return
}

func JSONResponse(w http.ResponseWriter, result interface{}, err string) (e error) {
	if result == nil {
		resp := ResponseJSON{
			Status: http.StatusInternalServerError,
			Error:  err,
		}

		j, _ := json.Marshal(resp)
		w.WriteHeader(resp.Status)
		w.Header().Set("Content-Type", "application/json")
		w.Write(j)
	} else {
		resp := ResponseJSON{
			Status: http.StatusOK,
		}

		j, _ := json.Marshal(resp)
		w.WriteHeader(resp.Status)
		w.Header().Set("Content-Type", "application/json")
		w.Write(j)
	}
	return
}
