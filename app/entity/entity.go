package entity

type Response struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

type ResponseWithHTTPStatus struct {
	HttpStatus int      `json:"http_status"`
	Details    Response `json:"details"`
}
