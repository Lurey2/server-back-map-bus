package conf

var Succes = "success"
var Error = "error"

type Response struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
	Error  ErrorType   `json:"error"`
	Code   uint        `json:"code"`
}

type ErrorType struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}
