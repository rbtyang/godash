package astdash

import (
	"github.com/rbtyang/godash/logdash"
	"go/ast"
	"go/parser"
	"go/token"
)

/*
NewAst is a ...

@Editor robotyang at 2023
*/
func NewAst(options ...*Option) *Ast {
	a := new(Ast)
	for _, opt := range options { // 应用 options设置 的选项
		opt.apply(a)
	}
	return a
}

func (a *Ast) ParseFile(inputPath string) (err error) {
	fset := token.NewFileSet()

	f, err := parser.ParseFile(fset, inputPath, nil, parser.ParseComments)
	if err != nil {
		logdash.Errorf("ParseFile ParseFile err:%v", err)
		return
	}

	a.Package = f.Name.Name
	for _, i := range f.Imports {
		a.Imports = append(a.Imports, a.ParseImport(i))
	}

	err = a.ParseScopes(f.Scope)
	if err != nil {
		logdash.Errorf("ParseFile ParseScope err:%v", err)
		return
	}

	err = a.ParseDecls(f.Decls)
	if err != nil {
		logdash.Errorf("ParseFile ParseDecls err:%v", err)
		return err
	}

	return
}

func (a *Ast) ParseScopes(Scope *ast.Scope) (err error) {
	structs := map[int]interface{}{}    //用来排序
	interfaces := map[int]interface{}{} //用来排序
	for _, obj := range Scope.Objects {
		if obj.Kind != ast.Typ {
			continue
		}
		if _, ok := obj.Decl.(*ast.TypeSpec); !ok {
			continue
		}

		typeSpec := obj.Decl.(*ast.TypeSpec).Type
		switch typeSpec.(type) {
		case *ast.StructType:
			if !(obj.Name[0] >= 'A' && obj.Name[0] <= 'Z') {
				//如果不是 可以导出的结构体，不进行处理，proto.pb.go中 存在该类的 内部结构体
				continue
			}
			typeSpec := obj.Decl.(*ast.TypeSpec)
			continueFlag := false
			for _, filter := range a.cfg.st {
				if !filter(typeSpec) {
					continueFlag = true
					break
				}
			}
			if continueFlag {
				continue
			}
			structs[int(typeSpec.Pos())] = a.ParseStruct(typeSpec)
			//a.Structs = append(a.Structs, a.ParseStruct(typeSpec))
		case *ast.InterfaceType:
			interfaces[int(typeSpec.Pos())] = a.ParseInterface(obj.Decl.(*ast.TypeSpec))
			//a.Interfaces = append(a.Interfaces, a.ParseInterface(obj.Decl.(*ast.TypeSpec)))
		}
	}

	if len(interfaces) > 0 {
		slice := getSortSlice(interfaces)
		for _, v := range slice {
			a.Interfaces = append(a.Interfaces, v.(Interface))
		}
	}
	if len(structs) > 0 {
		slice := getSortSlice(structs)
		for _, v := range slice {
			a.Structs = append(a.Structs, v.(Struct))
		}
	}

	return
}

func (a *Ast) ParseDecls(Decls []ast.Decl) (err error) {
	for _, decl := range Decls {
		switch decl.(type) {
		case *ast.FuncDecl:
			fun := decl.(*ast.FuncDecl)

			f := a.ParseFunc(fun.Type)
			if fun.Recv != nil {
				recv := a.ParseField(fun.Recv.List[0])
				f.Recv = &recv
			}
			f.Name = getIdent(fun.Name)
			f.Docs = getCommentGroup(fun.Doc)
			continueFlag := false
			for _, filter := range a.cfg.fun {
				if !filter(&f) {
					continueFlag = true
					break
				}
			}
			if continueFlag {
				continue
			}
			a.Funcs = append(a.Funcs, f)
		}
	}
	return nil
}

func (a *Ast) ParseImport(i *ast.ImportSpec) (ret Import) {
	ret.Path = i.Path.Value
	ret.Comments = getCommentGroup(i.Comment)
	ret.Docs = getCommentGroup(i.Doc)
	ret.Name = getIdentName(i.Name)
	ret.Name = getImportName(ret.Name, ret.Path)
	return
}

func (a *Ast) ParseStruct(typ *ast.TypeSpec) (ret Struct) {
	st := typ.Type.(*ast.StructType)
	ret.Comments = getCommentGroup(typ.Comment)
	ret.Docs = getCommentGroup(typ.Doc)
	ret.Name = typ.Name.String()

	for _, field := range st.Fields.List {
		if len(field.Names) == 0 {
			continue
		}
		continueFlag := false
		for _, filter := range a.cfg.field {
			if !filter(field) {
				continueFlag = true
				break
			}
		}
		if continueFlag {
			continue
		}
		ret.Fields = append(ret.Fields, a.ParseField(field))
	}
	return
}

func (a *Ast) ParseField(field *ast.Field) (ret Field) {
	name := getIdentsName(field.Names)
	ret.Name = name
	ret.Docs = getCommentGroup(field.Doc)
	ret.Comments = getCommentGroup(field.Comment)
	ret.Tag = getTag(field.Tag)
	ret.Type = handleType(field.Type)
	return
}

func (a *Ast) ParseInterface(typ *ast.TypeSpec) (ret Interface) {
	ret.Comments = getCommentGroup(typ.Comment)
	ret.Docs = getCommentGroup(typ.Doc)
	ret.Name = typ.Name.String()

	it := typ.Type.(*ast.InterfaceType)
	if it.Incomplete == true { //如果为true则是空接口类型
		return
	}

	for _, method := range it.Methods.List {
		f, ok := method.Type.(*ast.FuncType)
		if ok != true {
			logdash.Panicf("interface just support function")
		}
		funcT := a.ParseFunc(f)
		funcT.Name = getIdent(method.Names[0])
		funcT.Docs = getCommentGroup(method.Doc)
		funcT.Comments = getCommentGroup(method.Comment)
		ret.Funcs = append(ret.Funcs, funcT)

	}
	return
}

func (a *Ast) ParseFunc(f *ast.FuncType) (ret Func) {
	funcT := Func{}
	if f.Params != nil {
		for _, param := range f.Params.List {
			funcT.Params = append(funcT.Params, a.ParseField(param))
		}
	}
	if f.Results != nil {
		for _, result := range f.Results.List {
			funcT.Results = append(funcT.Results, a.ParseField(result))
		}
	}
	return funcT
}
