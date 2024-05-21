# mpip

`mpip` 是一个用 Go 编写的命令行工具，用于安装和卸载 Python 包，并自动更新 `requirements.txt` 文件。

## 功能

- 安装 Python 包并将其添加到 `requirements.txt`
- 卸载 Python 包并从 `requirements.txt` 中移除

## 安装

首先，克隆此仓库：

```sh
git clone https://github.com/litongjava/mpip.git
cd mpip
```

然后，编译 `mpip` 工具：

```sh
go build -o mpip main.go
```
这将在当前目录下生成一个名为 `mpip` 的可执行文件。请将mpip添加到path

```
go install
```
自动添加到go-path中
## 使用

### 安装包

使用 `mpip install <package>` 命令安装 Python 包，并将包名添加到 `requirements.txt` 文件中：

```sh
mpip install <package>
```

例如，要安装 `requests` 包：

```sh
mpip install requests
```

### 卸载包

使用 `mpip uninstall <package>` 命令卸载 Python 包，并从 `requirements.txt` 文件中移除包名：

```sh
mpip uninstall <package>
```

例如，要卸载 `requests` 包：

```sh
mpip uninstall requests
```

## 注意事项

- 确保你的环境中已经安装了 `pip`，并且可以从命令行访问。
- `requirements.txt` 文件会在当前目录中自动创建或更新。

## 贡献

如果你有任何改进建议或发现了问题，请提交 pull request 或 issue。

## 许可证

此项目使用 MIT 许可证。有关更多信息，请参见 [LICENSE](LICENSE) 文件。
