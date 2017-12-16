---
layout: default
title: 02-godoc
author: lijiaocn
createdate: 2017/12/16 15:39:52
changedate: 2017/12/16 16:44:54
categories:
tags:
keywords:
description: 

---

* auto-gen TOC:
{:toc}

# 使用godoc

godoc是go的一个命令，提供了可以在本地浏览的go文档。

## 安装

godoc默认位于go的安装包中，与go命令位于同一个目录中，安装了go以后就可以直接使用。

	$ godoc -h
	usage: godoc package [name ...]
	        godoc -http=:6060
	...

## 使用

运行下面的命令，即可启动一个可以在本地访问的godoc网站：

	$ godoc -http=:6060

用浏览器打开[http://127.0.0.1:6060/][1]，就可以看到一个运行在本地的godoc站点。

本书中引用godoc连接的时候，`默认链接到运行在本地的godoc`。

## 浏览

这个运行在本地的godoc站点的内容与go的主页[golang.org][2]中的内容相同，主要由五部分组成；

	Documents
	Packages
	The Project
	Help
	Blog

[Documents][3]中包含的信息最全，需要仔细阅读。例如[Command Documentation][4]、[The Go Programming Language Specification][5]。

特别是[The Go Programming Language Specification][5]，它对go语言语法的说明是最细致精确、最权威的。

[Packages][6]是go语言的package文档。

[Project][7]介绍了go语言项目。

[Help][8]给出寻求帮助的途径。

[Blog][9]是go项目的博客，介绍了go的发展、新特性，以及每个版本的性能情况等内容。

## 参考

1. [godoc][1]
2. [golang.org][2]
3. [go Documents][3]
4. [go Command Documentation][4]
5. [The Go Programming Language Specification][5]
6. [go Packages][6]
7. [go Project][7]
8. [go Help][8]
9. [go Blog][9]

[1]: http://127.0.0.1:6060/ "godoc" 
[2]: https://golang.org "golang.org" 
[3]: http://127.0.0.1:6060/doc/ "go Documents"
[4]: http://127.0.0.1:6060/doc/cmd "go Command Documentation"
[5]: http://127.0.0.1:6060/ref/spec "The Go Programming Language Specification"
[6]: http://127.0.0.1:6060/pkg/  "go Packages"
[7]: http://127.0.0.1:6060/project/ "go Project"
[8]: http://127.0.0.1:6060/help/ "go Help"
[9]: http://127.0.0.1:6060/blog/  "go Blog"
