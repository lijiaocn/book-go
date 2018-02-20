---
layout: post
title: "01-安装Go"
author: lijiaocn
createdate: 2017/12/16 15:39:52
changedate: 2018/02/20 15:53:22
categories: chapter01
tags:
keywords:
description: 

---

* auto-gen TOC:
{:toc}

## 安装Go

下面是在linux系统上安装go的步骤，在mac上安装也是类似的过程。

## 下载软件包

到golang官方网站的[下载页][1]下载软件包，golang的网站`https://golang.org`需要翻墙访问。

如果不能翻墙，可以到golang中国的[下载页][2]中下载，golang中国的网址是`https://www.golangtc.com`。

下载页列出每个版本的软件包的信息：

	go1.9.2.src.tar.gz           Source                         16MB   665f184bf8ac89986cfd5a4460736976f60b57df6b320ad71ad4cef53bb143dc
	go1.9.2.darwin-amd64.tar.gz  Archive     macOS     x86-64   98MB   73fd5840d55f5566d8db6c0ffdd187577e8ebe650c783f68bd27cbf95bde6743
	go1.9.2.darwin-amd64.pkg     Installer   macOS     x86-64   97MB   8b4f6ae6deae1150d2e341d02c247fd18a99af387516540eeb84702ffd76d3a1
	go1.9.2.linux-386.tar.gz     Archive     Linux     x86      88MB   574b2c4b1a248e58ef7d1f825beda15429610a2316d9cbd3096d8d3fa8c0bc1a
	go1.9.2.linux-amd64.tar.gz   Archive     Linux     x86-64   99MB   de874549d9a8d8d8062be05808509c09a88a248e77ec14eb77453530829ac02b
	go1.9.2.linux-armv6l.tar.gz  Archive     Linux     ARMv6    86MB   8a6758c8d390e28ef2bcea511f62dcb43056f38c1addc06a8bc996741987e7bb
	go1.9.2.windows-386.zip      Archive     Windows   x86      92MB   35d3be5d7b97c6d11ffb76c1b19e20a824e427805ee918e82c08a2e5793eda20
	go1.9.2.windows-386.msi      Installer   Windows   x86      79MB   020ea4a53093dd98b5ad074c4e493ff52be0aa71eee89dc24ca7783cb528de97
	go1.9.2.windows-amd64.zip    Archive     Windows   x86-64   104MB  682ec3626a9c45b657c2456e35cadad119057408d37f334c6c24d88389c2164c
	go1.9.2.windows-amd64.msi    Installer   Windows   x86-64   90MB   daeb761aa6fdb22dc3954fd911963b347c44aa5c6ba974b9c01be7cbbd6922ba

`src`表示源码，`darwin`表示mac系统，`linux`表示linux系统，`windows`表示windows系统，`386`表示32位平台，`amd64`表示64位平台。

这里以linux版本为例，选择了`linux`和`amd64`，即运行在64位平台上的linux系统。

>现在的CPU几乎都是64位的，64位运行效率更高，所以应当优先选择amd64，除非非常确定是要在32平台上安装go。

下载了安装包之后首先做一下校验（这是一个良好的习惯！）。

	$ sha256sum go1.9.2.linux-amd64.tar.gz
	de874549d9a8d8d8062be05808509c09a88a248e77ec14eb77453530829ac02b  go1.9.2.linux-amd64.tar.gz

得到的校验码与网站上公布的校验码核对无误之后，才可以放心的使用。

> 为什么要对下载的文件做校验？
>
> 互联网其实充满了各种秘密的，我们从网络上下载下来的文件很有可能是被掉过包的。
> 很多人都有做这种事情的动力，譬如黑客、运营商、捣蛋分子。
> 使用一个被掉过包的文件是非常危险的，因为我们无法预知它被做了什么手脚。
> 因此，一定要养成从官网上下载文件的习惯，而且下载完成之后要做校验。

## 解压软件包

将下载的软件包解压到指定目录中，为go单独建立一个目录。

	# mkdir -p /opt/go/1.9.2
	# tar -C /opt/go/1.9.2/ -xzf go1.9.2.linux-amd64.tar.gz

解压完成后，在新建的目标目录中，有一个go目录，这个目录里存放了go的所有文件。

	$ ls /opt/go/1.9.2/
	go

## 添加go命令

go命令位于`/opt/go/1.9.2/bin/`目录中：

	$ ls /opt/go/1.9.2/go/bin/
	go  godoc  gofmt

可以直接使用：

	$ /opt/go/1.9.2/go/bin/go version
	go version go1.9.2 linux/amd64

为了方便，最好将`/opt/go/1.9.2/go/bin/`添加到环境变量`PATH`中：

	$ echo 'export PATH=$PATH:/opt/go/1.9.2/go/bin/' >>~/.bash_profile
	$ source ~/.bash_profile

这样就可以直接在命令行执行go命令：

	$ go version
	go version go1.4.2 linux/amd64

## 参考

1. [golang.org download][1]
2. [golangtc.com download][2]

[1]: https://golang.org/dl/ "golang.org download"
[2]: https://www.golangtc.com/download "golangtc.com download"
