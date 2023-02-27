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

/*
init is a ...

@Editor robotyang at 2023
*/
func init() {
	goEnv = os.Getenv("GO_ENV")
}

/*
IsEnv is a ...

@Editor robotyang at 2023
*/
func IsEnv(env string) bool {
	return goEnv == env
}

/*
IsDev is a ...

@Editor robotyang at 2023
*/
func IsDev() bool {
	return IsEnv(Dev)
}

/*
IsTest is a ...

@Editor robotyang at 2023
*/
func IsTest() bool {
	return IsEnv(Test)
}

/*
IsBeta is a ...

@Editor robotyang at 2023
*/
func IsBeta() bool {
	return IsEnv(Beta)
}

/*
IsProd is a ...

@Editor robotyang at 2023
*/
func IsProd() bool {
	return IsEnv(Prod)
}
