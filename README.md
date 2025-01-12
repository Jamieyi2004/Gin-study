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

你需要将 Go 的 `bin` 目录添加到你的 `PATH` 环境变量中。可以通过编辑 `~/.bashrc`、`~/.zshrc` 或其他 shell 配置文件来实现这一点。例如，对于 Bash 用户：

```shell
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
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

## Git 配置
```bash
git config --global user.name "Jamieyi2004" # 配置用户名
git config --global user.email "Jamieyi2004@gmail.com" # 配置用户邮箱
ssh-keygen -t rsa -C "Jamieyi2004@gmail.com"

git checkout [branch-name] # 切换分支
git push -u origin main:main # 设置上游 不同分支可以有不同上游

```