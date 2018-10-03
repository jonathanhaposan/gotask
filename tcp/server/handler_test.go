package server

import (
	"testing"

	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestHandleLoginPositive(t *testing.T) {
	mockRedis()

	d, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Unexpected error %+v\n", err)
	}
	defer d.Close()
	db = d

	user := User{Username: "asd", Password: "asd"}
	request := TCPRequest{User: user}

	rows := sqlmock.NewRows([]string{"id", "username", "nickname", "password", "picture"}).
		AddRow(1, "asd", "asd", "asd", "/asd/asd.png")

	mock.ExpectQuery("SELECT (.+) FROM (.+)").WillReturnRows(rows)

	resp := handleLogin(request)

	if len(resp.Error) > 0 {
		t.Errorf("Error Not Expected")
	}
}

func TestHandleLoginNegativeUsernameNotFound(t *testing.T) {
	mockRedis()

	d, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Unexpected error %+v\n", err)
	}
	defer d.Close()
	db = d

	user := User{Username: "asd", Password: "asd"}
	request := TCPRequest{User: user}

	rows := sqlmock.NewRows([]string{"id", "username", "nickname", "password", "picture"}).
		AddRow(1, "asdxxx", "asd", "asd", "/asd/asd.png")

	mock.ExpectQuery("SELECT (.+) FROM (.+)").WillReturnRows(rows)

	resp := handleLogin(request)

	if len(resp.Error) > 0 {
		t.Errorf("Error Not Expected")
	}
}

func TestHandleLoginNegativeWrongPass(t *testing.T) {
	mockRedis()

	d, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Unexpected error %+v\n", err)
	}
	defer d.Close()
	db = d

	user := User{Username: "asd", Password: "asd"}
	request := TCPRequest{User: user}

	rows := sqlmock.NewRows([]string{"id", "username", "nickname", "password", "picture"}).
		AddRow(1, "asd", "asd", "asdxxxx", "/asd/asd.png")

	mock.ExpectQuery("SELECT (.+) FROM (.+)").WillReturnRows(rows)

	resp := handleLogin(request)

	if len(resp.Error) > 0 {
		t.Errorf("Error Not Expected")
	}
}

func TestHandleLoginNegativeErrorRedis(t *testing.T) {
	d, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Unexpected error %+v\n", err)
	}
	defer d.Close()
	db = d

	user := User{Username: "asd", Password: "asd"}
	request := TCPRequest{User: user}

	rows := sqlmock.NewRows([]string{"id", "username", "nickname", "password", "picture"}).
		AddRow(1, "asd", "asd", "asd", "/asd/asd.png")

	mock.ExpectQuery("SELECT (.+) FROM (.+)").WillReturnRows(rows)

	resp := handleLogin(request)

	if len(resp.Error) > 0 {
		t.Errorf("Error Not Expected")
	}

}
