package server

type TCPRequest struct {
	RequestType      int
	Cookie           string
	HasActiveSession bool
	Error            string
	User             User
	UploadedPicture  UploadedPicture
}

type User struct {
	ID       int64
	Username string
	Password string
	Nickname string
	Picture  string
}

type UploadedPicture struct {
	File     []byte
	FileType string
	FileExt  string
	FileSize int64
}
