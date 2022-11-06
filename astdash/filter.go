package astdash

import (
	"go/ast"
	"strings"
)

//返回true可以继续执行，返回false过滤掉这个数据
type (
	FieldFilter  func(*ast.Field) bool
	StructFilter func(*ast.TypeSpec) bool
	FuncFilter   func(*Func) bool
)

func WithFieldFilter(filter ...FieldFilter) *Option {
	return &Option{
		apply: func(a *Ast) {
			a.cfg.field = append(a.cfg.field, filter...)
		},
	}
}
func WithStructFilter(filter ...StructFilter) *Option {
	return &Option{
		apply: func(a *Ast) {
			a.cfg.st = append(a.cfg.st, filter...)
		},
	}
}

func WithFuncFilter(filter ...FuncFilter) *Option {
	return &Option{
		apply: func(a *Ast) {
			a.cfg.fun = append(a.cfg.fun, filter...)
		},
	}
}

type FilterFuncOpt struct {
	FuncName string
	Recv     *Field
}

//只返回符合的函数列表
func FilterFuncList(funcList []FilterFuncOpt) FuncFilter {
	return func(f *Func) bool {
		for _, v := range funcList {
			if f.Name != v.FuncName {
				continue
			}
			if v.Recv == nil {
				return true
			}
			if f.Recv == nil {
				continue
			}
			if v.Recv.Name != "" && f.Recv.Name != v.Recv.Name {
				continue
			}
			if !v.Recv.Type.Cmp(f.Recv.Type) {
				continue
			}
			return true
		}
		return false
	}
}

//新版proto内部字段需要过滤
var protoField = []string{
	"state",
	"sizeCache",
	"unknownFields",
}

func FilterProtoInner(field *ast.Field) bool {
	name := field.Names[0].Name
	if strings.HasPrefix(name, "XXX_") {
		//老版proto生成的内部字段是XXX_开头的 需要过滤了
		return false
	}
	for _, v := range protoField {
		if v == name {
			return false
		}
	}
	return true
}

//过滤小写字段
func FilterInnerField(field *ast.Field) bool {
	name := field.Names[0].Name
	if len(name) == 0 {
		return false
	}
	if !(name[0] >= 'A' && name[0] <= 'Z') {
		return false
	}
	return true
}

func FilterInnerSt(st *ast.TypeSpec) bool {
	name := st.Name.String()
	if len(name) == 0 {
		return false
	}
	if !(name[0] >= 'A' && name[0] <= 'Z') {
		return false
	}
	return true
}

func FilterProtoSt(st *ast.TypeSpec) bool {
	name := st.Name.String()
	if strings.HasPrefix(name, "Unimplemented") && strings.HasSuffix(name, "Server") {
		return false
	}
	return true
}
