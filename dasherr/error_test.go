package dasherr_test

import (
	"errors"
	"fmt"
	"github.com/rbtyang/godash/dasherr"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

/*
TestParseCode is a ...
*/
func TestCommon(t *testing.T) {
	{
		err := errors.New("数据库出错")
		ers := dasherr.New(err)
		recv := ers.Error()
		assert.Equal(t, "数据库出错", recv)
	}
	{
		err := errors.New("数据库出错")
		ers := dasherr.Err(err)
		recv := ers.Error()
		assert.Equal(t, "数据库出错", recv)
	}
	{
		{
			ers := dasherr.Code(123)
			recv := ers.Error()
			assert.Equal(t, "未知错误", recv)
		}
		{
			ers := dasherr.Code(10000)
			recv := ers.Error()
			assert.Equal(t, "未知错误", recv)
		}
	}
	{
		ers := dasherr.New()
		ers.Msg("数据库出错").Msg("基础数据出错").Msg("业务接口错误").Log("代码问题")
		{
			recv := ers.Error()
			assert.Equal(t, "业务接口错误", recv)
		}
		{
			recv := fmt.Sprintf("%s", ers)
			assert.Equal(t, "业务接口错误", recv)
		}
		{
			recv := strings.Join(ers.Logs, "->")
			assert.Equal(t, "数据库出错->基础数据出错->业务接口错误->代码问题", recv)
		}
	}
	{
		ers := dasherr.New()
		ers.Msg("数据库出错")
		ers.Msg("基础数据出错")
		ers = ers.Msg("业务接口错误").Log("代码问题")
		{
			recv := ers.Error()
			assert.Equal(t, "业务接口错误", recv)
		}
		{
			recv := fmt.Sprintf("%s", ers)
			assert.Equal(t, "业务接口错误", recv)
		}
		{
			recv := strings.Join(ers.Logs, "->")
			assert.Equal(t, "数据库出错->基础数据出错->业务接口错误->代码问题", recv)
		}
	}
}

/*
TestParseCode is a ...
*/
func TestWithPre(t *testing.T) {
	{
		ers := dasherr.Pre("外部接口错误")
		ers.Msg("未配置秘钥")
		recv := ers.Error()
		assert.Equal(t, "外部接口错误: 未配置秘钥", recv)
	}
	{
		ers := dasherr.Pre("外部接口错误")
		ers.Msg("未配置秘钥").Msg("秘钥为空")
		recv := ers.Error()
		assert.Equal(t, "外部接口错误: 秘钥为空", recv)
		recv2 := strings.Join(ers.Logs, "->")
		assert.Equal(t, "外部接口错误: 未配置秘钥->外部接口错误: 秘钥为空", recv2)
	}
	{
		ers := dasherr.Pre("外部接口错误")
		ers.Err(errors.New("外部接口错误")).Msg("未配置秘钥")
		recv := ers.Error()
		assert.Equal(t, "外部接口错误: 未配置秘钥", recv)
	}
	{
		ers := dasherr.Pre("外部接口错误")
		ers.Code(dasherr.CodeCanceled)
		recv := ers.Error()
		assert.Equal(t, "外部接口错误: 操作被取消", recv)
	}
}
