[toc]
## 简介
Golang 作为开源编程语言，在 2012 年编程语言 Go 1 发布之后，一直在不断发展. 并伴随着云原生而广泛进入各大公司. 在互联网发展至今，软件系统不断膨胀. 各个编程语言也都逐步支持了依赖包管理工具，从而加快软件迭代速度，提升整体研发效率. 依赖包管理工具能够快速的使用公司内部的依赖模块或者开源库, 极大的提升软件团队开发效率，减少迭代成本. 

Go 语言的包管理工具也伴随着开源社区的讨论、贡献不断进化，从最初单一的 GOPATH 目录的"**GOPATH 模式**"，到加入 **vendor** 目录，用于将依赖包与工程保存到同一个目录树下，最终官方结合社区版本的包管理工具正式将 **go mod** 融入 go 语言官方版本中，更好的支持多版本的依赖管理. 

后文将主要针对当期包管理机制中的 **GOPATH、vendor、go mod** 进行介绍，区分它们在不同情况下所使用依赖包.

## Go package 管理发展史
### Go 包管理迭代历程
复用代码一直是软件开发中提升效率的重要方法，大型工程的开发离不开积累的各种开源、闭源的依赖库. go 语言也不例外，go 的包管理随 go 的诞生后，经过了一系列的迭代. 迭代历程如下:

1. 2012年3月 Go 1 发布，此时没有版本的概念
2. 2013年 Golang 团队在 FAQ 中提议开发者保证相同 import path 的兼容性，后来成为一纸空文
3. 2013年10月 Godep
4. 2014年7月 glide
5. 2014年 有人提出 external packages 的概念，在项目的目录下增加一个 vendor 目录来存放外部的包
6. 2015年8月 Go 1.5 实验性质加入 vendor 机制
7. 2015年 有人提出了采用语义化版本的草案
8. 2016年2月 Go 1.6 vendor 机制 默认开启
9. 2016年5月 Go 团队的 Peter Bourgon 建立委员会，讨论依赖管理工具，也就是后面的 dep
10. **2016年8月 Go 1.7: vendor 目录永远启用**
11. 2017年1月 Go 团队发布 Dep，作为准官方试验
12. 2018年8月 Go 1.11发布 Modules 作为官方试验
13. 2019年2月 Go 1.12发布 Modules 默认为 auto
14. **2019年9月 Go 1.13 版本默认开启 Go Mod 模式**
<img src="../img/go 包管理发展史.png">


### 发展里程碑

在 go 的包管理历程中，比较重要的是分别”**GOPATH 模式**“、**vendor 特性**、 以及正式在 go 1.13 中加入的 **go mod** .  **GOPATH、Vendor、Go module** 作为官方支持机制，也是当前最常用的几个包管理方法，这些机制在最新的版本中都会被使用到. 通过 **GO111MODULE 、go.mod 、 build -mod** 等参数可以决定使用 go mod 模式或者 GOPATH 模式，并选择寻找依赖的位置(**vendor、GOPATH/src 或者 GOPATH/pkg/mod**).

<img src="../img/go 包管理发展里程碑.jpeg">

#### GOPATH 模式
GOPATH 目录是所有工程的公共依赖包目录，所有需要编译的 go 工程的依赖包都放在 GOPATH 目录下. 但这样不同的工程可能需要不同依赖包、不同版本的依赖包，若多个工程使用同一台机器编译，使用同一个 GOPATH 目录，将会使得 GOPATH 目录越来越臃肿，需要在编译工程时，将所有依赖包下载到 GOPATH 目录.

#### Vendor 特性
为了解决 GOPATH 模式下，多个工程需要共享 GOPATH 目录，无法适用于各个工程对于不同版本的依赖包的使用，不便于更新某个依赖包. go 1.6 之后开启了 vendor 目录. 每个工程可以将依赖包直接放到工程子目录 vendor 中，这样不同的工程可以存放自己需要的各种依赖包到 vendor 目录中，互不影响，当工程放到其它机器上进行开发、编译时，也不在需要在花时间下载所有需要的依赖了.  vendor 目录解决了工程依赖打包的问题，可将依赖与工程一起打包，减少下载依赖的时间, 它同时被 "GOPATH 模式" 和 "Go Mod 模式" 支持.

#### Go Module 包管理
vendor 目录下的依赖包还是需要手动加入，也没有依赖包的版本记录，那么 vendor 下的依赖包的进行升级更新也还是有困难，这些对于开发者来说都不是很友好. 并且也已经有社区针对性的做出了 go 的自动管理工具. 而 go 语言之外的其它编程也基本都已经有自动化的包管理工具了.
经过官方与社区的迭代，在 1.13 版本后，包管理工具 go mod 正式被 go 官方并默认开启，成为官方的自动管理工具. go mod 解决了必须将工程放在 GOPATH 目录下的问题，可以在 go.mod 中配置工程的 module 名. 在  go.mod 文件中记录了不同版本依赖包的信息，支持自动下载指定的依赖包，可以自动扫描工程依赖的包信息加入 go.mod.  在 go get 依赖包后也会自动加入 go.mod ，同时 go mod 也兼容了 vendor 目录，可以使用 go mod vendor 命令自动将依赖包放在工程子目录的 vendor 目录下. go mod 对 vendor 目录的支持可以很方便的将老的 GOPATH 模式的工程转移到新的包管理工具 go mod 中.



GOPATH、go vendor、go mod 是 go 包管理发展中三个重要阶段. 即使新版本的 golang 都默认使用 go mod 进行管理，但是这三个特性都依然保存，GOPATH 作为最基础的包管理方式，当 go mod 没有开启时，依然可以使用 GOPATH 目录, 而 vendor 特性则用于将依赖保存在工程内，将工程与依赖包独立的打包.

## GOPATH 模式应用
对于 golang 的工程，若没有开启 go mod，则工程必须放在 GOPATH/src 目录下，工程本身也将作为一个依赖包，可以被其它 GOPATH/src 目录下的工程引用. vendor 特性则作为 GOPATH 模式的一个补充.  GOPATH、vendor 目录下均可以存放 go 的第三方依赖包，那么当 vendor 目录和 GOPATH 目录均存在依赖包时，是如何进行选择的?

在 "GOPATH 模式"下，执行 go build 或 go run 时，在 vendor 目录、GOPATH 目录、GOROOT 目录都可能存在依赖库(标准库、第三方库等)，将依次按照如下的目录过程寻找引用的依赖:

 1. 在当前目录下的 vendor 目录查找依赖的 package
 2. 当前目录不存在 vendor 目录，则去上一级目录寻找
 3. 重复步骤 2 直到进入 $GOPATH/src 目录
 4. 没有在 vendor 目录中查找到依赖包，则进入 $GOROOT 目录查找依赖包
 5. $GOROOT 目录也没有依赖包，则进入 $GOPATH 目录寻找依赖包
<img src="../img/go 依赖包寻找顺序.png">

## Go Mod 模式应用
go mod 在 1.11 版本中试验性加入 go ，在 1.13 版本后正式作为官方的包管理工具. 对于它的使用，需要了解 go.mod 文件、 GO111MODULE 变量、go build -mod 命令以及 GOPATH/pkg/mod ，它们决定了依赖的版本、依赖包存储位置以及编译过程的依赖包.

### 环境变量 GO111MODULE
go 通过环境变量 GO111MODULE 的 3 个值 off、on、auto 来决定是否使用 go mod. 这三个值影响分别如下:

 - **GO111MODULE=off**：关闭 go modules 功能，在编译的时候仍旧在 $GOPATH/src 或者 vendor 目录中寻找依赖. 这种包管理模式为 "GOPATH 模式".
 - **GO111MODULE=on**：开启 go modules 功能，在编译时不会在 $GOPATH/src 中寻找依赖. 将在项目根目录生成 go.mod 文件。同时，依赖包不再存放在 $GOPATH/src 目录，而是存放在 $GOPATH/pkg/mod 目录，多个项目可以共享缓存的 modules。
 - **GO111MODULE=auto**：默认值，在 go v1.13 及之后版本中，如果工程目录下包含 go.mod 文件或者位于包含 go.mod 文件的目录下，则开启 go modules 功能. 在 go v1.11 中 auto 值需要工程在 GOPATH/src 之外的目录中才会开启 go mod，以确保兼容性.
### go mod 命令行
#### go mod init
使用 go mod 的工程，其目录下都包含有 go.mod 文件. 对于一个新的工程，若要开启 go mod 管理包依赖. 则需要创建 go.mod 文件，go 1.11 版本之后提供了 go mod 相关命令操作 go mod. 对于初始化 go mod 管理，可以使用 go mod init 命令. 对于  go mod init，可以对其指定工程的 module 名. 或者当其处于 GOPATH/src 目录下时，可以使用默认值.

#### GOPATH 目录中

当使用 go mod 的工程放在 GOPATH/src 目录下，可以直接用 go mod init 进行初始化，将自动检测 $GOPATH/src 后的目录作为包的 module. 命令如下即可:
``` basic
     go mod init
```

GOPAHT 目录外

工程在 GOPATH 之外使用 go mod，在进行 mod 初始化时，需要给当前工程指定 moudle 目录. 如下示例:

go mod init github.com/repo/package 
在 go mod 包管理过程中，常使用下面三个命令 init、tidy、vendor，它们可以完成 go mod 的初始化、依赖文件检查更新、以及自动建立  vendor 目录.

``` basic
   // 常用命令
   go mod init        // 初始化 go.mod，将开启 mod 使用
   go mod tidy        // 添加或者删除 modules，取决于依赖的引用
   go mod vendor      // 复制依赖到 vendor 目录下
    
   // 其它命令
   go mod download  // 下载 module 到本地
   go mod edit     //  编辑 go.mod
   go mod graph    //  打印 modules 依赖图
   go mod verify   //  验证依赖
   go mod why      //  解释依赖使用
```

 

go.mod 文件解析
使用 go mod init 进行 go 工程的 mod 初始化后，在目录下将生成 go.mod 文件，它记录了当前工程的 module 名，使用的 go 版本，依赖的 go package 以及进行过 replace 替换的工程. 如下是 go.mod 文件记录包依赖的样例:
```basic
module github.com/repo/package
 
go 1.14
 
require (
    github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
    github.com/satori/go.uuid v1.2.1-0.20181028125025-b2ce2384e17b
    gopkg.in/check.v1 v1.0.0-20200902074654-038fdea0a05b // indirect
)
 
replace golang.org/x/crypto v0.0.0-20181203042331-505ab145d0a9 => github.com/golang/crypto v0.0.0-20181203042331-505ab145d0a9
```

一般来说，并不需要手动修改 go.mod 文件，通过在工程里使用 go get 命令、go mod tidy 命令，都将会把依赖自动记录到 go.mod 文件中.
```basic
go mod tidy // 添加工程中使用到的 go 依赖包，并删除未使用过得依赖包.
go get github.com/repo/package@branch // 下载指定版本的 go 依赖包
```
GOPATH 目录下除了 src 存放工程(依赖) 源代码，还有 bin、pkg 俩目录，其中 bin 目录存在的是执行 go install 后的可执行文件，而 pkg 则存放的依赖包编译的中间文件，一般文件后缀是 .a，用于工程的编译，一般 pkg/ 目录下首先是平台目录如 pkg/linux_amd64. 而在引入 go mod 后，GOPATH/pkg/ 目录下会增加 mod 目录(GOPATH/pkg/mod)存放各版本的依赖包的缓存文件，它将优先用于代码的编译, 当 pkg/mod 目录下不存在指定的依赖包缓存时，才会依次去 vendor 目录、GOPATH/src 目录拷贝指定版本的工程到 pkg/mod 下，再进行编译.

### go build -mod 编译模式选择
开启 GO111MODULE=on 后，go build 将使用 mod 模式寻找依赖包进行编译，GOPATH/src 目录下的依赖将是无效的. 其中 go build 可以携带 -mod 的 flag 用于选择不同模式的 mod 编译. 包括 -mod=vendor、-mod=mod、-mod=readonly. 默认情况下，使用的是 -mod=readonly 模式. 但在 go 1.14及以上的版本中，如果目录中出现了 vendor 目录，将默认使用 -mod=vendor 模式进行编译. 三种 go build 的 flag 如下：

#### -mod=readonly
只读模式，如果待引入的 package 不在 go.mod 文件的列表中. 不会修改 go.mod  ，而是报错. 此外，若模块的 checksum 不在 go.sum 中也会报错. 这种模式可以在编译时候避免隐式修改 go.mod.

#### -mod=vendor
 mod=vendor 模式下. 将使用工程的 vendor 目录下的 package 而不是 mod cache( GOPATH/pkg/mod) 目录. 该模式下编译，将不会检查 go.mod 文件下的包版本. 但是会检查 vendor 目录下的 modules.txt(由 go mod vendor 生成). 在 go.1.14 及更高版本，若存在 vendor 目录，将优先使用 vendor 模式.

#### -mod=mod
mod=vendor 模式下，将使用 module cache，即使存在 vendor 目录，也会使用 GOPATH/pkg/mod 下的package，若 package 不存在，将自动下载指定版本的 package.

<img src="../img/go vendor、go mod.png">

## docker 编译镜像
无论 "GOPATH 模式" 还是 "GO Module 模式" 管理 golang 的依赖，它们也只能管理使用 golang 编写的依赖工程. 若使用 cgo 引入了 c/c++ 等的依赖，或者需要使用不同版本的 go 进行编译. 这时就需要依赖编译的基础系统环境了. 使用 docker 镜像可以管理不同版本的 golang 编译环境.  通过制作基础的编译镜像，镜像中安装基础的 golang 环境，并安装部分第三方 lib 依赖， 则可以在任意安装了 docker runtime 环境的机器上使用基础镜像完成 go 工程的编译，生成二进制文件. 再使用进制文件进一步构建业务服务的镜像.

结合 docker 编译镜像以及 go mod 包管理，可以很好的完成对 golang 依赖包、依赖环境的管理

<img src="../img/docker 镜像编译.png">

## 参考资料
1. [一张图看懂 go 包管理发展史](https://www.cyningsun.com/09-07-2019/package-management.html)
2. [入坑Go语言（二）—— 包机制的理解](https://www.jianshu.com/p/bc2bcfaf2a0f)
3. [Go Vendor简介](https://juejin.im/post/6860377811488604168)
4. [go mod 使用](https://juejin.im/post/6844903798658301960)
5. [GO项目目录下bin,pkg,src从何而来](https://hosword.github.io/2015/10/28/GO%E9%A1%B9%E7%9B%AE%E7%9B%AE%E5%BD%95%E4%B8%8Bbin%E3%80%81pkg%E3%80%81src%E4%BB%8E%E4%BD%95%E8%80%8C%E6%9D%A5/)
6. [Go语言包管理简史](https://www.imooc.com/article/292880)
7. [vendor directory in Go](https://medium.com/@gophertuts/vendor-directory-in-go-723de6cab46a)
8. [Go modules](https://systemdump.io/posts/2018-07-22-go-modules)
9. [Go go mod 終於不會再被GOPATH綁死了](https://tedmax100.github.io/2019/10/09/Go-go-mod-%E7%B5%82%E6%96%BC%E4%B8%8D%E6%9C%83%E5%86%8D%E8%A2%ABGOPATH%E7%B6%81%E6%AD%BB%E4%BA%86/)
10. [Go的包管理工具（三）：Go Modules](https://juejin.im/post/6844903791502819341)
11. [关于Go Modules，看这一篇文章就够了](https://zhuanlan.zhihu.com/p/105556877)
12. [拜拜了，GOPATH君！新版本Golang的包管理入门教程](https://zhuanlan.zhihu.com/p/60703832)
13. [Golang包管理工具govendor的使用&go mod](https://www.jianshu.com/p/ac06dcb34d39)
14. [golang 源码: go help modules](https://github.com/golang/go/blob/master/src/cmd/go/internal/modload/help.go)