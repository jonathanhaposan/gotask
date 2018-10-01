package server

type TCPRequest struct {
	RequestType      int
	Cookie           string
	HasActiveSession bool
	User             User
}

type User struct {
	Username string
	Password string
	Nickname string
	Picture  string
}
