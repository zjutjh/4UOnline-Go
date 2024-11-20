## 项目说明


### 项目结构
```
4u-go
├── LICENSE                     # 许可证文件
├── README.md                   # 项目概述和说明文件
├── app                         # 应用主目录
│    ├── apiException            # 自定义API异常处理模块
│    ├── config                  # 应用配置模块
│    ├── controllers             # 控制器层，负责处理业务逻辑
│    ├── midwares                # 中间件，用于请求处理的拦截和过滤
│    ├── models                  # 数据模型定义
│    ├── services                # 服务层，包含业务逻辑
│    └── utils                   # 工具包，存放通用的辅助函数
├── config                      # 配置文件目录
│    ├── api                     # 外部服务的API相关配置
│    ├── config                  # 全局viper配置
│    ├── database                # 数据库相关配置
│    ├── redis                   # Redis相关配置
│    ├── router                  # 路由配置
│    ├── session                 # 会话管理配置
│    └── wechat                  # 微信配置
├── config.example.yaml         # 配置文件示例
├── docs                        # 项目文档目录
│    └── README.md               # 文档说明文件
├── go.mod                      # Go模块依赖配置文件
├── go.sum                      # Go模块依赖的校验文件
├── logs                        # 日志目录
└── main.go                     # 项目主入口

```

### 如何运行
1. 克隆该项目的开发分支
```
git clone -b dev https://github.com/zjutjh/4UOnline-Go.git
```
2. 更改配置文件，并按注释要求填写database、redis、user和wechat的配置(user(用户中心)配置询问部长团，并要提供个人学号，wechat不为空就行)
```
mv config.example.yaml config.yaml
```
3. 首次运行后端时，在初始化数据库后，在config表插入两条`encryptKey`值为`16位的整数倍的字符串`和`initKey`值为`True`的记录
```
go run main.go
```
4. 后续直接运行就行
```
go run main.go
```
5. 每次提交commit前，先运行以下代码检查后端
使用[golangci-lint](https://golangci-lint.run/)(v1.62.0)检查代码
```
golangci-lint run --config .golangci.yml
```
6. 打包后端到服务器运行
```
SET CGO_ENABLE=0
SET GOOS=linux
SET GOARCH=amd64
go build -o 4u main.go
```

