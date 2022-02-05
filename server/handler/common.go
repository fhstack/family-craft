package handler

import "github.com/fhstack/family-craft/server/consts"

type Response struct {
	Code    int
	Message string
	Content interface{}
}

func BuildRespByErr(err error) *Response {
	return &Response{
		Code:    consts.ErrCode_Unknown,
		Message: err.Error(),
	}
}

func BuildRespByCode(code int) *Response {
	return &Response{
		Code:    code,
		Message: consts.GetCodeMsg(code),
	}
}

func BuildParamIllegalResp(errMsg string) *Response {
	r := &Response{
		Code: consts.ErrCode_Param_Illegal,
	}
	r.Message = consts.GetCodeMsg(r.Code)
	if errMsg != "" {
		r.Message = errMsg
	}
	return r
}

func BuildSuccResp(c interface{}) *Response {
	return &Response{
		Code:    consts.Code_SUCC,
		Content: c,
	}
}
