package server

import (
	"testing"

	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestHandleLogin(t *testing.T) {
	testHandleLoginNegativeErrorRedis(t)
	testHandleLoginNegativeWrongPass(t)
	testHandleLoginNegativeUsernameNotFound(t)
	testHandleLoginPositive(t)
}

func testHandleLoginPositive(t *testing.T) {
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

func testHandleLoginNegativeUsernameNotFound(t *testing.T) {
	mockRedis()

	d, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Unexpected error %+v\n", err)
	}
	defer d.Close()
	db = d

	user := User{Username: "asd", Password: "asd"}
	request := TCPRequest{User: user}

	rows := sqlmock.NewRows([]string{"id", "username", "nickname", "password", "picture"})

	mock.ExpectQuery("SELECT (.+) FROM (.+)").WithArgs(user.Username).WillReturnRows(rows)

	resp := handleLogin(request)

	if len(resp.Error) == 0 {
		t.Errorf("Error were Expected")
	}
}

func testHandleLoginNegativeWrongPass(t *testing.T) {
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

	if len(resp.Error) == 0 {
		t.Errorf("Error were Expected")
	}
}

func testHandleLoginNegativeErrorRedis(t *testing.T) {
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
