package webserver

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/jonathanhaposan/gotask/tcp/server"
)

func validateLogin(username, password string, reqType int) (data server.TCPRequest, err error) {
	if len(username) == 0 {
		err = errors.New("username empty")
		return
	}

	if len(password) == 0 {
		err = errors.New("password empty")
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

func errorJSONResponse(w http.ResponseWriter, err string) {
	resp := ResponseJSON{
		Status: http.StatusInternalServerError,
		Error:  err,
	}

	j, _ := json.Marshal(resp)
	w.WriteHeader(resp.Status)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}
