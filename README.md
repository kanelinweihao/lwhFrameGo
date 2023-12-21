<!-- README -->
# LwhFrameGo

## 1.项目简述
| 条目   | 说明 |
|:--|:--|
| 中文名称 | 林为豪自用框架GO |
| 英文名称 | LwhFrameGo |
| 项目作者 | 林为豪 |
| 启动时间 | 2023.08 |
| 编程语言 | GO |
| 开发框架 | 原生 |
| 项目类型 | 自用框架 |
| 简要描述 | 自制自用框架,复用以便快速开发 |

## 2.外部依赖

| 路径 | 备注  |
| :-- |:----|
| github.com/akavel/rsrc | rsrc打包指定图标 |
| github.com/go-playground/validator/v10 | validator |
| github.com/go-sql-driver/mysql | mysql_base |
| github.com/jmoiron/sqlx | mysql_sqlx |
| github.com/redis/go-redis/v9 | redis |
| github.com/shopspring/decimal | decimal高精度 |
| github.com/xuri/excelize/v2 | excel |
| golang.org/x/crypto/ssh | ssh |
| gopkg.in/gomail.v2 | email |

## 3.内部架构
- | 根目录
    - app | 程序文件
        - boot | 启动入口
        - conf | 配置信息
        - controller | 控制器
        - dict | 字典
        - router | 路由器
        - service | 主要逻辑
        - utils | 工具
    - res | 静态资源文件
        - env | 环境配置
        - ico | 图标
        - privateKey | 私钥
        - sql | 数据表语句
        - version | 版本信息
        - view | 视图
    - tmp | 临时文件
        - bak | 备份
        - downloads | 已下载
        - logs | 日志
        - uploads | 已上传
    - .editorconfig | 编辑器格式规范
    - .gitignore | git忽略文件
    - favicon.syso | 图标文件
    - go.mod | GO依赖目录
    - go.sum | GO依赖哈希
    - main.go | 总入口
    - main_test.go | 测试入口
    - readme.md | 项目描述

## 4.常用命令
#### 设置环境变量
GO依赖管理使用官方自带的MODULE,而不是早已过时的GOPATH
```
go env -w GO111MODULE=on
```
代理使用七牛云的,避免下载依赖时因被墙而失败
```
go env -w GOPROXY=https://goproxy.cn,direct
```
不使用CGO,保证打包出的执行文件无任何外部依赖
```
go env -w CGO_ENABLED=0
```
指定编译操作系统为windows
```
go env -w GOOS=windows
```
指定编译操作架构为windows的amd64架构
```
go env -w GOARCH=amd64
```

#### 初始化本项目
创建目录
```
mkdir lwhFrameGo
cd ./lwhFrameGo
```
初始化项目
```
go mod init github.com/kanelinweihao/lwhFrameGo
```
创建图标资源文件
```
rsrc -manifest ./res/ico/main.manifest -ico ./res/ico/icon_go.ico -o favicon.syso
```

#### 执行代码
运行代码
```
go run main.go
```
逻辑测试
```
go test
go test -v -run="TestGetCountCodeLine"
```
基准测试
```
go test -bench=.
```

#### 编译打包
格式化代码
```
go fmt ./...
```
更新依赖
```
go get -u ./...
```
格式化依赖,补充已使用依赖,删除未使用依赖
```
go mod tidy
```
打印所有依赖项,以json格式
```
go list -m -json all
```
执行编译
```
GOOS=windows GOARCH=amd64 go build -o LwhFrameGo.exe
GOOS=linux GOARCH=amd64 go build -o LwhFrameGo.linux
```
