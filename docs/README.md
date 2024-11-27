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

### 如何参与开发

1. 克隆该项目并切换到dev分支，然后切出自己的分支进行开发

```shell
git clone https://github.com/zjutjh/4UOnline-Go.git
cd 4UOnline-Go

git checkout dev
git checkout -b <Github用户名>/dev
git push
```

在开发过程中，请确保自己分支的进度与`dev`分支同步

```shell
git pull origin
git merge dev  // 将dev分支的更改合并到自己的分支
```

若你实在无法解决冲突，可以尝试将自己的分支重置到`dev`分支

```shell
git reset --hard origin/dev
git push --force
```

2. 复制示例配置，并按注释要求填写配置文件（`user`配置询问部长团，并要提供个人学号）

```shell
/* Linux */
cp config.example.yaml config.yaml

/* Windows */
copy config.example.yaml config.yaml
```

在配置数据库后，向 config 表插入如下两条记录来完成初始化

| key | value |
|---|---|
| encryptKey | *16位的整数倍的字符串 |
| initKey | True |

3. 启动程序

```shell
go run main.go
```

4. 每次提交 commit 前，先运行以下命令来格式化代码并检查规范（需要安装 [gci](https://github.com/daixiang0/gci) 和 [golangci-lint](https://golangci-lint.run/)）

```
gofmt -w .
gci write . -s standard -s default
golangci-lint run --config .golangci.yml
```

5. 打包后端到服务器运行

```
SET CGO_ENABLE=0
SET GOOS=linux
SET GOARCH=amd64
go build -o 4u main.go
```
