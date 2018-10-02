package webserver

type ResponseJSON struct {
	Status int         `json:"status"`
	Error  string      `json:"error"`
	Result interface{} `json:"result"`
}
