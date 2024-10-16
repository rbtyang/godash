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

## 开发规范

- 必须使用 `commitizen` 插件提交规范的 commit msg
  - [官方插件仓库](https://github.com/commitizen/cz-cli)
  - [大佬介绍文档](https://www.jianshu.com/p/d264f88d13a4)
  - [Jetbrains插件](https://plugins.jetbrains.com/plugin/9861-git-commit-template)
- 必须编写 详细的规范的 包注释、方法注释，格式参考已有的注释，并能通过 `pkgsite` 正常渲染
  - [官方工具仓库](https://github.com/golang/pkgsite)
  - [大佬介绍文档](https://cloud.tencent.com/developer/article/1959696?from=10910)
- 必须设计 每个方法/类 的单测用例（至少4个case），建议编写 基准测试用例（benchmark），类包 必须设计 `example` 示例用例
- 函数设计 请充分利用 go并发优势，比如可以提供 常规版本和并发版本 给使用者自由选择
- 注意考虑 函数应用过程中的 并发性能和并发安全。
- 必须通过 整个项目的单测，参考 [./godash.go](./godash.go)

## 联系我们

- 微信号 `RobotYang7` (加我进微信群，需备注“来自`Github`”)。

<img width="350" src="./doc/微信二维码.jpg">

## 好评赞赏

- 如果该项目对你有帮助，希望获得您肯定的赞赏， 我们将有动力提供更多支持！

<img width="350" src="./doc/微信赞赏码.jpg">


--- 

## 推荐的包

#### 后端

- 引用的包
```shell
github.com/axgle/mahonia //字符集转换
github.com/bwmarrin/snowflake //雪花ID生成，参考 godash\dashrand\snow_test.go
github.com/go-playground/validator/v10 //结构和参数字段规则验证器
github.com/go-redis/redis/v8 //redis客户端
github.com/goinggo/mapstructure //Map转结构体，参考 godash\dashconv\mapstructure.go
github.com/google/uuid //UUID生成，参考 godash\dashrand\uuid_test.go
github.com/shirou/gopsutil/v3 //读取系统性能指标
	- "github.com/shirou/gopsutil/v3/cpu"
	- "github.com/shirou/gopsutil/v3/disk"
	- "github.com/shirou/gopsutil/v3/mem"
github.com/spf13/cast //各种类型转换
github.com/stretchr/testify //单测工具
  - "github.com/stretchr/testify/assert" //类型断言，参考 各个_test.go
golang.org/x/exp //go官方实验性或废弃的包
	- "golang.org/x/exp/constraints" //泛型的类型约束
golang.org/x/text //go官方文本处理的补充Go库
	- "golang.org/x/text/encoding/simplifiedchinese" //简体中文编解码器
	- "golang.org/x/text/transform" //各种字符集的阅读器和编写器的封装
google.golang.org/grpc //gprc库
	- "google.golang.org/grpc/codes" //grpc标准规范错误码
	- "google.golang.org/grpc/status" //grpc状态对象
gopkg.in/yaml.v3 //读取yaml配置文件
```

- 推荐的包
```shell
github.com/fatih/structs //结构体转Map
github.com/imdario/mergo //结构体合并，参考 godash\dashconv\mergo_test.go
github.com/spf13/viper //配置文件读取和管理
golang.org/x/sync //go官方sync包的扩展
	- "golang.org/x/sync/errgroup" //用于并发执行多个任务‌，参考 godash\dashdemo\errgroup_test.go
github.com/deckarep/golang-set/v2 //用map实现的集合及其相关操作
```

#### 前端

- 引用的包
```shell
crypto-js //加解密、哈希
mocha //单测
assert //断言
```

- 推荐的包
```shell
```

## 友情链接

- [iThings](https://github.com/i4de/ithings)