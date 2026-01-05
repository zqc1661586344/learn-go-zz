# 构建项目流程
## 创建项目
创建并进入项目路径，`mkdir my-project && cd my-project`

## 使用go mod初始化
使用`go mod init github.com/username/my-project`初始化项目

## 创建代码必要文件夹
使用命令创建一些必要文件夹，例如`mkdir -p cmd pkg script`

## 创建并编辑Makefile文件
Makefile中内容如下
```
# 变量定义
APP_NAME = appname
GO = go
GOFMT = gofmt
BUILD_TIME = $(shell date +%Y-%m-%d_%T)

# 构建参数
LDFLAGS = -w
GOFLAGS = all="-N -l"

# 默认执行
default: run

# 运行项目
run:
	$(GO) run cmd/main.go

# 构建项目
build:
	$(GO) build -gcflags=$(GOFLAGS) -ldflags=$(LDFLAGS) -o cmd/$(APP_NAME) ./pkg

# 项目依赖
deps:
	go get 
```

## 一个简单的项目跟路径如下
```
.
├── cmd // 用于放置程序入口
│   └── main.go
├── docs // 放置一些项目必要文件，如：编码规范
├── go.mod
├── knowledge // 知识积累文件
├── Makefile  // 项目构建文件
├── pkg // 业务代码放置路径
│   └── codepractice
│       └── practice_code_01.go
├── README.md // 项目介绍
└── script // 一些脚本放置路径，比如：用例脚本
```

# 注意事项
## 如何添加第三方库？
- 在代码中import第三方库，然后执行`go mod tidy`命令，会自动将第三方库下载到$GOPATH/pkg/mod路径下供所有项目使用，同时第三方库也会自动添加到go.mod和go.sum文件中
- 直接使用`go get {第三方库}`安装，也会将第三方库信息自动添加到go.mod和go.sum文件中

## 如何将第三方库添加到当前项目的vendor路径下？
适用于离线环境和多人共同开发环境
- 执行完上吗的添加第三方库的操作命令后，执行`go mod vendor`命令
