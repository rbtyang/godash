[toc]

> Golang常用工具、函数、方法、帮助类的封装

## 已实现工具

##### arrdash
- 数组工具

##### convdash
- 类型转换

##### cryptdash
- 加密解密

##### demodash
- 我的模板

##### encodash
- 编码解码

##### envdash
- 环境变量

##### errdash
- 错误类

##### filedash
- 文件工具

##### hashdash
- 哈希工具

##### jsondash
- 序列化

##### logdash
- 日志工具

##### pagedash
- 分页类

##### randdash
- 随机生成

##### strdash
- 字符处理

##### timedash
- 时间函数

##### validdash
- 验证器

--- 

## TODO工具

##### xxx

--- 

## 推荐的包
```shell script
"github.com/stretchr/testify/assert" //单测类型断言，参考 各个_test.go
"github.com/spf13/cast" //各种类型转换
"github.com/fatih/structs" //结构体转Map
"github.com/goinggo/mapstructure" //Map转结构体，参考 godash\convdash\mapstructure.go
"github.com/imdario/mergo" //结构体合并，参考 godash\convdash\mergo_test.go
"github.com/spf13/viper" //配置文件读取和管理
"golang.org/x/sync/errgroup" //并发编程，参考 godash\demodash\errgroup_test.go
"github.com/bwmarrin/snowflake" //雪花ID生成，参考 godash\randdash\snow_test.go
"github.com/google/uuid" //UUID生成，参考 godash\randdash\uuid_test.go
```
