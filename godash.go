/*
Package godash 丰富易用的、带详细说明和用例的 Golang 工具库（或者函数集、方法集、帮助类）。
它收集和封装了一系列常用函数集，比如 数组函数集、加解密函数集等。
简单来说，就是类似于大家自己封装在项目中的 utils 或 helper 工具包。

Package godash is a rich, easy-to-use Golang tool library (or set of functions, methods, help classes) with detailed instructions and use cases.
It collects and encapsulates a series of common function sets. Such as array function set, encryption and decryption function set, etc.
In short, it is similar to the utils or helper toolkits that you have packaged in your project.
*/
package godash

//执行所有包的单测
//go:generate go test ./...

//执行指定包的单测
//go:generate go test -v ./dashdemo/...

//执行指定主函数（需要是main包，且有main函数）
//go:generate go run ./gen.go
