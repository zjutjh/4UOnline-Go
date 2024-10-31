# 设置环境变量
CGO_ENABLED=0
GOOS=linux
GOARCH=amd64

# Go 文件
TARGET=main

# 默认目标
all: build

# 构建目标
build:
	@echo "Building $(TARGET)..."
	go build -o $(TARGET) $(TARGET).go

# 编译为 Linux 目标
build-linux:
	@echo "Building $(TARGET) for $(GOOS)/$(GOARCH)..."
	GOOS=$(GOOS) GOARCH=$(GOARCH) CGO_ENABLED=$(CGO_ENABLED) go build -o $(TARGET) $(TARGET).go

# 清理生成的文件
clean:
	@echo "Cleaning up..."
	rm -f $(TARGET)

# 运行程序
run: build
	@echo "Running $(TARGET)..."
	./$(TARGET)
