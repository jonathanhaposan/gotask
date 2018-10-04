package server

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/alicebob/miniredis"
	"github.com/gomodule/redigo/redis"
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

func TestGetUserCookie(t *testing.T) {
	testGetUserCookiePositive(t)
	testGetUserCookieNegative(t)
}

func mockRedis() *miniredis.Miniredis {
	s, err := miniredis.Run()
	if err != nil {
		panic(err)
	}

	redisPool = &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,

		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", s.Addr())
			if err != nil {
				log.Println("Redis Error", err)
				return nil, err
			}
			return c, err
		},

		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			if err != nil {
				log.Println("Error Ping Redis", err)
				return err
			}
			return err
		},
	}

	return s
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

	mock.ExpectQuery("SELECT (.+) FROM (.+)").WithArgs(user.Username).WillReturnRows(rows)

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

	mock.ExpectQuery("SELECT (.+) FROM (.+)").WithArgs(user.Username).WillReturnRows(rows)

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

func testGetUserCookiePositive(t *testing.T) {
	server := mockRedis()
	defer server.Close()

	c, err := redis.Dial("tcp", server.Addr())
	if err != nil {
		t.Errorf("Error dial Dummy")
	}

	_, err = c.Do("SETEX", "123", 10, "value")
	if err != nil {
		t.Errorf("Error Set Dummy")
	}

	request := TCPRequest{
		Cookie: "123",
	}

	_, err = getUserCookie(request)
	if err != nil {
		t.Errorf("Error were not Expected")
	}

	server.CheckGet(t, "123", "value")
}

func testGetUserCookieNegative(t *testing.T) {
	server := mockRedis()
	defer server.Close()

	request := TCPRequest{
		Cookie: "123",
	}

	_, err := getUserCookie(request)
	if err == nil {
		t.Errorf("Error were Expected")
	}
}

func TestSetUserCookie(t *testing.T) {
	server := mockRedis()
	defer server.Close()

	user := User{}

	expected, _ := json.Marshal(user)

	cookie, err := setUserCookie(user)
	if err != nil {
		t.Errorf("Error were not Expected")
	}

	server.CheckGet(t, cookie, string(expected))

	server.FastForward(1201 * time.Second)
	if server.Exists(cookie) {
		t.Errorf("Key should disappear")
	}
}

func testDeleteUserCookiePos(t *testing.T) {
	server := mockRedis()
	defer server.Close()

	request := TCPRequest{
		Cookie: "123",
	}

	server.Set(request.Cookie, "some value")
	server.SetTTL(request.Cookie, 1200*time.Second)

	err := deleteUserCookie(request.Cookie)
	if err != nil {
		t.Error("Error not expected")
	}
}
