# WtmpViewer

`WtmpViewer` 是一个使用 Go 语言开发的工具，用于解析和查看 Linux 系统中的 `wtmp` 日志文件，并从 `secure` 和 `auth.log` 日志中查找成功的登录记录。该工具支持跨平台编译，可以在多个操作系统上运行。

## 功能简介

- **查看 `wtmp` 日志：** 解析 `wtmp` 文件，显示用户登录记录。
- **检查 `secure` 日志：** 扫描指定目录下所有以 `secure` 开头的文件，查找并输出 SSH 成功登录的记录。
- **检查 `auth.log` 日志：** 查找 `auth.log` 日志文件中的成功登录记录。

## 目录结构

```azure
├── cmd
│ ├── root.go # 定义了根命令和基础配置
│ └── check_secure.go # 定义了检查 secure 日志文件的命令
├── internal
│ └── logparser
│ ├── secure.go # 定义了 secure 日志解析功能
│ └── wtmp.go # 定义了 wtmp 文件解析功能
├── main.go # 程序入口
├── build.sh # 自动跨平台编译脚本
└── README.md # 项目说明文件
```


## 安装与编译

### 依赖

- Go 1.16 及以上版本

### 手动编译

你可以使用以下命令手动编译 `WtmpViewer`：

#### 编译为 Linux 64-bit 可执行文件

```bash
GOOS=linux GOARCH=amd64 go build -o wtmpviewer
```
#### 编译为 Windows 64-bit 可执行文件
```bash
GOOS=windows GOARCH=amd64 go build -o wtmpviewer.exe
```
#### 编译为 macOS 64-bit 可执行文件
```bash
GOOS=darwin GOARCH=amd64 go build -o wtmpviewer
```

#### 使用方法
WtmpViewer 提供了多种命令用于解析不同类型的日志文件。以下是基本的使用示例：

##### 查看 wtmp 日志
要查看 wtmp 日志，运行以下命令：

```bash
./wtmpviewer wtmp --file /var/log/wtmp
```

##### 检查 secure 日志
要扫描某个目录下所有以 secure 开头的文件并查找成功登录的记录，运行以下命令：

```bash
./wtmpviewer check-secure --directory /var/log
```

#### 检查 auth.log 日志
（假设你添加了 auth.log 的解析功能）要查找 auth.log 中的成功登录记录，运行：

```bash
./wtmpviewer check-authlog --file /var/log/auth.log
```

### 贡献
欢迎贡献代码！请提交 Pull Request 或报告 Issue。

### 许可
本项目基于 MIT 许可开源。详情请查看 LICENSE 文件。