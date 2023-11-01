package utils

type Response struct {
	Code    int32       `json:"code"`
	Flag    bool        `json:"flag"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}
