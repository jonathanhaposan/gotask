package server

import (
	"encoding/json"
	"log"
	"net"

	"github.com/gomodule/redigo/redis"
	uuid "github.com/satori/go.uuid"
)

func HandleRequest(conn net.Conn) {
	var response TCPRequest

	requestTCP, err := ReadTCPData(conn)
	if err != nil {
		log.Println("Error read TCP data")
		return
	}
	defer conn.Close()

	log.Printf("%+v\n", requestTCP)

	switch requestTCP.RequestType {
	case RequestLogin:
		response = handleLogin(requestTCP)
	case RequestEdit:
		response = handleUpload(requestTCP)
	case RequestCheckCookie:
		response = getUserCookie(requestTCP)
	}

	err = SendTCPData(conn, response)
	if err != nil {
		log.Println("Error read TCP data")
		return
	}
}

func handleLogin(data TCPRequest) (resp TCPRequest) {
	result := getUserLoginFromDB(data.User)
	log.Printf("%+v\n", result)

	if len(result.Username) == 0 {
		log.Println("Username not found")
		return
	}

	if result.Password != data.User.Password {
		log.Println("Wrong password")
		return
	}

	resp.User = result

	resp = data
	cookie, err := setUserCookie(result)
	if err != nil {
		log.Println("err login", err)
		return
	}
	resp.Cookie = cookie
	return
}

func handleUpload(data TCPRequest) (resp TCPRequest) {
	// var picture multipart.File
	// if data.User.PictureMIME != nil {
	// 	picture = data.User.PictureMIME
	// 	defer picture.Close()

	// 	buff := make([]byte, 512) // docs tell that it take only first 512 bytes into consideration
	// 	if _, err := picture.Read(buff); err != nil {
	// 		fmt.Println(err) // do something with that error
	// 		return
	// 	}
	// 	fmt.Println(http.DetectContentType(buff)) // do something based on your detection.

	// }

	return
}

func getUserCookie(data TCPRequest) (resp TCPRequest) {
	conn := redisPool.Get()
	defer conn.Close()

	result, err := redis.Bytes(conn.Do("GET", data.Cookie))
	if err != nil {
		log.Println("Error get cookie from redis", err)
		return
	}

	if result != nil {
		resp.HasActiveSession = true
		log.Printf("This %s cookie is found", data.Cookie)
	}

	user := User{}
	json.Unmarshal(result, &user)

	resp.User = user

	log.Printf("get %+v\n", user)
	log.Printf("get2 %+v\n", result)
	return
}

func setUserCookie(user User) (cookie string, err error) {
	sessionCookie, _ := uuid.NewV4()

	b, _ := json.Marshal(user)

	conn := redisPool.Get()
	defer conn.Close()

	_, err = conn.Do("SETEX", sessionCookie.String(), 120, string(b))
	if err != nil {
		log.Println("Error set cookie from redis:", err)
		return
	}

	cookie = sessionCookie.String()
	return
}
