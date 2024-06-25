package dashast

import (
	"go/ast"
	"sort"
	"strings"
)

/*
getCommentGroup @Editor robotyang at 2023

# getCommentGroup 提取注释
*/
func getCommentGroup(cg *ast.CommentGroup) (ret []string) {
	if cg == nil {
		return
	}
	for _, v := range cg.List {
		ret = append(ret, v.Text)
	}
	return
}

/*
getIdent @Editor robotyang at 2023
*/
func getIdent(id *ast.Ident) (ret string) {
	if id == nil {
		return ""
	}
	return id.Name
}

/*
@Editor robotyang at 2023

getTag 提取tag
*/
func getTag(cg *ast.BasicLit) (ret string) {
	if cg == nil {
		return
	}
	return cg.Value
}

/*
getImportName @Editor robotyang at 2023
*/
func getImportName(name, path string) string {
	if name != "" {
		return name
	}
	s := strings.Split(path, "/")
	if len(s) == 1 { //如果没找到
		name = path
		//将双引号去除
		return name[1 : len(name)-1]
	}
	name = s[len(s)-1]
	//将双引号去除
	return name[0 : len(name)-1]
}

/*
getIdentName @Editor robotyang at 2023
*/
func getIdentName(i *ast.Ident) string {
	if i == nil {
		return ""
	}
	return i.String()
}

/*
getIdentsName @Editor robotyang at 2023
*/
func getIdentsName(i []*ast.Ident) string {
	if len(i) == 0 {
		return ""
	}
	return i[0].String()
}

/*
@Editor robotyang at 2023

CheckInnerStruct 检查是不是 内部使用的结构体
*/
func CheckInnerStruct(name string) bool {
	_, t := innerType[name]
	return !t
}

/*
getSortSlice @Editor robotyang at 2023
*/
func getSortSlice(inMap map[int]any) []any {
	if len(inMap) == 0 {
		return nil
	}

	idxSlice := make([]int, 0, len(inMap))
	for idx := range inMap {
		idxSlice = append(idxSlice, idx)
	}
	sort.Ints(idxSlice)

	valSlice := make([]any, 0, len(inMap))
	for _, v := range idxSlice {
		valSlice = append(valSlice, inMap[v])
	}
	return valSlice
}
