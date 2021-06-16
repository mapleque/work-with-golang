在Mac安装Go语言执行环境
====

Go官方提供了详细的[安装文档](https://go-zh.org/doc/install)。

下面重点介绍使用homebrew安装Go并管理版本的方法。

Homebrew是一个面向MacOS的包管理工具，官网 https://brew.sh/ 有详细安装使用方法说明。

在国内使用Homebrew，建议更改源：
```sh
cd "$(brew --repo)"
git remote set-url origin https://mirrors.ustc.edu.cn/brew.git 

cd "$(brew --repo)/Library/Taps/homebrew/homebrew-core"
git remote set-url origin https://mirrors.ustc.edu.cn/homebrew-core.git 

brew update
```

查看可用的Go版本：
```sh
brew search go
```

这里将可以看到一系列可以安装的Go版本：
```sh
go go@1.10 go@1.9 go@1.8 go@1.4
```

通常直接安装最新版本即可：
```sh
brew install go
```

安装完成后，即可在命令行执行go命令了：
```sh
go version
```

当go更新了新版本，需要升级的时候，执行：
```sh
brew upgrade go
```

这个命令将安装Homebrew所管理的最新版本的go，并替换掉原来安装的go。

这里建议用户使用go1.13以上版本，因为从这个版本开始，
go module已经作为官方包管理工具集成并配置，且默认是开启的。

安装完成后，用户就可以编辑hello world文件并执行了。

```shell
mkdir hello-world
cd hello-world
# 选用适合的编辑器编辑main.go文件
go run main.go
```

```golang
package main

import "fmt"

func main() {
	fmt.Println("hello world!")
}
```

对于国内用户，建议添加go proxy设置：
```shell
go env -w GOPROXY=https://goproxy.cn
```
