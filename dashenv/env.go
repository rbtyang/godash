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
init @Editor robotyang at 2023

# init 初始化
*/
func init() {
	curKey = "GO_ENV"
	curVal = os.Getenv(curKey)
}

/*
Init @Editor robotyang at 2023

# Init 使用自定义 环境变量key 进行初始化
*/
func Init(key string) {
	curKey = key
	curVal = os.Getenv(curKey)
}

/*
isEnv @Editor robotyang at 2023

# isEnv 当前是否 env环境
*/
func isEnv(env string) bool {
	return curVal == env
}

/*
IsDev @Editor robotyang at 2023

# IsDev 当前是否 开发环境
*/
func IsDev() bool {
	return isEnv(Dev)
}

/*
IsTest @Editor robotyang at 2023

# IsTest 当前是否 测试环境
*/
func IsTest() bool {
	return isEnv(Test)
}

/*
IsRelease @Editor robotyang at 2023

# IsRelease 当前是否 预发布环境
*/
func IsRelease() bool {
	return isEnv(Release)
}

/*
IsProd @Editor robotyang at 2023

# IsProd 当前是否 生产环境
*/
func IsProd() bool {
	return isEnv(Prod)
}
