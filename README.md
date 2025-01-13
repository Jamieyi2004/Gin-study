# Gin-study
## 1 下载Go & Gin
### 1.1 下载最新的 Go 版本
使用官方提供的二进制文件（推荐），这是最直接的方法，适用于希望确保安装的是最新稳定版 Go 的用户。
首先，访问 [Go 官方下载页面](https://golang.org/dl/) 找到最新版本的下载链接。你可以使用 `wget` 或 `curl` 来下载 tarball 文件。这里假设我们下载的是 Linux 64-bit 版本：

```shell
cd /tmp
wget https://dl.google.com/go/go$(go version | awk '{print $3}' | cut -d'.' -f2,3).linux-amd64.tar.gz
# 或者如果你没有安装 go，可以手动指定版本号，例如：
# wget https://dl.google.com/go/go1.21.0.linux-amd64.tar.gz
```

### 1.2 解压并安装

解压缩下载的 tarball 文件，并将其移动到 `/usr/local` 目录下：

```shell
sudo tar -C /usr/local -xzf /tmp/go*.linux-amd64.tar.gz
```

### 1.3 设置环境变量
添加或更新以下行到你的 shell 配置文件中:
```shell
# 设置 GOROOT 为 Go 的安装目录
export GOROOT=/usr/local/go

# 添加 Go bin 到 PATH
export PATH=$PATH:$GOROOT/bin

# 设置 GOPATH 为你的工作区目录
# 注意：从 Go 1.8 开始，GOPATH 不再强制要求，但从 Go modules 引入后，它仍然是有用的。
# 如果你不指定 GOPATH，Go 会默认使用 ~/go 作为工作区。
export GOPATH=$HOME/go

# 添加 GOPATH/bin 到 PATH，以便可以运行你在 GOPATH 中安装的工具
export PATH=$PATH:$GOPATH/bin

source ~/.bashrc
```

对于 Zsh 用户：

```shell
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.zshrc
source ~/.zshrc
```

### 1.4 验证安装

最后，验证 Go 是否安装成功：

```shell
go version
```

你应该看到类似如下的输出，表示安装成功：

```
go version go1.21.0 linux/amd64
```

### 1.5 安装Gin
Go版本不能低于 1.21。
```shell
go mod init github.com/Jamieyi2004/Gin-study

export GOPROXY=https://mirrors.aliyun.com/goproxy/,direct

go get -u github.com/gin-gonic/gin
```

## 2 Git 配置
```bash
git config --global user.name "Jamieyi2004" # 配置用户名
git config --global user.email "Jamieyi2004@gmail.com" # 配置用户邮箱
ssh-keygen -t rsa -C "Jamieyi2004@gmail.com"

git checkout main # 切换分支
git add .
git commit -m "message"
git push -u origin main:main # 设置上游 不同分支可以有不同上游
git push # 设置上游以后 以后就可以省略
```

## 3 概念
- 路由（Routing）是指 Web 应用根据 HTTP 请求的 URL 和方法将其映射到相应处理函数的过程。

- 路由引擎（Router Engine）是 Web 框架中的一个组件，它负责解析传入的请求并根据预定义的规则将请求分发给适当的处理器或控制器。

- 中间件（Middleware）是在应用程序处理管道中执行的一段代码，它可以在请求到达最终目的地之前和响应返回客户端之前对它们进行处理，通常用于执行诸如身份验证、日志记录等通用任务。

- URL（统一资源定位符）的结构通常由协议、域名（或IP地址）、端口（可选）、路径、查询参数和片段标识符组成，格式如下：协议://域名:端口/路径?查询参数#片段标识符。

- map[string]interface{} 是 Go 语言中的一种数据结构，它表示一个键为字符串、值为任意类型的映射（即字典）。这里的 interface{} 是 Go 的空接口类型，它可以存储任何类型的值。因此，map[string]interface{} 可以用来创建一个动态的键值对集合，其中每个键都是字符串，而对应的值可以是任意类型的数据。

- `go mod init go-dev` 是 Go 语言中用于初始化一个新的模块的命令。具体来说：

    - `go mod`：是 Go 1.11 引入的模块支持工具，用来管理项目依赖。它替代了旧版本的 Go 中使用的 `GOPATH` 工作空间和 `vendor` 文件夹来处理依赖关系。
    - `init`：是 `go mod` 的子命令之一，用于初始化一个新的 Go 模块，在当前目录下创建一个 `go.mod` 文件。
    - `go-dev`：这是你为你的 Go 模块指定的模块路径名称。它可以是你项目的唯一标识符（例如，基于托管代码的Git仓库URL），也可以是一个描述性的名称，取决于你如何组织和分发你的代码。

    执行 `go mod init go-dev` 命令后，Go 会在当前目录创建一个名为 `go.mod` 的文件，该文件会包含类似如下的内容：

    ```
    module go-dev

    go 1.x
    ```

    这行声明表示这个 Go 模块的名称是 `go-dev`，并且指定了模块所兼容的最小 Go 版本（这里的 `1.x` 会根据你的Go版本自动填充）。

    `go.mod` 文件对于追踪项目直接依赖的其他模块及其版本非常重要。当你添加新的依赖时，这些信息会被记录在这个文件中，并且 Go 工具链会使用这个文件来确保所有开发者以及构建环境都使用相同的依赖版本。

    如果你打算分享或发布你的模块，推荐使用更具体的模块路径，比如包含你在代码托管平台上的用户名和仓库名，例如 `github.com/yourusername/go-dev`。这样可以确保模块路径的唯一性，并使得其他人更容易导入和使用你的模块。

- 在编程和软件工程中，“模块”是一个宽泛的概念，通常指的是一个独立的、可重用的代码单元或组件，它具有明确的功能性，并且可以与其他模块组合起来构建更大的系统。模块化是软件设计的一种重要原则，通过将程序划分为多个模块，可以提高代码的可维护性、可读性和可测试性。
**Go 语言中的模块**：
   在 Go 中，模块是一组相关的包（packages），这些包一起工作以提供一组功能。从 Go 1.11 开始引入了 `go mod` 工具来支持模块，允许开发者更灵活地管理项目依赖。每个模块都有一个唯一的模块路径（通常是托管该模块的 Git 仓库 URL），并且有一个 `go.mod` 文件来记录模块的元数据及其依赖关系。使用模块可以帮助解决 Go 项目的依赖管理和版本控制问题。

- 在 Go 语言中，包（package）和文件夹（directory）之间的关系通常是1:1的映射。这意味着一个包通常存在于一个单独的文件夹下，该文件夹下的所有源代码文件都属于同一个包。这种组织方式有助于保持代码的清晰性和可维护性，并且是 Go 社区广泛接受的标准。每个 Go 源文件的第一行应该是 package 声明，用来指定该文件所属的包。

- Go 编译器通过 go.mod 文件来确定代码所在的模块。当在一个项目中运行 Go 命令（如 go build、go run 或 go test）时，编译器会从当前工作目录开始查找 go.mod 文件，沿着文件系统的父目录逐级向上搜索，直到找到一个 go.mod 文件为止。这个过程被称为“模块根”的发现。一旦找到了 go.mod 文件，该文件所在目录就被认为是模块的根目录，并且所有位于此目录及其子目录下的 Go 包都被认为是属于同一个模块。   

- go build 更适合于生产环境部署和构建过程，你会得到一个可执行文件，你可以随后运行它；而 go run 则更适合于开发和测试阶段，你会编译并立即运行这个程序，你会看到输出但不会有持久的二进制文件留下。

- Go 编译器会递归地编译所有依赖的包。这意味着不仅 hello.go 会被编译，而且它所导入的所有包（包括你自己写的包）以及这些包所依赖的任何其他包也会被编译。这个过程确保了整个依赖链都被正确处理，并且生成的可执行文件包含了运行所需的所有代码。

- 在 Go 语言中，只有 main 包可以包含一个名为 main 的函数，并且这个 main 函数是程序的入口点。其他包不可以定义自己的 main 函数。如果尝试在一个非 main 包中定义 main 函数，编译器会报错，因为这违反了 Go 语言的规定。

- go.mod 文件定义了 Go 模块的元数据和依赖项，而 go.sum 文件记录了模块依赖项的版本及其校验和，确保依赖的完整性与可重复构建。
