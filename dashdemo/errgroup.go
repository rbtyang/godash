package dashdemo

import (
	"fmt"
)

/*
ErrGroup @Editor robotyang at 2023

# ErrGroup 请查看同级目录下的单测文件 errgroup_test.go
*/
func ErrGroup(user string) string {
	return fmt.Sprintf("%v Ni Hao", user)
}
