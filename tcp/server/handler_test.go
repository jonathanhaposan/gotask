package server

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestHandleLogin(t *testing.T) {
	testHandleLoginNegativeErrorRedis(t)
	testHandleLoginNegativeWrongPass(t)
	testHandleLoginNegativeUsernameNotFound(t)
	testHandleLoginPositive(t)
	testHandleLoginNegativeQueryError(t)
}

func TestHandleUpload(t *testing.T) {
	testHandleUploadNegative(t)
	testHandleUploadPotive(t)
	testHandleUploadNegativeOnlyNickname(t)
	testHandleUploadPotiveOnlyNickname(t)
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

func testHandleLoginNegativeQueryError(t *testing.T) {
	mockRedis()

	d, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Unexpected error %+v\n", err)
	}
	defer d.Close()
	db = d

	user := User{Username: "asd", Password: "asd"}
	request := TCPRequest{User: user}

	mock.ExpectQuery("SELECT (.+) FROM (.+)").WillReturnError(fmt.Errorf("error"))

	resp := handleLogin(request)

	if len(resp.Error) == 0 {
		t.Errorf("Error were Expected")
	}
}

func testHandleUploadPotiveOnlyNickname(t *testing.T) {
	d, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Unexpected error %+v\n", err)
	}
	defer d.Close()
	db = d

	user := User{ID: 1, Nickname: "asda"}
	request := TCPRequest{User: user}

	mock.ExpectExec("UPDATE user").WithArgs(user.Nickname, user.ID).WillReturnResult(sqlmock.NewResult(1, 1))

	resp := handleUpload(request)
	if len(resp.Error) > 0 {
		t.Errorf("Error were not expected")
	}
}

func testHandleUploadNegativeOnlyNickname(t *testing.T) {
	d, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Unexpected error %+v\n", err)
	}
	defer d.Close()
	db = d

	user := User{ID: 1, Nickname: "asda"}
	request := TCPRequest{User: user}

	mock.ExpectExec("UPDATE user").WithArgs(user.Nickname, user.ID).WillReturnError(fmt.Errorf("Error"))

	resp := handleUpload(request)
	if len(resp.Error) == 0 {
		t.Errorf("Error were expected")
	}
}

func testHandleUploadPotive(t *testing.T) {
	d, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Unexpected error %+v\n", err)
	}
	defer d.Close()
	db = d

	file, err := ioutil.ReadFile(imageDirectory + "/ipsum.png")
	if err != nil {
		t.Fatal(err)
	}

	uploadPic := UploadedPicture{File: file, FileExt: ".png"}
	user := User{ID: 1, Username: "temp", Nickname: "temp"}
	request := TCPRequest{User: user, UploadedPicture: uploadPic}
	url := imageURL + request.User.Username + request.UploadedPicture.FileExt

	mock.ExpectExec("UPDATE user").WithArgs(user.Nickname, url, user.ID).WillReturnResult(sqlmock.NewResult(1, 1))

	resp := handleUpload(request)
	if len(resp.Error) > 0 {
		t.Errorf("Error were not expected")
	}

	os.Remove(imageDirectory + "/temp.png")
}

func testHandleUploadNegative(t *testing.T) {
	d, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Unexpected error %+v\n", err)
	}
	defer d.Close()
	db = d

	uploadPic := UploadedPicture{File: []byte{1}, FileExt: ".png"}
	user := User{ID: 1, Username: "temp", Nickname: "temp"}
	request := TCPRequest{User: user, UploadedPicture: uploadPic}
	url := imageURL + request.User.Username + request.UploadedPicture.FileExt

	mock.ExpectExec("UPDATE user").WithArgs(user.Nickname, url, user.ID).WillReturnError(fmt.Errorf("error"))

	resp := handleUpload(request)
	if len(resp.Error) == 0 {
		t.Errorf("Error not expected")
	}
}
