# godash - 丰富易用的、带详细说明和用例的 Golang 工具库、函数集、方法集
> A rich, easy-to-use Golang library of tools, functions, and methods with detailed instructions and use cases

[![Latest](https://img.shields.io/badge/latest-v0.0.1-blue.svg)](https://github.com/Andrew-M-C/go.jsonvalue/tree/v0.0.1)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://opensource.org/license/mit/)

[简体中文](README.md) | [ENGLISH](README_EN.md)

## 使用文档

- 请移步到[官方平台文档](https://pkg.go.dev/github.com/rbtyang/godash)，可以查看更多详细的说明，比如 包、入参、出参 的作用等等。

## 贡献伙伴

- 感谢每一位兢兢业业无私奉献的伙伴，感谢您为开源社区作出的贡献！

<a href="github.com/rbtyang/godash/graphs/contributors">
  <img src="https://contributors-img.web.app/image?repo=rbtyang/godash" />
</a>

## 联系我们

- 微信号 `RobotYang7` (加我进微信群，需备注“来自github”)。

<img width="350" src="./doc/微信二维码.jpg">

## 好评赞赏

- 如果该项目对你有帮助，希望获得您肯定的赞赏， 我们将有动力提供更多支持！

<img width="350" src="./doc/微信赞赏码.jpg">


--- 

## 推荐的包

#### 后端

```shell script
"github.com/stretchr/testify/assert" //单测类型断言，参考 各个_test.go
"github.com/spf13/cast" //各种类型转换
"github.com/fatih/structs" //结构体转Map
"github.com/goinggo/mapstructure" //Map转结构体，参考 godash\dashconv\mapstructure.go
"github.com/imdario/mergo" //结构体合并，参考 godash\dashconv\mergo_test.go
"github.com/spf13/viper" //配置文件读取和管理
"golang.org/x/sync/errgroup" //并发编程，参考 godash\dashdemo\errgroup_test.go
"github.com/bwmarrin/snowflake" //雪花ID生成，参考 godash\dashrand\snow_test.go
"github.com/google/uuid" //UUID生成，参考 godash\dashrand\uuid_test.go
```

#### 前端

```shell script
crypto-js //加解密、哈希
assert //断言
mocha //单测
```