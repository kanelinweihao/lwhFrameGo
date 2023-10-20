<!-- README -->
# LwhUtils

## 1.项目简述

| 条目 | 说明 |
| :- |:-----------------------|
| 中文名称 | 林为豪自用框架GO |
| 英文名称 | LwhFrameGo |
| 项目作者 | 林为豪 |
| 启动时间 | 2023.08 |
| 编程语言 | GO |
| 开发框架 | 原生 |
| 项目类型 | 自用框架 |
| 简要描述 | 自制自用框架,复用以便快速开发 |

## 2.文件结构

- | 根目录
    - app | 程序代码
        - api | 后台逻辑
        - conf | 配置
        - factory | 工厂
        - utils | 工具
        - web | 前台页面
    - res | 静态资源文件
        - env | 环境配置
        - ico | 图标
        - privateKey | 私钥
        - view | 视图
    - tmp | 临时文件
        - bak | 备份
        - downloads | 待下载
        - logs | 日志
        - uploads | 待上传
    - .editorconfig | 格式规范
    - .gitignore | git忽略文件
    - go.mod | GO依赖目录
    - go.sum | GO依赖哈希
    - main.go | 总入口
    - readme.md | 项目描述
    - rsrc.syso | 图标文件

## 3.常用命令

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
go mod init go.lwh.com/linweihao/lwhFrameGo
```

#### 引入外部依赖

mysql_1
```
go get -u github.com/go-sql-driver/mysql
```
mysql_2
```
go get -u github.com/jmoiron/sqlx
```
ssh
```
go get -u golang.org/x/crypto/ssh
```
redis
```
go get -u github.com/go-redis/redis
```
decimal 高精度
```
go get -u github.com/shopspring/decimal
```
excel
```
go get -u github.com/xuri/excelize/v2
```
rsrc 打包指定图标
```
go get -u github.com/akavel/rsrc
```

#### 编译打包

创建图标资源文件
```
rsrc -manifest ./res/ico/main.manifest -ico ./res/ico/icon_go.ico -o rsrc.syso
```
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
GOOS=windows GOARCH=amd64 go build -o 林为豪自用框架GO.exe
```

#### 执行源码

```
go run main.go
```

