package dasherr

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/rbtyang/godash/dashlog"
	"google.golang.org/grpc/status"
)

//错误类 ----------------------------------------------------

type RpcError interface {
	GRPCStatus() *status.Status
}

// error struct
type Errdash struct {
	Pres  string   //错误信息前缀
	Codes uint32   //错误码
	Msgs  string   //错误信息（用户看）
	Logs  []string //日志信息（开发看）
	Stack []string //调用堆栈（开发看）
}

/*
New 实例化错误类

@Editor robotyang at 2023
*/
func New(err ...error) *Errdash {
	dsErr := &Errdash{
		Codes: CodeUnknown,
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

/*
Err is a ...

@Editor robotyang at 2023
*/
func (m *Errdash) Err(err error) *Errdash {
	switch err.(type) {
	case nil:
		m.Codes = CodeUnknown
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
			dashlog.Infof("Errdash.Err() RpcError s.Details(): %#v", s.Details())
		}
		m.Codes = code
		m.Msgs = m.withPre(msg)
		m.Logs = append(m.Logs, m.withPre(msg))
	default:
		m.Codes = CodeUnknown
		m.Msgs = m.withPre(err.Error())
		m.Logs = append(m.Logs, m.withPre(err.Error()))
	}

	return m
}

/*
Pre is a ...

@Editor robotyang at 2023
*/
func (m *Errdash) Pre(prefix string) *Errdash {
	m.Pres = prefix
	return m
}

/*
Pref is a ...

@Editor robotyang at 2023
*/
func (m *Errdash) Pref(format string, a ...interface{}) *Errdash {
	return m.Pre(fmt.Sprintf(format, a...))
}

/*
withPre 错误加前缀

@Editor robotyang at 2023
*/
func (m *Errdash) withPre(errStr string) string {
	if m.Pres != "" {
		return fmt.Sprintf("%v: %v", m.Pres, errStr)
	} else {
		return errStr
	}
}

/*
Code 设置错误码

@Editor robotyang at 2023
*/
func (m *Errdash) Code(code uint32) *Errdash {
	msg := GetCodeMsg(code)
	m.Codes = code
	m.Msgs = m.withPre(msg)
	m.Logs = append(m.Logs, m.withPre(msg))
	return m
}

/*
Msg 设置 错误消息（用户看）

@Editor robotyang at 2023
*/
func (m *Errdash) Msg(msg string) *Errdash {
	if msg == "" {
		msg = GetCodeMsg(m.Codes)
	}
	m.Msgs = m.withPre(msg)
	m.Logs = append(m.Logs, m.withPre(msg))
	return m
}

/*
Msgf 设置 错误消息（用户看）

@Editor robotyang at 2023
*/
func (m *Errdash) Msgf(format string, a ...interface{}) *Errdash {
	return m.Msg(fmt.Sprintf(format, a...))
}

/*
Log 设置 日志消息（开发看）

@Editor robotyang at 2023
*/
func (m *Errdash) Log(log string) *Errdash {
	m.Logs = append(m.Logs, m.withPre(log))
	return m
}

/*
Logf 设置 日志消息（开发看）

@Editor robotyang at 2023
*/
func (m *Errdash) Logf(format string, a ...interface{}) *Errdash {
	return m.Log(fmt.Sprintf(format, a...))
}

/*
Error 必须实现的接口

@Editor robotyang at 2023
*/
func (m *Errdash) Error() string {
	return m.Msgs
}
