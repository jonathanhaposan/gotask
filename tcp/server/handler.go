package server

import (
	"io/ioutil"
	"log"
	"net"
)

func HandleRequest(conn net.Conn) {
	var response TCPRequest

	requestTCP, err := ReadTCPData(conn)
	if err != nil {
		log.Println("Error read TCP data")
		return
	}
	defer conn.Close()

	switch requestTCP.RequestType {
	case RequestLogin:
		response = handleLogin(requestTCP)
	case RequestEdit:
		response = handleUpload(requestTCP)
	case RequestCheckCookie:
		response, _ = getUserCookie(requestTCP)
	}

	err = SendTCPData(conn, response)
	if err != nil {
		log.Println("Error read TCP data")
		return
	}
}

func handleLogin(data TCPRequest) (resp TCPRequest) {
	result, err := getUserLoginFromDB(data.User)
	if err != nil {
		log.Println("Error Get DB", err)
		resp.Error = err.Error()
		return
	}

	if len(result.Username) == 0 {
		log.Println("Username not found")
		resp.Error = err.Error()
		return
	}

	if result.Password != data.User.Password {
		log.Println("Wrong password")
		resp.Error = err.Error()
		return
	}

	cookie, err := setUserCookie(result)
	if err != nil {
		log.Println("err login", err)
		resp.Error = err.Error()
		return
	}

	resp = data
	resp.User = result
	resp.Cookie = cookie
	return
}

func handleUpload(data TCPRequest) (resp TCPRequest) {
	if len(data.UploadedPicture.File) > 0 {
		path := imageDirectory + "/" + data.User.Username + data.UploadedPicture.FileExt
		url := imageURL + data.User.Username + data.UploadedPicture.FileExt
		err := updateUserDetail(data.User, url)
		if err != nil {
			log.Println("Error update data", err)
			resp.Error = err.Error()
			return
		}

		err = ioutil.WriteFile(path, data.UploadedPicture.File, 0666)
		if err != nil {
			log.Println("Error write file", err)
			resp.Error = err.Error()
			return
		}
	} else {
		err := updateUserNickname(data.User)
		if err != nil {
			log.Println("Error update nickname", err)
			resp.Error = err.Error()
			return
		}
	}

	return
}
