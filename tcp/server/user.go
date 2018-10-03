package server

import (
	"log"
)

func getUserDetail() {

}

func getUserLoginFromDB(user User) (result User, err error) {
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
