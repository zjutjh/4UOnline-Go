# 开启CGO
CGO_ENABLED=1

# Go 文件
TARGET=main

# 默认目标
all: build

# 检测 GCC
check-gcc:
	@command -v gcc >/dev/null 2>&1 || { echo "Error: gcc is not installed. Please install gcc and try again."; exit 1; }

# 构建目标
build:
	check-gcc
	@echo "Building $(TARGET)..."
	go build -v -o $(TARGET) .

# 清理生成的文件
clean:
	@echo "Cleaning up..."
	rm -f $(TARGET)

# 运行程序
run: build
	@echo "Running $(TARGET)..."
	./$(TARGET)
	
# 格式化代码并检查风格
fmt:
	@echo "Formatting Go files..."
	gofmt -w .
	gci write . -s standard -s default
	@echo "Running Lints..."
	golangci-lint run