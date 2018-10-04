package webserver

import (
	"net"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jonathanhaposan/gotask/tcp/server"
	"github.com/julienschmidt/httprouter"
)

func getRequest(t testing.TB, url string) *http.Request {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}
	return req
}

func Test_handlerGetLogin(t *testing.T) {
	templateDirectory = "../file/asset/html"
	r := getRequest(t, "/login")

	rw := httptest.NewRecorder()
	handlerGetLogin(rw, r, httprouter.Params{})
}

func Test_handlerGetProfile(t *testing.T) {
	templateDirectory = "../file/asset/html"
	r := getRequest(t, "/profile")

	rw := httptest.NewRecorder()
	handlerGetProfile(rw, r, httprouter.Params{})
}

func Test_checkUserSession(t *testing.T) {
	request := server.TCPRequest{}
	httpCookie := http.Cookie{Value: "unique"}

	clientMock, serverMock := net.Pipe()

	go func() {
		conn := clientMock
		defer conn.Close()

		checkUserSession(conn, &httpCookie)

	}()

	for {
		conn := serverMock
		defer conn.Close()

		_, err := server.ReadTCPData(conn)
		if err != nil {
			t.Fatal(err)
		}

		request.Cookie = "unique"
		request.HasActiveSession = true

		err = server.SendTCPData(conn, request)
		if err != nil {
			t.Fatal(err)
		}

		return // Done
	}
}
