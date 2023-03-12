package dashenv

import "os"

const (
	Dev     = "dev"     //开发环境
	Test    = "test"    //测试环境
	Release = "release" //预发布环境
	Prod    = "prod"    //生产环境
)

var curKey string //当前环境 curKey
var curVal string //当前环境 curVal

/*
init 初始化

@Editor robotyang at 2023
*/
func init() {
	curKey = "GO_ENV"
	curVal = os.Getenv(curKey)
}

/*
Init 使用自定义 环境变量key 进行初始化

@Editor robotyang at 2023
*/
func Init(key string) {
	curKey = key
	curVal = os.Getenv(curKey)
}

/*
isEnv 当前是否 env环境

@Editor robotyang at 2023
*/
func isEnv(env string) bool {
	return curVal == env
}

/*
IsDev 当前是否 开发环境

@Editor robotyang at 2023
*/
func IsDev() bool {
	return isEnv(Dev)
}

/*
IsTest 当前是否 测试环境

@Editor robotyang at 2023
*/
func IsTest() bool {
	return isEnv(Test)
}

/*
IsRelease 当前是否 预发布环境

@Editor robotyang at 2023
*/
func IsRelease() bool {
	return isEnv(Release)
}

/*
IsProd 当前是否 生产环境

@Editor robotyang at 2023
*/
func IsProd() bool {
	return isEnv(Prod)
}
