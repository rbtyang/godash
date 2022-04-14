package envdash

import "os"

const (
	Dev  = "dev"
	Test = "test"
	Beta = "beta"
	Prod = "prod"
)

//是否开发环境
var goEnv string

func init() {
	goEnv = os.Getenv("GO_ENV")
}

func IsEnv(env string) bool {
	return goEnv == env
}

func IsDev() bool {
	return IsEnv(Dev)
}

func IsTest() bool {
	return IsEnv(Test)
}

func IsBeta() bool {
	return IsEnv(Beta)
}

func IsProd() bool {
	return IsEnv(Prod)
}
