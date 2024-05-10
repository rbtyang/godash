package dashrun

import (
	"bytes"
	"fmt"
	"runtime"
	"strings"
)

/*
PanicTrace  @Editor robotyang at 2023

# PanicTrace 获取异常的堆栈信息
*/
func PanicTrace(kb int) []byte {
	s := []byte("/src/runtime/panic.go")
	e := []byte("\ngoroutine ")
	line := []byte("\n")
	stack := make([]byte, kb<<10) // 4KB
	length := runtime.Stack(stack, true)
	start := bytes.Index(stack, s)
	stack = stack[start:length]
	start = bytes.Index(stack, line) + 1
	stack = stack[start:]
	end := bytes.LastIndex(stack, line)
	if end != -1 {
		stack = stack[:end]
	}
	end = bytes.Index(stack, e)
	if end != -1 {
		stack = stack[:end]
	}
	stack = bytes.TrimRight(stack, "\n")
	return stack
}

/*
LastCallerPlace @Editor robotyang at 2023

# LastCallerPlace 获取包外的最后一个调用位置

@Param curFileName 调用者所在文件路径，通过 _, curFileName, _, _ = runtime.Caller(0) 获得
*/
func LastCallerPlace(curFileName string) string {
	stack := ""
	for i := 2; i < 8; i++ { //最多向外延伸6层堆栈
		_, callFileName, line, _ := runtime.Caller(i)
		if callFileName != curFileName {
			stack = fmt.Sprintf("%v:%v", callFileName, line)
			break
		}
	}
	return stack
}

/*
LastCallerFuncName @Editor robotyang at 2023

# LastCallerFuncName 获取调用者的方法名

@Param short 是否只返回 不带包路径的 方法名
*/
func LastCallerFuncName(short bool) string {
	pc, _, _, ok := runtime.Caller(2)
	if !ok {
		return "unknown"
	}
	f := runtime.FuncForPC(pc)
	fName := f.Name() //git.xxx.cn/ooo/zzzzz/app/order.AddCustomer
	if short {
		fName = fName[strings.LastIndexAny(fName, "/\\")+1:] //order.AddCustomer
	}
	return fName
}
