# 变量定义
APP_NAME = gozz
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
	$(GO) run cmd/codepractice/main.go

# 构建项目
build:
	$(GO) build -gcflags=$(GOFLAGS) -ldflags=$(LDFLAGS) -o cmd/$(APP_NAME) ./pkg

# 项目依赖
deps:
	go get 