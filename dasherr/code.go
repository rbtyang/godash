package dasherr

import (
	"github.com/rbtyang/godash/dashlog"
	"github.com/spf13/cast"
	"google.golang.org/grpc/codes"
	"math"
)

// 常规错误码 errdash code [10000~20000)
const (
	CodeOK                 = uint32(0)           //成功
	CodeUnknown            = 9999 + uint32(iota) //未知错误
	CodeCanceled                                 // 操作被取消
	CodeInvalidArgument                          // 无效参数
	CodeDeadlineExceeded                         // 操作超过截止日期
	CodeNotFound                                 // 请求实体未找到
	CodeAlreadyExists                            // 请求实体已存在
	CodePermissionDenied                         // 权限拒绝
	CodeResourceExhausted                        // 资源已经耗尽
	CodeFailedPrecondition                       // 先决条件失败
	CodeAborted                                  // 操作已中止
	CodeOutOfRange                               // 操作值溢出
	CodeUnimplemented                            // 未实现操作
	CodeInternal                                 // 内部错误
	CodeUnavailable                              // 服务当前不可用
	CodeDataLoss                                 // 数据丢失或损坏
	CodeUnauthenticated                          // 未认证或凭证无效
	CodeExternalGrpcError                        // 外部通信错误
	CodeMax                = 19999               // Dasherr Max Code
)

// errdash code => msgtext
var code2text = map[uint32]string{
	CodeOK:                 "成功",
	CodeUnknown:            "未知错误",
	CodeCanceled:           "操作被取消",
	CodeInvalidArgument:    "无效参数",
	CodeDeadlineExceeded:   "操作超过截止日期",
	CodeNotFound:           "请求实体未找到",
	CodeAlreadyExists:      "请求实体已存在",
	CodePermissionDenied:   "权限拒绝",
	CodeResourceExhausted:  "资源已经耗尽",
	CodeFailedPrecondition: "先决条件失败",
	CodeAborted:            "操作已中止",
	CodeOutOfRange:         "操作值溢出",
	CodeUnimplemented:      "未实现操作",
	CodeInternal:           "内部错误",
	CodeUnavailable:        "服务当前不可用",
	CodeDataLoss:           "数据丢失或损坏",
	CodeUnauthenticated:    "未认证或凭证无效",
	CodeExternalGrpcError:  "外部通信错误",
	CodeMax:                "Dasherr Max Code",
}

// grpc code => errdash code
var grpc2code = map[codes.Code]uint32{
	codes.OK:                 CodeOK,
	codes.Canceled:           CodeCanceled,
	codes.Unknown:            CodeUnknown,
	codes.InvalidArgument:    CodeInvalidArgument,
	codes.DeadlineExceeded:   CodeDeadlineExceeded,
	codes.NotFound:           CodeNotFound,
	codes.AlreadyExists:      CodeAlreadyExists,
	codes.PermissionDenied:   CodePermissionDenied,
	codes.ResourceExhausted:  CodeResourceExhausted,
	codes.FailedPrecondition: CodeFailedPrecondition,
	codes.Aborted:            CodeAborted,
	codes.OutOfRange:         CodeOutOfRange,
	codes.Unimplemented:      CodeUnimplemented,
	codes.Internal:           CodeInternal,
	codes.Unavailable:        CodeUnavailable,
	codes.DataLoss:           CodeDataLoss,
	codes.Unauthenticated:    CodeUnauthenticated,
}

/*
RegisterCode @Editor robotyang at 2023

# RegisterCode 注册错误字典

@Param dict 要注册的错误字段（建议尽量注册 >= 20000 的 code）
*/
func RegisterCode(dict map[uint32]string) {
	for code, errMsg := range dict {
		code2text[code] = errMsg
	}
}

/*
ParseCode @Editor robotyang at 2023

# ParseCode 解析 code 对应的 errdash code

@Param code 可能是 grpc code、errdash code 的 uint32 或 codes.Code;
*/
func ParseCode(code interface{}) uint32 {
	var cCode codes.Code = math.MaxUint32
	switch code.(type) {
	case nil:
		cCode = codes.Internal
		dashlog.Errorf("ParseCode code is nil, code:%#v", code)
	case codes.Code:
		cCode = code.(codes.Code)
	default:
		co := cast.ToUint32(code)
		if co == uint32(0) {
			dashlog.Errorf("ParseCode code is not uint32, code:%#v", code)
			cCode = codes.Internal
		} else {
			cCode = codes.Code(co)
		}
	}
	if co, ok := grpc2code[cCode]; ok {
		return co //存在映射
	} else {
		return cast.ToUint32(code) //不存在映射，则认为是 errdash code，原样返回
	}
}

/*
GetCodeMsg @Editor robotyang at 2023

# GetCodeMsg 转义 code 为对应中文含义

@Param code 可能是 grpc code、errdash code 的 uint32 或 codes.Code;
*/
func GetCodeMsg(code uint32) string {
	if txt, ok := code2text[code]; ok {
		return txt
	} else {
		dashlog.Errorf("GetCodeMsg code not defined text, code:%#v", code)
		return code2text[CodeUnknown]
	}
}
