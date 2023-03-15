/*
Package godash is a rich function collection package.
It collects and encapsulates a series of common function sets. Such as array function set, encryption and decryption function set, etc.
In short, it is similar to the utils or helper toolkits that you have packaged in your project.

Package godash 是一个丰富的函数集合包，
它收集和封装了一系列常用函数集，比如 数组函数集、加解密函数集等等。
简单来说，就是类似于大家自己封装在项目中的 utils 或 helper 工具包。
*/
package godash

//执行所有包的单测
//go:generate go test ./...

//执行指定包的单测
//go:generate go test -v ./dashdemo/...

//执行指定主函数（需要是main包，且有main函数）
//go:generate go run ./gen.go
