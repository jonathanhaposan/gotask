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
		log.Println("Error Query:", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&result.ID, &result.Username, &result.Nickname, &result.Password, &result.Picture)
		if err != nil {
			log.Println("Error Scan:", err)
			return
		}
	}

	return
}

func updateUserDetail(user User, url string) (err error) {
	query := `UPDATE user SET nickname = ?, picture = ? WHERE id = ?`

	_, err = db.Exec(query, user.Nickname, url, user.ID)
	if err != nil {
		log.Println("Error Query:", err)
		return
	}

	return
}

func updateUserNickname(user User) (err error) {
	query := `UPDATE user SET nickname = ? WHERE id = ?`

	_, err = db.Exec(query, user.Nickname, user.ID)
	if err != nil {
		log.Println("Error Query:", err)
		return
	}

	return
}

func getUserCookie(data TCPRequest) (resp TCPRequest, err error) {
	conn := redisPool.Get()
	defer conn.Close()

	result, err := redis.Bytes(conn.Do("GET", data.Cookie))
	if err != nil {
		log.Println("Error get cookie from redis", err)
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
		log.Println("Error set cookie from redis:", err)
		return
	}

	cookie = sessionCookie.String()
	return
}
