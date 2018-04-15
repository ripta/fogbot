package client

type ResponseError struct {
	Message string `json:"message"`
	Detail  string `json:"detail"`
	Code    string `json:"code"`
}

func (e ResponseError) Error() string {
	return e.Message
}
