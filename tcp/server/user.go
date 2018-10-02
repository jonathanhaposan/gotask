package server

import (
	"log"
)

func getUserDetail() {

}

func getUserLoginFromDB(user User) (result User) {
	query := `SELECT id, username, password, nickname, picture FROM user WHERE username=?`
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

func updateUserDetail() {

}
