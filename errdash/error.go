package errdash

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
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
//每一次都需要初始化新的，避免Code和Msg因为后续覆盖而对不上号
func New(err ...error) *Errdash {
	mErr := &Errdash{
		Codes: ErrDefault,
	}
	for _, er := range err {
		_ = mErr.Err(er)
	}
	return mErr
}

func Err(err ...error) *Errdash {
	return New(err...)
}

func Code(code uint32) *Errdash {
	return New().Code(code)
}

func Pre(msg string) *Errdash {
	return New().Pre(msg)
}

func Pref(format string, a ...interface{}) *Errdash {
	return New().Pref(format, a...)
}

func Msg(msg string) *Errdash {
	return New().Msg(msg)
}

func Msgf(format string, a ...interface{}) *Errdash {
	return New().Msgf(format, a...)
}

func Log(log string) *Errdash {
	return New().Log(log)
}

func Logf(format string, a ...interface{}) *Errdash {
	return New().Logf(format, a...)
}

func (m *Errdash) Err(err error) *Errdash {
	switch err.(type) {
	case nil:
		m.Codes = ErrDefault
		m.Msgs = m.withPre("内部错误")
		m.Logs = append(m.Logs, m.withPre("错误类err为nil"))
	case *Errdash:
		m.Codes = err.(*Errdash).Codes
		m.Msgs = err.(*Errdash).Msgs
		m.Logs = append(m.Logs, err.(*Errdash).Logs...)
	case validator.ValidationErrors:
		m.Codes = ErrParam
		m.Msgs = err.(validator.ValidationErrors).Error()
		m.Logs = append(m.Logs, err.(validator.ValidationErrors).Error())
	case RpcError: //如果是grpc类型的错误
		s, _ := status.FromError(err)
		if s.Code() != codes.Unknown {
			m.Codes = ErrDefault
			m.Msgs = m.withPre(err.Error())
			m.Logs = append(m.Logs, m.withPre(err.Error()))
			return m
		} else { //只有自定义的错误，grpc会返回unknown错误码
			m.Codes = uint32(s.Code())
			m.Msgs = m.withPre(s.Message())
			m.Logs = append(m.Logs, m.withPre(err.Error()))
			if len(s.Details()) > 0 {
				log.Printf("[errdash.Details] %#v", s.Details())
			}
		}
	default:
		m.Codes = ErrDefault
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
	msg := TransCodeMsg(code)
	m.Codes = code
	m.Msgs = m.withPre(msg)
	m.Logs = append(m.Logs, m.withPre(msg))
	return m
}

//设置 错误消息（用户看）
func (m *Errdash) Msg(msg string) *Errdash {
	if msg == "" {
		msg = TransCodeMsg(m.Codes)
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
