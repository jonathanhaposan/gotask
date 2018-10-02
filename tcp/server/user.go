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

func updateUserDetail(user User) (err error) {
	query := `UPDATE user SET nickname = ?, picture = ? WHERE id = ?`

	_, err = db.Exec(query, user.Nickname, user.Picture)
	if err != nil {
		log.Println("Error Query:", err)
		return
	}

	return
}

func savePictureToServer(raw UploadedPicture) (err error) {

	// file := raw.File
	// defer file.Close()

	// f, err := os.OpenFile(imageDirectory+"/"+user.Username, os.O_WRONLY|os.O_CREATE, 0666)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// defer f.Close()
	// io.Copy(f, file)

	return
}
