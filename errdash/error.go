package errdash

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/rbtyang/godash/logdash"
	"google.golang.org/grpc/status"
)

//来访登记错误类 ----------------------------------------------------

type RpcError interface {
	GRPCStatus() *status.Status
}

// error struct
type Errdash struct {
	Pres  string   //错误信息前缀
	Codes uint32   //错误码
	Msgs  string   //错误信息（用户看）
	Logs  []string //日志信息（开发看）
}

//使用前请先阅读 errUtil/README.md；

/*
New 每一次都需要初始化新的，避免Code和Msg因为后续覆盖而对不上号

@Editor robotyang at 2023
*/
func New(err ...error) *Errdash {
	dsErr := &Errdash{
		Codes: CodeDefault,
	}
	for _, er := range err {
		_ = dsErr.Err(er)
	}
	return dsErr
}

/*
Err is a ...

@Editor robotyang at 2023
*/
func Err(err ...error) *Errdash {
	return New(err...)
}

/*
Code is a ...

@Editor robotyang at 2023
*/
func Code(code uint32) *Errdash {
	return New().Code(code)
}

/*
Pre is a ...

@Editor robotyang at 2023
*/
func Pre(msg string) *Errdash {
	return New().Pre(msg)
}

/*
Pref is a ...

@Editor robotyang at 2023
*/
func Pref(format string, a ...interface{}) *Errdash {
	return New().Pref(format, a...)
}

/*
Msg is a ...

@Editor robotyang at 2023
*/
func Msg(msg string) *Errdash {
	return New().Msg(msg)
}

/*
Msgf is a ...

@Editor robotyang at 2023
*/
func Msgf(format string, a ...interface{}) *Errdash {
	return New().Msgf(format, a...)
}

/*
Log is a ...

@Editor robotyang at 2023
*/
func Log(log string) *Errdash {
	return New().Log(log)
}

/*
Logf is a ...

@Editor robotyang at 2023
*/
func Logf(format string, a ...interface{}) *Errdash {
	return New().Logf(format, a...)
}

func (m *Errdash) Err(err error) *Errdash {
	switch err.(type) {
	case nil:
		m.Codes = CodeDefault
		m.Msgs = m.withPre("内部错误")
		m.Logs = append(m.Logs, m.withPre("错误类err为nil"))
	case *Errdash:
		m.Codes = err.(*Errdash).Codes
		m.Msgs = err.(*Errdash).Msgs
		m.Logs = append(m.Logs, err.(*Errdash).Logs...)
	case validator.ValidationErrors:
		m.Codes = CodeInvalidArgument
		m.Msgs = err.(validator.ValidationErrors).Error()
		m.Logs = append(m.Logs, err.(validator.ValidationErrors).Error())
	case RpcError: //如果是grpc类型的错误
		s, _ := status.FromError(err)
		code := ParseCode(s.Code()) //解析 code 对应的 errdash code;
		msg := s.Message()
		if msg == "" {
			msg = GetCodeMsg(code) //转义 code 为对应中文含义;
		}
		if len(s.Details()) > 0 { //记录 详情日志
			logdash.Infof("Errdash.Err() RpcError s.Details(): %#v", s.Details())
		}
		m.Codes = code
		m.Msgs = m.withPre(msg)
		m.Logs = append(m.Logs, m.withPre(msg))
	default:
		m.Codes = CodeDefault
		m.Msgs = m.withPre(err.Error())
		m.Logs = append(m.Logs, m.withPre(err.Error()))
	}

	return m
}

func (m *Errdash) Pre(prefix string) *Errdash {
	m.Pres = prefix
	return m
}

func (m *Errdash) Pref(format string, a ...interface{}) *Errdash {
	return m.Pre(fmt.Sprintf(format, a...))
}

//错误加前缀
func (m *Errdash) withPre(errStr string) string {
	if m.Pres != "" {
		return fmt.Sprintf("%v: %v", m.Pres, errStr)
	} else {
		return errStr
	}
}

//Code 一般是 gRpc 要求的 正整数
//但也可能是 不规范的 PHP负数 retCode
func (m *Errdash) Code(code uint32) *Errdash {
	msg := GetCodeMsg(code)
	m.Codes = code
	m.Msgs = m.withPre(msg)
	m.Logs = append(m.Logs, m.withPre(msg))
	return m
}

//设置 错误消息（用户看）
func (m *Errdash) Msg(msg string) *Errdash {
	if msg == "" {
		msg = GetCodeMsg(m.Codes)
	}
	m.Msgs = m.withPre(msg)
	m.Logs = append(m.Logs, m.withPre(msg))
	return m
}

//设置 日志消息（开发看）
func (m *Errdash) Msgf(format string, a ...interface{}) *Errdash {
	return m.Msg(fmt.Sprintf(format, a...))
}

//设置 日志消息（开发看）
func (m *Errdash) Log(log string) *Errdash {
	m.Logs = append(m.Logs, m.withPre(log))
	return m
}

//设置 日志消息（开发看）
func (m *Errdash) Logf(format string, a ...interface{}) *Errdash {
	return m.Log(fmt.Sprintf(format, a...))
}

//必须实现的接口
func (m *Errdash) Error() string {
	return ""
}
