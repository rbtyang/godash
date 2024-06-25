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

// Dasherr error struct
type Dasherr struct {
	Pres  string   //错误信息前缀
	Codes uint32   //错误码
	Msgs  string   //错误信息（用户看）
	Logs  []string //日志信息（开发看）
	Stack []string //调用堆栈（开发看）
}

/*
New @Editor robotyang at 2023

# New 实例化错误类
*/
func New(err ...error) *Dasherr {
	dsErr := &Dasherr{
		Codes: CodeUnknown,
	}
	for _, er := range err {
		_ = dsErr.Err(er)
	}
	return dsErr
}

/*
Err @Editor robotyang at 2023

# Err is a ...
*/
func Err(err ...error) *Dasherr {
	return New(err...)
}

/*
Code @Editor robotyang at 2023

# Code is a ...
*/
func Code(code uint32) *Dasherr {
	return New().Code(code)
}

/*
Pre @Editor robotyang at 2023

# Pre is a ...
*/
func Pre(msg string) *Dasherr {
	return New().Pre(msg)
}

/*
Pref @Editor robotyang at 2023

# Pref is a ...
*/
func Pref(format string, a ...any) *Dasherr {
	return New().Pref(format, a...)
}

/*
Msg @Editor robotyang at 2023

# Msg is a ...
*/
func Msg(msg string) *Dasherr {
	return New().Msg(msg)
}

/*
Msgf @Editor robotyang at 2023

# Msgf is a ...
*/
func Msgf(format string, a ...any) *Dasherr {
	return New().Msgf(format, a...)
}

/*
Log @Editor robotyang at 2023

# Log is a ...
*/
func Log(log string) *Dasherr {
	return New().Log(log)
}

/*
Logf @Editor robotyang at 2023

# Logf is a ...
*/
func Logf(format string, a ...any) *Dasherr {
	return New().Logf(format, a...)
}

/*
Err @Editor robotyang at 2023

# Err is a ...
*/
func (m *Dasherr) Err(err error) *Dasherr {
	switch err.(type) {
	case nil:
		m.Codes = CodeUnknown
		m.Msgs = m.withPre("内部错误")
		m.Logs = append(m.Logs, m.withPre("错误类err为nil"))
	case *Dasherr:
		m.Codes = err.(*Dasherr).Codes
		m.Msgs = err.(*Dasherr).Msgs
		m.Logs = append(m.Logs, err.(*Dasherr).Logs...)
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
			dashlog.Infof("Dasherr.Err() RpcError s.Details(): %#v", s.Details())
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
Pre @Editor robotyang at 2023

# Pre is a ...
*/
func (m *Dasherr) Pre(prefix string) *Dasherr {
	m.Pres = prefix
	return m
}

/*
Pref @Editor robotyang at 2023

# Pref is a ...
*/
func (m *Dasherr) Pref(format string, a ...any) *Dasherr {
	return m.Pre(fmt.Sprintf(format, a...))
}

/*
withPre @Editor robotyang at 2023

# withPre 错误加前缀
*/
func (m *Dasherr) withPre(errStr string) string {
	if m.Pres != "" {
		return fmt.Sprintf("%v: %v", m.Pres, errStr)
	} else {
		return errStr
	}
}

/*
Code @Editor robotyang at 2023

# Code 设置错误码
*/
func (m *Dasherr) Code(code uint32) *Dasherr {
	msg := GetCodeMsg(code)
	m.Codes = code
	m.Msgs = m.withPre(msg)
	m.Logs = append(m.Logs, m.withPre(msg))
	return m
}

/*
Msg @Editor robotyang at 2023

# Msg 设置 错误消息（用户看）
*/
func (m *Dasherr) Msg(msg string) *Dasherr {
	if msg == "" {
		msg = GetCodeMsg(m.Codes)
	}
	m.Msgs = m.withPre(msg)
	m.Logs = append(m.Logs, m.withPre(msg))
	return m
}

/*
Msgf @Editor robotyang at 2023

# Msgf 设置 错误消息（用户看）
*/
func (m *Dasherr) Msgf(format string, a ...any) *Dasherr {
	return m.Msg(fmt.Sprintf(format, a...))
}

/*
Log @Editor robotyang at 2023

# Log 设置 日志消息（开发看）
*/
func (m *Dasherr) Log(log string) *Dasherr {
	m.Logs = append(m.Logs, m.withPre(log))
	return m
}

/*
Logf @Editor robotyang at 2023

# Logf 设置 日志消息（开发看）
*/
func (m *Dasherr) Logf(format string, a ...any) *Dasherr {
	return m.Log(fmt.Sprintf(format, a...))
}

/*
Error @Editor robotyang at 2023

# Error 必须实现的接口
*/
func (m *Dasherr) Error() string {
	return m.Msgs
}
