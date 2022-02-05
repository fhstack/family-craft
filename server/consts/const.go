package consts

const (
	Code_SUCC = 0

	ErrCode_Unknown       = 1
	ErrCode_Param_Illegal = 2

	ErrCode_Locking = 100
)

var errMsgMap = map[int]string{
	ErrCode_Unknown: "未知错误",

	ErrCode_Param_Illegal: "参数错误",

	ErrCode_Locking: "锁正在使用中哦",
}

func GetCodeMsg(code int) string {
	return errMsgMap[code]
}
