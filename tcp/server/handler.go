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
		log.Printf("[server][HandleRequest]Failed read incoming data. %+v\n", err)
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
		log.Printf("[server][HandleRequest]Failed send outcoming data. %+v\n", err)
		return
	}
}

func handleLogin(data TCPRequest) (resp TCPRequest) {
	result, err := getUserLoginFromDB(data.User)
	if err != nil {
		log.Printf("[server][handleLogin]Failed get user data from db. %+v\n", err)
		resp.Error = err.Error()
		return
	}

	if len(result.Username) == 0 {
		resp.Error = "Username not found"
		return
	}

	if result.Password != data.User.Password {
		resp.Error = "Wrong password"
		return
	}

	cookie, err := setUserCookie(result)
	if err != nil {
		log.Printf("[server][handleLogin]Failed set user cookie. %+v\n", err)
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
			log.Printf("[server][handleUpload]Failed update nickname and picture. %+v\n", err)
			resp.Error = err.Error()
			return
		}

		err = ioutil.WriteFile(path, data.UploadedPicture.File, 0666)
		if err != nil {
			log.Printf("[server][handleUpload]Failed when write file. %+v\n", err)
			resp.Error = err.Error()
			return
		}

		data.User.Picture = url
	} else {
		err := updateUserNickname(data.User)
		if err != nil {
			log.Printf("[server][handleUpload]Failed update nickname. %+v\n", err)
			resp.Error = err.Error()
			return
		}
	}

	resp = data
	err := deleteUserCookie(data.Cookie)
	if err != nil {
		log.Printf("[server][handleUpload]Failed delete user cookie. %+v\n", err)
		return
	}

	newCookie, err := setUserCookie(resp.User)
	if err != nil {
		log.Printf("[server][handleUpload]Failed set user cookie to redis. %+v\n", err)
		return
	}

	resp.Cookie = newCookie
	return
}
