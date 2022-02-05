package handler

type Response struct {
	Code    int
	Message string
	Content interface{}
}

func BuildRespByErr(err error) *Response {
	return &Response{
		Code:    -1,
		Message: err.Error(),
	}
}

func BuildResp(c interface{}) *Response {
	return &Response{
		Content: c,
	}
}
