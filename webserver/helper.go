package webserver

import (
	"errors"

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
