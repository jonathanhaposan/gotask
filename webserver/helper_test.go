package webserver

import (
	"net/http/httptest"
	"testing"

	"github.com/jonathanhaposan/gotask/tcp/server"
)

func Test_errorJSONRespose(t *testing.T) {
	rw := httptest.NewRecorder()
	e := "error"

	if err := errorJSONResponse(rw, e); err != nil {
		t.Errorf("Error were not expected")
	}
}

func Test_validateLogin(t *testing.T) {
	username := ""
	password := ""

	_, err := validateLogin(username, password, server.RequestLogin)
	if err == nil {
		t.Errorf("Error were expected")
	}

	username = "asd"
	_, err = validateLogin(username, password, server.RequestLogin)
	if err == nil {
		t.Errorf("Error were expected")
	}

	username = "asd"
	password = "asd"
	_, err = validateLogin(username, password, server.RequestLogin)
	if err != nil {
		t.Errorf("Error were not expected")
	}
}
