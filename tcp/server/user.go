package server

import (
	"encoding/json"
	"log"

	"github.com/gomodule/redigo/redis"
	uuid "github.com/satori/go.uuid"
)

func getUserDetail() {

}

func getUserLoginFromDB(user User) (result User, err error) {
	result = User{}
	query := `SELECT id, username, nickname, password, picture FROM user WHERE username=?`
	rows, err := db.Query(query, user.Username)
	if err != nil {
		log.Printf("[server][getUserLoginFromDB]Error when executing query. %+v\n", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&result.ID, &result.Username, &result.Nickname, &result.Password, &result.Picture)
		if err != nil {
			log.Printf("[server][getUserLoginFromDB]Error when scan query result. %+v\n", err)
			return
		}
	}
	return
}

func updateUserDetail(user User, url string) (err error) {
	query := `UPDATE user SET nickname = ?, picture = ? WHERE id = ?`

	_, err = db.Exec(query, user.Nickname, url, user.ID)
	if err != nil {
		log.Printf("[server][updateUserDetail]Error when executing update query. %+v\n", err)
		return
	}
	return
}

func updateUserNickname(user User) (err error) {
	query := `UPDATE user SET nickname = ? WHERE id = ?`

	_, err = db.Exec(query, user.Nickname, user.ID)
	if err != nil {
		log.Printf("[server][updateUserNickname]Error when executing update query. %+v\n", err)
		return
	}
	return
}

func getUserCookie(data TCPRequest) (resp TCPRequest, err error) {
	conn := redisPool.Get()
	defer conn.Close()

	result, err := redis.Bytes(conn.Do("GET", data.Cookie))
	if err != nil {
		log.Printf("[server][getUserCookie]Error get data from redis. %+v\n", err)
		resp.Error = err.Error()
		return
	}

	if result != nil {
		resp.HasActiveSession = true
	}

	user := User{}
	json.Unmarshal(result, &user)

	resp.User = user
	return
}

func setUserCookie(user User) (cookie string, err error) {
	sessionCookie, _ := uuid.NewV4()

	b, _ := json.Marshal(user)

	conn := redisPool.Get()
	defer conn.Close()

	_, err = conn.Do("SETEX", sessionCookie.String(), 1200, string(b))
	if err != nil {
		log.Printf("[server][setUserCookie]Error set data to redis. %+v\n", err)
		return
	}

	cookie = sessionCookie.String()
	return
}

func deleteUserCookie(cookie string) (err error) {
	conn := redisPool.Get()
	defer conn.Close()

	_, err = conn.Do("DEL", cookie)
	if err != nil {
		log.Printf("[server][deleteUserCookie]Error delete data from redis. %+v\n", err)
		return
	}

	return
}
