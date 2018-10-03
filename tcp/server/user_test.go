package server

import (
	"fmt"
	"testing"

	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestGetUserLoginFromDB(t *testing.T) {
	testGetUserLoginFromDBPositive(t)
	testGetUserLoginFromDBNegative(t)
}
func TestUpdateUserDetail(t *testing.T) {
	testUpdateUserDetailPositive(t)
	testUpdateUserDetailNegative(t)
}
func TestUpdateUserNickname(t *testing.T) {
	testUpdateUserNicknamePositive(t)
	testUpdateUserNicknameNegative(t)
}

func testGetUserLoginFromDBPositive(t *testing.T) {
	d, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Unexpected error %+v\n", err)
	}
	defer d.Close()
	db = d

	user := User{Nickname: "asd"}

	rows := sqlmock.NewRows([]string{"id", "username", "nickname", "password", "picture"}).
		AddRow(1, "asd", "asd", "asd", "/asd/asd.png")

	mock.ExpectQuery("SELECT (.+) FROM (.+)").WillReturnRows(rows)

	_, err = getUserLoginFromDB(user)
	if err != nil {
		t.Errorf("Error was not expected: %+v\n", err)
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectation: %+v\n", err)
	}
}
func testGetUserLoginFromDBNegative(t *testing.T) {
	d, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Unexpected error %+v\n", err)
	}
	defer d.Close()
	db = d

	user := User{Nickname: "asd"}

	rows := sqlmock.NewRows([]string{"id", "username", "nickname", "password", "picture"}).
		AddRow("trouble maker", 123, "asd", "asd", "/asd/asd.png")

	mock.ExpectQuery("SELECT (.+) FROM (.+)").WillReturnRows(rows)

	_, err = getUserLoginFromDB(user)
	if err == nil {
		t.Errorf("Error was expected: %+v\n", err)
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectation: %+v\n", err)
	}
}

func testUpdateUserDetailPositive(t *testing.T) {
	d, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Unexpected error %+v\n", err)
	}
	defer d.Close()
	db = d

	user := User{ID: 1, Nickname: "asda"}
	url := "/image/img.jpg"

	mock.ExpectExec("UPDATE user").WithArgs(user.Nickname, url, user.ID).WillReturnResult(sqlmock.NewResult(1, 1))

	if err = updateUserDetail(user, url); err != nil {
		t.Errorf("Error was not expected. %+v\n", err)
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectation: %+v\n", err)
	}
}
func testUpdateUserDetailNegative(t *testing.T) {
	d, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Unexpected error %+v\n", err)
	}
	defer d.Close()
	db = d

	user := User{ID: 1, Nickname: "asda"}
	url := "/image/img.jpg"

	mock.ExpectExec("UPDATE user").WithArgs(user.Nickname, url, user.ID).WillReturnError(fmt.Errorf("Error"))

	if err = updateUserDetail(user, url); err == nil {
		t.Errorf("Error was expected. %+v\n", err)
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectation: %+v\n", err)
	}
}

func testUpdateUserNicknamePositive(t *testing.T) {
	d, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Unexpected error %+v\n", err)
	}
	defer d.Close()
	db = d

	user := User{ID: 1, Nickname: "asda"}

	mock.ExpectExec("UPDATE user").WithArgs(user.Nickname, user.ID).WillReturnResult(sqlmock.NewResult(1, 1))

	if err = updateUserNickname(user); err != nil {
		t.Errorf("Error was not expected. %+v\n", err)
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectation: %+v\n", err)
	}
}
func testUpdateUserNicknameNegative(t *testing.T) {
	d, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Unexpected error %+v\n", err)
	}
	defer d.Close()
	db = d

	user := User{ID: 1, Nickname: "asda"}

	mock.ExpectExec("UPDATE user").WithArgs(user.Nickname, user.ID).WillReturnError(fmt.Errorf("Error"))

	if err = updateUserNickname(user); err == nil {
		t.Errorf("Error was expected. %+v\n", err)
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectation: %+v\n", err)
	}
}
