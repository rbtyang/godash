package astdash

import (
	"fmt"
	"go/ast"
	"log"
)

const (
	Unsupport     = iota //未支持的类型
	Ident                //string int 非指针的结构体
	StructType           //结构体中的匿名结构体或 外面定义的结构体类型
	StarExpr             //指针类型的参数
	SelectorExpr         //别的包的类型或表达式类型参数
	Ellipsis             //...可变参数类型
	ArrayType            //切片或数组
	ChanType             //通道类型
	MapType              //map类型
	InterfaceType        // interface类型
	FuncType             //函数类型
)

var m = map[int]string{
	Unsupport:    "未支持",
	Ident:        "非指针",
	StructType:   "结构体",
	StarExpr:     "指针",
	SelectorExpr: "表达式",
	ArrayType:    "切片",
}

var innerType = map[string]struct{}{
	"bool":       {},
	"int":        {},
	"int8":       {},
	"int16":      {},
	"int32":      {},
	"int64":      {},
	"uint":       {},
	"uint8":      {},
	"uint16":     {},
	"uint32":     {},
	"uint64":     {},
	"uintptr":    {},
	"float32":    {},
	"float64":    {},
	"complex64":  {},
	"complex128": {},
	"chan":       {},
	"func":       {},
	"interface":  {},
	"map":        {},
	"slice":      {},
	"string":     {},
	"struct":     {},
	"byte":       {},
	"rune":       {},
}

type (
	Ast struct {
		Package    string
		Imports    []Import
		Structs    []Struct
		Interfaces []Interface
		Funcs      []Func
		cfg        config
	}
	config struct {
		field []FieldFilter
		st    []StructFilter
		fun   []FuncFilter
	}
	Interface struct {
		Name     string   //
		Comments []string //在后面的备注
		Docs     []string //在上方的备注
		Funcs    []Func   //函数列表
	}
	Func struct {
		Name     string
		Recv     *Field   //接收器 func前的参数
		Docs     []string //在上方的备注
		Comments []string //在后面的备注
		Params   []Field  //入参
		Results  []Field  //返回参数
	}
	Import struct {
		Name     string   //路径的别名
		Path     string   //路径
		Comments []string //在后面的备注
		Docs     []string //在上方的备注
	}
	//结构体
	Struct struct {
		Name     string   //结构体的名字
		Comments []string //在后面的备注
		Docs     []string //在上方的备注
		Fields   []Field  //字段
	}
	//结构体字段
	Field struct {
		Name     string   //参数名
		Type     Type     //参数类型
		Comments []string //在后面的备注
		Docs     []string //在上方的备注
		Tag      string   //原来的tag
	}
	Type struct {
		Name  []string //如果是 a.b 则name 为 a,b
		Kind  int      //type类型
		Types []Type   //如果是复杂类型 则在Types中进行嵌套 如map
		Inner bool     //最低位如果为true则是内部结构体,需要修改名字,只针对结构体
	}
	Option struct {
		apply func(*Ast)
	}
)

func (s Struct) GetFieldWithName(name string) *Field {
	for _, f := range s.Fields {
		if f.Name == name {
			return &f
		}
	}
	return nil
}

func (t Type) String() string {
	return fmt.Sprintf("%v type:%s tpyes:%v inner:%v", t.Name, m[t.Kind], t.Types, t.Inner)
}

//内部字段添加后缀
func (t Type) InnerAddSuffix(suffix string) Type {
	if t.Inner == true {
		t.Name[len(t.Name)-1] += suffix
		return t
	}
	for _, t2 := range t.Types {
		t2.InnerAddSuffix(suffix)
	}
	return t
}

func (i Interface) Copy() (new Interface) {
	new.Name = i.Name
	new.Comments = i.Comments
	new.Docs = i.Docs
	for _, f := range i.Funcs {
		var newF Func
		var params []Field
		var results []Field
		for _, field := range f.Params {
			params = append(params, field.Copy())
		}
		for _, field := range f.Results {
			results = append(results, field.Copy())
		}
		newF.Params = params
		newF.Results = results
		newF.Comments = f.Comments
		newF.Docs = f.Docs
		newF.Name = f.Name
		new.Funcs = append(new.Funcs, newF)
	}
	return
}

func (t Type) Copy() (new Type) {
	for _, v := range t.Name {
		new.Name = append(new.Name, v)
	}
	new.Kind = t.Kind
	new.Inner = t.Inner
	for i := range t.Types {
		new.Types = append(new.Types, t.Types[i].Copy())
	}
	return
}

func (t Field) Copy() (new Field) {
	for _, v := range t.Docs {
		new.Docs = append(new.Docs, v)
	}
	for _, v := range t.Comments {
		new.Comments = append(new.Comments, v)
	}
	new.Name = t.Name
	new.Type = t.Type.Copy()
	new.Tag = t.Tag
	return
}

func (t Type) Fmt() string {
	switch t.Kind {
	case Ident: //string int 非指针的结构体
		return t.Name[0]
	case StarExpr: //指针类型的参数
		return t.Name[0] + t.Types[0].Fmt()
	case SelectorExpr: //别的包的类型或表达式类型参数
		return t.Name[0] + "." + t.Name[1]
	case ArrayType: //切片或数组
		return t.Name[0] + t.Types[0].Fmt()
	case MapType: //map
		return t.Name[0] + "[" + t.Types[0].Fmt() + "]" + t.Types[1].Fmt()
	default:
		log.Fatalf("type fmt not supported: %v", t.Kind)
	}
	return ""
}

//比较是否类型相同
func (t Type) Cmp(t2 Type) bool {
	if t.Kind != t2.Kind || len(t.Name) != len(t2.Name) || len(t.Types) != len(t2.Types) {
		return false
	}
	for i := range t.Name {
		if t.Name[i] != t2.Name[i] {
			return false
		}
	}
	for i := range t.Types {
		if !t.Types[i].Cmp(t2.Types[i]) {
			return false
		}
	}
	return true
}

func (t Type) IsInnerStruct() bool {
	if t.Inner == true {
		return true
	}
	for _, t2 := range t.Types {
		if t2.IsInnerStruct() {
			return true
		}
	}
	return false
}

func handleType(expr ast.Expr) (ret Type) {
	switch expr.(type) {
	case *ast.Ident: //string int 非指针的命名结构体
		tp := expr.(*ast.Ident)
		ret = Type{
			Name:  []string{tp.Name},
			Kind:  Ident,
			Inner: CheckInnerStruct(tp.Name),
		}
	case *ast.StarExpr: //指针类型的参数
		tp := expr.(*ast.StarExpr)
		ret.Kind = StarExpr
		ret.Name = []string{"*"}
		ret.Types = []Type{
			handleType(tp.X),
		}
	case *ast.ArrayType: //切片类型的参数 暂时不支持数组
		tp := expr.(*ast.ArrayType)
		ret.Kind = ArrayType
		ret.Name = []string{"[]"}
		ret.Types = []Type{
			handleType(tp.Elt),
		}
	case *ast.SelectorExpr:
		tp := expr.(*ast.SelectorExpr)
		ret.Kind = SelectorExpr
		ret.Name = []string{tp.X.(*ast.Ident).Name, tp.Sel.Name}
	case *ast.Ellipsis:
		tp := expr.(*ast.Ellipsis)
		ret.Kind = Ellipsis
		ret.Name = []string{"..."}
		ret.Types = []Type{
			handleType(tp.Elt),
		}
	case *ast.MapType:
		tp := expr.(*ast.MapType)
		ret.Kind = MapType
		ret.Name = []string{"map"}
		ret.Types = []Type{
			handleType(tp.Key),
			handleType(tp.Value),
		}
	case *ast.InterfaceType:
		ret.Kind = InterfaceType
	case *ast.FuncType:
		ret.Kind = FuncType
	default:
		log.Fatalf("field.Type not support type:%#v", expr)
		return
	}
	return
}
