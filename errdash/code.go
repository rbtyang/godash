package errdash

// 常规错误码
const (
	ErrSuccess = 0
	ErrDefault = 10000000
	ErrParam   = 10000001
)

var errDict = map[uint32]string{
	ErrSuccess: "成功",
	ErrDefault: "默认错误",
	ErrParam:   "参数错误",
}

// 注册错误字典
func RegisterDict(dict map[uint32]string) {
	for code, errMsg := range dict {
		errDict[code] = errMsg
	}
}

// 转义code为对应中文含义
func TransCodeMsg(code uint32) string {
	return errDict[code]
}
