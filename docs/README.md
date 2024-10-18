## 项目说明


### 项目结构


### 如何运行
1. 克隆该项目
```
git clone https://github.com/zjutjh/4UOnline-Go.git
```
2. 更改配置文件，并按注释要求填写
```
mv conf/config.yaml.example conf/config.yaml
```
3. 运行后端
```
go run main.go
```
4. 检查后端
使用[golangci-lint](https://golangci-lint.run/)检查代码
```
golangci-lint run --config .golangci.yml
```
5. 打包后端
```
go build -o 4u main.go
```

